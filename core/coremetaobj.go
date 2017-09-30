package core

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/covrom/gonec/names"
)

// VMMetaObj корневая структура для системных функциональных структур Го, доступных из языка Гонец
// поля и методы должны отличаться друг от друга без учета регистра
// например, Set и set - в вирт. машине будут считаться одинаковыми, будет использоваться последнее по индексу
type VMMetaObj struct {
	vmMetaCacheM map[int]VMFunc
	vmMetaCacheF map[int]VMValuer

	vmOriginal VMMetaObject
}

func (v *VMMetaObj) vmval() {}

func (v *VMMetaObj) VMInit(m VMMetaObject) {
	// исходная структура
	v.vmOriginal = m
}

func (v *VMMetaObj) Interface() interface{} {
	// возвращает ссылку на структуру, от которой был вызван метод VMInit
	//rv:=*(*VMMetaObject)(v.vmOriginal)
	return v.vmOriginal
}

func (v *VMMetaObj) String() string {
	b, err := json.Marshal(v.vmOriginal)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (x *VMMetaObj) Hash() VMString {
	b := []byte(x.String())
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, HashBytes(b))
	return VMString(hex.EncodeToString(h))
}

func (v *VMMetaObj) VMRegisterMethod(name string, m VMMethod) {
	if v.vmMetaCacheM == nil {
		v.vmMetaCacheM = make(map[int]VMFunc)
	}
	namtyp := names.UniqueNames.Set(name)
	v.vmMetaCacheM[namtyp] = func(meth VMMethod) VMFunc {
		return VMFunc(meth)
	}(m)
}

func (v *VMMetaObj) VMRegisterField(name string, m VMValuer) {
	if v.vmMetaCacheF == nil {
		v.vmMetaCacheF = make(map[int]VMValuer)
	}
	switch m.(type) {
	case *VMInt, *VMString, *VMBool,
		*VMChan, *VMDecimal, *VMStringMap,
		*VMSlice, *VMTime, *VMTimeDuration:

		namtyp := names.UniqueNames.Set(name)
		v.vmMetaCacheF[namtyp] = m
	default:
		panic("Поле не может быть зарегистрировано")
	}
}

func (v *VMMetaObj) VMIsField(name int) bool {
	_, ok := v.vmMetaCacheF[name]
	return ok
}

func (v *VMMetaObj) VMGetField(name int) VMValuer {
	if r, ok := v.vmMetaCacheF[name]; ok {
		switch rv := r.(type) {
		case *VMInt:
			return *rv
		case *VMString:
			return *rv
		case *VMBool:
			return *rv
		case *VMChan:
			return *rv
		case *VMDecimal:
			return *rv
		case *VMStringMap:
			return *rv
		case *VMSlice:
			return *rv
		case *VMTime:
			return *rv
		case *VMTimeDuration:
			return *rv
		}
	}
	panic("Невозможно получить значение поля")
}

func (v *VMMetaObj) VMSetField(name int, val VMValuer) {

	if r, ok := v.vmMetaCacheF[name]; ok {
		switch rv := r.(type) {
		case *VMInt:
			*rv = VMInt(val.(VMNumberer).Int())
			return
		case *VMString:
			*rv = VMString(val.(VMStringer).String())
			return
		case *VMBool:
			*rv = VMBool(val.(VMBooler).Bool())
			return
		case *VMChan:
			*rv = val.(VMChan)
			return
		case *VMDecimal:
			*rv = val.(VMNumberer).Decimal()
			return
		case *VMStringMap:
			*rv = val.(VMStringMaper).StringMap()
			return
		case *VMSlice:
			*rv = val.(VMSlicer).Slice()
			return
		case *VMTime:
			*rv = val.(VMDateTimer).Time()
			return
		case *VMTimeDuration:
			*rv = val.(VMDurationer).Duration()
			return
		}
	}

	panic("Невозможно установить значение поля")
}

// VMGetMethod генерит функцию,
// которая возвращает либо одно значение и ошибку, либо массив значений интерпретатора VMSlice
func (v *VMMetaObj) VMGetMethod(name int) (VMFunc, bool) {

	// fmt.Println(name)

	rv, ok := v.vmMetaCacheM[name]
	return rv, ok
}

func (v *VMMetaObj) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		return VMNil, errors.New("Операция между значениями невозможна")
	case SUB:
		return VMNil, errors.New("Операция между значениями невозможна")
	case MUL:
		return VMNil, errors.New("Операция между значениями невозможна")
	case QUO:
		return VMNil, errors.New("Операция между значениями невозможна")
	case REM:
		return VMNil, errors.New("Операция между значениями невозможна")
	case EQL:
		switch yy := y.(type) {
		case *VMMetaObj:
			eq := v.Hash() == yy.Hash()
			return VMBool(eq), nil
		}
		return VMNil, errors.New("Операция между значениями невозможна")
	case NEQ:
		switch yy := y.(type) {
		case *VMMetaObj:
			eq := v.Hash() != yy.Hash()
			return VMBool(eq), nil
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
		return VMNil, errors.New("Операция между значениями невозможна")
	case LOR:
		return VMNil, errors.New("Операция между значениями невозможна")
	case AND:
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

func (v *VMMetaObj) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMString:
		// сериализуем в json
		b, err := json.Marshal(v.vmOriginal)
		if err != nil {
			return VMNil, err
		}
		return VMString(string(b)), nil
		// case ReflectVMInt:
		// case ReflectVMTime:
		// case ReflectVMBool:
		// case ReflectVMDecimal:
		// case ReflectVMSlice:
		// case ReflectVMStringMap: // получится только через Структура(Строка(объект))
	}

	return VMNil, errors.New("Приведение к типу невозможно")
}

// TODO: маршаллинг исходной структуры, как у VMTime!!!
