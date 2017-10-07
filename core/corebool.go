package core

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/covrom/gonec/names"
	"github.com/shopspring/decimal"
)

// VMInt для ускорения работы храним целочисленное представление отдельно от decimal
type VMBool bool

var ReflectVMBool = reflect.TypeOf(VMBool(true))

func (x VMBool) vmval() {}

func (x VMBool) Interface() interface{} {
	return bool(x)
}

func (x *VMBool) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case bool:
		*x = VMBool(vv)
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		*x = VMBool(rv.Bool()) // выдаст панику, если это не булево
	}
}

func (x VMBool) String() string {
	if x {
		return "true"
	} else {
		return "false"
	}
}

func (x VMBool) Int() int64 {
	if x {
		return 1
	} else {
		return 0
	}
}

func (x VMBool) Float() float64 {
	if x {
		return 1.0
	} else {
		return 0.0
	}
}

func (x VMBool) Decimal() VMDecimal {
	if x {
		return VMDecimal(decimal.New(1, 0))
	} else {
		return VMDecimal(decimal.New(0, 0))
	}
}

func (x VMBool) Bool() bool {
	return bool(x)
}

func (x VMBool) BinaryType() VMBinaryType {
	return VMBOOL
}

func (x VMBool) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func ParseVMBool(s string) (VMBool, error) {
	switch names.FastToLower(s) {
	case "true", "истина":
		return true, nil
	case "false", "ложь":
		return false, nil
	}
	return false, VMErrorNotConverted
}

func (x VMBool) EvalUnOp(op rune) (VMValuer, error) {
	switch op {
	// case '-':
	// case '^':
	case '!':
		return VMBool(!bool(x)), nil
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMBool) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		return VMNil, VMErrorIncorrectOperation
	case SUB:
		return VMNil, VMErrorIncorrectOperation
	case MUL:
		return VMNil, VMErrorIncorrectOperation
	case QUO:
		return VMNil, VMErrorIncorrectOperation
	case REM:
		return VMNil, VMErrorIncorrectOperation
	case EQL:
		switch yy := y.(type) {
		case VMBool:
			return VMBool(bool(x) == bool(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case NEQ:
		switch yy := y.(type) {
		case VMBool:
			return VMBool(bool(x) != bool(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GTR:
		return VMNil, VMErrorIncorrectOperation
	case GEQ:
		return VMNil, VMErrorIncorrectOperation
	case LSS:
		return VMNil, VMErrorIncorrectOperation
	case LEQ:
		return VMNil, VMErrorIncorrectOperation
	case OR:
		return VMNil, VMErrorIncorrectOperation
	case LOR:
		switch yy := y.(type) {
		case VMBool:
			return VMBool(bool(x) || bool(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case AND:
		return VMNil, VMErrorIncorrectOperation
	case LAND:
		switch yy := y.(type) {
		case VMBool:
			return VMBool(bool(x) && bool(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case POW:
		return VMNil, VMErrorIncorrectOperation
	case SHR:
		return VMNil, VMErrorIncorrectOperation
	case SHL:
		return VMNil, VMErrorIncorrectOperation
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMBool) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMBool:
		return x, nil
	case ReflectVMString:
		return VMString(x.String()), nil
	case ReflectVMInt:
		return VMInt(x.Int()), nil
	// case ReflectVMTime:
	case ReflectVMDecimal:
		return x.Decimal(), nil
		// case ReflectVMSlice:
		// case ReflectVMStringMap:
	}

	return VMNil, VMErrorNotConverted
}

func (x VMBool) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if bool(x) {
		buf.WriteByte(byte(0))
	} else {
		buf.WriteByte(byte(1))
	}
	return buf.Bytes(), nil
}

func (x *VMBool) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	if v, err := buf.ReadByte(); err != nil {
		return err
	} else {
		if v == 0 {
			*x = VMBool(false)
		} else {
			*x = VMBool(true)
		}
	}
	return nil
}

func (x VMBool) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMBool) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(x))
}

func (x *VMBool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var b bool
	err := json.Unmarshal(data, &b)
	if err != nil {
		return err
	}
	*x = VMBool(b)
	return nil
}

func (x VMBool) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *VMBool) UnmarshalText(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	b, err := ParseVMBool(string(data))
	if err != nil {
		return err
	}
	*x = b
	return nil
}
