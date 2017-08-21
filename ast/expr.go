package ast

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

// уникальные названия переменных, индекс используется в AST-дереве
type EnvNames struct {
	sync.RWMutex
	names   map[string]int
	handles map[int]string
	iter    int
}

func NewEnvNames() *EnvNames {
	en := EnvNames{
		names:   make(map[string]int, 200),
		handles: make(map[int]string, 200),
		iter:    1,
	}
	return &en
}

func (en *EnvNames) Set(n string) int {
	ns := strings.ToLower(n)
	en.RLock()
	if i, ok := en.names[ns]; ok {
		en.RUnlock()
		return i
	}
	en.RUnlock()
	en.Lock()
	i := en.iter
	en.names[ns] = i
	en.handles[i] = n
	en.iter++
	en.Unlock()
	return i
}

func (en *EnvNames) Get(i int) string {
	en.RLock()
	defer en.RUnlock()
	if s, ok := en.handles[i]; ok {
		return s
	} else {
		panic(fmt.Sprintf("Не найден идентификатор переменной id=%d", i))
	}
}

// все переменные
var UniqueNames = NewEnvNames()

type Nullable interface {
	null()
	String()
}

type NullType struct {
	Nullable
}

func (x *NullType) null()          {}
func (x *NullType) String() string { return "NULL" }

var NullVar = NullType{}

// Expr provides all of interfaces for expression.
type Expr interface {
	Pos
	expr()
}

// ExprImpl provide commonly implementations for Expr.
type ExprImpl struct {
	PosImpl // ExprImpl provide Pos() function.
}

// expr provide restraint interface.
func (x *ExprImpl) expr() {}

// NumberExpr provide Number expression.
type NumberExpr struct {
	ExprImpl
	Lit string
}

// StringExpr provide String expression.
type StringExpr struct {
	ExprImpl
	Lit string
}

// ArrayExpr provide Array expression.
type ArrayExpr struct {
	ExprImpl
	Exprs []Expr
}

// PairExpr provide one of Map key/value pair.
type PairExpr struct {
	ExprImpl
	Key   string
	Value Expr
}

// MapExpr provide Map expression.
type MapExpr struct {
	ExprImpl
	MapExpr map[string]Expr
}

// IdentExpr provide identity expression.
type IdentExpr struct {
	ExprImpl
	Lit string
	Id  int
}

// UnaryExpr provide unary minus expression. ex: -1, ^1, ~1.
type UnaryExpr struct {
	ExprImpl
	Operator string
	Expr     Expr
}

// AddrExpr provide referencing address expression.
type AddrExpr struct {
	ExprImpl
	Expr Expr
}

// DerefExpr provide dereferencing address expression.
type DerefExpr struct {
	ExprImpl
	Expr Expr
}

// ParenExpr provide parent block expression.
type ParenExpr struct {
	ExprImpl
	SubExpr Expr
}

// BinOpExpr provide binary operator expression.
type BinOpExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

type TernaryOpExpr struct {
	ExprImpl
	Expr Expr
	Lhs  Expr
	Rhs  Expr
}

// CallExpr provide calling expression.
type CallExpr struct {
	ExprImpl
	Func     interface{}
	Name     int //string
	SubExprs []Expr
	VarArg   bool
	Go       bool
}

// AnonCallExpr provide anonymous calling expression. ex: func(){}().
type AnonCallExpr struct {
	ExprImpl
	Expr     Expr
	SubExprs []Expr
	VarArg   bool
	Go       bool
}

// MemberExpr provide expression to refer menber.
type MemberExpr struct {
	ExprImpl
	Expr Expr
	Name int //string
}

// ItemExpr provide expression to refer Map/Array item.
type ItemExpr struct {
	ExprImpl
	Value Expr
	Index Expr
}

// SliceExpr provide expression to refer slice of Array.
type SliceExpr struct {
	ExprImpl
	Value Expr
	Begin Expr
	End   Expr
}

