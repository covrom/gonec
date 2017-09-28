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
	var binbuf bytes.Buffer
	binary.Write(&binbuf, binary.BigEndian, x)
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, HashBytes(binbuf.Bytes()))
	return VMString(hex.EncodeToString(h))
}

func (x VMSlice) SortDefault() {
	sort.Sort(VMSliceUpSort(x))
}

func (x VMSlice) Сортировать() {
	x.SortDefault()
}

func (x VMSlice) СортироватьУбыв() {
	sort.Sort(VMSliceDownSort(x))
}

func (x VMSlice) Обратить() {
	for left, right := 0, len(x)-1; left < right; left, right = left+1, right-1 {
		x[left], x[right] = x[right], x[left]
	}
}

func (x VMSlice) Скопировать() VMSlice {
	rv := make(VMSlice, len(x))
	copy(rv, x)
	return rv
}

func (x VMSlice) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		// добавляем второй слайс в конец первого
		switch yy := y.(type) {
		case VMSlice:
			return append(x, yy...), nil
		}
		return append(x, y), nil
		// return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:
		// удаляем из первого слайса любые элементы второго слайса, встречающиеся в первом
		switch yy := y.(type) {
		case VMSlice:
			// проходим слайс и переставляем ненайденные в вычитаемом слайсе элементы
			il := 0
			for i := range x {
				fnd := false
				for j := range yy {
					if xop, ok := x[i].(VMOperationer); ok {
						if yop, ok := yy[j].(VMOperationer); ok {
							cmp, err := xop.EvalBinOp(EQL, yop)
							if err == nil {
								if rcmp, ok := cmp.(VMBool); ok {
									if bool(rcmp) {
										fnd = true
										break
									}
								}
							}
						}
					}
				}
				if !fnd {
					if il < i {
						x[il] = x[i]
						il++
					}
				}
			}
			return x[:il], nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		switch yy := y.(type) {
		case VMSlice:
			eq := true
			for i := range x {
				for j := range yy {
					if xop, ok := x[i].(VMOperationer); ok {
						if yop, ok := yy[j].(VMOperationer); ok {
							cmp, err := xop.EvalBinOp(EQL, yop)
							if err == nil {
								if rcmp, ok := cmp.(VMBool); ok {
									if !bool(rcmp) {
										eq = false
										goto eqlret
									}
								} else {
									eq = false
									goto eqlret
								}
							} else {
								eq = false
								goto eqlret
							}
						}
					}
				}
			}
		eqlret:
			return VMBool(eq), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case VMSlice:
			eq := true
			for i := range x {
				for j := range yy {
					if xop, ok := x[i].(VMOperationer); ok {
						if yop, ok := yy[j].(VMOperationer); ok {
							cmp, err := xop.EvalBinOp(EQL, yop)
							if err == nil {
								if rcmp, ok := cmp.(VMBool); ok {
									if !bool(rcmp) {
										eq = false
										goto neqlret
									}
								} else {
									eq = false
									goto neqlret
								}
							} else {
								eq = false
								goto neqlret
							}
						}
					}
				}
			}
		neqlret:
			return VMBool(!eq), nil
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
					if xop, ok := x[i].(VMOperationer); ok {
						if yop, ok := yy[j].(VMOperationer); ok {
							cmp, err := xop.EvalBinOp(EQL, yop)
							if err == nil {
								if rcmp, ok := cmp.(VMBool); ok {
									if bool(rcmp) {
										fnd = true
										break
									}
								}
							}
						}
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
			rv := x[:0]
			for i := range x {
				fnd := false
				for j := range yy {
					if xop, ok := x[i].(VMOperationer); ok {
						if yop, ok := yy[j].(VMOperationer); ok {
							cmp, err := xop.EvalBinOp(EQL, yop)
							if err == nil {
								if rcmp, ok := cmp.(VMBool); ok {
									if bool(rcmp) {
										fnd = true
										break
									}
								}
							}
						}
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

// TODO:

// func (x VMTimeDuration) MarshalText() ([]byte, error) {
// 	var buf bytes.Buffer
// 	buf.WriteString(time.Duration(x).String())
// 	return buf.Bytes(), nil
// }

// func (x *VMTimeDuration) UnmarshalText(data []byte) error {
// 	d, err := time.ParseDuration(string(data))
// 	if err != nil {
// 		return err
// 	}
// 	*x = VMTimeDuration(d)
// 	return nil
// }

// func (x VMTimeDuration) MarshalJSON() ([]byte, error) {
// 	b, err := x.MarshalText()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte("\"" + string(b) + "\""), nil
// }

// func (x *VMTimeDuration) UnmarshalJSON(data []byte) error {
// 	if string(data) == "null" {
// 		return nil
// 	}
// 	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
// 		data = data[1 : len(data)-1]
// 	}
// 	return x.UnmarshalText(data)
// }

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
			return vi.Раньше(vj)
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
			return vi.Позже(vj)
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
