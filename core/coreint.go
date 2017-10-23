package core

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/covrom/decnum"
)

// VMInt для ускорения работы храним целочисленное представление отдельно от decimal
type VMInt int64

var ReflectVMInt = reflect.TypeOf(VMInt(0))

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

func (x VMInt) DecNum() VMDecNum {
	return VMDecNum{num: decnum.Zero()}
}

func (x VMInt) InvokeNumber() (VMNumberer, error) {
	return x, nil
}

func (x VMInt) Bool() bool {
	return x > 0
}

func (x VMInt) BinaryType() VMBinaryType {
	return VMINT
}

func (x VMInt) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMInt) Time() VMTime {
	return VMTime(time.Unix(int64(x), 0))
}

func (x VMInt) Duration() VMTimeDuration {
	return VMTimeDuration(time.Duration(int64(x) * int64(VMSecond)))
}

func ParseVMInt(s string) (VMInt, error) {
	var i64 int64
	var err error
	if strings.HasPrefix(s, "0x") {
		i64, err = strconv.ParseInt(s[2:], 16, 64)
	} else {
		i64, err = strconv.ParseInt(s, 10, 64)
	}
	if err != nil {
		return VMInt(0), err
	}
	return VMInt(i64), nil
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

func (x VMInt) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) + int64(yy)), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Add(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SUB:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) - int64(yy)), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Sub(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) * int64(yy)), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Mul(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case QUO:
		switch yy := y.(type) {
		case VMInt:
			return NewVMDecNumFromInt64(int64(x)).Div(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Div(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case REM:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) % int64(yy)), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Mod(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case EQL:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(int64(x) == int64(yy)), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Equal(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case NEQ:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(int64(x) != int64(yy)), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).NotEqual(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GTR:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(int64(x) > int64(yy)), nil
		case VMDecNum:
			return VMBool(yy.Less(NewVMDecNumFromInt64(int64(x)))), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GEQ:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(int64(x) >= int64(yy)), nil
		case VMDecNum:
			cmp := NewVMDecNumFromInt64(int64(x)).num.GreaterEqual(yy.num)
			return VMBool(cmp), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LSS:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(int64(x) < int64(yy)), nil
		case VMDecNum:
			return VMBool(NewVMDecNumFromInt64(int64(x)).Less(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LEQ:
		switch yy := y.(type) {
		case VMInt:
			return VMBool(int64(x) <= int64(yy)), nil
		case VMDecNum:
			cmp := NewVMDecNumFromInt64(int64(x)).num.LessEqual(yy.num)
			return VMBool(cmp), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case OR:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) | int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LOR:
		return VMNil, VMErrorIncorrectOperation
	case AND:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) & int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LAND:
		return VMNil, VMErrorIncorrectOperation
	case POW:
		switch yy := y.(type) {
		case VMInt:
			// TODO: переделать на math
			return NewVMDecNumFromInt64(int64(x)).Pow(NewVMDecNumFromInt64(int64(yy))), nil
		case VMDecNum:
			return NewVMDecNumFromInt64(int64(x)).Pow(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SHR:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) >> uint64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SHL:
		switch yy := y.(type) {
		case VMInt:
			return VMInt(int64(x) << uint64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMInt) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMInt:
		return x, nil
	case ReflectVMTime:
		// приведение к дате - исходное число в секундах
		return x.Time(), nil
	case ReflectVMTimeDuration:
		return x.Duration(), nil
	case ReflectVMBool:
		return VMBool(x.Bool()), nil
	case ReflectVMString:
		return VMString(x.String()), nil
	case ReflectVMDecNum:
		return x.DecNum(), nil
	}
	return VMNil, VMErrorNotConverted
}

// маршаллинг нужен для того, чтобы encoding не использовал рефлексию

func (x VMInt) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, int64(x))
	return buf.Bytes(), nil
}

func (x *VMInt) UnmarshalBinary(data []byte) error {
	var i int64
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &i); err != nil {
		return err
	}
	*x = VMInt(i)
	return nil
}

func (x VMInt) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMInt) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMInt) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(x.String())
	return buf.Bytes(), nil
}

func (x *VMInt) UnmarshalText(data []byte) error {
	i64, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*x = VMInt(i64)
	return nil
}

func (x VMInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(x))
}

func (x *VMInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var i64 int64
	err := json.Unmarshal(data, &i64)
	if err != nil {
		return err
	}
	*x = VMInt(i64)
	return nil
}
