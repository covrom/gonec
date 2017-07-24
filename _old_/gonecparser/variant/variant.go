package variant

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// универсальный вариантный тип данных
// TODO: переписать на интерфейсы

const (
	NONE = iota
	NULL
	UNDEF
	NUM
	DATE
	STR
	BOOL
)

type Variant struct {
	typ  int //может быть одна из констант
	num  big.Float
	date time.Time
	str  string
	boo  bool
}

func (v Variant) isVariant() bool { return true }

func NewVariant() *Variant{
	return &Variant{typ:NONE}
}

func (v Variant) String() string {
	switch v.typ {
	case NULL:
		return "NULL"
	case UNDEF:
		return "Неопределено"
	case NUM:
		return v.num.String()
	case DATE:
		return v.date.Format(time.RFC3339)
	case STR:
		return v.str
	case BOOL:
		if v.boo {
			return "Истина"
		}
		return "Ложь"
	default:
		return "NONE"
	}
}

func (v *Variant) SetFrom(s *Variant) {
	v.typ = s.typ
	v.str = s.str
	v.date = s.date
	v.num = s.num
	v.boo = s.boo
}

func (v *Variant) SetString(s string) {
	v.typ = STR
	v.str = s
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *Variant) SetDate(d time.Time) {
	v.typ = DATE
	v.str = ""
	v.date = d
	v.num = big.Float{}
	v.boo = false
}

func (v *Variant) SetNum(n big.Float) {
	v.typ = NUM
	v.str = ""
	v.date = time.Time{}
	v.num = n
	v.boo = false
}

func (v *Variant) SetUNDEF() {
	v.typ = UNDEF
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *Variant) SetNULL() {
	v.typ = NULL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *Variant) SetTrue() {
	v.typ = BOOL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = true
}

func (v *Variant) SetFalse() {
	v.typ = BOOL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *Variant) SetBool(b bool) {
	v.typ = BOOL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = b
}

func (v Variant) GetValue() (typ int, val interface{}) {
	typ = v.typ
	switch typ {
	case DATE:
		val = v.date
	case STR:
		val = v.str
	case NUM:
		val = v.num
	case BOOL:
		val = v.boo
	default:
		val = nil
	}
	return
}

func (v Variant) IsNone() bool {
	return v.typ == NONE
}

func (v Variant) IsNULL() bool {
	return v.typ == NULL
}

func (v Variant) IsUNDEF() bool {
	return v.typ == UNDEF
}

func (v Variant) IsDate() bool {
	return v.typ == DATE
}

func (v Variant) IsNum() bool {
	return v.typ == NUM
}

func (v Variant) IsString() bool {
	return v.typ == STR
}

func (v Variant) IsBool() bool {
	return v.typ == BOOL
}

func (v Variant) IsZero() bool {
	switch v.typ {
	case BOOL:
		return v.boo == false
	case NUM:
		return v.num.Cmp(big.NewFloat(0)) == 0
	case DATE:
		return v.date.IsZero()
	case STR:
		return v.str == ""
	default:
		return true
	}
}

func (v Variant) GetTypeString() string {
	switch v.typ {
	case NULL:
		return "NULL"
	case UNDEF:
		return "Неопределено"
	case NUM:
		return "Число"
	case DATE:
		return "Дата"
	case STR:
		return "Строка"
	case BOOL:
		return "Булево"
	default:
		return "NONE"
	}
}

func (v Variant) Cmp(to Variant) (int, error) {
	if v.typ != to.typ {
		return 0, fmt.Errorf("Несравниваемые типы %v и %v", v.GetTypeString(), to.GetTypeString())
	}
	switch v.typ {
	case NULL:
		return 0, nil
	case UNDEF:
		return 0, nil
	case NUM:
		return v.num.Cmp(&to.num), nil
	case DATE:
		if v.date.Equal(to.date) {
			return 0, nil
		}
		if v.date.Before(to.date) {
			return -1, nil
		}
		if v.date.After(to.date) {
			return 1, nil
		}
		return 0, errors.New("Невозможно сравнить даты")
	case STR:
		return strings.Compare(v.str, to.str), nil
	case BOOL:
		if (!v.boo) && to.boo {
			return -1, nil
		}
		if v.boo == to.boo {
			return 0, nil
		}
		if v.boo && (!to.boo) {
			return 1, nil
		}
	}
	return 0, errors.New("Невозможно сравнить значения")
}
