package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"runtime"

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

// ЗапуститьСервер (Адрес, ТипПротокола) (КаналЗавершения VMChan) - запускает сервер в зависимости от выбранного типа
// Первый аргумент - адрес и порт в формате как для Го http, т.е. "[addr]:port"
// Второй аргумент - код протокола, строка
//   Допустимые значения:
//     "bin" - бинарный протокол Гонец через net/tcp, обмен только объектами VMStringMap (со вложенными VMSlice и другими типами интерпретатора)
//     ...[остальные в разработке]
// Возвращает канал закрытия (запись в него остановит сервер, чтение сигнализирует об остановке)
// В исходный канал поступают соединения VMConn
func (x VMChan) ЗапуститьСервер(args VMSlice, rets *VMSlice) error {
	adr, ok := args[0].(VMString)
	if !ok {
		return errors.New("Первый аргумент должен быть строкой")
	}
	p, ok := args[1].(VMString)
	if !ok {
		return errors.New("Второй аргумент должен быть строкой определенного вида")
	}

	// в этот канал посылает ошибку либо сам сервер, если он отстрелен по ошибке,
	// либо в него можно послать любое значение, и тогда сервер отстрелится корректно
	chDone := make(VMChan)

	switch string(p) {
	case "bin":
		//бинарный протокол Гонец через net/tcp, обмен объектами VMStringMap со вложенными VMSlice и другими типами интерпретатора
		go ServeGncBin(string(adr), x, chDone)
	default:
		return errors.New("Неизвестный тип протокола")
	}

	rets.Append(chDone)

	return nil
}

//ServeGncBin - бинарный протокол Гонец через net/tcp, обмен объектами VMStringMap со вложенными VMSlice и другими типами интерпретатора
// получает запрос из сети и передает его интерпретацию в виде VMStringMap по каналу
// передает ошибку по каналу done, если произошла ошибка
// завершает работу, если получает любое значение по каналу done (и ретранслирует его), либо если он становится закрыт

// КаналЗапросов=Новый Канал
// КаналЗакрытия = КаналЗапросов.ЗапуститьСервер(":9330","bin")
// // В канал запросов приходят соединения с клиентами
// Пока Истина Цикл
//   Выбор:
//   Когда Соед = <-КаналЗапросов:
//     Старт Функция (Соединение)
//       Запрос = <-Соединение.КаналОтКлиента
//       Если Запрос<>Неопределено Тогда
//         Если Запрос["Запрос"]="Чо да как?" Тогда
//           Соединение.КаналНаКлиента <- {"id":Запрос.id, "Результат":"ok", "НовыйИд":УникальныйИдентификатор()}
//         КонецЕсли
//       КонецЕсли
//     КонецФункции (Соед)
//   Когда <-КаналЗакрытия:
//     Прервать
//   КонецВыбора
//   ОбработатьГорутины()
// КонецЦикла

// КаналЗапросов=Новый Канал
// КаналЗакрытия, КаналОтСервера = КаналЗапросов.ЗапуститьКлиент("127.0.0.1:9330","bin")
// Ид = УникальныйИдентификатор()
// КаналЗапросов <- {"id":Ид, "Запрос":"Чо да как?"}
// // Ждем нужный ответ
// Пока Истина Цикл
//   Выбор:
//   Когда б= <-КаналОтСервера:
//     Если б.id=Ид Тогда
//       Сообщить("Новый Ид",б["НовыйИд"])
//       Прервать
//     КонецЕсли
//   Когда <-КаналЗакрытия:
//     Прервать
//   КонецВыбора
//   ОбработатьГорутины()
// КонецЦикла

