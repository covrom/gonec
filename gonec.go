package gonec

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/covrom/gonec/gonecscan"
)

// APIPath содержит путь к api интерпретатора
const APIPath = "/gonec"

type clientConnection struct {
	IP string
}

// TODO: номера строк исходного кода для литералов
type token struct {
	toktype, category rune
	literal           string
}

type interpreter struct {
	sync.RWMutex
	clientConnections []clientConnection
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

	// TODO: синхронно запускается код модуля, но он может создавать вэб-сервера и горутины, которые будут работать и после возврата

	//tokens, err := i.Lexer(r, w)
	_, err = i.Lexer(r, w)

	if err != nil {
		return
	}

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

	aExport
	aTrue
	aFalse

	oComment
	oLineFeed
	oLabelStart
	oLabelEnd
	oSemi
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

	tNull
	tBool
	tDate
	tNum
	tStr
	tUndef
	tType

	iIllegal

	//для категорий
	defIdentifier
	defKeyword
	defOperator
	defDelimiter
	defPoint
	defValueInt
	defValueFloat
	defValueDate
	defValueString
	defUnknown
)

var keywordMap = map[string]rune{
	"если":      rIf,
	"if":        rIf,
	"тогда":     rThen,
	"then":      rThen,
	"иначеесли": rElsIf,
	"elsif":     rElsIf,
	"иначе":     rElse,
	"else":      rElse,
	"конецесли": rEndIf,
	"endif":     rEndIf,
	"для":       rFor,
	"for":       rFor,
	"каждого":   rEach,
	"each":      rEach,
	"из":        rIn,
	"in":        rIn,
	"по":        rTo,
	"to":        rTo,
	"пока":      rWhile,
	"while":     rWhile,
	"цикл":      rDo,
	"do":        rDo,
	"конеццикла":     rEndDo,
	"enddo":          rEndDo,
	"процедура":      rProcedure,
	"procedure":      rProcedure,
	"функция":        rFunction,
	"function":       rFunction,
	"конецпроцедуры": rEndProcedure,
	"endprocedure":   rEndProcedure,
	"конецфункции":   rEndFunction,
	"endfunction":    rEndFunction,
	"перем":          rVar,
	"var":            rVar,
	"перейти":        rGoto,
	"goto":           rGoto,
	"возврат":        rReturn,
	"return":         rReturn,
	"продолжить":     rContinue,
	"continue":       rContinue,
	"прервать":       rBreak,
	"break":          rBreak,
	"и":              rAnd,
	"and":            rAnd,
	"или":            rOr,
	"or":             rOr,
	"не":             rNot,
	"not":            rNot,
	"попытка":        rTry,
	"try":            rTry,
	"исключение": rExcept,
	"except":     rExcept,
	"вызватьисключение": rRaise,
	"raise":        rRaise,
	"конецпопытки": rEndTry,
	"endtry":       rEndTry,
	"новый":        rNew,
	"new":          rNew,
	"выполнить":    rExecute,
	"execute":      rExecute,

	"экспорт": aExport,
	"export":  aExport,
	"истина":  aTrue,
	"true":    aTrue,
	"ложь":    aFalse,
	"false":   aFalse,

	"null":         tNull,
	"булево":       tBool,
	"boolean":      tBool,
	"дата":         tDate,
	"date":         tDate,
	"число":        tNum,
	"number":       tNum,
	"строка":       tStr,
	"string":       tStr,
	"неопределено": tUndef,
	"undefined":    tUndef,
	"тип":          tType,
	"type":         tType,
}

var operMap = map[string]rune{
	"//": oComment,    //Двумя знаками косая черта начинается комментарий. Комментарием считается весь текст от символа до конца текущей строки
	"|":  oLineFeed,   //Используется только в строковых константах в начале строки и означает, что данная строка является продолжением предыдущей (перенос строки)
	"~":  oLabelStart, //Начало метки оператора
	":":  oLabelEnd,   //Окончание метки оператора
	"(":  oLBr,
	")":  oRBr, //В круглые скобки заключается список параметров методов, процедур, функций и конструкторов. Также они используются в выражениях встроенного языка
	"[":  oLSqBr,
	"]":  oRSqBr,  //С помощью оператора квадратные скобки производится обращение к свойствам объекта по строковому представлению имени свойства.Также возможно обращение к элементам коллекций по индексу или другому параметру
	",":  oComma,  //Разделяет параметры в списке параметров методов, процедур, функций и конструкторов
	"\"": oDQuote, //Обрамляет строковые литералы
	"'":  oSQuote, //Обрамляет литералы даты
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

var delimMap = map[string]rune{
	";": oSemi, //Символ разделения операторов
}

var pointMap = map[string]rune{
	".": oPoint, //Десятичная точка в числовых литералах. Разделитель, используемый для обращения к свойствам и методам объектов встроенного языка
}

// В общем случае формат оператора языка следующий:
// ~метка:Оператор[(параметры)] [ДобКлючевоеСлово];

func (i *interpreter) Lexer(r io.Reader, w io.Writer) (tokens []token, err error) {
	//лексический анализ
	var s gonecscan.Scanner

	s.Error = func(s *gonecscan.Scanner, msg string) {
		err = errors.New(msg)
	}

	s.Init(r)

	var tok rune

	for tok != gonecscan.EOF {
		tok = s.Scan()
		if err != nil {
			return
		}

		nt := token{literal: s.TokenText()}
		ntlit := strings.ToLower(nt.literal)
		var ok bool
		switch tok {
		case gonecscan.Ident:
			nt.toktype, ok = keywordMap[ntlit]
			if !ok {
				nt.category = defIdentifier
			} else {
				nt.category = defKeyword
			}
		case gonecscan.String:
			// TODO: строки возвращаиюся вместе с промежуточными переносами и комментариями - нужно дополнительно очищать
			nt.category = defValueString
		case gonecscan.Int:
			nt.category = defValueInt
		case gonecscan.Float:
			nt.category = defValueFloat
		case gonecscan.Date:
			nt.category = defValueDate
		default:
			nt.toktype, ok = operMap[ntlit]
			if !ok {
				nt.toktype, ok = delimMap[ntlit]
				if !ok {
					nt.toktype, ok = pointMap[ntlit]
					if !ok {
						nt.category = defUnknown
					} else {
						nt.category = defPoint
					}
				} else {
					nt.category = defDelimiter
				}
			} else {
				nt.category = defOperator
			}
		}

		tokens = append(tokens,nt)
		// TODO:


	}

	return nil, nil
}