// FuncExpr provide function expression.
type FuncExpr struct {
	ExprImpl
	Name   int //string
	Stmts  []Stmt
	Args   []int //string
	VarArg bool
}

// LetExpr provide expression to let variable.
type LetExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

// LetsExpr provide multiple expression of let.
type LetsExpr struct {
	ExprImpl
	Lhss     []Expr
	Operator string
	Rhss     []Expr
}

// AssocExpr provide expression to assoc operation.
type AssocExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

// NewExpr provide expression to make new instance.
type NewExpr struct {
	ExprImpl
	Type int //string
}

// ConstExpr provide expression for constant variable.
type ConstExpr struct {
	ExprImpl
	Value string
}

type ChanExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

type Type struct {
	Name int //string
}

type TypeCast struct {
	ExprImpl
	Type     int
	TypeExpr Expr // должен быть строкой
	CastExpr Expr
}

type MakeExpr struct {
	ExprImpl
	Type     int  //string
	TypeExpr Expr // должен быть строкой
}

type MakeChanExpr struct {
	ExprImpl
	// Type     int //string
	SizeExpr Expr
}

type MakeArrayExpr struct {
	ExprImpl
	// Type    int //string
	LenExpr Expr
	CapExpr Expr
}

// хранит реальное значение, рассчитанное на этапе оптимизации AST
type NativeExpr struct {
	ExprImpl
	Value reflect.Value
}

// Функции для работы с литералами
func InvokeConst(v string, defval reflect.Value) reflect.Value {
	switch strings.ToLower(v) {
	case "истина":
		return reflect.ValueOf(true)
	case "ложь":
		return reflect.ValueOf(false)
	case "null":
		return reflect.ValueOf(NullVar)
	}
	return defval
}

func InvokeNumber(lit string, defval reflect.Value) (reflect.Value, error) {
	if strings.Contains(lit, ".") || strings.Contains(lit, "e") {
		v, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return defval, err
		}
		return reflect.ValueOf(float64(v)), nil
	}
	var i int64
	var err error
	if strings.HasPrefix(lit, "0x") {
		i, err = strconv.ParseInt(lit[2:], 16, 64)
	} else {
		i, err = strconv.ParseInt(lit, 10, 64)
	}
	if err != nil {
		return defval, err
	}
	return reflect.ValueOf(i), nil
}

func EvalUnOp(op string, v, defval reflect.Value) (reflect.Value, error) {
	switch op {
	case "-":
		if v.Kind() == reflect.Float64 {
			return reflect.ValueOf(-v.Float()), nil
		}
		return reflect.ValueOf(-v.Int()), nil
	case "^":
		return reflect.ValueOf(^ToInt64(v)), nil
	case "!":
		return reflect.ValueOf(!ToBool(v)), nil
	default:
		return defval, fmt.Errorf("Неизвестный оператор")
	}
}

