package bincode

import (
	"fmt"
	"reflect"

	"github.com/covrom/gonec/ast"
)

type BinStmt interface {
	ast.Pos
	binstmt()
}

type BinStmtImpl struct {
	ast.PosImpl
	fmt.Stringer
}

func (x *BinStmtImpl) binstmt() {}

type BinCode []BinStmt

func (v BinCode) String() string {
	s := ""
	for _, e := range v {
		s += fmt.Sprintf("%v\n", e)
	}
	return s
}

//////////////////////
// команды байткода
//////////////////////

const (
	_    = iota
	ADD  // +
	SUB  // -
	MUL  // *
	QUO  // /
	REM  // %
	EQL  // ==
	NEQ  // !=
	GTR  // >
	GEQ  // >=
	LSS  // <
	LEQ  // <=
	OR   // |
	LOR  // ||
	AND  // &
	LAND // &&
	POW  //**
	SHL  // <<
	SHR  // >>
)

var OperMap = map[string]int{
	"+":  ADD,  // +
	"-":  SUB,  // -
	"*":  MUL,  // *
	"/":  QUO,  // /
	"%":  REM,  // %
	"==": EQL,  // ==
	"!=": NEQ,  // !=
	">":  GTR,  // >
	">=": GEQ,  // >=
	"<":  LSS,  // <
	"<=": LEQ,  // <=
	"|":  OR,   // |
	"||": LOR,  // ||
	"&":  AND,  // &
	"&&": LAND, // &&
	"**": POW,  //**
	"<<": SHL,  // <<
	">>": SHR,  // >>
}

var OperMapR = map[int]string{
	ADD:  "+",  // +
	SUB:  "-",  // -
	MUL:  "*",  // *
	QUO:  "/",  // /
	REM:  "%",  // %
	EQL:  "==", // ==
	NEQ:  "!=", // !=
	GTR:  ">",  // >
	GEQ:  ">=", // >=
	LSS:  "<",  // <
	LEQ:  "<=", // <=
	OR:   "|",  // |
	LOR:  "||", // ||
	AND:  "&",  // &
	LAND: "&&", // &&
	POW:  "**", //**
	SHL:  "<<", // <<
	SHR:  ">>", // >>
}

type BinLOAD struct {
	BinStmtImpl

	Reg int
	Val interface{}
}

func (v BinLOAD) String() string {
	return fmt.Sprintf("LOAD r%d, %#v", v.Reg, v.Val)
}

type BinMV struct {
	BinStmtImpl

	RegFrom int
	RegTo   int
}

func (v BinMV) String() string {
	return fmt.Sprintf("MV r%d, r%d", v.RegTo, v.RegFrom)
}

type BinEQUAL struct {
	BinStmtImpl

	Reg  int // результат сравнения помещаем сюда, тип "булево"
	Reg1 int
	Reg2 int
}

func (v BinEQUAL) String() string {
	return fmt.Sprintf("EQUAL r%d, r%d == r%d", v.Reg, v.Reg1, v.Reg2)
}

type BinCASTNUM struct {
	BinStmtImpl

	Reg int
}

func (v BinCASTNUM) String() string {
	return fmt.Sprintf("CAST r%d, NUMBER", v.Reg)
}

type BinMAKESLICE struct {
	BinStmtImpl

	Reg int
	Len int
	Cap int
}

func (v BinMAKESLICE) String() string {
	return fmt.Sprintf("MAKESLICE r%d, LEN %d, CAP %d", v.Reg, v.Len, v.Cap)
}

type BinSETIDX struct {
	BinStmtImpl

	Reg    int
	Index  int
	RegVal int
}

func (v BinSETIDX) String() string {
	return fmt.Sprintf("SETIDX r%d[%d], r%d", v.Reg, v.Index, v.RegVal)
}

type BinMAKEMAP struct {
	BinStmtImpl

	Reg int
	Len int
}

func (v BinMAKEMAP) String() string {
	return fmt.Sprintf("MAKEMAP r%d, LEN %d", v.Reg, v.Len)
}

type BinSETKEY struct {
	BinStmtImpl

	Reg    int
	Key    string
	RegVal int
}

func (v BinSETKEY) String() string {
	return fmt.Sprintf("SETKEY r%d[%q], r%d", v.Reg, v.Key, v.RegVal)
}

type BinGET struct {
	BinStmtImpl

	Reg    int
	Id     int
	Dotted bool // содержит точку "."
}

func (v BinGET) String() string {
	return fmt.Sprintf("GET r%d, %q", v.Reg, ast.UniqueNames.Get(v.Id))
}

type BinSET struct {
	BinStmtImpl

	Id  int // id переменной
	Reg int // регистр со значением
}

