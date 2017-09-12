package core

import (
	"reflect"
	"time"
	"unsafe"

	"github.com/shopspring/decimal"

	"github.com/covrom/gonec/ast"
)

// VMMetaObj корневая структура для системных функциональных структур Го, доступных из языка Гонец
// поля и методы должны отличаться друг от друга без учета регистра
// например, Set и set - в вирт. машине будут считаться одинаковыми, будет использоваться последнее по индексу
type VMMetaObj struct {
	vmMetaCacheM map[int]int
	vmMetaCacheF map[int][]int
	vmReflectPtr unsafe.Pointer
}

func (v *VMMetaObj) vmval() {}

func (v *VMMetaObj) Interface() interface{} {
	// возвращает ссылку на структуру, от которой был вызван метод VMCacheMembers
	return *(*VMMetaStructer)(v.vmReflectPtr)
}

// VMCacheMembers кэширует все поля и методы ссылки на объединенную структуру в VMMetaStructer
// Вызывать эту функцию надо так:
// v:=&struct{VMMetaObj, a int}{}; v.VMCacheMembers(v)
func (v *VMMetaObj) VMCacheMembers(vv VMMetaStructer) {
	v.vmMetaCacheM = make(map[int]int)
	v.vmMetaCacheF = make(map[int][]int)
	v.vmReflectPtr = unsafe.Pointer(&vv)

	typ := reflect.TypeOf(vv)

	// пишем в кэш индексы полей и методов для структур

	// методы
	nm := typ.NumMethod()
	for i := 0; i < nm; i++ {
		meth := typ.Method(i)
		// только экспортируемые
		if meth.PkgPath == "" {
			namtyp := ast.UniqueNames.Set(meth.Name)
			v.vmMetaCacheM[namtyp] = meth.Index
		}
	}

	// поля
	nm = typ.NumField()
	for i := 0; i < nm; i++ {
		field := typ.Field(i)
		// только экспортируемые неанонимные поля
		if field.PkgPath == "" && !field.Anonymous {
			namtyp := ast.UniqueNames.Set(field.Name)
			v.vmMetaCacheF[namtyp] = field.Index
		}
	}
}

func (v *VMMetaObj) VMIsField(name int) bool {
	_, ok := v.vmMetaCacheF[name]
	return ok
}

func (v *VMMetaObj) VMGetField(name int) VMInterfacer {
	vv := reflect.ValueOf(v.Interface()).FieldByIndex(v.vmMetaCacheF[name])
	if x, ok := vv.Interface().(VMInterfacer); ok {
		return x
	}
	return ReflectToVMValue(vv)
}

func (v *VMMetaObj) VMSetField(name int, val VMInterfacer) {
	vv := reflect.ValueOf(v.Interface()).FieldByIndex(v.vmMetaCacheF[name])
	if !vv.CanSet() {
		panic("Невозможно установить значение поля только для чтения")
	}
	if _, ok := vv.Interface().(VMInterfacer); ok {
		vv.Set(reflect.ValueOf(val))
		return
	}
	if reflect.TypeOf(val).AssignableTo(vv.Type()) {
		vv.Set(reflect.ValueOf(val.Interface()))
		return
	}
	if reflect.TypeOf(val).ConvertibleTo(vv.Type()) {
		vv.Set(reflect.ValueOf(val).Convert(vv.Type()))
		return
	}
	panic("Невозможно установить значение поля")
}

// VMGetMethod генерит функцию,
// которая возвращает либо одно значение и ошибку, либо массив значений интерпретатора VMSlice
func (v *VMMetaObj) VMGetMethod(name int) VMMeth {
	vv := reflect.ValueOf(v.Interface()).Method(v.vmMetaCacheM[name])
	return VMMeth(
		func(args VMSlicer) (VMInterfacer, error) {
			a := args.Slice()
			x := make([]reflect.Value, len(a))
			for i := range a {
				x[i] = reflect.ValueOf(a[i])
			}
			r := vv.Call(x)
			switch len(r) {
			case 0:
				return nil, nil
			case 1:
				return ReflectToVMValue(r[0]), nil
			case 2:
				if e, ok := r[1].Interface().(error); ok {
					return ReflectToVMValue(r[0]), e
				}
				fallthrough
			default:
				rvm := make(VMSlice, len(r))
				for i := range r {
					rvm[i] = ReflectToVMValue(r[i])
				}
				return rvm, nil
			}
		})
}

// ReflectToVMValue преобразовывает значение Го в наиболее подходящий тип значения для вирт. машшины
func ReflectToVMValue(rv reflect.Value) VMInterfacer {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return VMInt(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return VMInt(rv.Uint())
	case reflect.String:
		return VMString(rv.String())
	case reflect.Bool:
		return VMBool(rv.Bool())
	case reflect.Float32, reflect.Float64:
		return VMDecimal(decimal.NewFromFloat(rv.Float()))
	case reflect.Chan:
		// проверяем, может это VMChaner
		if x, ok := rv.Interface().(VMChaner); ok {
			return x
		}
	case reflect.Array, reflect.Slice:
		// проверяем, может это VMSlicer
		if x, ok := rv.Interface().(VMSlicer); ok {
			return x
		}
	case reflect.Map:
		// проверяем, может это VMStringMaper
		if x, ok := rv.Interface().(VMStringMaper); ok {
			return x
		}
	case reflect.Func:
		// проверяем, может это VMFuncer
		if x, ok := rv.Interface().(VMFuncer); ok {
			return x
		}
	case reflect.Struct:
		switch v := rv.Interface().(type) {
		case decimal.Decimal:
			return VMDecimal(v)
		case time.Time:
			return VMTime(v)
		case VMNumberer:
			return v
		case VMDateTimer:
			return v
		case VMMetaStructer:
			return v
		}
	}
	panic("Невозможно привести к типу интерпретатора")
}
