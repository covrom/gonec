package vm

import (
	"math/big"
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
