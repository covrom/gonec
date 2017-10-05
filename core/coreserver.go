package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"

	"github.com/covrom/gonec/names"
	uuid "github.com/satori/go.uuid"
)

var (
	VMErrorServerNowOnline   = errors.New("Сервер уже запущен")
	VMErrorServerOffline     = errors.New("Сервер уже остановлен")
	VMErrorIncorrectClientId = errors.New("Неверный идентификатор соединения")
	VMErrorIncorrectMessage  = errors.New("Неверный формат сообщения")
	VMErrorWrongType         = errors.New("По каналу TCP можно передавать только структуры")
)

type VMConn struct {
	conn   net.Conn
	id     int
	closed bool
	uid    string
}

func (c *VMConn) vmval() {}

func (c *VMConn) Interface() interface{} {
	return c.conn
}

func (c *VMConn) String() string {
	if c.closed {
		return fmt.Sprintf("Соединение с клиентом (закрыто)")
	}
	return fmt.Sprintf("Соединение с клиентом %s", c.conn.RemoteAddr())
}

func (c *VMConn) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	case "получить":
		return VMFuncMustParams(0, c.Получить), true
	case "отправить":
		return VMFuncMustParams(1, c.Отправить), true
	case "закрыт":
		return VMFuncMustParams(0, c.Закрыт), true
	}

	return nil, false
}

func (x *VMConn) Handle(f VMFunc) {
	args := make(VMSlice, 1)
	rets := make(VMSlice, 0)
	args[0] = x
	f(args, &rets)
}

func (x *VMConn) Receive() (VMStringMap, error) {

	rv := make(VMStringMap)
	var buf bytes.Buffer

	_, err := io.CopyN(&buf, x.conn, 24)

	if err != nil {
		if err == io.EOF {
			x.closed = true
			x.conn.Close()
		}
		return rv, err
	}

	b := buf.Bytes()

	// проверяем целостность полученного сообщения
	// сначала идет заголовок 8 байт "gonectcp"
	// затем хэш тела, он идет 8-ю байтами
	// затем длина тела 8 байт
	// затем тело, шифрованное по AES128

	if len(b) < 24 {
		return rv, VMErrorIncorrectMessage
	}
	cstr := b[:8] // gonectcp
	if !bytes.Equal(cstr, []byte("gonectcp")) {
		return rv, VMErrorIncorrectMessage
	}
	hashbts, _ := binary.Uvarint(b[8:16]) // hash
	lenb, _ := binary.Uvarint(b[16:24])   // len

	buf.Reset()
	_, err = io.CopyN(&buf, x.conn, int64(lenb))
	if err != nil {
		if err == io.EOF {
			x.closed = true
			x.conn.Close()
		}
		return rv, err
	}

	b = buf.Bytes()

	// хэш зашифрованного
	if HashBytes(b) != hashbts {
		return rv, VMErrorIncorrectMessage
	}
	// проверили хэш, все ок - получаем VMStringMap
	bd, err := DecryptAES128(b)
	if err != nil {
		return rv, err
	}

	if err := (&rv).UnmarshalBinary(bd); err != nil {
		return rv, err
	}
	return rv, nil
}

func (x *VMConn) Получить(args VMSlice, rets *VMSlice) error {
	v, err := x.Receive()
	rets.Append(v)
	return err // при ошибке вызовет исключение, нужно обрабатывать в попытке
}

func (x *VMConn) Send(val VMStringMap) error {

	b, err := val.MarshalBinary()
	if err != nil {
		return err
	}

	be, err := EncryptAES128(b)
	if err != nil {
		return err
	}

	//хэш зашифрованного
	var hb [24]byte
	copy(hb[:8], []byte("gonectcp"))
	binary.PutUvarint(hb[8:16], HashBytes(be))
	binary.PutUvarint(hb[16:24], uint64(len(b)))

	_, err = io.Copy(x.conn, bytes.NewBuffer(hb[:]))
	if err != nil {
		if err == io.EOF {
			x.closed = true
			x.conn.Close()
		}
		return err
	}

	_, err = io.Copy(x.conn, bytes.NewReader(be))
	if err != nil {
		if err == io.EOF {
			x.closed = true
			x.conn.Close()
		}
		return err
	}
	return nil
}

func (x *VMConn) Отправить(args VMSlice, rets *VMSlice) error {
	v, ok := args[0].(VMStringMap)
	if !ok {
		return VMErrorWrongType
	}
	return x.Send(v) // при ошибке вызовет исключение, нужно обрабатывать в попытке
}

func (x *VMConn) Закрыт(args VMSlice, rets *VMSlice) error {
	rets.Append(VMBool(x.closed))
	return nil
}

// VMServer - сервер протоколов взаимодействия, предоставляет базовый обработчик для TCP, RPC-JSON и HTTP соединений
// данный объект не может сериализоваться и не может участвовать в операциях с операндами
type VMServer struct {
	VMMetaObj //должен передаваться по ссылке, поэтому это будет объект метаданных

	mu       sync.RWMutex
	addr     string // [addr]:port
	protocol string // tcp, json, http
	done     chan error
	health   chan bool
	clients  []*VMConn // каждому соединению присваивается GUID
	lnr      net.Listener
	maxconn  int
}