func EvalBinOp(op string, lhsV, rhsV, defval reflect.Value) (reflect.Value, error) {
	switch op {

	// TODO: расширить возможные варианты

	case "+":
		if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
			return reflect.ValueOf(ToString(lhsV) + ToString(rhsV)), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() != reflect.Array && rhsV.Kind() != reflect.Slice) {
			return reflect.Append(lhsV, rhsV), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() == reflect.Array || rhsV.Kind() == reflect.Slice) {
			return reflect.AppendSlice(lhsV, rhsV), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(ToFloat64(lhsV) + ToFloat64(rhsV)), nil
		}
		return reflect.ValueOf(ToInt64(lhsV) + ToInt64(rhsV)), nil
	case "-":
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(ToFloat64(lhsV) - ToFloat64(rhsV)), nil
		}
		return reflect.ValueOf(ToInt64(lhsV) - ToInt64(rhsV)), nil
	case "*":
		if lhsV.Kind() == reflect.String && (rhsV.Kind() == reflect.Int || rhsV.Kind() == reflect.Int32 || rhsV.Kind() == reflect.Int64) {
			return reflect.ValueOf(strings.Repeat(ToString(lhsV), int(ToInt64(rhsV)))), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(ToFloat64(lhsV) * ToFloat64(rhsV)), nil
		}
		return reflect.ValueOf(ToInt64(lhsV) * ToInt64(rhsV)), nil
	case "/":
		return reflect.ValueOf(ToFloat64(lhsV) / ToFloat64(rhsV)), nil
	case "%":
		return reflect.ValueOf(ToInt64(lhsV) % ToInt64(rhsV)), nil
	case "==":
		return reflect.ValueOf(Equal(lhsV, rhsV)), nil
	case "!=":
		return reflect.ValueOf(Equal(lhsV, rhsV) == false), nil
	case ">":
		return reflect.ValueOf(ToFloat64(lhsV) > ToFloat64(rhsV)), nil
	case ">=":
		return reflect.ValueOf(ToFloat64(lhsV) >= ToFloat64(rhsV)), nil
	case "<":
		return reflect.ValueOf(ToFloat64(lhsV) < ToFloat64(rhsV)), nil
	case "<=":
		return reflect.ValueOf(ToFloat64(lhsV) <= ToFloat64(rhsV)), nil
	case "|":
		return reflect.ValueOf(ToInt64(lhsV) | ToInt64(rhsV)), nil
	case "||":
		if ToBool(lhsV) {
			return lhsV, nil
		}
		return rhsV, nil
	case "&":
		return reflect.ValueOf(ToInt64(lhsV) & ToInt64(rhsV)), nil
	case "&&":
		if ToBool(lhsV) {
			return rhsV, nil
		}
		return lhsV, nil
	case "**":
		if lhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(math.Pow(ToFloat64(lhsV), ToFloat64(rhsV))), nil
		}
		return reflect.ValueOf(int64(math.Pow(ToFloat64(lhsV), ToFloat64(rhsV)))), nil
	case ">>":
		return reflect.ValueOf(ToInt64(lhsV) >> uint64(ToInt64(rhsV))), nil
	case "<<":
		return reflect.ValueOf(ToInt64(lhsV) << uint64(ToInt64(rhsV))), nil
	default:
		return defval, fmt.Errorf("Неизвестный оператор")
	}
}

// ToString converts all reflect.Value-s into string.
func ToString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	if !v.IsValid() {
		return "Неопределено"
	}
	if v.Kind() == reflect.Bool {
		if v.Bool() {
			return "Истина"
		} else {
			return "Ложь"
		}
	}
	return fmt.Sprint(v.Interface())
}

// ToBool converts all reflect.Value-s into bool.
func ToBool(v reflect.Value) bool {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0.0
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		vlow := strings.ToLower(v.String())
		if vlow == "true" || vlow == "истина" {
			return true
		}
		if ToInt64(v) != 0 {
			return true
		}
	}
	return false
}

// toFloat64 converts all reflect.Value-s into float64.
func ToFloat64(v reflect.Value) float64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	}
	return 0.0
}

func IsNil(v reflect.Value) bool {
	if !v.IsValid() || v.Kind().String() == "unsafe.Pointer" {
		return true
	}
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
		return true
	}
	return false
}

func IsNum(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// equal returns true when lhsV and rhsV is same value.
func Equal(lhsV, rhsV reflect.Value) bool {
	lhsIsNil, rhsIsNil := IsNil(lhsV), IsNil(rhsV)
	if lhsIsNil && rhsIsNil {
		return true
	}
	if (!lhsIsNil && rhsIsNil) || (lhsIsNil && !rhsIsNil) {
		return false
	}
	if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
		lhsV = lhsV.Elem()
	}
	if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
		rhsV = rhsV.Elem()
	}
	if !lhsV.IsValid() || !rhsV.IsValid() {
		return true
	}
	if IsNum(lhsV) && IsNum(rhsV) {
		if rhsV.Type().ConvertibleTo(lhsV.Type()) {
			rhsV = rhsV.Convert(lhsV.Type())
		}
	}
	if lhsV.CanInterface() && rhsV.CanInterface() {
		return reflect.DeepEqual(lhsV.Interface(), rhsV.Interface())
	}
	return reflect.DeepEqual(lhsV, rhsV)
}