func (v BinSET) String() string {
	return fmt.Sprintf("SET %q, r%d", ast.UniqueNames.Get(v.Id), v.Reg)
}

type BinSETMEMBER struct {
	BinStmtImpl

	Reg    int // регистр со структтурой или мапой
	Id     int // id поля структуры или мапы
	RegVal int // регистр со значением
}

func (v BinSETMEMBER) String() string {
	return fmt.Sprintf("SETMEMBER r%d.%q, r%d", v.Reg, ast.UniqueNames.Get(v.Id), v.RegVal)
}

type BinSETNAME struct {
	BinStmtImpl

	Reg int // регистр с именем (строкой), сюда же возвращается id имени, записанного в ast.UniqueNames.Set()
}

func (v BinSETNAME) String() string {
	return fmt.Sprintf("SETNAME r%d", v.Reg)
}

type BinSETITEM struct {
	BinStmtImpl

	Reg      int
	RegIndex int
	RegVal   int
}

func (v BinSETITEM) String() string {
	return fmt.Sprintf("SETITEM r%d[r%d], r%d", v.Reg, v.RegIndex, v.RegVal)
}

type BinSETSLICE struct {
	BinStmtImpl

	Reg      int
	RegBegin int
	RegEnd   int
	RegVal   int
}

func (v BinSETSLICE) String() string {
	return fmt.Sprintf("SETSLICE r%d[r%d:r%d], r%d", v.Reg, v.RegBegin, v.RegEnd, v.RegVal)
}

type BinUNARY struct {
	BinStmtImpl

	Reg int
	Op  rune // - ! ^
}

func (v BinUNARY) String() string {
	return fmt.Sprintf("UNARY %sr%d", string(v.Op), v.Reg)
}

type BinADDR struct {
	BinStmtImpl

	Reg int
}

func (v BinADDR) String() string {
	return fmt.Sprintf("ADDR r%d", v.Reg)
}

type BinUNREF struct {
	BinStmtImpl

	Reg int
}

func (v BinUNREF) String() string {
	return fmt.Sprintf("UNREF r%d", v.Reg)
}

type BinLABEL struct {
	BinStmtImpl

	Label int
}

func (v BinLABEL) String() string {
	return fmt.Sprintf("L%d:", v.Label)
}

type BinJMP struct {
	BinStmtImpl

	JumpTo int
}

func (v BinJMP) String() string {
	return fmt.Sprintf("JMP L%d", v.JumpTo)
}

type BinJTRUE struct {
	BinStmtImpl

	Reg    int
	JumpTo int
}

func (v BinJTRUE) String() string {
	return fmt.Sprintf("JTRUE r%d, L%d", v.Reg, v.JumpTo)
}

type BinJFALSE struct {
	BinStmtImpl

	Reg    int
	JumpTo int
}

func (v BinJFALSE) String() string {
	return fmt.Sprintf("JFALSE r%d, L%d", v.Reg, v.JumpTo)
}

type BinOPER struct {
	BinStmtImpl

	RegL int // сюда же помещается результат
	RegR int
	Op   int
}

func (v BinOPER) String() string {
	return fmt.Sprintf("OP r%d, %q, r%d", v.RegL, OperMapR[v.Op], v.RegR)
}

type BinCALL struct {
	BinStmtImpl

	Name int // либо вызов по имени из ast.UniqueNames, если Name != 0
	// либо вызов обработчика (Name==0), напр. для анонимной функции
	// (выражение типа func, или ссылка или интерфейс с ним, находится в reg, а параметры начиная с reg+1)
	NumArgs int // число аргументов, которое надо взять на входе из регистров (<=7) или массива (Reg)
	RegArgs int // первый регистр из числа регистров с параметрами (параметров<=7) или регистр с массивом аругментов (>7)

	// в последнем регистре (из серии, если <=7, или в RegArgs, если >7) передан
	// массив аргументов переменной длины, и это приемлемо для вызываемой функции (оператор "...")
	// здесь надо быть аккуратным при числе аргументов >7
	// таким массивом будет только последний аргумент
	VarArg bool

	Go bool // признак необходимости запуска в новой горутине
}

func (v BinCALL) String() string {
	if v.Name == 0 {
		return fmt.Sprintf("CALL ANON r%d, ARGS r%d, ARGS_COUNT %d, VARARG %v, GO %v", v.RegArgs, v.RegArgs+1, v.NumArgs, v.VarArg, v.Go)
	}
	return fmt.Sprintf("CALL %q, ARGS r%d, ARGS_COUNT %d, VARARG %v, GO %v", ast.UniqueNames.Get(v.Name), v.RegArgs, v.NumArgs, v.VarArg, v.Go)
}

