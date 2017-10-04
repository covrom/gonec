package names

import (
	"bytes"
	"fmt"
	"sync"
)

// все переменные
var UniqueNames = NewEnvNames()

// уникальные названия переменных, индекс используется в AST-дереве
type EnvNames struct {
	mu      sync.RWMutex
	Names   map[string]int
	Handles []string
	Handlow []string
	Iter    int
}

func NewEnvNames() *EnvNames {
	en := EnvNames{
		Names:   make(map[string]int, 200),
		Handles: make([]string, 2, 200),
		Handlow: make([]string, 2, 200),
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
	for en.Iter >= len(en.Handles) {
		en.Handles = append(en.Handles, "")
	}
	for en.Iter >= len(en.Handlow) {
		en.Handlow = append(en.Handlow, "")
	}
	en.mu.Unlock()
	return i
}

func (en *EnvNames) Get(i int) string {
	en.mu.RLock()
	defer en.mu.RUnlock()
	if i >= 0 && i < len(en.Handles) {
		return en.Handles[i]
	} else {
		panic(fmt.Sprintf("Не найден идентификатор переменной id=%d", i))
	}
}

func (en *EnvNames) GetLowerCase(i int) string {
	en.mu.RLock()
	defer en.mu.RUnlock()
	if i >= 0 && i < len(en.Handlow) {
		return en.Handlow[i]
	} else {
		panic(fmt.Sprintf("Не найден идентификатор переменной id=%d", i))
	}
}

func (en *EnvNames) GetLowerCaseOk(i int) (s string, ok bool) {
	en.mu.RLock()
	defer en.mu.RUnlock()
	if i >= 0 && i < len(en.Handlow) {
		return en.Handlow[i], true
	} else {
		return "", false
	}
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
	for en.Iter >= len(en.Handles) {
		en.Handles = append(en.Handles, "")
	}
	for en.Iter >= len(en.Handlow) {
		en.Handlow = append(en.Handlow, "")
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
