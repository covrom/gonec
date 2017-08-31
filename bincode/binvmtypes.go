package bincode

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/covrom/gonec/ast"
	envir "github.com/covrom/gonec/env"
)

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

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	Message string
	Pos     ast.Position
}

var (
	BreakError     = errors.New("Неверное применение оператора Прервать")
	ContinueError  = errors.New("Неверное применение оператора Продолжить")
	ReturnError    = errors.New("Неверное применение оператора Возврат")
	InterruptError = errors.New("Выполнение прервано")
	// JumpError      = errors.New("Переход на метку")
)

// NewStringError makes error interface with message.
func NewStringError(pos ast.Pos, err string) error {
	if pos == nil {
		return &Error{Message: err, Pos: ast.Position{1, 1}}
	}
	return &Error{Message: err, Pos: pos.Position()}
}

// NewErrorf makes error interface with message.
func NewErrorf(pos ast.Pos, format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...), Pos: pos.Position()}
}

// NewError makes error interface with message.
// This doesn't overwrite last error.
func NewError(pos ast.Pos, err error) error {
	if err == nil {
		return nil
	}
	if err == BreakError || err == ContinueError || err == ReturnError {
		return err
	}
	// if pe, ok := err.(*parser.Error); ok {
	// 	return pe
	// }
	if ee, ok := err.(*Error); ok {
		return ee
	}
	return &Error{Message: err.Error(), Pos: pos.Position()}
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

type CatchFunc func() string

// Func is function interface to reflect functions internaly.
type Func func(args ...reflect.Value) (reflect.Value, error)

func (f Func) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

func ToFunc(f Func) reflect.Value {
	return reflect.ValueOf(f)
}

// Регистры виртуальной машины

type VMRegs struct {
	Reg       []interface{} // регистры значений
	Labels    map[int]int   // [label]=index в BinCode
	TryLabel  []int         // последний элемент - это метка на текущий обработчик CATCH
	TryRegErr []int         // последний элемент - это регистр с ошибкой текущего обработчика
}

const initlenregs = 20

func NewVMRegs(stmts BinCode) *VMRegs {
	//собираем мапу переходов
	lbls := make(map[int]int)
	for i, stmt := range stmts {
		if s, ok := stmt.(*BinLABEL); ok {
			lbls[s.Label] = i
		}
	}
	return &VMRegs{
		Reg:       make([]interface{}, initlenregs),
		Labels:    lbls,
		TryLabel:  make([]int, 0, 5),
		TryRegErr: make([]int, 0, 5),
	}
}

func (v *VMRegs) Set(reg int, val interface{}) {
	if reg < len(v.Reg) {
		v.Reg[reg] = val
	} else {
		for reg >= len(v.Reg) {
			v.Reg = append(v.Reg, make([]interface{}, initlenregs)...)
		}
		v.Reg[reg] = val
	}
}

func (v *VMRegs) PushTry(reg, label int) {
	v.TryRegErr = append(v.TryRegErr, reg)
	v.TryLabel = append(v.TryLabel, label)
}

func (v *VMRegs) TopTryLabel() int {
	l := len(v.TryLabel)
	if l == 0 {
		return -1
	}
	return v.TryLabel[l-1]
}

func (v *VMRegs) PopTry() (reg int, label int) {
	l := len(v.TryLabel)
	if l == 0 {
		return -1, -1
	}
	reg = v.TryRegErr[l-1]
	v.TryRegErr = v.TryRegErr[0 : l-1]
	label = v.TryLabel[l-1]
	v.TryLabel = v.TryLabel[0 : l-1]
	return
}

type VMSlice []interface{}

type VMStringMap map[string]interface{}

func InvokeNumber(lit string) (interface{}, error) {
	if strings.Contains(lit, ".") || strings.Contains(lit, "e") || strings.Contains(lit, "E") {
		v, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return v, err
		}
		return v, nil
	}
	var i int64
	var err error
	if strings.HasPrefix(lit, "0x") {
		i, err = strconv.ParseInt(lit[2:], 16, 64)
	} else {
		i, err = strconv.ParseInt(lit, 10, 64)
	}
	if err != nil {
		return i, err
	}
	return i, nil
}

func ToString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	if v == nil {
		return "Неопределено"
	}
	if b, ok := v.(bool); ok {
		if b {
			return "Истина"
		} else {
			return "Ложь"
		}
	}
	return fmt.Sprint(v)
}

func ToBool(v interface{}) bool {

	switch v.(type) {
	case float32, float64:
		return v.(float64) != 0.0
	case int, int32, int64:
		return v.(int64) != 0
	case bool:
		return v.(bool)
	case string:
		vlow := strings.ToLower(v.(string))
		if vlow == "true" || vlow == "истина" {
			return true
		}
		if ToInt64(v) != 0 {
			return true
		}
	}
	return false
}

func ToFloat64(v interface{}) float64 {
	switch v.(type) {
	case float32, float64:
		return v.(float64)
	case int, int32, int64:
		return float64(v.(int64))
	}
	return 0.0
}