type BinGETMEMBER struct {
	BinStmtImpl

	Reg  int
	Name int
}

func (v BinGETMEMBER) String() string {
	return fmt.Sprintf("GETMEMBER r%d, %q", v.Reg, ast.UniqueNames.Get(v.Name))
}

type BinGETIDX struct {
	BinStmtImpl

	Reg      int
	RegIndex int
}

func (v BinGETIDX) String() string {
	return fmt.Sprintf("GETIDX r%d[r%d]", v.Reg, v.RegIndex)
}

type BinGETSUBSLICE struct {
	BinStmtImpl

	Reg      int
	BeginReg int
	EndReg   int
}

func (v BinGETSUBSLICE) String() string {
	return fmt.Sprintf("SLICE r%d[r%d : r%d]", v.Reg, v.BeginReg, v.EndReg)
}

type BinFUNC struct {
	BinStmtImpl

	Reg      int // регистр, в который сохраняется значение определяемой функции типа func
	Name     int
	Code     BinCode
	Args     []int // идентификаторы параметров
	VarArg   bool
	ReturnTo int //метка инструкции возврата из функции
}

func (v BinFUNC) String() string {
	s := ""
	for _, a := range v.Args {
		if s != "" {
			s += ", "
		}
		s += ast.UniqueNames.Get(a)
	}
	vrg := ""
	if v.VarArg {
		vrg = "..."
	}
	return fmt.Sprintf("FUNC r%d, %q, (%s %s)\n{\n%v}\n", v.Reg, ast.UniqueNames.Get(v.Name), s, vrg, v.Code)
}

type BinCASTTYPE struct {
	BinStmtImpl

	Reg     int
	TypeReg int
}

func (v BinCASTTYPE) String() string {
	return fmt.Sprintf("CAST r%d AS TYPE r%d", v.Reg, v.TypeReg)
}

type BinMAKE struct {
	BinStmtImpl

	Reg int // здесь id типа, и сюда же пишем новое значение
}

func (v BinMAKE) String() string {
	return fmt.Sprintf("MAKE r%d AS TYPE r%d", v.Reg, v.Reg)
}

type BinMAKECHAN struct {
	BinStmtImpl

	Reg int // тут размер буфера (0=без буфера), сюда же помещается созданный канал
}

func (v BinMAKECHAN) String() string {
	return fmt.Sprintf("MAKECHAN r%d SIZE r%d", v.Reg, v.Reg)
}

type BinMAKEARR struct {
	BinStmtImpl

	Reg    int // тут длина, сюда же помещается слайс
	RegCap int
}

func (v BinMAKEARR) String() string {
	return fmt.Sprintf("MAKESLICE r%d, LEN r%d, CAP r%d", v.Reg, v.Reg, v.RegCap)
}

type BinCHANRECV struct {
	BinStmtImpl
	// с ожиданием
	Reg    int // канал
	RegVal int // сюда помещается результат
}

func (v BinCHANRECV) String() string {
	return fmt.Sprintf("<-CHAN r%d, r%d", v.RegVal, v.Reg)
}

type BinCHANSEND struct {
	BinStmtImpl
	// с ожиданием
	Reg    int // канал
	RegVal int // значение
}

func (v BinCHANSEND) String() string {
	return fmt.Sprintf("CHAN<- r%d, r%d", v.Reg, v.RegVal)
}

type BinISKIND struct {
	BinStmtImpl

	Reg  int          // значение для проверки, сюда же возвращается bool
	Kind reflect.Kind // категория типа значения в reg
}

func (v BinISKIND) String() string {
	return fmt.Sprintf("ISKIND r%d, %s", v.Reg, v.Kind)
}

type BinTRY struct {
	BinStmtImpl

	Reg    int // регистр, куда будет помещаться error во время выполнения последующего кода
	JumpTo int // метка блока обработки ошибки
}

func (v BinTRY) String() string {
	return fmt.Sprintf("TRY r%d, CATCH L%d", v.Reg, v.JumpTo)
}

type BinCATCH struct {
	BinStmtImpl

	Reg    int
	JumpTo int
}

func (v BinCATCH) String() string {
	return fmt.Sprintf("CATCH r%d, NOERR L%d", v.Reg, v.JumpTo)
}

type BinPOPTRY struct {
	BinStmtImpl

	CatchLabel int // снимаем со стека исключений конструкцию с этим регистром
}

func (v BinPOPTRY) String() string {
	return fmt.Sprintf("POPTRY L%d", v.CatchLabel)
}

type BinFOREACH struct {
	BinStmtImpl

	Reg        int // регистр для итерационного выбора из него значений
	RegIter    int // в этот регистр будет записываться итератор
	BreakLabel int
}

