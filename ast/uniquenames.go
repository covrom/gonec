package ast

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// уникальные названия переменных, индекс используется в AST-дереве
type EnvNames struct {
	mu      sync.RWMutex
	Names   map[string]int
	Handles map[int]string
	Handlow map[int]string
	Iter    int
}

func NewEnvNames() *EnvNames {
	en := EnvNames{
		Names:   make(map[string]int, 200),
		Handles: make(map[int]string, 200),
		Handlow: make(map[int]string, 200),
		Iter:    1,
	}
	return &en
}

func (en *EnvNames) Set(n string) int {
	ns := strings.ToLower(n)
	en.mu.RLock()
	if i, ok := en.Names[ns]; ok {
		en.mu.RUnlock()
		return i
	}
	en.mu.RUnlock()
	en.mu.Lock()
	i := en.Iter
	en.Names[ns] = i
	en.Handles[i] = n
	en.Handlow[i] = ns
	en.Iter++
	en.mu.Unlock()
	return i
}

func (en *EnvNames) Get(i int) string {
	en.mu.RLock()
	defer en.mu.RUnlock()
	if s, ok := en.Handles[i]; ok {
		return s
	} else {
		panic(fmt.Sprintf("Не найден идентификатор переменной id=%d", i))
	}
}

func (en *EnvNames) GetLowerCase(i int) string {
	en.mu.RLock()
	defer en.mu.RUnlock()
	if s, ok := en.Handlow[i]; ok {
		return s
	} else {
		panic(fmt.Sprintf("Не найден идентификатор переменной id=%d", i))
	}
}

func (en *EnvNames) GetLowerCaseOk(i int) (s string, ok bool) {
	en.mu.RLock()
	defer en.mu.RUnlock()
	s, ok = en.Handlow[i]
	return
}

func (en *EnvNames) SetToId(n string, i int) {
	ns := strings.ToLower(n)
	en.mu.Lock()
	en.Names[ns] = i
	en.Handles[i] = n
	en.Handlow[i] = ns
	if en.Iter <= i {
		en.Iter = i + 1 // гарантированно следующий
	}
	en.mu.Unlock()
}

// все переменные
var UniqueNames = NewEnvNames()

// TODO: упростить, перенести в binfuncs

var StructMethodIndexes = struct {
	Cache map[int]int // pkg.typename.methname из UniqueNames
}{
	Cache: make(map[int]int, 200),
}

func MethodByNameCI(v reflect.Value, name int) (reflect.Value, error) {
	tv := v.Type()
	var fullname string
	// методы ссылок - вычисляем отдельно
	if tv.Kind() == reflect.Ptr {
		fullname = tv.Elem().PkgPath() + "." + tv.Elem().Name() + ".*" + UniqueNames.GetLowerCase(name)
	} else {
		fullname = tv.PkgPath() + "." + tv.Name() + "." + UniqueNames.GetLowerCase(name)
	}
	// fmt.Println("GET METHOD: " + fullname)
	if idx, ok := StructMethodIndexes.Cache[UniqueNames.Set(fullname)]; ok {
		return v.Method(idx), nil
	}
	return reflect.Value{}, fmt.Errorf("Метод %s не найден", fullname)
}

var StructFieldIndexes = struct {
	Cache map[int][]int // pkg.typename.fieldname из UniqueNames
}{
	Cache: make(map[int][]int, 200),
}

func FieldByNameCI(v reflect.Value, name int) (reflect.Value, error) {
	tv := v.Type()
	// у ссылок поля не будут найдены - это правильно
	fullname := tv.PkgPath() + "." + tv.Name() + "." + UniqueNames.GetLowerCase(name)
	// fmt.Println("GET FIELD: " + fullname)
	if idx, ok := StructFieldIndexes.Cache[UniqueNames.Set(fullname)]; ok {
		return v.FieldByIndex(idx), nil
	}
	return reflect.Value{}, fmt.Errorf("Поле %s не найдено", fullname)
}