func ServeGncBin(addr string, chv, done VMChan) {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		done <- VMString(fmt.Sprint(err))
		return
	}

	defer ln.Close()

	// мониторинг остановки
	go func() {
		select {
		case e, ok := <-done:
			// закрываем сервер
			ln.Close()
			if ok {
				done <- e // ретранслируем
			}
			return
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			done <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
			return
		}
		go func(cn net.Conn, ch, cl VMChan) {
			
			// TODO: переделать на VMConn

			// в этом протоколе происходит обмен структурами VMStringMap с сериализацией в binary формат
			var buf bytes.Buffer
			_, err := io.CopyN(&buf, cn, 21)

			if err != nil {
				cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
				cn.Write([]byte("er"))
				return
			}
			b := buf.Bytes()
			// проверяем целостность полученного сообщения
			// проверяем хэш, он идет первыми 8-ю байтами
			// затем идет заголовок 5 байт "gonec"
			// затем идет тело
			if len(b) < 21 {
				// ошибка? ну и ладно, ничего в канал не отправим
				cn.Write([]byte("er"))
				return
			}
			cstr := b[:5] // "gonec"
			if string(cstr) != "gonec" {
				cn.Write([]byte("er"))
				return
			}
			hashbts, _ := binary.Uvarint(b[5:13]) // hash
			lenb, _ := binary.Uvarint(b[13:21])   // hash

			buf.Reset()
			_, err = io.CopyN(&buf, cn, int64(lenb))
			if err != nil {
				cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
				cn.Write([]byte("er"))
				return
			}
			b = buf.Bytes()

			if HashBytes(b) != hashbts {
				cn.Write([]byte("er"))
				return
			}
			// проверили хэш, все ок - получаем VMStringMap
			rv := make(VMStringMap)
			if err := (&rv).UnmarshalBinary(b); err != nil {
				// ошибка? ну и ладно, ничего в канал не отправим
				cn.Write([]byte("er"))
				return
			}
			ch <- rv // все ок - отправили VMStringMap в канал
			cn.Write([]byte("ok"))
		}(conn, chv, done)
		runtime.Gosched()
	}
}

// DialGncBin отправляет запросы из канала ch на сервер по адресу addr и возвращает ответы в канал cret
// Если произошла ошибка подключения, она отправляется в канал cl, просмотр канала ch и отправка сообщений на сервер прекращается
// Если получит любое значение в канал cl, то прекратит просматривать канал ch и перестанет отправлять запросы на сервер
func DialGncBin(addr string, ch, cl VMChan) (cret VMChan) {
	cret = make(VMChan)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
		return
	}
	defer conn.Close()

	go func(cr, cl VMChan, cn net.Conn) {
		// получаем ответы от сервера в cr, строками
		for {
			select {
			case e := <-cl:
				cl <- e
				return
			default:
				var buf bytes.Buffer
				_, err := io.CopyN(&buf, cn, 2)
				if err != nil {
					cl <- VMString(fmt.Sprint(err))
					return
				}
				rv := string(buf.Bytes())
				if rv == "er" {
					cr <- VMNil
				}
			}
			runtime.Gosched()
		}
	}(cret, cl, conn)

	for {
		// ждем значение к отправке
		select {
		case v := <-ch:
			// отправляем только VMStringMap
			if vv, ok := v.(VMStringMap); ok {
				go func(cn net.Conn, val VMStringMap) {
					b, err := val.MarshalBinary()
					if err != nil {
						cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
						return
					}
					hb := make([]byte, 8)
					binary.PutUvarint(hb, HashBytes(b))

					_, err = cn.Write([]byte("gonec"))
					if err != nil {
						cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
						return
					}
					_, err = cn.Write(hb)
					if err != nil {
						cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
						return
					}
					//пишем длину
					binary.PutUvarint(hb, uint64(len(b)))
					_, err = cn.Write(hb)
					if err != nil {
						cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
						return
					}

					buf := bytes.NewReader(b)
					_, err = io.Copy(cn, buf)
					if err != nil {
						cl <- VMString(fmt.Sprint(err)) // сигнализируем остальным горутинам (в т.ч. вызывающей), что этот сервер отстрелился
						return
					}
				}(conn, vv)
			} else {
				cl <- VMString("Можно отправлять только структуры")
				return // выходим
			}
		case e := <-cl:
			cl <- e // ретранслируем
			return  // выходим
		}
		runtime.Gosched()
	}
}
