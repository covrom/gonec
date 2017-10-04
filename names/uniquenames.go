package names

import (
	"bytes"
	"fmt"
	"sync"
)

// TODO: переделать на atomic инкремент Iter

// все переменные
var UniqueNames = NewEnvNames()

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
	ns := FastToLower(n)
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
	ns := FastToLower(n)
	en.mu.Lock()
	en.Names[ns] = i
	en.Handles[i] = n
	en.Handlow[i] = ns
	if en.Iter <= i {
		en.Iter = i + 1 // гарантированно следующий
	}
	en.mu.Unlock()
}

func FastToLower(s string) string {
	rs := bytes.NewBuffer(make([]byte, 0, len(s)))
	for _, rn := range s {
		switch {
		case (rn >= 'А' && rn <= 'Я') || (rn >= 'A' && rn <= 'Z'):
			rs.WriteRune(rn + 0x20)
		case rn == 'Ё':
			rs.WriteRune('ё')
		default:
			rs.WriteRune(rn)
		}
	}
	return rs.String()
}
