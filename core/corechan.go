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

// VMChan - канал для передачи любого типа вирт. машины
type VMChan chan VMValuer

func (x VMChan) vmval() {}

func (x VMChan) Interface() interface{} {
	return x
}

func (x VMChan) Send(v VMValuer) {
	x <- v
}

func (x VMChan) Recv() (VMValuer, bool) {
	rv, ok := <-x
	return rv, ok
}

func (x VMChan) TrySend(v VMValuer) (ok bool) {
	select {
	case x <- v:
		ok = true
	default:
		ok = false
	}
	return
}

func (x VMChan) TryRecv() (v VMValuer, ok bool, notready bool) {
	select {
	case v, ok = <-x:
		notready = false
	default:
		ok = false
		notready = true
	}
	return
}

func (x VMChan) Close() { close(x) }

func (x VMChan) Size() int { return cap(x) }

func (x VMChan) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!
	switch names.UniqueNames.GetLowerCase(name) {
	case "закрыть":
		return VMFuncMustParams(0, x.Закрыть), true
	case "размер":
		return VMFuncMustParams(0, x.Размер), true
	case "запуститьсервер":
		return VMFuncMustParams(2, x.ЗапуститьСервер), true
	}
	return nil, false
}

func (x VMChan) Закрыть(args VMSlice, rets *VMSlice) error {
	x.Close()
	return nil
}

func (x VMChan) Размер(args VMSlice, rets *VMSlice) error {
	rets.Append(VMInt(x.Size()))
	return nil
}

// ЗапуститьСервер (Адрес, ТипПротокола) (Канал остановки VMChan) - запускает сервер в зависимости от выбранного типа
// Первый аргумент - адрес и порт в формате как для Го http, т.е. "[addr]:port"
// Второй аргумент - код протокола, строка
//   Допустимые значения:
//     "bin" - бинарный протокол Гонец через net/tcp, обмен только объектами VMStringMap (со вложенными VMSlice и другими типами интерпретатора)
//     ...[остальные в разработке]
// Возвращает канал, закрытие которого приведет к остановке сервера
func (x VMChan) ЗапуститьСервер(args VMSlice, rets *VMSlice) error {
	adr, ok := args[0].(VMString)
	if !ok {
		return errors.New("Первый аргумент должен быть строкой")
	}
	p, ok := args[1].(VMString)
	if !ok {
		return errors.New("Второй аргумент должен быть строкой определенного вида")
	}

	// в этот канал посылает сигнал VMNil либо сам сервер, если он отстрелен по ошибке,
	// либо в него можно послать такой сигнал, и тогда сервер отстрелится корректно
	// оба канала могут работать на запись, поэтому, их закрывать нельзя, чтобы не было паники в горутинах
	chClose := make(VMChan)

	switch string(p) {

	case "bin":
		//бинарный протокол Гонец через net/tcp, обмен объектами VMStringMap со вложенными VMSlice и другими типами интерпретатора
		go ServeGncBin(string(adr), x, chClose)
	default:
		return errors.New("Неизвестный тип протокола")
	}

	rets.Append(chClose)
	return nil
}

//ServeGncBin - бинарный протокол Гонец через net/tcp, обмен объектами VMStringMap со вложенными VMSlice и другими типами интерпретатора
// получает запрос из сети и передает его интерпретацию в виде VMStringMap по каналу ch
// передает ошибку по каналу cl, если произошла ошибка
// завершает работу, если получает любое значение по каналу cl (и ретранслирует его)
func ServeGncBin(addr string, ch, cl VMChan) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
		return
	}
	defer ln.Close()
	go func(cl VMChan) {
		select {
		case e := <-cl:
			// закрываем сервер
			ln.Close()
			cl <- e // ретранслируем
			return
		}
	}(cl)
	for {
		conn, err := ln.Accept()
		if err != nil {
			cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
			return
		}
		go func(cn net.Conn, ch, cl VMChan) {
			// в этом протоколе происходит обмен структурами VMStringMap с сериализацией в binary формат
			var buf bytes.Buffer
			_, err := io.Copy(&buf, cn)
			if err != nil {
				cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
				cn.Write([]byte("error"))
				return
			}
			b := buf.Bytes()
			// проверяем целостность полученного сообщения
			// проверяем хэш, он идет первыми 8-ю байтами
			// затем идет заголовок 5 байт "gonec"
			// затем идет тело
			if len(b) < 13 {
				// ошибка? ну и ладно, ничего в канал не отправим
				cn.Write([]byte("error"))
				return
			}
			hashbts, _ := binary.Uvarint(b[:8]) // hash
			cstr := b[8:13]                     // "gonec"
			b = b[13:]
			if string(cstr) != "gonec" || len(b) == 0 {
				cn.Write([]byte("error"))
				return
			}
			if HashBytes(b) != hashbts {
				cn.Write([]byte("error"))
				return
			}
			// проверили хэш, все ок - получаем VMStringMap
			rv := make(VMStringMap)
			if err := (&rv).UnmarshalBinary(b); err != nil {
				// ошибка? ну и ладно, ничего в канал не отправим
				cn.Write([]byte("error"))
				return
			}
			ch <- rv // все ок - отправили VMStringMap в канал
		}(conn, ch, cl)
	}
}

func DialGncBin(addr string, ch, cerr VMChan) {

}
