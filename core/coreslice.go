package core

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
	"sort"
	"strings"

	"github.com/covrom/gonec/names"
	"github.com/shopspring/decimal"
)

type VMSlice []VMValuer

var ReflectVMSlice = reflect.TypeOf(make(VMSlice, 0))

func (x VMSlice) vmval() {}

func (x VMSlice) Interface() interface{} {
	return x
}

func (x VMSlice) Slice() VMSlice {
	return x
}

func (x VMSlice) BinaryType() VMBinaryType {
	return VMSLICE
}

func (x VMSlice) Args() []interface{} {
	ai := make([]interface{}, len(x))
	for i := range x {
		ai[i] = x[i]
	}
	return ai
}

func (x *VMSlice) Append(a ...VMValuer) {
	*x = append(*x, a...)
}

func (x VMSlice) Length() VMInt {
	return VMInt(len(x))
}

func (x VMSlice) IndexVal(i VMValuer) VMValuer {
	if ii, ok := i.(VMInt); ok {
		return x[int(ii)]
	}
	panic("Индекс должен быть целым числом")
}

func (x VMSlice) Hash() VMString {
	b, err := x.MarshalBinary()
	if err != nil {
		panic(err)
	}
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, HashBytes(b))
	return VMString(hex.EncodeToString(h))
}

func (x VMSlice) SortDefault() {
	sort.Sort(VMSliceUpSort(x))
}

func (x VMSlice) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	case "сортировать":
		return VMFuncMustParams(0, x.Сортировать), true
	case "сортироватьубыв":
		return VMFuncMustParams(0, x.СортироватьУбыв), true
	case "обратить":
		return VMFuncMustParams(0, x.Обратить), true
	case "скопировать":
		return VMFuncMustParams(0, x.Скопировать), true
	}

	return nil, false
}

func (x VMSlice) Сортировать(args VMSlice, rets *VMSlice) error {
	x.SortDefault()
	return nil
}

func (x VMSlice) СортироватьУбыв(args VMSlice, rets *VMSlice) error {
	sort.Sort(VMSliceDownSort(x))
	return nil
}

func (x VMSlice) Обратить(args VMSlice, rets *VMSlice) error {
	for left, right := 0, len(x)-1; left < right; left, right = left+1, right-1 {
		x[left], x[right] = x[right], x[left]
	}
	return nil
}

func (x VMSlice) CopyRecursive() VMSlice {
	rv := make(VMSlice, len(x))
	for i, v := range x {
		switch vv := v.(type) {
		case VMSlice:
			rv[i]=vv.CopyRecursive()
		case VMStringMap:
			rv[i]=vv.CopyRecursive()
		default:
			rv[i]=v
		}
	}
	return rv
}

func (x VMSlice) Скопировать(args VMSlice, rets *VMSlice) error { //VMSlice {
	rv := make(VMSlice, len(x))
	copy(rv, x)
	for i,v:=range rv{
		switch vv := v.(type) {
		case VMSlice:
			rv[i]=vv.CopyRecursive()
		case VMStringMap:
			rv[i]=vv.CopyRecursive()
		}		
	}
	rets.Append(rv)
	return nil
}

