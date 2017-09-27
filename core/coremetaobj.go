package core

import (
	"reflect"
	"strings"

	"github.com/covrom/gonec/names"
)

// VMMetaObj корневая структура для системных функциональных структур Го, доступных из языка Гонец
// поля и методы должны отличаться друг от друга без учета регистра
// например, Set и set - в вирт. машине будут считаться одинаковыми, будет использоваться последнее по индексу
type VMMetaObj struct {
	vmMetaCacheM map[int]VMFunc
	vmMetaCacheF map[int][]int
	vmOriginal   VMMetaObject
}

func (v *VMMetaObj) vmval() {}

func (v *VMMetaObj) Interface() interface{} {
	// возвращает ссылку на структуру, от которой был вызван метод VMCacheMembers
	//rv:=*(*VMMetaObject)(v.vmOriginal)
	return v.vmOriginal
}

// VMCacheMembers кэширует все русскоязычные поля и методы для ссылки на объединенную структуру, переданную в VMMetaObject
// Поля и методы объектных структур на английском языке недоступны в коде на языке Гонец, что обеспечивает защиту внутренней реализации
// Вызывать эту функцию надо так:
// v:=&struct{VMMetaObj, a int}{}; v.VMCacheMembers(v)
func (v *VMMetaObj) VMCacheMembers(vv VMMetaObject) {
	v.vmMetaCacheM = make(map[int]VMFunc)
	v.vmMetaCacheF = make(map[int][]int)
	v.vmOriginal = vv

	// rv := reflect.ValueOf(v.vmOriginal)
	typ := reflect.TypeOf(v.vmOriginal)

	// пишем в кэш индексы полей и методов для структур

	ruslang := "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

	// методы
	nm := typ.NumMethod()
	for i := 0; i < nm; i++ {
		meth := typ.Method(i)

		// только экспортируемые методы с русскими буквами
		if meth.PkgPath == "" && strings.ContainsAny(meth.Name, ruslang) {
			namtyp := names.UniqueNames.Set(meth.Name)

			// fmt.Println(i, meth.Name, namtyp, meth.Func)

			v.vmMetaCacheM[namtyp] = func(vfunc reflect.Value) VMFunc {
				return VMFunc(
					func(args VMSlice, rets *VMSlice) error {
						x := make([]reflect.Value, len(args)+1)

						// receiver
						x[0] = reflect.ValueOf(v.vmOriginal)

						for i := range args {
							x[i+1] = reflect.ValueOf(args[i])
						}

						r := vfunc.Call(x)
						switch len(r) {
						case 0:
							return nil
						case 1:
							if x, ok := r[0].Interface().(VMValuer); ok {
								rets.Append(x)
								return nil
							}
							rets.Append(ReflectToVMValue(r[0]))
							return nil
						case 2:
							if e, ok := r[1].Interface().(error); ok {
								rets.Append(ReflectToVMValue(r[0]))
								return e
							}
							fallthrough
						default:
							*rets = make(VMSlice, len(r))
							for i := range r {
								(*rets)[i] = ReflectToVMValue(r[i])
							}
							return nil
						}
					})
			}(meth.Func)
		}
	}

	// поля
	ityp := typ.Elem()
	nm = ityp.NumField()
	for i := 0; i < nm; i++ {
		field := ityp.Field(i)
		// только экспортируемые неанонимные поля с русскими буквами
		if field.PkgPath == "" && !field.Anonymous && strings.ContainsAny(field.Name, ruslang) {

			// fmt.Println(field.Name)

			namtyp := names.UniqueNames.Set(field.Name)
			v.vmMetaCacheF[namtyp] = field.Index
		}
	}
}

func (v *VMMetaObj) VMIsField(name int) bool {
	_, ok := v.vmMetaCacheF[name]
	return ok
}

func (v *VMMetaObj) VMGetField(name int) VMInterfacer {

	// fmt.Println("GET " + names.UniqueNames.Get(name))

	rv := reflect.ValueOf(v.Interface()).Elem()

	// fmt.Println(rv.Type())

	vv := rv.FieldByIndex(v.vmMetaCacheF[name])

	// поля с типом вирт. машины вернем сразу
	if x, ok := vv.Interface().(VMInterfacer); ok {
		return x
	}
	return ReflectToVMValue(vv)
}

func (v *VMMetaObj) VMSetField(name int, val VMInterfacer) {

	// fmt.Println("SET " + names.UniqueNames.Get(name))

	rv := reflect.ValueOf(v.Interface()).Elem()

	// fmt.Println(rv.Type())

	vv := rv.FieldByIndex(v.vmMetaCacheF[name])
	if !vv.CanSet() {
		panic("Невозможно установить значение поля только для чтения")
	}
	// поля с типом вирт. машины присваиваем без конверсии
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
func (v *VMMetaObj) VMGetMethod(name int) (VMFunc, bool) {

	// fmt.Println(name)

	rv, ok := v.vmMetaCacheM[name]
	return rv, ok
}
