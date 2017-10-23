package core

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
	"time"

	"github.com/covrom/decnum"
)

// VMDecNum с плавающей токой, для финансовых расчетов высокой точности (decnum)
type VMDecNum struct {
	num decnum.Quad
}

var VMDecNumZero = VMDecNum{num: decnum.Zero()}
var ReflectVMDecNum = reflect.TypeOf(VMDecNumZero)
var VMDecNumOne = NewVMDecNumFromInt64(1)
var VMDecNumNegOne = NewVMDecNumFromInt64(-1)

func (x VMDecNum) vmval() {}

func (x VMDecNum) Interface() interface{} {
	return x.num
}

func (x *VMDecNum) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int8:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int16:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int32:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case int64:
		*x = VMDecNum{num: decnum.FromInt64(vv)}
	case uint:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint8:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint16:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint32:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uint64:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case uintptr:
		*x = VMDecNum{num: decnum.FromInt64(int64(vv))}
	case float32:
		*x = VMDecNum{num: decnum.FromFloat(float64(vv))}
	case float64:
		*x = VMDecNum{num: decnum.FromFloat(vv)}
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Float32 || rv.Kind() == reflect.Float64 {
			*x = VMDecNum{num: decnum.FromFloat(rv.Float())}
		} else {
			*x = VMDecNum{num: decnum.FromInt64(rv.Int())}
		}
	}
}

func (x VMDecNum) String() string {
	return x.num.String()
}

func (x VMDecNum) Int() int64 {
	i, err := x.num.ToInt64(decnum.RoundDown) //целая часть, без округления
	if err != nil {
		return 0
	}
	return i
}

func (x VMDecNum) RoundHalfUp() int64 {
	i, err := x.num.ToInt64(decnum.RoundHalfUp) //целая часть, округление вверх, если модуль>0.5
	if err != nil {
		return 0
	}
	return i
}

func (x VMDecNum) Float() float64 {
	i, err := x.num.ToFloat64()
	if err != nil {
		return i
	}
	return i
}

func (x VMDecNum) DecNum() VMDecNum {
	return x
}

func (x VMDecNum) InvokeNumber() (VMNumberer, error) {
	return x, nil
}

func (x VMDecNum) Bool() bool {
	return x.num.IsPositive()
}

func (x VMDecNum) BinaryType() VMBinaryType {
	return VMDECNUM
}

func (x VMDecNum) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMDecNum) Time() VMTime {
	intpart := x.Int()
	rem := x.num.Sub(decnum.FromInt64(intpart))
	nano := rem.Mul(decnum.FromInt64(int64(VMSecond)))
	nanopart, _ := nano.ToInt64(decnum.RoundHalfUp)
	return VMTime(time.Unix(intpart, nanopart))
}

func (x VMDecNum) Duration() VMTimeDuration {
	i, _ := x.num.Mul(decnum.FromInt64(int64(VMSecond))).ToInt64(decnum.RoundHalfUp)
	return VMTimeDuration(time.Duration(i))
}

func ParseVMDecNum(s string) (VMDecNum, error) {
	d, err := decnum.FromString(s)
	return VMDecNum{num: d}, err
}

func (x VMDecNum) Add(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Add(d2.num)}
}

func (x VMDecNum) Sub(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Sub(d2.num)}
}

func (x VMDecNum) Mul(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Mul(d2.num)}
}

func (x VMDecNum) Div(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Div(d2.num)}
}

func (x VMDecNum) Mod(d2 VMDecNum) VMDecNum {
	return VMDecNum{num: x.num.Mod(d2.num)}
}

// Pow работает только в диапазоне и точности чисел float64
func (x VMDecNum) Pow(d2 VMDecNum) VMDecNum {
	v1, err := x.num.ToFloat64()
	if err != nil {
		panic(err)
	}
	v2, err := d2.num.ToFloat64()
	if err != nil {
		panic(err)
	}
	rv := math.Pow(v1, v2)
	return VMDecNum{num: decnum.FromFloat(rv)}
}

func (x VMDecNum) Equal(d2 VMDecNum) VMBool {
	return VMBool(x.num.Equal(d2.num))
}

func (x VMDecNum) NotEqual(d2 VMDecNum) VMBool {
	return VMBool(!x.num.Equal(d2.num))
}

func (x VMDecNum) Less(d2 VMDecNum) bool {
	// x <  d2
	return x.num.Less(d2.num)
}

func NewVMDecNumFromInt64(x int64) VMDecNum {
	return VMDecNum{num: decnum.FromInt64(x)}
}

func (x VMDecNum) EvalUnOp(op rune) (VMValuer, error) {
	switch op {
	case '-':
		return VMDecNum{num: x.num.Neg()}, nil
	case '!':
		return VMBool(!x.Bool()), nil
	default:
		return VMNil, VMErrorUnknownOperation
	}
}

func (x VMDecNum) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMInt:
			return x.Add(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Add(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SUB:
		switch yy := y.(type) {
		case VMInt:
			return x.Sub(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Sub(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return x.Mul(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Mul(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case QUO:
		switch yy := y.(type) {
		case VMInt:
			return x.Div(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Div(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case REM:
		switch yy := y.(type) {
		case VMInt:
			return x.Mod(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Mod(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case EQL:
		switch yy := y.(type) {
		case VMInt:
			return x.Equal(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Equal(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case NEQ:
		switch yy := y.(type) {
		case VMInt:
			return x.NotEqual(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.NotEqual(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GTR:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(NewVMDecNumFromInt64(int64(yy)).Less(x)), nil
		case VMDecNum:
			return VMBool(yy.Less(x)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GEQ:
		switch yy := y.(type) {
		case VMInt:
			cmp := x.num.GreaterEqual(decnum.FromInt64(int64(yy)))
			return VMBool(cmp), nil
		case VMDecNum:
			cmp := x.num.GreaterEqual(yy.num)
			return VMBool(cmp), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LSS:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(x.Less(NewVMDecNumFromInt64(int64(yy)))), nil
		case VMDecNum:
			return VMBool(x.Less(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LEQ:
		switch yy := y.(type) {
		case VMInt:
			cmp := x.num.LessEqual(decnum.FromInt64(int64(yy)))
			return VMBool(cmp), nil
		case VMDecNum:
			cmp := x.num.LessEqual(yy.num)
			return VMBool(cmp), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case OR:
		return VMNil, VMErrorIncorrectOperation
	case LOR:
		return VMNil, VMErrorIncorrectOperation
	case AND:
		return VMNil, VMErrorIncorrectOperation
	case LAND:
		return VMNil, VMErrorIncorrectOperation
	case POW:
		switch yy := y.(type) {
		case VMInt:
			return x.Pow(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return x.Pow(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SHR:
		return VMNil, VMErrorIncorrectOperation
	case SHL:
		return VMNil, VMErrorIncorrectOperation
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMDecNum) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMDecNum:
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
	return VMNil, VMErrorNotConverted
}

func (x VMDecNum) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, x.num)
	return buf.Bytes(), err
}

func (x *VMDecNum) UnmarshalBinary(data []byte) error {
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &x.num)
	return err
}

func (x VMDecNum) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMDecNum) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMDecNum) MarshalJSON() ([]byte, error) {
	return x.num.MarshalJSON()
}

func (x *VMDecNum) UnmarshalJSON(data []byte) error {
	return x.num.UnmarshalJSON(data)
}

func (x VMDecNum) MarshalText() ([]byte, error) {
	return []byte(x.num.String()), nil
}

func (x *VMDecNum) UnmarshalText(data []byte) error {
	var err error
	x.num, err = decnum.FromString(string(data))
	return err
}
