package core

import (
	"bytes"
	"encoding/binary"
	"errors"
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
	data   VMValuer
}

func (c *VMConn) vmval() {}

func (c *VMConn) Interface() interface{} {
	return c.conn
}

func (c *VMConn) String() string {
	if c.closed {
		return fmt.Sprintf("Соединение (закрыто)")
	}
	return fmt.Sprintf("Соединение с %s", c.conn.RemoteAddr())
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
	case "параметры":
		return VMFuncMustParams(0, c.Параметры), true
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

type binTCPHead struct {
	Signature [8]byte //[8]byte{'g', 'o', 'n', 'e', 'c', 't', 'c', 'p'}
	Hash      uint64  //хэш зашифрованного тела
	Len       int64   //длина тела
}

func (x *VMConn) Receive() (VMStringMap, error) {

	rv := make(VMStringMap)
	var buf bytes.Buffer

	var head binTCPHead

	err := binary.Read(x.conn, binary.LittleEndian, &head)
	if err != nil {
		if err == io.EOF {
			x.closed = true
			x.conn.Close()
			err = VMErrorEOF
		}
		return rv, err
	}

	// проверяем целостность полученного сообщения
	// сначала идет заголовок
	// затем тело, шифрованное по AES128

	if head.Signature != [8]byte{'g', 'o', 'n', 'e', 'c', 't', 'c', 'p'} {
		return rv, errors.New(VMErrorIncorrectMessage.Error() + " - неверная сигнатура")
	}

	buf.Reset()
	_, err = io.CopyN(&buf, x.conn, head.Len)
	if err != nil {
		if err == io.EOF {
			x.closed = true
			x.conn.Close()
		}
		return rv, err
	}

	b := buf.Bytes()

	// хэш зашифрованного
	hb := HashBytes(b)
	if hb != head.Hash {
		// log.Println("in", hb, b)
		return rv, errors.New(VMErrorIncorrectMessage.Error() + " - не совпал хэш")
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
	hs := HashBytes(be)

	head := binTCPHead{
		Signature: [8]byte{'g', 'o', 'n', 'e', 'c', 't', 'c', 'p'},
		Hash:      hs,
		Len:       int64(len(be)),
	}

	// log.Println("out", hs, be)

	err = binary.Write(x.conn, binary.LittleEndian, head)
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

func (x *VMConn) Параметры(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(x.data)
	return nil
}
