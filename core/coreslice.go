package core

import (
	"sort"
	"strings"

	"github.com/shopspring/decimal"
)

// TODO: для слайса реализовать поведение массива при изменении размеров (т.е. изменение размера передается в переменную)

type VMSlice []VMValuer

func (x VMSlice) vmval() {}

func (x VMSlice) Interface() interface{} {
	return x
}

func (x VMSlice) Slice() VMSlice {
	return x
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

func (x VMSlice) SortDefault() {
	sort.Sort(VMSliceDefaultSort(x))
}

func (x VMSlice) Сортировать() {
	x.SortDefault()
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

// TODO: маршаллинг и String!!!

type VMSliceDefaultSort VMSlice

func (x VMSliceDefaultSort) Len() int { return len(x) }

func (x VMSliceDefaultSort) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x VMSliceDefaultSort) Less(i, j int) bool {

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
