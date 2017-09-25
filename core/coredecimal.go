package core

import (
	"fmt"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

// VMDecimal с плавающей токой, для финансовых расчетов высокой точности (decimal)
type VMDecimal decimal.Decimal

var ReflectVMDecimal = reflect.TypeOf(VMDecimal{})

func (x VMDecimal) vmval() {}

func (x VMDecimal) Interface() interface{} {
	return decimal.Decimal(x)
}

func (x *VMDecimal) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int8:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int16:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int32:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int64:
		*x = VMDecimal(decimal.New(vv, 0))
	case uint:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint8:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint16:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint32:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint64:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uintptr:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case float32:
		*x = VMDecimal(decimal.NewFromFloat(float64(vv)))
	case float64:
		*x = VMDecimal(decimal.NewFromFloat(vv))
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Float32 || rv.Kind() == reflect.Float64 {
			*x = VMDecimal(decimal.NewFromFloat(rv.Float()))
		} else {
			*x = VMDecimal(decimal.New(rv.Int(), 0)) // выдаст панику, если это не число
		}
	}
}

func (x VMDecimal) String() string {
	return decimal.Decimal(x).String()
}

func (x VMDecimal) Int() int64 {
	return decimal.Decimal(x).IntPart()
}

func (x VMDecimal) Float() float64 {
	f64, ok := decimal.Decimal(x).Float64()
	if !ok {
		panic("Невозможно получить значение с плавающей запятой 64 бит")
	}
	return f64
}

func (x VMDecimal) Decimal() VMDecimal {
	return x
}

func (x VMDecimal) InvokeNumber() (VMNumberer, error) {
	return x, nil
}

func (x VMDecimal) Bool() bool {
	return decimal.Decimal(x).GreaterThan(decimal.Zero)
}

func (x VMDecimal) MakeChan(size int) VMChaner {

	return make(VMChan, size)
}

func (x VMDecimal) Time() VMTime {
	intpart := decimal.Decimal(x).IntPart()
	nanopart := decimal.Decimal(x).Sub(decimal.New(intpart, 0)).Mul(decimal.New(int64(VMSecond), 0)).IntPart()
	return VMTime(time.Unix(intpart, nanopart))
}

func (x VMDecimal) Duration() VMTimeDuration {
	return VMTimeDuration(time.Duration(decimal.Decimal(x).Mul(decimal.New(int64(VMSecond), 0)).IntPart()))
}

func ParseVMDecimal(s string) (VMDecimal, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return VMDecimal{}, err
	}
	return VMDecimal(d), nil
}

func (x VMDecimal) Add(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Add(decimal.Decimal(d2)))
}

func (x VMDecimal) Sub(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Sub(decimal.Decimal(d2)))
}

func (x VMDecimal) Mul(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Mul(decimal.Decimal(d2)))
}

func (x VMDecimal) Div(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Div(decimal.Decimal(d2)))
}

func (x VMDecimal) Mod(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Mod(decimal.Decimal(d2)))
}

func (x VMDecimal) Pow(d2 VMDecimal) VMDecimal {
	return VMDecimal(decimal.Decimal(x).Pow(decimal.Decimal(d2)))
}

func (x VMDecimal) Equal(d2 VMDecimal) VMBool {
	return VMBool(decimal.Decimal(x).Equal(decimal.Decimal(d2)))
}

func (x VMDecimal) NotEqual(d2 VMDecimal) VMBool {
	return VMBool(!decimal.Decimal(x).Equal(decimal.Decimal(d2)))
}

func (x VMDecimal) Cmp(d2 VMDecimal) int {
	//     -1 if x <  d2
	//      0 if x == d2
	//     +1 if x >  d2
	return decimal.Decimal(x).Cmp(decimal.Decimal(d2))
}

func NewVMDecimalFromInt64(x int64) VMDecimal {
	return VMDecimal(decimal.New(x, 0))
}

func (x VMDecimal) EvalUnOp(op rune) (VMValuer, error) {
	switch op {
	case '-':
		return VMDecimal(decimal.Decimal(x).Neg()), nil
	case '!':
		return VMBool(!x.Bool()), nil
	default:
		return VMNil, fmt.Errorf("Неизвестный оператор")
	}
}

func (x VMDecimal) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMInt:
			return x.Add(NewVMDecimalFromInt64(int64(yy))), nil
		case VMDecimal:
			return x.Add(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SUB:
		switch yy := y.(type) {
		case VMInt:
			return x.Sub(NewVMDecimalFromInt64(int64(yy))), nil
		case VMDecimal:
			return x.Sub(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return x.Mul(NewVMDecimalFromInt64(int64(yy))), nil
		case VMDecimal:
			return x.Mul(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case QUO:
		switch yy := y.(type) {
		case VMInt:
			return x.Div(NewVMDecimalFromInt64(int64(yy))), nil
		case VMDecimal:
			return x.Div(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case REM:
		switch yy := y.(type) {
		case VMInt:
			return x.Mod(NewVMDecimalFromInt64(int64(yy))), nil
		case VMDecimal:
			return x.Mod(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case EQL:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(x.Equal(NewVMDecimalFromInt64(int64(yy)))), nil
		case VMDecimal:
			return x.Equal(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(x.NotEqual(NewVMDecimalFromInt64(int64(yy)))), nil
		case VMDecimal:
			return x.NotEqual(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case GTR:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(x.Cmp(NewVMDecimalFromInt64(int64(yy))) == 1), nil
		case VMDecimal:
			return VMBool(x.Cmp(yy) == 1), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case GEQ:
		switch yy := y.(type) {
		case VMInt:
			cmp := x.Cmp(NewVMDecimalFromInt64(int64(yy)))
			return VMBool(cmp == 1 || cmp == 0), nil
		case VMDecimal:
			cmp := x.Cmp(yy)
			return VMBool(cmp == 1 || cmp == 0), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LSS:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(x.Cmp(NewVMDecimalFromInt64(int64(yy))) == -1), nil
		case VMDecimal:
			return VMBool(x.Cmp(yy) == -1), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LEQ:
		switch yy := y.(type) {
		case VMInt:
			cmp := x.Cmp(NewVMDecimalFromInt64(int64(yy)))
			return VMBool(cmp == -1 || cmp == 0), nil
		case VMDecimal:
			cmp := x.Cmp(yy)
			return VMBool(cmp == -1 || cmp == 0), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case OR:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LOR:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case AND:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case LAND:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case POW:
		switch yy := y.(type) {
		case VMInt:
			return x.Pow(NewVMDecimalFromInt64(int64(yy))), nil
		case VMDecimal:
			return x.Pow(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SHR:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SHL:
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	}
	return VMNil, fmt.Errorf("Неизвестная операция")
}

func (x VMDecimal) ConvertToType(nt reflect.Type, skipCollections bool) (VMValuer, error) {
	switch nt {
	case ReflectVMDecimal:
		return x, nil
	case ReflectVMTime:
		return x.Time(), nil
	case ReflectVMTimeDuration:
		return x.Duration(), nil
	case ReflectVMBool:
		return VMBool(x.Bool()), nil
	case ReflectVMString:
		return VMString(x.String()), nil
	case ReflectVMInt:
		return VMInt(x.Int()), nil
	}
	return VMNil, fmt.Errorf("Приведение к типу невозможно")
}
