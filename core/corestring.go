package core

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/covrom/decnum"
)

// VMString строки
type VMString string

var ReflectVMString = reflect.TypeOf(VMString(""))

func (x VMString) vmval() {}

func (x VMString) Interface() interface{} {
	return string(x)
}

func (x VMString) String() string {
	return string(x)
}

func (x VMString) Length() VMInt {
	return VMInt(utf8.RuneCountInString(string(x)))
}

func (x VMString) IndexVal(i VMValuer) VMValuer {
	if ii, ok := i.(VMInt); ok {
		return VMString(string([]rune(string(x))[int(ii)]))
	}
	panic("Индекс должен быть числом")
}

func (x VMString) Int() int64 {
	var i64 int64
	var err error
	if strings.HasPrefix(string(x), "0x") {
		i64, err = strconv.ParseInt(string(x)[2:], 16, 64)
	} else {
		i64, err = strconv.ParseInt(string(x), 10, 64)
	}
	if err != nil {
		panic(err)
	}
	return i64
}

func (x VMString) Float() float64 {
	f64, err := strconv.ParseFloat(string(x), 64)
	if err != nil {
		panic(err)
	}
	return f64
}

func (x VMString) Decimal() VMDecNum {
	d, err := decnum.FromString(string(x))
	if err != nil {
		panic(err)
	}
	return VMDecNum{num: d}
}

func (x VMString) InvokeNumber() (v VMNumberer, err error) {
	if strings.ContainsAny(string(x), ".eE") {
		v, err = ParseVMDecNum(string(x))
	} else {
		v, err = ParseVMInt(string(x))
	}
	return
}

func (x VMString) BinaryType() VMBinaryType {
	return VMSTRING
}

func (x VMString) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMString) Hash() VMString {
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, HashBytes([]byte(x)))
	return VMString(hex.EncodeToString(h))
}

func (x VMString) Time() VMTime {
	t, err := time.Parse(time.RFC3339, string(x))
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("2006-01-02T15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("2006-01-02 15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", string(x))
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006 15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("20060102150405", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("20060102", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("2006-01-02", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse(time.RFC1123, string(x))
	if err == nil {
		return VMTime(t)
	}
	panic("Неверный формат даты и времени")
}

func (x VMString) Bool() bool {
	r, _ := ParseVMBool(string(x))
	return r.Bool()
}

func (x VMString) Slice() VMSlice {
	var rm VMSlice
	if err := json.Unmarshal([]byte(x), &rm); err != nil {
		panic(err)
	}
	return rm
}

func (x VMString) StringMap() VMStringMap {
	var rm VMStringMap
	if err := json.Unmarshal([]byte(x), rm); err != nil {
		panic(err)
	}
	return rm
}

func (x VMString) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMString:
			return VMString(string(x) + string(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SUB:
		switch yy := y.(type) {
		case VMString:
			return VMString(strings.Replace(string(x), string(yy), "", -1)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return VMString(strings.Repeat(string(x), int(yy))), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case QUO:
		return VMNil, VMErrorIncorrectOperation
	case REM:
		return VMNil, VMErrorIncorrectOperation
	case EQL:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Equal([]byte(x), []byte(yy))), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case NEQ:
		switch yy := y.(type) {
		case VMString:
			return VMBool(!bytes.Equal([]byte(x), []byte(yy))), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GTR:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Compare([]byte(x), []byte(yy)) == 1), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GEQ:
		switch yy := y.(type) {
		case VMString:
			cmp := bytes.Compare([]byte(x), []byte(yy))
			return VMBool(cmp == 1 || cmp == 0), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LSS:
		switch yy := y.(type) {
		case VMString:
			return VMBool(bytes.Compare([]byte(x), []byte(yy)) == -1), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LEQ:
		switch yy := y.(type) {
		case VMString:
			cmp := bytes.Compare([]byte(x), []byte(yy))
			return VMBool(cmp == -1 || cmp == 0), nil
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
		return VMNil, VMErrorIncorrectOperation
	case SHR:
		return VMNil, VMErrorIncorrectOperation
	case SHL:
		return VMNil, VMErrorIncorrectOperation
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMString) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMString:
		return x, nil
	case ReflectVMInt:
		return VMInt(x.Int()), nil
	case ReflectVMTime:
		return x.Time(), nil
	case ReflectVMBool:
		return VMBool(x.Bool()), nil
	case ReflectVMDecNum:
		return x.Decimal(), nil
	case ReflectVMSlice:
		return VMSliceFromJson(string(x))
	case ReflectVMStringMap:
		return VMStringMapFromJson(string(x))
	}

	// попробуем десериализировать структуру из json
	if nt.Kind() == reflect.Struct {
		//парсим json из строки и пытаемся получить указатель на структуру
		rm := reflect.New(nt).Interface()
		if err := json.Unmarshal([]byte(x), rm); err != nil {
			return VMNil, err
		}
		if rv, ok := rm.(VMValuer); ok {
			if vobj, ok := rv.(VMMetaObject); ok {
				vobj.VMInit(vobj)
				vobj.VMRegister()
				return vobj, nil
			} else {
				return nil, VMErrorIncorrectStructType
				//return rv, nil
			}
		}
		return VMNil, VMErrorIncorrectStructType
	}
	return VMNil, VMErrorNotConverted
}

func (x VMString) MarshalBinary() ([]byte, error) {
	return []byte(string(x)), nil
}

func (x *VMString) UnmarshalBinary(data []byte) error {
	*x = VMString(string(data))
	return nil
}

func (x VMString) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMString) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(x))
}

func (x *VMString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var us string
	err := json.Unmarshal(data, &us)
	if err != nil {
		return err
	}
	*x = VMString(us)
	return nil
}

func (x VMString) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *VMString) UnmarshalText(data []byte) error {
	*x = VMString(data)
	return nil
}