func (v BinFOREACH) String() string {
	return fmt.Sprintf("FOREACH r%d, ITER r%d, BREAK TO L%d", v.Reg, v.RegIter, v.BreakLabel)
}

type BinNEXT struct {
	BinStmtImpl

	Reg int // выбираем из этого регистра следующее значение и помещаем в регистр RegVal
	// это может быть очередное значение из слайса или из канала, зависит от типа значения в Reg
	RegVal  int
	RegIter int // регистр с итератором, инициализированным FOREACH
	JumpTo  int // переход в случае, если нет очередного значения (достигнут конец выборки)
	// туда же переходим по Прервать
}

func (v BinNEXT) String() string {
	return fmt.Sprintf("NEXT r%d, FROM r%d, ITER r%d, ENDLOOP L%d", v.RegVal, v.Reg, v.RegIter, v.JumpTo)
}

type BinPOPFOR struct {
	BinStmtImpl

	Reg int // снимаем со стека циклов конструкцию с этим регистром
}

func (v BinPOPFOR) String() string {
	return fmt.Sprintf("POPFOR r%d", v.Reg)
}

type BinFORNUM struct {
	BinStmtImpl

	Reg        int // регистр для итерационного значения
	RegFrom    int // регистр с начальным значением
	RegTo      int // регистр с конечным значением
	BreakLabel int
}

func (v BinFORNUM) String() string {
	return fmt.Sprintf("FORNUM r%d, FROM r%d, TO r%d, BREAK TO L%d", v.Reg, v.RegFrom, v.RegTo, v.BreakLabel)
}

type BinNEXTNUM struct {
	BinStmtImpl

	Reg    int // следующее значение итератора
	JumpTo int // переход в случае, если значение после увеличения стало больше, чем ранее определенное в RegTo
	// туда же переходим по Прервать
}

func (v BinNEXTNUM) String() string {
	return fmt.Sprintf("NEXTNUM r%d, ENDLOOP L%d", v.Reg, v.JumpTo)
}

type BinWHILE struct {
	BinStmtImpl

	Reg        int // регистр для условия
	BreakLabel int
}

func (v BinWHILE) String() string {
	return fmt.Sprintf("WHILE r%d, BREAK TO L%d", v.Reg, v.BreakLabel)
}

type BinBREAK struct {
	BinStmtImpl
}

func (v BinBREAK) String() string {
	return fmt.Sprintf("BREAK")
}

type BinCONTINUE struct {
	BinStmtImpl
}

func (v BinCONTINUE) String() string {
	return fmt.Sprintf("CONTINUE")
}

type BinRET struct {
	BinStmtImpl

	Reg int
}

func (v BinRET) String() string {
	return fmt.Sprintf("RETURN r%d", v.Reg)
}

type BinTHROW struct {
	BinStmtImpl

	Reg int
}

func (v BinTHROW) String() string {
	return fmt.Sprintf("THROW r%d", v.Reg)
}

type BinMODULE struct {
	BinStmtImpl

	Name int
	Code BinCode
}

func (v BinMODULE) String() string {
	return fmt.Sprintf("MODULE %s\n{\n%v}\n", ast.UniqueNames.Get(v.Name), v.Code)
}

type BinERROR struct {
	BinStmtImpl

	Error string
}

func (v BinERROR) String() string {
	return fmt.Sprintf("ERROR %q", v.Error)
}

type BinTRYRECV struct {
	BinStmtImpl

	Reg    int // на входе канал, на выходе тоже
	RegVal int // получаемое значение
	RegOk  int // успешное чтение, или не было чтения, или в Reg не канал
}

func (v BinTRYRECV) String() string {
	return fmt.Sprintf("TRYRECV r%d, OK r%d", v.Reg, v.RegOk)
}

type BinTRYSEND struct {
	BinStmtImpl

	Reg    int // на входе канал, на выходе тоже
	RegVal int // регистр со значением для отправки
	RegOk  int // успешно передано в канал, или не было передачи, или в Reg не канал
}

func (v BinTRYSEND) String() string {
	return fmt.Sprintf("TRYSEND r%d, r%d, OK r%d", v.Reg, v.RegVal, v.RegOk)
}

type BinGOSHED struct {
	BinStmtImpl
}

func (v BinGOSHED) String() string {
	return fmt.Sprintf("GOSHED")
}

type BinINC struct {
	BinStmtImpl

	Reg int
}

func (v BinINC) String() string {
	return fmt.Sprintf("INC r%d", v.Reg)
}

type BinDEC struct {
	BinStmtImpl

	Reg int
}

func (v BinDEC) String() string {
	return fmt.Sprintf("DEC r%d", v.Reg)
}
