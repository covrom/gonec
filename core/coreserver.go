package core

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"sync"

	uuid "github.com/satori/go.uuid"
)

var (
	VMErrorServerNowOnline = errors.New("Сервер уже запущен!")
	VMErrorServerOffline   = errors.New("Сервер уже остановлен!")
)

// VMServer - сервер протоколов взаимодействия, предоставляет базовый обработчик для TCP, RPC-JSON и HTTP соединений
// данный объект не может сериализоваться и не может участвовать в операциях с операндами
type VMServer struct {
	VMMetaObj //должен передаваться по ссылке, поэтому это будет объект метаданных

	mu       sync.RWMutex
	addr     string // [addr]:port
	protocol string // tcp, json, http
	done     chan error
	clients  map[string]net.Conn // каждому соединению присваивается GUID
	lnr      net.Listener
}

func (x *VMServer) String() string {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return fmt.Sprintf("Сервер %s %s", x.protocol, x.addr)
}

func (x *VMServer) IsOnline() bool {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return x.lnr != nil
}

func (x *VMServer) Open(proto, addr string) (err error) {
	// запускаем сервер
	x.mu.RLock()
	if x.lnr != nil {
		x.mu.RUnlock()
		return VMErrorServerNowOnline
	}
	x.mu.RUnlock()

	x.mu.Lock()
	defer x.mu.Unlock()

	x.done = make(chan error, 1) //буфер на одну ошибку, для закрытия
	x.clients = make(map[string]net.Conn)

	x.addr = addr
	x.protocol = proto

	switch proto {
	case "tcp":
		x.lnr, err = net.Listen("tcp", addr)
		if err != nil {
			x.lnr = nil
			return err
		}

		go func() {
			defer func() {
				x.mu.Lock()
				x.lnr = nil
				x.mu.Unlock()
			}()
			for {
				x.mu.RLock()
				ln := x.lnr
				x.mu.RUnlock()
				conn, err := ln.Accept()
				if err != nil {
					ln.Close() // отстрел сервера может произойти как принудительно, так и по внешнему событию,
					//поэтому, дублируем очистку ресурсов как в Close
					x.done <- err
					return
				}
				x.mu.Lock()
				x.clients[uuid.NewV1().String()] = conn
				x.mu.Unlock()

				runtime.Gosched()
			}
		}()

		// TODO: просмотр и удаление закрытых коннектов с клиентами, обработка открытых коннектов с помощью callback
	}

	return nil
}

// Close закрываем все, останавливаем горутины, всегда возвращаем ошибку закрытия
func (x *VMServer) Close() error {
	x.mu.RLock()
	ln := x.lnr
	x.mu.RUnlock()

	if ln != nil {
		ln.Close()
	}

	err,ok := <-x.done // дождемся ошибки из горутины, или возьмем ее, если она уже была
	if ok{
		// канал не закрыт
		close(x.done)   // чтобы при вызове второй раз не зависнуть
	}else{
		err = VMErrorServerOffline
	}
	x.mu.Lock()
	x.lnr = nil
	x.mu.Unlock()
	return err
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