func ToInt64(v interface{}) int64 {
	switch v.(type) {
	case float32, float64:
		return int64(v.(float64))
	case int, int32, int64:
		return v.(int64)
	case string:
		s := v.(string)
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return i
		}
	}
	return 0
}

func IsNum(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
		return true
	}
	return false
}

func Equal(lhsV, rhsV interface{}) bool {
	if lhsV == rhsV {
		return true
	}
	// if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
	// 	lhsV = lhsV.Elem()
	// }
	// if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
	// 	rhsV = rhsV.Elem()
	// }
	// if !lhsV.IsValid() || !rhsV.IsValid() {
	// 	return true
	// }
	if IsNum(lhsV) && IsNum(rhsV) {
		if reflect.TypeOf(rhsV).ConvertibleTo(reflect.TypeOf(lhsV)) {
			rhsV = reflect.ValueOf(rhsV).Convert(reflect.TypeOf(lhsV)).Interface()
		}
	}
	return reflect.DeepEqual(lhsV, rhsV)
}

func GetMember(v reflect.Value, name int, stmt ast.Pos) (reflect.Value, error) {

	m, _ := ast.MethodByNameCI(v, name)
	// ошибку не обрабатываем, т.к. ищем поле
	if !m.IsValid() {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Struct {
			var err error
			m, err = ast.FieldByNameCI(v, name)
			if err != nil || !m.IsValid() {
				return envir.NilValue, NewStringError(stmt, "Метод или поле не найдено: "+ast.UniqueNames.Get(name))
			}
		} else if v.Kind() == reflect.Map {
			m = v.MapIndex(reflect.ValueOf(ast.UniqueNames.Get(name)))
			if !m.IsValid() {
				return envir.NilValue, NewStringError(stmt, "Значение по ключу не найдено")
			}
		} else {
			return envir.NilValue, NewStringError(stmt, "У значения нет полей")
		}
	}
	return m, nil
}

func EvalBinOp(op int, lhsV, rhsV reflect.Value) (interface{}, error) {
	switch op {

	// TODO: математика множеств и графов

	case ADD:
		if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
			return ToString(lhsV) + ToString(rhsV), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() != reflect.Array && rhsV.Kind() != reflect.Slice) {
			return reflect.Append(lhsV, rhsV).Interface(), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() == reflect.Array || rhsV.Kind() == reflect.Slice) {
			return reflect.AppendSlice(lhsV, rhsV).Interface(), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return ToFloat64(lhsV) + ToFloat64(rhsV), nil
		}
		return ToInt64(lhsV) + ToInt64(rhsV), nil
	case SUB:
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return ToFloat64(lhsV) - ToFloat64(rhsV), nil
		}
		return ToInt64(lhsV) - ToInt64(rhsV), nil
	case MUL:
		if lhsV.Kind() == reflect.String && (rhsV.Kind() == reflect.Int || rhsV.Kind() == reflect.Int32 || rhsV.Kind() == reflect.Int64) {
			return strings.Repeat(ToString(lhsV), int(ToInt64(rhsV))), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return ToFloat64(lhsV) * ToFloat64(rhsV), nil
		}
		return ToInt64(lhsV) * ToInt64(rhsV), nil
	case QUO:
		return ToFloat64(lhsV) / ToFloat64(rhsV), nil
	case REM:
		return ToInt64(lhsV) % ToInt64(rhsV), nil
	case EQL:
		return Equal(lhsV, rhsV), nil
	case NEQ:
		return Equal(lhsV, rhsV) == false, nil
	case GTR:
		return ToFloat64(lhsV) > ToFloat64(rhsV), nil
	case GEQ:
		return ToFloat64(lhsV) >= ToFloat64(rhsV), nil
	case LSS:
		return ToFloat64(lhsV) < ToFloat64(rhsV), nil
	case LEQ:
		return ToFloat64(lhsV) <= ToFloat64(rhsV), nil
	case OR:
		return ToInt64(lhsV) | ToInt64(rhsV), nil
	case LOR:
		if x := ToBool(lhsV); x {
			return x, nil
		} else {
			return ToBool(rhsV), nil
		}
	case AND:
		return ToInt64(lhsV) & ToInt64(rhsV), nil
	case LAND:
		if x := ToBool(lhsV); x {
			return ToBool(rhsV), nil
		} else {
			return x, nil
		}
	case POW:
		if lhsV.Kind() == reflect.Float64 {
			return math.Pow(ToFloat64(lhsV), ToFloat64(rhsV)), nil
		}
		return int64(math.Pow(ToFloat64(lhsV), ToFloat64(rhsV))), nil
	case SHR:
		return ToInt64(lhsV) >> uint64(ToInt64(rhsV)), nil
	case SHL:
		return ToInt64(lhsV) << uint64(ToInt64(rhsV)), nil
	default:
		return nil, fmt.Errorf("Неизвестный оператор")
	}
}
