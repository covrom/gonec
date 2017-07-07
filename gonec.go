package gonec

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// APIPath содержит путь к api интерпретатора
const APIPath = "/gonec"

type clientConnection struct {
	IP string
}

type interpreter struct {
	sync.RWMutex
	clientConnections []clientConnection
	query             []byte
}

// Interpreter возвращает новый интерпретатор
func Interpreter() *interpreter {
	return &interpreter{}
}

// handlerMain обрабатывает входящие запросы к интерпретатору через POST-запросы
// тело запроса - это код для интерпретации
func (i *interpreter) handlerMain(w http.ResponseWriter, r *http.Request) {

	i.RLock()
	//лимит количества одновременных подключений к одному интерпретатору
	overconn := len(i.clientConnections) >= 2
	i.RUnlock()

	if overconn {
		http.Error(w, "Слишком много запросов обрабатывается в данный момент", http.StatusForbidden)
		return
	}

	clconn := clientConnection{
		IP: r.RemoteAddr,
	}

	i.Lock()
	i.clientConnections = append(i.clientConnections, clconn)
	i.Unlock()

	defer func(cc clientConnection) {
		i.Lock()
		for n := range i.clientConnections {
			if i.clientConnections[n] == cc {
				i.clientConnections = append(i.clientConnections[:n], i.clientConnections[n+1:]...)
				break
			}
		}
		i.Unlock()
	}(clconn)

	if r.ContentLength > 1<<26 {
		http.Error(w, "Слишком большой запрос", http.StatusForbidden)
		return
	}

	switch r.Method {

	case http.MethodPost:

		defer r.Body.Close()
		//интерпретируется код и возвращается вывод как простой текст
		w.Header().Set("Content-Type", "text/plain")
		err := i.ParseAndRun(r.Body, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
}

// Run запускает микросервис интерпретатора по адресу и порту
func (i *interpreter) Run(srv string) {
	http.HandleFunc(APIPath, i.handlerMain)
	log.Fatal(http.ListenAndServe(srv, nil))
}

// ParseAndRun разбирает запрос и запускает интерпретацию
func (i *interpreter) ParseAndRun(r io.Reader, w io.Writer) (err error) {
	//если r==nil, то разбирается старый запрос, иначе заполняется новый
	if r != nil {
		i.query, err = ioutil.ReadAll(r)
		if err != nil {
			return
		}
	}

	//TODO: синхронно запускается код модуля, но он может создавать вэб-сервера и горутины, которые будут работать и после возврата

	return nil
}

// Приведенные далее ключевые слова являются зарезервированными и не могут использоваться в качестве создаваемых имен переменных, реквизитов объектов конфигурации и объявляемых процедур и функций.
// В данном варианте языка каждое из ключевых слов имеет два представления – русское и английское.
const (
	iEOF = -(iota + 1)
	rIf
	rThen
	rElsIf
	rElse
	rEndIf
	rFor
	rEach
	rIn
	rTo
	rWhile
	rDo
	rEndDo
	rProcedure
	rFunction
	rEndProcedure
	rEndFunction
	rVar
	rGoto
	rReturn
	rContinue
	rBreak
	rAnd
	rOr
	rNot
	rTry
	rExcept
	rRaise
	rEndTry
	rNew
	rExecute

	oComment
	oLineFeed
	oLabelStart
	oLabelEnd
	oEOp
	oLBr
	oRBr
	oLSqBr
	oRSqBr
	oComma
	oDQuote
	oSQuote
	oPoint
	oAdd
	oSub
	oMul
	oDiv
	oRemDiv
	oGt
	oGe
	oLt
	oLe
	oEq
	oNe
)

var tokenMap = map[string]rune{
	"Если":      rIf,
	"If":        rIf,
	"Тогда":     rThen,
	"Then":      rThen,
	"ИначеЕсли": rElsIf,
	"ElsIf":     rElsIf,
	"Иначе":     rElse,
	"Else":      rElse,
	"КонецЕсли": rEndIf,
	"EndIf":     rEndIf,
	"Для":       rFor,
	"For":       rFor,
	"Каждого":   rEach,
	"Each":      rEach,
	"Из":        rIn,
	"In":        rIn,
	"По":        rTo,
	"To":        rTo,
	"Пока":      rWhile,
	"While":     rWhile,
	"Цикл":      rDo,
	"Do":        rDo,
	"КонецЦикла":     rEndDo,
	"EndDo":          rEndDo,
	"Процедура":      rProcedure,
	"Procedure":      rProcedure,
	"Функция":        rFunction,
	"Function":       rFunction,
	"КонецПроцедуры": rEndProcedure,
	"EndProcedure":   rEndProcedure,
	"КонецФункции":   rEndFunction,
	"EndFunction":    rEndFunction,
	"Перем":          rVar,
	"Var":            rVar,
	"Перейти":        rGoto,
	"Goto":           rGoto,
	"Возврат":        rReturn,
	"Return":         rReturn,
	"Продолжить":     rContinue,
	"Continue":       rContinue,
	"Прервать":       rBreak,
	"Break":          rBreak,
	"И":              rAnd,
	"And":            rAnd,
	"Или":            rOr,
	"Or":             rOr,
	"Не":             rNot,
	"Not":            rNot,
	"Попытка":        rTry,
	"Try":            rTry,
	"Исключение": rExcept,
	"Except":     rExcept,
	"ВызватьИсключение": rRaise,
	"Raise":        rRaise,
	"КонецПопытки": rEndTry,
	"EndTry":       rEndTry,
	"Новый":        rNew,
	"New":          rNew,
	"Выполнить":    rExecute,
	"Execute":      rExecute,
}

var operMap = map[string]rune{
	"//": oComment,    //Двумя знаками косая черта начинается комментарий. Комментарием считается весь текст от символа до конца текущей строки
	"|":  oLineFeed,   //Используется только в строковых константах в начале строки и означает, что данная строка является продолжением предыдущей (перенос строки)
	"~":  oLabelStart, //Начало метки оператора
	":":  oLabelEnd,   //Окончание метки оператора
	";":  oEOp,        //Символ разделения операторов
	"(":  oLBr,
	")":  oRBr, //В круглые скобки заключается список параметров методов, процедур, функций и конструкторов. Также они используются в выражениях встроенного языка
	"[":  oLSqBr,
	"]":  oRSqBr,  //С помощью оператора квадратные скобки производится обращение к свойствам объекта по строковому представлению имени свойства.Также возможно обращение к элементам коллекций по индексу или другому параметру
	",":  oComma,  //Разделяет параметры в списке параметров методов, процедур, функций и конструкторов
	"\"": oDQuote, //Обрамляет строковые литералы
	"'":  oSQuote, //Обрамляет литералы даты
	".":  oPoint,  //Десятичная точка в числовых литералах. Разделитель, используемый для обращения к свойствам и методам объектов встроенного языка
	"+":  oAdd,    //Операция сложения. Операция конкатенации строк
	"-":  oSub,    //Операция вычитания
	"*":  oMul,    //Операция умножения
	"/":  oDiv,    //Операция деления
	"%":  oRemDiv, //Получение остатка от деления. Допускается использование дробных значений делимого и делителя
	">":  oGt,     //Логическая операция Больше
	">=": oGe,     //Логическая операция Больше или равно
	"<":  oLt,     //Логическая операция Меньше
	"<=": oLe,     //Логическая операция Меньше или равно
	"=":  oEq,     //Операция присваивания. Логическая операция Равно
	"<>": oNe,     //Логическая операция Не равно
}