// toInt64 converts all reflect.Value-s into int64.
func ToInt64(v reflect.Value) int64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.String:
		s := v.String()
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return int64(i)
		}
	}
	return 0
}

func TypeCastConvert(rv reflect.Value, nt reflect.Type, skipCollections bool, defval reflect.Value) (reflect.Value, error) {
	rvkind := rv.Kind()
	if skipCollections && (rvkind == reflect.Array || rvkind == reflect.Slice ||
		rvkind == reflect.Map || rvkind == reflect.Struct || rvkind == reflect.Chan) {
		return rv, nil
	}
	if rvkind == reflect.Interface || rvkind == reflect.Ptr {
		rv = rv.Elem()
		rvkind = rv.Kind()
	}
	// учитываем случай двойной вложенности указателя или интерфейса в указателе
	if rvkind == reflect.Interface || rvkind == reflect.Ptr {
		rv = rv.Elem()
		rvkind = rv.Kind()
	}
	if rvkind == nt.Kind() {
		return rv, nil
	}

	switch rvkind {
	case reflect.Array, reflect.Slice:
		switch nt.Kind() {
		case reflect.String:
			// сериализуем в json
			b, err := json.Marshal(rv.Interface())
			if err != nil {
				return rv, err
			}
			return reflect.ValueOf(string(b)), nil
		default:
			// преобразуем в такой же слайс, но с типизированными значениями, и копируем их с новым типом
			rs := reflect.MakeSlice(reflect.SliceOf(nt), rv.Len(), rv.Cap())
			for i := 0; i < rv.Len(); i++ {
				iv := rv.Index(i)
				// конверсия вложенных массивов и структур не производится
				rsi, err := TypeCastConvert(iv, nt, true, defval)
				if err != nil {
					return rv, err
				}
				sv := rs.Index(i)
				if sv.CanSet() {
					sv.Set(rsi)
				}
				//rs = reflect.Append(rs, rsi)
			}
			return rs, nil
		}
	case reflect.Chan:
		// возвращаем новый канал с типизированными значениями и прежним размером буфера
		return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, nt), rv.Cap()), nil
	case reflect.Map:
		switch nt.Kind() {
		case reflect.String:
			// сериализуем в json
			b, err := json.Marshal(rv.Interface())
			if err != nil {
				return rv, err
			}
			return reflect.ValueOf(string(b)), nil
		case reflect.Struct:
			// для приведения в структурные типы - можно использовать мапу для заполнения полей
			rs := reflect.New(nt) // указатель на новую структуру
			//заполняем экспортируемые неанонимные поля, если их находим в мапе
			for i := 0; i < nt.NumField(); i++ {
				f := nt.Field(i)
				if f.PkgPath == "" && !f.Anonymous {
					setv := reflect.Indirect(rv.MapIndex(reflect.ValueOf(f.Name)))
					if setv.Kind() == reflect.Interface {
						setv = setv.Elem()
					}
					fv := rs.Elem().FieldByName(f.Name)
					if setv.IsValid() && fv.IsValid() && fv.CanSet() {
						if fv.Kind() != setv.Kind() && setv.Type().ConvertibleTo(fv.Type()) {
							setv = setv.Convert(fv.Type())
						}
						fv.Set(setv)
					}
				}
			}
			return rs, nil
		}
	case reflect.String:
		switch nt.Kind() {
		case reflect.Float64:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
			f, err := strconv.ParseFloat(ToString(rv), 64)
			if err == nil {
				return reflect.ValueOf(f), nil
			}
		case reflect.Array, reflect.Slice:
			//парсим json из строки и пытаемся получить массив
			var rm []interface{}
			if err := json.Unmarshal([]byte(ToString(rv)), &rm); err != nil {
				return rv, err
			}
			return reflect.ValueOf(rm), nil
		case reflect.Map:
			//парсим json из строки и пытаемся получить мапу
			var rm map[string]interface{}
			if err := json.Unmarshal([]byte(ToString(rv)), rm); err != nil {
				return rv, err
			}
			return reflect.ValueOf(rm), nil
		case reflect.Struct:
			//парсим json из строки и пытаемся получить указатель на структуру
			rm := reflect.New(nt).Interface()
			if err := json.Unmarshal([]byte(ToString(rv)), rm); err != nil {
				return rv, err
			}
			return reflect.ValueOf(rm), nil
		case reflect.Int64:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
			i, err := strconv.ParseInt(ToString(rv), 10, 64)
			if err == nil {
				return reflect.ValueOf(i), nil
			}
			f, err := strconv.ParseFloat(ToString(rv), 64)
			if err == nil {
				return reflect.ValueOf(int64(f)), nil
			}
		case reflect.Bool:
			s := strings.ToLower(ToString(rv))
			if s == "истина" {
				return reflect.ValueOf(true), nil
			}
			if rv.Type().ConvertibleTo(reflect.TypeOf(1.0)) && rv.Convert(reflect.TypeOf(1.0)).Float() > 0.0 {
				return reflect.ValueOf(true), nil
			}
			b, err := strconv.ParseBool(s)
			if err == nil {
				return reflect.ValueOf(b), nil
			}
			return reflect.ValueOf(false), nil
		default:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
		}
	case reflect.Bool:
		switch nt.Kind() {
		case reflect.String:
			if ToBool(rv) {
				return reflect.ValueOf("Истина"), nil
			} else {
				return reflect.ValueOf("Ложь"), nil
			}
		case reflect.Int64:
			if ToBool(rv) {
				return reflect.ValueOf(int64(1)), nil
			} else {
				return reflect.ValueOf(int64(0)), nil
			}
		case reflect.Float64:
			if ToBool(rv) {
				return reflect.ValueOf(float64(1.0)), nil
			} else {
				return reflect.ValueOf(float64(0.0)), nil
			}
		}
	case reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// числа конвертируются стандартно
		if rv.Type().ConvertibleTo(nt) {
			return rv.Convert(nt), nil
		}
	case reflect.Struct:
		if t, ok := rv.Interface().(time.Time); ok {
			// это дата/время - конвертируем в секунды (целые или с плавающей запятой) или в формат RFC3339
			switch nt.Kind() {
			case reflect.String:
				return reflect.ValueOf(t.Format(time.RFC3339)), nil
			case reflect.Int64:
				return reflect.ValueOf(t.Unix()), nil
			case reflect.Float64:
				return reflect.ValueOf(float64(t.UnixNano()) / 1e9), nil
			}
		} else {
			switch nt.Kind() {
			case reflect.Map:
				// структура может быть приведена в мапу
				rs := make(map[string]interface{})
				rtyp := rv.Type()
				for i := 0; i < rtyp.NumField(); i++ {
					f := rtyp.Field(i)
					fv := rv.Field(i)
					if f.PkgPath == "" && !f.Anonymous {
						rs[f.Name] = fv.Interface()
					}
				}
				return reflect.ValueOf(rs), nil
			case reflect.String:
				// сериализуем структуру в json
				b, err := json.Marshal(rv.Interface())
				if err != nil {
					return rv, err
				}
				return reflect.ValueOf(string(b)), nil

			}
		}
	}
	return defval, fmt.Errorf("Приведение типа недопустимо")
}

func MethodByNameCI(v reflect.Value, name int) reflect.Value {
	tv := v.Type()
	for i := 0; i < tv.NumMethod(); i++ {
		meth := tv.Method(i)
		if UniqueNames.Set(meth.Name) == name {
			return v.Method(i)
		}
	}
	return reflect.Value{}
}

func FieldByNameCI(v reflect.Value, name int) reflect.Value {
	tv := v.Type()
	for i := 0; i < tv.NumField(); i++ {
		f := tv.Field(i)
		if f.PkgPath == "" && !f.Anonymous && UniqueNames.Set(f.Name) == name {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}
