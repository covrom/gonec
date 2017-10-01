package core

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type VMStringMap map[string]VMValuer

var ReflectVMStringMap = reflect.TypeOf(make(VMStringMap, 0))

func (x VMStringMap) vmval() {}

func (x VMStringMap) Interface() interface{} {
	return x
}

func (x VMStringMap) StringMap() VMStringMap {
	return x
}

func (x VMStringMap) Len() VMInt {
	return VMInt(len(x))
}

func (x VMStringMap) Index(i VMValuer) VMValuer {
	if s, ok := i.(VMString); ok {
		return x[string(s)]
	}
	panic("Индекс должен быть строкой")
}

func (x VMStringMap) BinaryType() VMBinaryType {
	return VMSTRINGMAP
}

func (x VMStringMap) Hash() VMString {
	b, err := x.MarshalBinary()
	if err != nil {
		panic(err)
	}
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, HashBytes(b))
	return VMString(hex.EncodeToString(h))
}

func (x VMStringMap) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		// новые добавляются, а существующие обновляются
		switch yy := y.(type) {
		case VMStringMap:
			rv := make(VMStringMap)
			for k, v := range x {
				if _, ok := yy[k]; !ok {
					rv[k] = v
				}
			}
			for k, v := range yy {
				rv[k] = v
			}
			return rv, nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:
		switch yy := y.(type) {
		case VMStringMap:
			rv := make(VMStringMap)
			for k, v := range x {
				if _, ok := yy[k]; !ok {
					rv[k] = v
				}
			}
			return rv, nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		switch yy := y.(type) {
		case VMStringMap:
			rv := make(VMStringMap)
			for k, v := range x {
				if _, ok := yy[k]; !ok {
					rv[k] = v
				}
			}
			for k := range yy {
				// if _, ok := rv[k]; ok {
				delete(rv, k)
				// }
			}
			return rv, nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		switch yy := y.(type) {
		case VMStringMap:
			if len(x) != len(yy) {
				return VMBool(false), nil
			}
			for k, v := range x {
				if yv, ok := yy[k]; ok {
					if EqualVMValues(v, yv) {
						return VMBool(false), nil
					}
				} else {
					return VMBool(false), nil
				}
			}
			return VMBool(true), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case VMStringMap:
			if len(x) != len(yy) {
				return VMBool(true), nil
			}
			for k, v := range x {
				if yv, ok := yy[k]; ok {
					if EqualVMValues(v, yv) {
						return VMBool(true), nil
					}
				} else {
					return VMBool(true), nil
				}
			}
			return VMBool(false), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case GTR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case GEQ:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LSS:
		return VMNil, errors.New("Операция между значениями невозможна")
	case LEQ:
		return VMNil, errors.New("Операция между значениями невозможна")
	case OR:
		// только добавляем только те, которых еще нет, в отличие от ADD, где существующие обновляются
		switch yy := y.(type) {
		case VMStringMap:
			rv := make(VMStringMap)
			for k, v := range x {
				rv[k] = v
			}
			for k, v := range yy {
				if _, ok := rv[k]; !ok {
					rv[k] = v
				}
			}
			return rv, nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case LOR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case AND:
		// оставляем только те элементы, которые есть в обоих структурах
		switch yy := y.(type) {
		case VMStringMap:
			rv := make(VMStringMap)
			for k, v := range x {
				if _, ok := yy[k]; ok {
					rv[k] = v
				}
			}
			return rv, nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case LAND:
		return VMNil, errors.New("Операция между значениями невозможна")
	case POW:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SHR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SHL:
		return VMNil, errors.New("Операция между значениями невозможна")
	}
	return VMNil, errors.New("Неизвестная операция")
}

func (x VMStringMap) ConvertToType(nt reflect.Type) (VMValuer, error) {

	// fmt.Println(nt)

	switch nt {
	case ReflectVMString:
		// сериализуем в json
		b, err := json.Marshal(x)
		if err != nil {
			return VMNil, err
		}
		return VMString(string(b)), nil
	// case ReflectVMInt:
	// case ReflectVMTime:
	// case ReflectVMBool:
	// case ReflectVMDecimal:
	// case ReflectVMSlice:
	case ReflectVMStringMap:
		return x, nil
	}

	if nt.Kind() == reflect.Struct {
		rv := reflect.ValueOf(x)
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
					if fv.Kind() != setv.Kind() {
						if setv.Type().ConvertibleTo(fv.Type()) {
							setv = setv.Convert(fv.Type())
						} else {
							return nil, fmt.Errorf("Поле структуры имеет другой тип")
						}
					}
					fv.Set(setv)
				}
			}
		}
		if vv, ok := rs.Interface().(VMValuer); ok {
			if vobj, ok := vv.(VMMetaObject); ok {
				vobj.VMInit(vobj)
				vobj.VMRegister()
				return vobj, nil
			} else {
				return nil, errors.New("Невозможно использовать данный тип структуры")
				//return vv, nil
			}
		} else {
			return nil, errors.New("Невозможно использовать данный тип в интерпретаторе")
		}

	}

	return VMNil, errors.New("Приведение к типу невозможно")
}

func (x VMStringMap) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uint64(len(x)))
	for i := range x {
		if v, ok := x[i].(VMBinaryTyper); ok {
			bb, err := v.MarshalBinary()
			if err != nil {
				return nil, err
			}
			binary.Write(&buf, binary.LittleEndian, uint64(len(i)))
			buf.Write([]byte(i))
			buf.WriteByte(byte(v.BinaryType()))
			binary.Write(&buf, binary.LittleEndian, uint64(len(bb)))
			buf.Write(bb)
		} else {
			return nil, errors.New("Значение не может быть преобразовано в бинарный формат")
		}
	}
	return buf.Bytes(), nil
}

func (x *VMStringMap) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	var l, li, lv uint64
	if err := binary.Read(buf, binary.LittleEndian, &l); err != nil {
		return err
	}
	rv := make(VMStringMap, int(l))
	for i := 0; i < int(l); i++ {
		if err := binary.Read(buf, binary.LittleEndian, &li); err != nil {
			return err
		}
		bi := buf.Next(int(li))
		if err := binary.Read(buf, binary.LittleEndian, &lv); err != nil {
			return err
		}
		if tt, err := buf.ReadByte(); err != nil {
			return err
		} else {
			bb := buf.Next(int(lv))
			vv, err := VMBinaryType(tt).ParseBinary(bb)
			if err != nil {
				return err
			}
			rv[string(bi)] = vv
		}
	}
	*x = rv
	return nil
}

func (x VMStringMap) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMStringMap) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

// TODO:

// func (x VMStringMap) MarshalText() ([]byte, error) {
// 	var buf bytes.Buffer
// 	buf.WriteString(time.Duration(x).String())
// 	return buf.Bytes(), nil
// }

// func (x *VMStringMap) UnmarshalText(data []byte) error {
// 	d, err := time.ParseDuration(string(data))
// 	if err != nil {
// 		return err
// 	}
// 	*x = VMTimeDuration(d)
// 	return nil
// }

// func (x VMStringMap) MarshalJSON() ([]byte, error) {
// 	b, err := x.MarshalText()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte("\"" + string(b) + "\""), nil
// }

// func (x *VMStringMap) UnmarshalJSON(data []byte) error {
// 	if string(data) == "null" {
// 		return nil
// 	}
// 	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
// 		data = data[1 : len(data)-1]
// 	}
// 	return x.UnmarshalText(data)
// }
