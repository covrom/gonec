package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/covrom/gonec/names"
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
	case "закрыто":
		return VMFuncMustParams(0, c.Закрыто), true
	case "идентификатор":
		return VMFuncMustParams(0, c.Идентификатор), true
	}

	return nil, false
}

func (x *VMConn) Handle(f VMFunc) {
	args := make(VMSlice, 1)
	rets := make(VMSlice, 0)
	args[0] = x
	var env *Env // сюда вернется окружение вызываемой функции
	err := f(args, &rets, &env)
	// закрываем по окончании обработки
	x.conn.Close()
	x.closed = true
	if err != nil && env.Valid {
		env.Println(err)
	}
}

func (x *VMConn) Идентификатор(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(VMString(x.uid))
	return nil
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
	hashbts := binary.LittleEndian.Uint64(b[8:16]) // hash
	lenb := binary.LittleEndian.Uint64(b[16:24])   // len

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

func (x *VMConn) Получить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
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
	binary.LittleEndian.PutUint64(hb[8:16], HashBytes(be))
	binary.LittleEndian.PutUint64(hb[16:24], uint64(len(b)))

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

func (x *VMConn) Отправить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMStringMap)
	if !ok {
		return VMErrorNeedSlice
	}
	return x.Send(v) // при ошибке вызовет исключение, нужно обрабатывать в попытке
}

func (x *VMConn) Закрыто(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(VMBool(x.closed))
	return nil
}