func (x VMSlice) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMSlice:
			// добавляем второй слайс в конец первого
			return append(x, yy...), nil
		case VMValuer:
			return append(x, yy), nil
		}
		return append(x, y), nil
		// return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:

		// TODO: оптимизировать для сортированных слайсов VMSliceUpSort, VMSliceDownSort

		// удаляем из первого слайса любые элементы второго слайса, встречающиеся в первом
		switch yy := y.(type) {
		case VMSlice:
			// проходим слайс и переставляем ненайденные в вычитаемом слайсе элементы
			rv := make(VMSlice, len(x))
			il := 0
			for i := range x {
				fnd := false
				for j := range yy {
					if EqualVMValues(x[i], yy[j]) {
						fnd = true
						break
					}
				}
				if !fnd {
					rv[il] = x[i]
					il++
				}
			}
			return rv[:il], nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		// оставляем только элементы, которые есть в первом и нет во втором и есть во втором но нет в первом
		// эквивалентно (С1 | С2) - (С1 & С2), или (С1-С2)|(С2-С1), или С2-(С1-С2), внешнее соединение

		switch yy := y.(type) {
		case VMSlice:
			rvx := make(VMSlice, len(x))
			rvy := make(VMSlice, len(yy))
			// С1-С2
			il := 0
			for i := range x {
				fnd := false
				for j := range yy {
					if EqualVMValues(x[i], yy[j]) {
						fnd = true
						break
					}
				}
				if !fnd {
					// оставляем
					rvx[il] = x[i]
					il++
				}
			}

			rvx = rvx[:il]

			// С2-(С1-C2)
			il = 0
			for j := range yy {
				fnd := false
				for i := range x {
					if EqualVMValues(x[i], yy[j]) {
						fnd = true
						break
					}
				}
				if !fnd {
					// оставляем
					rvy[il] = yy[j]
					il++
				}
			}

			rvy = rvy[:il]

			return append(rvx, rvy...), nil
		}

		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		// равенство по глубокому равенству элементов
		switch yy := y.(type) {
		case VMSlice:
			if len(x) != len(yy) {
				return VMBool(false), nil
			}
			for i := range x {
				for j := range yy {
					if !EqualVMValues(x[i], yy[j]) {
						return VMBool(false), nil
					}
				}
			}
			return VMBool(true), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case VMSlice:
			if len(x) != len(yy) {
				return VMBool(true), nil
			}
			for i := range x {
				for j := range yy {
					if !EqualVMValues(x[i], yy[j]) {
						return VMBool(true), nil
					}
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
		// добавляем в конец первого слайса только те элементы второго слайса, которые не встречаются в первом
		switch yy := y.(type) {
		case VMSlice:
			rv := x[:]
			for j := range yy {
				fnd := false
				for i := range x {
					if EqualVMValues(x[i], yy[j]) {
						fnd = true
						break
					}
				}
				if !fnd {
					rv = append(rv, yy[j])
				}
			}
			return rv, nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case LOR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case AND:
		// оставляем только те элементы, которые есть в обоих слайсах
		switch yy := y.(type) {
		case VMSlice:
			rv := make(VMSlice, 0, len(x))
			for i := range x {
				fnd := false
				for j := range yy {
					if EqualVMValues(x[i], yy[j]) {
						fnd = true
						break
					}
				}
				if fnd {
					rv = append(rv, x[i])
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

func (x VMSlice) ConvertToType(nt reflect.Type) (VMValuer, error) {
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
	case ReflectVMSlice:
		return x, nil
		// case ReflectVMStringMap:
	}

	return VMNil, errors.New("Приведение к типу невозможно")
}

func (x VMSlice) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uint64(len(x)))
	for i := range x {
		if v, ok := x[i].(VMBinaryTyper); ok {
			bb, err := v.MarshalBinary()
			if err != nil {
				return nil, err
			}
			buf.WriteByte(byte(v.BinaryType()))
			binary.Write(&buf, binary.LittleEndian, uint64(len(bb)))
			buf.Write(bb)
		} else {
			return nil, errors.New("Значение не может быть преобразовано в бинарный формат")
		}
	}
	return buf.Bytes(), nil
}

func (x *VMSlice) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	var l, lv uint64
	if err := binary.Read(buf, binary.LittleEndian, &l); err != nil {
		return err
	}
	var rv VMSlice
	if x == nil || len(*x) < int(l) {
		rv = make(VMSlice, int(l))
	} else {
		rv = (*x)[:int(l)]
	}

	for i := 0; i < int(l); i++ {
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
			rv[i] = vv
		}
	}
	*x = rv
	return nil
}

func (x VMSlice) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMSlice) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMSlice) String() string {
	b, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (x VMSlice) MarshalText() ([]byte, error) {
	b, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (x *VMSlice) UnmarshalText(data []byte) error {
	sl, err := VMSliceFromJson(string(data))
	if err != nil {
		return err
	}
	*x = sl
	return nil
}

func (x VMSlice) MarshalJSON() ([]byte, error) {
	var err error
	rm := make([]json.RawMessage, len(x))
	for i, v := range x {
		rm[i], err = json.Marshal(v)
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(rm)
}

func (x *VMSlice) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	sl, err := VMSliceFromJson(string(data))
	if err != nil {
		return err
	}
	*x = sl
	return nil
}

// VMSliceUpSort - обертка для сортировки слайса по возрастанию
type VMSliceUpSort VMSlice

func (x VMSliceUpSort) Len() int      { return len(x) }
func (x VMSliceUpSort) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x VMSliceUpSort) Less(i, j int) bool {

	// числа
	if vi, ok := x[i].(VMInt); ok {
		if vj, ok := x[j].(VMInt); ok {
			return vi.Int() < vj.Int()
		}
		if vj, ok := x[j].(VMDecimal); ok {
			vii := decimal.New(int64(vi), 0)
			return vii.LessThan(decimal.Decimal(vj))
		}
	}

	if vi, ok := x[i].(VMDecimal); ok {
		if vj, ok := x[j].(VMInt); ok {
			vjj := decimal.New(int64(vj), 0)
			return decimal.Decimal(vi).LessThan(vjj)
		}
		if vj, ok := x[j].(VMDecimal); ok {
			return decimal.Decimal(vi).LessThan(decimal.Decimal(vj))
		}
	}

	// строки
	if vi, ok := x[i].(VMString); ok {
		if vj, ok := x[j].(VMString); ok {
			return strings.Compare(vi.String(), vj.String()) == -1
		}
		if vj, ok := x[j].(VMInt); ok {
			return strings.Compare(vi.String(), vj.String()) == -1
		}
		if vj, ok := x[j].(VMDecimal); ok {
			return strings.Compare(vi.String(), vj.String()) == -1
		}
	}

	// булево

	if vi, ok := x[i].(VMBool); ok {
		if vj, ok := x[j].(VMBool); ok {
			return !vi.Bool() && vj.Bool()
		}
	}

	// дата

	if vi, ok := x[i].(VMTime); ok {
		if vj, ok := x[j].(VMTime); ok {
			return vi.Before(vj)
		}
	}

	// длительность
	if vi, ok := x[i].(VMTimeDuration); ok {
		if vj, ok := x[j].(VMTimeDuration); ok {
			return int64(vi) < int64(vj)
		}
	}

	// прочее

	if vi, ok := x[i].(VMOperationer); ok {
		if vj, ok := x[j].(VMOperationer); ok {
			b, err := vi.EvalBinOp(LSS, vj)
			if err == nil {
				return b.(VMBool).Bool()
			}
		}
	}

	return false
}

// VMSliceDownSort по убыванию
type VMSliceDownSort VMSlice

func (x VMSliceDownSort) Len() int      { return len(x) }
func (x VMSliceDownSort) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x VMSliceDownSort) Less(i, j int) bool {

	// числа
	if vi, ok := x[i].(VMInt); ok {
		if vj, ok := x[j].(VMInt); ok {
			return vi.Int() > vj.Int()
		}
		if vj, ok := x[j].(VMDecimal); ok {
			vii := decimal.New(int64(vi), 0)
			return vii.GreaterThan(decimal.Decimal(vj))
		}
	}

	if vi, ok := x[i].(VMDecimal); ok {
		if vj, ok := x[j].(VMInt); ok {
			vjj := decimal.New(int64(vj), 0)
			return decimal.Decimal(vi).GreaterThan(vjj)
		}
		if vj, ok := x[j].(VMDecimal); ok {
			return decimal.Decimal(vi).GreaterThan(decimal.Decimal(vj))
		}
	}

	// строки
	if vi, ok := x[i].(VMString); ok {
		if vj, ok := x[j].(VMString); ok {
			return strings.Compare(vi.String(), vj.String()) == 1
		}
		if vj, ok := x[j].(VMInt); ok {
			return strings.Compare(vi.String(), vj.String()) == 1
		}
		if vj, ok := x[j].(VMDecimal); ok {
			return strings.Compare(vi.String(), vj.String()) == 1
		}
	}

	// булево

	if vi, ok := x[i].(VMBool); ok {
		if vj, ok := x[j].(VMBool); ok {
			return !(!vi.Bool() && vj.Bool())
		}
	}

	// дата

	if vi, ok := x[i].(VMTime); ok {
		if vj, ok := x[j].(VMTime); ok {
			return vi.After(vj)
		}
	}

	// длительность
	if vi, ok := x[i].(VMTimeDuration); ok {
		if vj, ok := x[j].(VMTimeDuration); ok {
			return int64(vi) > int64(vj)
		}
	}

	// прочее

	if vi, ok := x[i].(VMOperationer); ok {
		if vj, ok := x[j].(VMOperationer); ok {
			b, err := vi.EvalBinOp(GTR, vj)
			if err == nil {
				return b.(VMBool).Bool()
			}
		}
	}

	return false
}

// NewVMSliceFromStrings создает слайс вирт. машины []VMString из слайса строк []string на языке Го
func NewVMSliceFromStrings(ss []string) (rv VMSlice) {
	for i := range ss {
		rv = append(rv, VMString(ss[i]))
	}
	return
}
