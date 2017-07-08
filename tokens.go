package gonec

// TODO: номера строк исходного кода для литералов
type token struct {
	toktype, category rune
	literal           string
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
	//комментарии не учитываются интерпретатором, поэтому // не входит в операторы
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
