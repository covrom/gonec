package vm

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// универсальный вариантный тип данных

const (
	NONE = iota
	NULL
	UNDEF
	NUM
	DATE
	STR
	BOOL
)

type variant struct {
	typ  int //может быть одна из констант
	num  big.Float
	date time.Time
	str  string
	boo  bool
}

func (v variant) String() string {
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

func (v *variant) SetString(s string) {
	v.typ = STR
	v.str = s
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *variant) SetDate(d time.Time) {
	v.typ = DATE
	v.str = ""
	v.date = d
	v.num = big.Float{}
	v.boo = false
}

func (v *variant) SetNum(n big.Float) {
	v.typ = NUM
	v.str = ""
	v.date = time.Time{}
	v.num = n
	v.boo = false
}

func (v *variant) SetUNDEF() {
	v.typ = UNDEF
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *variant) SetNULL() {
	v.typ = NULL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *variant) SetTrue() {
	v.typ = BOOL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = true
}

func (v *variant) SetFalse() {
	v.typ = BOOL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = false
}

func (v *variant) SetBool(b bool) {
	v.typ = BOOL
	v.str = ""
	v.date = time.Time{}
	v.num = big.Float{}
	v.boo = b
}

func (v variant) GetValue() (typ int, val interface{}) {
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

func (v variant) IsNULL() bool {
	return v.typ == NULL
}

func (v variant) IsUNDEF() bool {
	return v.typ == UNDEF
}

func (v variant) IsDate() bool {
	return v.typ == DATE
}

func (v variant) IsNum() bool {
	return v.typ == NUM
}

func (v variant) IsString() bool {
	return v.typ == STR
}

func (v variant) IsBool() bool {
	return v.typ == BOOL
}

func (v variant) IsZero() bool {
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

func (v variant) GetTypeString() string {
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

func (v variant) Cmp(to variant) (int, error) {
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
		if (!v.boo)&&to.boo{return -1,nil}
		if v.boo==to.boo{return 0,nil}
		if v.boo&&(!to.boo){return 1,nil}
	default:
		return 0, errors.New("Невозможно сравнить значения")
	}
}