func (x *VMServer) String() string {
	return fmt.Sprintf("Сервер %s %s", x.protocol, x.addr)
}

func (x *VMServer) IsOnline() bool {
	return <-x.health
}

func (x *VMServer) healthSender() {
	for {
		select {
		case x.health <- true:
			runtime.Gosched()
		case e, ok := <-x.done:
			close(x.health)
			if ok {
				// перехватили ошибку, а канал не закрыт -> ретранслируем
				x.done <- e
			}
			return
		}
	}
}

func (x *VMServer) Open(proto, addr string, maxconn int, handler VMFunc) (err error) {
	// запускаем сервер
	if x.lnr != nil {
		return VMErrorServerNowOnline
	}

	x.done = make(chan error)
	x.health = make(chan bool)
	x.clients = make([]*VMConn, 0)

	x.addr = addr
	x.protocol = proto
	x.maxconn = maxconn

	switch proto {
	case "tcp":
		x.lnr, err = net.Listen("tcp", addr)
		if err != nil {
			x.lnr = nil
			return err
		}

		go x.healthSender()

		// запускаем воркер, который принимает команды по каналу управления
		// x.lnr может стать nil, поэтому, передаем сюда копию указателя
		go func(lnr net.Listener) {
			for {
				conn, err := lnr.Accept()
				if err != nil {
					x.done <- err
					return
				}

				x.mu.Lock()
				l := len(x.clients)
				if l < maxconn || maxconn == -1 {

					vcn := &VMConn{
						conn:   conn,
						id:     l,
						closed: false,
						uid:    uuid.NewV4().String(),
					}
					x.clients = append(x.clients, vcn)
					go vcn.Handle(handler)

				} else {
					conn.Close()
				}
				x.mu.Unlock()

				runtime.Gosched()
			}
		}(x.lnr)
	}
	return nil
}

// Close закрываем все ресурсы и всегда возвращаем ошибку,
// которая могла возникнуть на сервере, либо во время закрытия
func (x *VMServer) Close() error {
	if x.lnr != nil {
		x.lnr.Close()
	}
	err, ok := <-x.done // дождемся ошибки из горутины, или возьмем ее, если она уже была
	if ok {
		// канал не закрыт
		close(x.done)
	} else {
		err = VMErrorServerOffline
	}
	x.mu.Lock()
	x.lnr = nil
	// закрываем все клиентские соединения
	for i := range x.clients {
		if !x.clients[i].closed {
			x.clients[i].conn.Close()
			x.clients[i].closed = true
		}
	}
	x.clients = x.clients[:0]
	x.mu.Unlock()
	return err
}

func (x *VMServer) ClientsCount() int {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return len(x.clients)
}

func (x *VMServer) CloseClient(i int) (err error) {
	x.mu.Lock()
	defer x.mu.Unlock()
	l := len(x.clients)
	if i >= 0 && i < l {
		err = nil
		if !x.clients[i].closed {
			err = x.clients[i].conn.Close()
			x.clients[i].closed = true
		}
		return
	} else {
		return VMErrorIncorrectClientId
	}
}

func (x *VMServer) RemoveAllClosedClients() {
	x.mu.Lock()
	defer x.mu.Unlock()
	l := len(x.clients)
	for i := l - 1; i >= 0; i-- {
		if x.clients[i].closed {
			copy(x.clients[i:], x.clients[i+1:])
			nl := len(x.clients) - 1
			x.clients[nl].conn = nil
			x.clients = x.clients[:nl]
			for j := i; j < nl; j++ {
				x.clients[j].id--
			}
		}
	}
}

// ForEachClient запускает обработчики для каждого клиента, последовательно
func (x *VMServer) ForEachClient(f VMFunc) {
	x.mu.Lock()
	defer x.mu.Unlock()
	for _, cli := range x.clients {
		args := make(VMSlice, 1)
		rets := make(VMSlice, 0)
		args[0] = cli
		f(args, &rets)
	}
}

func (x *VMServer) VMRegister() {
	x.VMRegisterMethod("Закрыть", x.Закрыть)
	x.VMRegisterMethod("Работает", x.Работает)
	// tst.VMRegisterField("ПолеСтрока", &tst.ПолеСтрока)
}

// Закрыть возвращает настоящую причину закрытия, в том числе, ошибку отстрела сервера до вызова закрытия
func (x *VMServer) Закрыть(args VMSlice, rets *VMSlice) error {
	rets.Append(VMString(fmt.Sprint(x.Close())))
	return nil
}

func (x *VMServer) Работает(args VMSlice, rets *VMSlice) error {
	rets.Append(VMBool(x.IsOnline()))
	return nil
}

// TODO:
// функция обработать(соед)
//   сообщить(соед)
// конецфункции
// серв = Новый Сервер
// серв.Открыть("tcp", ":9990", 1000, обработать)