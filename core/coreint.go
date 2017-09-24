package core

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// VMInt для ускорения работы храним целочисленное представление отдельно от decimal
type VMInt int64

func (x VMInt) vmval() {}

func (x VMInt) Interface() interface{} {
	return int64(x)
}

func (x *VMInt) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMInt(vv)
	case int8:
		*x = VMInt(vv)
	case int16:
		*x = VMInt(vv)
	case int32:
		*x = VMInt(vv)
	case int64:
		*x = VMInt(vv)
	case uint:
		*x = VMInt(vv)
	case uint8:
		*x = VMInt(vv)
	case uint16:
		*x = VMInt(vv)
	case uint32:
		*x = VMInt(vv)
	case uint64:
		*x = VMInt(vv)
	case uintptr:
		*x = VMInt(vv)
	case float32:
		*x = VMInt(int64(vv))
	case float64:
		*x = VMInt(int64(vv))
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		*x = VMInt(rv.Int()) // выдаст панику, если это не число
	}
}

func (x VMInt) String() string {
	return strconv.FormatInt(int64(x), 10)
}

func (x VMInt) Int() int64 {
	return int64(x)
}

func (x VMInt) Float() float64 {
	return float64(x)
}

func (x VMInt) Decimal() VMDecimal {
	return VMDecimal(decimal.New(int64(x), 0))
}

func (x VMInt) Bool() bool {
	return x > 0
}

func (x VMInt) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMInt) Time() VMTime {
	return VMTime(time.Unix(int64(x), 0))
}

func (x *VMInt) Parse(s string) error {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*x = VMInt(i64)
	return nil
}

func (x VMInt) EvalUnOp(op rune) (VMValuer, error) {
	switch op {
	case '-':
		return VMInt(-int64(x)), nil
	case '^':
		return VMInt(^int64(x)), nil
	case '!':
		return VMBool(!x.Bool()), nil
	default:
		return VMNil, fmt.Errorf("Неизвестный оператор")
	}
}

func (x VMInt) Duration() VMTimeDuration {
	return VMTimeDuration(time.Duration(int64(x) * int64(VMSecond)))
}

func (x VMInt) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) + int64(yy)), nil
		case VMDecimal:
			return NewVMDecimalFromInt64(int64(x)).Add(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case SUB:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) - int64(yy)), nil
		case VMDecimal:
			return NewVMDecimalFromInt64(int64(x)).Sub(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) * int64(yy)), nil
		case VMDecimal:
			return NewVMDecimalFromInt64(int64(x)).Mul(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case QUO:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) / int64(yy)), nil
		case VMDecimal:
			return NewVMDecimalFromInt64(int64(x)).Div(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case REM:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) % int64(yy)), nil
		case VMDecimal:
			return NewVMDecimalFromInt64(int64(x)).Mod(yy), nil
		}
		return VMNil, fmt.Errorf("Операция между значениями невозможна")
	case EQL:
	case NEQ:
	case GTR:
	case GEQ:
	case LSS:
	case LEQ:
	case OR:
	case LOR:
	case AND:
	case LAND:
	case POW:
	case SHR:
	case SHL:
	}
	return VMNil, fmt.Errorf("Неизвестная операция")
}

// TODO:
func (x VMInt) ConvertToType(t reflect.Type, skipCollections bool) (VMValuer, error) {
	return VMNil, fmt.Errorf("Не реализовано")

}