package core

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// VMServer - сервер протоколов взаимодействия, предоставляет базовый обработчик для TCP, RPC-JSON и HTTP соединений
// данный объект не может сериализоваться и не может участвовать в операциях с операндами
type VMServer struct {
	VMMetaObj //должен передаваться по ссылке, поэтому это будет объект метаданных

	mu       sync.RWMutex
	addr     string // [addr]:port
	protocol string // tcp, tcpzip, tcptls, http, https
	done     chan error
	health   chan bool
	clients  []*VMConn // каждому соединению присваивается GUID
	lnr      net.Listener
	mux      *http.ServeMux
	srv      *http.Server
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

func (x *VMServer) Open(proto, addr string, maxconn int, handler VMFunc, data VMValuer, vsmHandlers VMStringMap) (err error) {
	// запускаем сервер
	if x.lnr != nil || x.srv != nil {
		return VMErrorServerNowOnline
	}

	x.done = make(chan error)
	x.health = make(chan bool)
	x.clients = make([]*VMConn, 0)

	x.addr = addr
	x.protocol = proto
	x.maxconn = maxconn

	switch proto {
	case "tcp", "tcpzip", "tcptls":
		gzipped := false
		if proto == "tcptls" {
			config := &tls.Config{
				Certificates: []tls.Certificate{TLSKeyPair},
			}
			x.lnr, err = tls.Listen("tcp", addr, config)
			if err != nil {
				return err
			}
		} else {
			x.lnr, err = net.Listen("tcp", addr)
			if err != nil {
				x.lnr = nil
				return err
			}
			if proto == "tcpzip" {
				gzipped = true
			}
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
						data:   data,
						gzip:   gzipped,
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
	case "http", "https":
		// TODO: https
		x.mux = http.NewServeMux()
		for k, v := range vsmHandlers {
			if f, ok := v.(VMFunc); ok {
				x.mux.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
					req := &VMHttpRequest{r: r}
					resp := &VMHttpResponse{w: w}
					args := make(VMSlice, 2)
					rets := make(VMSlice, 0)
					args[0] = resp
					args[1] = req
					var env *Env // сюда вернется окружение вызываемой функции
					err := f(args, &rets, &env)
					if err != nil && env.Valid {
						env.Println(err)
					}
				})
			}
		}
		x.srv = &http.Server{
			Addr:    addr,
			Handler: x.mux,
		}
		go x.healthSender()
		go func(s *http.Server) {
			err := s.ListenAndServe()
			x.done <- err
		}(x.srv)

	default:
		return VMErrorIncorrectProtocol
	}
	return nil
}

// Close закрываем все ресурсы и всегда возвращаем ошибку,
// которая могла возникнуть на сервере, либо во время закрытия
// !!! Эту процедуру нужно обязательно вызывать по окончании работы с сервером !!!
func (x *VMServer) Close() error {
	if x.lnr != nil {
		x.lnr.Close()
	}
	if x.srv != nil {
		x.srv.Close()
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
	x.srv = nil
	x.mux = nil
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
		var env *Env
		f(args, &rets, &env)
	}
}

func (x *VMServer) VMRegister() {
	x.VMRegisterMethod("Закрыть", x.Закрыть)
	x.VMRegisterMethod("Работает", x.Работает)
	x.VMRegisterMethod("Открыть", x.Открыть)
	// tst.VMRegisterField("ПолеСтрока", &tst.ПолеСтрока)
}

// Закрыть возвращает настоящую причину закрытия, в том числе, ошибку отстрела сервера до вызова закрытия
func (x *VMServer) Закрыть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(VMString(fmt.Sprint(x.Close())))
	return nil
}

func (x *VMServer) Работает(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(VMBool(x.IsOnline()))
	return nil
}

func (x *VMServer) Открыть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	if len(args) != 5 {
		return VMErrorNeedArgs(5)
	}
	p, ok := args[0].(VMString)
	if !ok {
		return errors.New("Первый аргумент должен быть строкой с типом канала")
	}
	adr, ok := args[1].(VMString)
	if !ok {
		return errors.New("Второй аргумент должен быть строкой-адресом")
	}
	lim, ok := args[2].(VMInt)
	if !ok {
		return errors.New("Третий аргумент должен быть числом-лимитом подключений")
	}
	var f VMFunc
	var vsm VMStringMap
	switch string(p) {
	case "http", "https":
		vsm, ok = args[3].(VMStringMap)
		if !ok {
			return errors.New("Четвертый аргумент должен быть структурой с функциями с одним аргументом-соединением, где ключ строкой - относительный путь URI")
		}
	default:
		f, ok = args[3].(VMFunc)
		if !ok {
			return errors.New("Четвертый аргумент должен быть функцией с одним аргументом-соединением")
		}
	}

	return x.Open(string(p), string(adr), int(lim), f, args[4], vsm)
}
