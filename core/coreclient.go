package core

import (
	"errors"
	"fmt"
	"sync"
)

type VMClient struct {
	VMMetaObj //должен передаваться по ссылке, поэтому это будет объект метаданных

	addr     string  // [addr]:port
	protocol string  // tcp, json, http
	conn     *VMConn // клиент tcp, каждому соединению присваивается GUID
}

func (x *VMClient) String() string {
	return fmt.Sprintf("Клиент %s %s", x.protocol, x.addr)
}

func (x *VMClient) IsOnline() bool {
	return x.conn != nil && !x.conn.closed
}

func (x *VMClient) Open(proto, addr string, handler VMFunc, data VMValuer, closeOnExitHandler bool) error {

	switch proto {
	case "tcp", "tcpzip", "tcptls", "http", "https":

		x.conn = NewVMConn(data)
		err := x.conn.Dial(proto, addr, handler, closeOnExitHandler)
		if err != nil {
			return err
		}

	default:
		return VMErrorIncorrectProtocol
	}
	return nil
}

func (x *VMClient) Close() {
	x.conn.Close()
}

func (x *VMClient) VMRegister() {
	x.VMRegisterMethod("Закрыть", x.Закрыть)
	x.VMRegisterMethod("Работает", x.Работает)
	x.VMRegisterMethod("Открыть", x.Открыть)     // асинхронно
	x.VMRegisterMethod("Соединить", x.Соединить) // синхронно

	// tst.VMRegisterField("ПолеСтрока", &tst.ПолеСтрока)
}

func (x *VMClient) Открыть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	if len(args) != 4 {
		return VMErrorNeedArgs(4)
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

	return x.Open(string(p), string(adr), f, args[3], true)
}

func (x *VMClient) Соединить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	if len(args) != 2 {
		return VMErrorNeedArgs(2)
	}
	p, ok := args[0].(VMString)
	if !ok {
		return errors.New("Первый аргумент должен быть строкой с типом канала")
	}
	adr, ok := args[1].(VMString)
	if !ok {
		return errors.New("Второй аргумент должен быть строкой с адресом")
	}

	var vcn *VMConn

	wg := &sync.WaitGroup{}
	wg.Add(1)

	f := VMFunc(func(args VMSlice, rets *VMSlice, envout *(*Env)) error {
		defer wg.Done()
		vcn = args[0].(*VMConn)
		return nil
	})

	err := x.Open(string(p), string(adr), f, VMNil, false)
	if err != nil {
		return err
	}
	wg.Wait()

	if vcn == nil {
		return errors.New("Соединение не было установлено")
	}

	rets.Append(vcn)

	return nil
}

func (x *VMClient) Закрыть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	x.Close()
	return nil
}

func (x *VMClient) Работает(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(VMBool(x.IsOnline()))
	return nil
}
