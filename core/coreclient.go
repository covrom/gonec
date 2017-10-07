package core

import (
	"errors"
	"fmt"
	"net"

	uuid "github.com/satori/go.uuid"
)

type VMClient struct {
	VMMetaObj //должен передаваться по ссылке, поэтому это будет объект метаданных

	addr     string  // [addr]:port
	protocol string  // tcp, json, http
	conn     *VMConn // каждому соединению присваивается GUID
}

func (x *VMClient) String() string {
	return fmt.Sprintf("Клиент %s %s", x.protocol, x.addr)
}

func (x *VMClient) IsOnline() bool {
	return x.conn != nil && !x.conn.closed
}

func (x *VMClient) Open(proto, addr string, handler VMFunc) error {
	switch proto {
	case "tcp":
		conn, err := net.Dial(proto, addr)
		if err != nil {
			return err
		}

		x.conn = &VMConn{
			conn:   conn,
			id:     -1,
			closed: false,
			uid:    uuid.NewV4().String(),
		}

		go x.conn.Handle(handler)

	default:
		return VMErrorIncorrectProtocol
	}
	return nil
}

func (x *VMClient) Close() {
	x.conn.conn.Close()
	x.conn.closed = true
}

func (x *VMClient) VMRegister() {
	x.VMRegisterMethod("Закрыть", x.Закрыть)
	x.VMRegisterMethod("Работает", x.Работает)
	x.VMRegisterMethod("Открыть", x.Открыть)
	// tst.VMRegisterField("ПолеСтрока", &tst.ПолеСтрока)
}

func (x *VMClient) Открыть(args VMSlice, rets *VMSlice) error {
	if len(args) != 3 {
		return VMErrorNeedArgs(3)
	}
	p, ok := args[0].(VMString)
	if !ok {
		return errors.New("Первый аргумент должен быть строкой с типом канала")
	}
	adr, ok := args[1].(VMString)
	if !ok {
		return errors.New("Второй аргумент должен быть строкой с адресом")
	}
	f, ok := args[2].(VMFunc)
	if !ok {
		return errors.New("Третий аргумент должен быть функцией с одним аргументом-соединением")
	}

	return x.Open(string(p), string(adr), f)
}

func (x *VMClient) Закрыть(args VMSlice, rets *VMSlice) error {
	x.Close()
	return nil
}

func (x *VMClient) Работает(args VMSlice, rets *VMSlice) error {
	rets.Append(VMBool(x.IsOnline()))
	return nil
}
