package vm

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/covrom/gonec/parser"
)

// Env provides interface to run VM. This mean function scope and blocked-scope.
// If stack goes to blocked-scope, it will make new Env.
type Env struct {
	name      string
	env       map[string]reflect.Value
	typ       map[string]reflect.Type
	parent    *Env
	interrupt *bool
	stdout    io.Writer
	sid       string
	sync.RWMutex
}

// NewEnv creates new global scope.
func NewEnv() *Env {
	b := false

	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    nil,
		interrupt: &b,
		stdout:    os.Stdout,
	}
}

// NewEnv creates new child scope.
func (e *Env) NewEnv() *Env {
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    e,
		name:      e.name,
		interrupt: e.interrupt,
		stdout:    e.stdout,
	}
}

// func NewPackage(n string, w io.Writer) *Env {
// 	b := false

// 	return &Env{
// 		env:       make(map[string]reflect.Value),
// 		typ:       make(map[string]reflect.Type),
// 		parent:    nil,
// 		name:      strings.ToLower(n),
// 		interrupt: &b,
// 		stdout:    w,
// 	}
// }

func (e *Env) NewPackage(n string) *Env {
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    e,
		name:      strings.ToLower(n),
		interrupt: e.interrupt,
		stdout:    e.stdout,
	}
}

// Destroy deletes current scope.
func (e *Env) Destroy() {
	e.Lock()
	defer e.Unlock()

	if e.parent == nil {
		return
	}
	for k, v := range e.parent.env {
		if v.IsValid() && v.Interface() == e {
			delete(e.parent.env, k)
		}
	}
	e.parent = nil
	e.env = nil
}

// NewModule creates new module scope as global.
func (e *Env) NewModule(n string) *Env {
	ni := strings.ToLower(n)
	m := &Env{
		env:    make(map[string]reflect.Value),
		parent: e,
		name:   ni,
		stdout: e.stdout,
	}
	e.Define(ni, m)
	return m
}

// SetName sets a name of the scope. This means that the scope is module.
func (e *Env) SetName(n string) {
	e.Lock()
	e.name = strings.ToLower(n)
	e.Unlock()
}

// GetName returns module name.
func (e *Env) GetName() string {
	e.RLock()
	defer e.RUnlock()

	return e.name
}

// Addr returns pointer value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Addr(k string) (reflect.Value, error) {
	e.RLock()
	defer e.RUnlock()

	if v, ok := e.env[strings.ToLower(k)]; ok {
		return v.Addr(), nil
	}
	if e.parent == nil {
		return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
	}
	return e.parent.Addr(k)
}

// TypeName определяет имя типа по типу значения
func (e *Env) TypeName(t reflect.Type) string {
	e.RLock()
	defer e.RUnlock()

	for k, v := range e.typ {
		if v == t {
			return k
		}
	}
	if e.parent == nil {
		return t.String()
	}
	return e.parent.TypeName(t)
}

// Type returns type which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Type(k string) (reflect.Type, error) {
	e.RLock()
	defer e.RUnlock()

	if v, ok := e.typ[strings.ToLower(k)]; ok {
		return v, nil
	}
	if e.parent == nil {
		return NilType, fmt.Errorf("Undefined type '%s'", k)
	}
	return e.parent.Type(k)
}

// Get returns value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Get(k string) (reflect.Value, error) {
	e.RLock()
	defer e.RUnlock()

	if v, ok := e.env[strings.ToLower(k)]; ok {
		return v, nil
	}
	if e.parent == nil {
		return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
	}
	return e.parent.Get(k)
}

// Set modifies value which specified as symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Set(k string, v interface{}) error {
	e.Lock()
	defer e.Unlock()
	ki := strings.ToLower(k)
	if _, ok := e.env[ki]; ok {
		val, ok := v.(reflect.Value)
		if !ok {
			val = reflect.ValueOf(v)
		}
		e.env[ki] = val
		return nil
	}
	if e.parent == nil {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	return e.parent.Set(ki, v)
}

// DefineGlobal defines symbol in global scope.
func (e *Env) DefineGlobal(k string, v interface{}) error {
	if e.parent == nil {
		return e.Define(k, v)
	}
	return e.parent.DefineGlobal(k, v)
}

// DefineType defines type which specifis symbol in global scope.
func (e *Env) DefineType(k string, t interface{}) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	global := e
	keys := []string{strings.ToLower(k)}

	e.RLock()
	for global.parent != nil {
		if global.name != "" {
			keys = append(keys, global.name)
		}
		global = global.parent
	}
	e.RUnlock()

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	typ, ok := t.(reflect.Type)
	if !ok {
		typ = reflect.TypeOf(t)
	}

	global.Lock()
	global.typ[strings.Join(keys, ".")] = typ
	global.Unlock()

	return nil
}

// Define defines symbol in current scope.
func (e *Env) Define(k string, v interface{}) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}

	e.Lock()
	e.env[strings.ToLower(k)] = val
	e.Unlock()

	return nil
}

// String return the name of current scope.
func (e *Env) String() string {
	e.RLock()
	defer e.RUnlock()

	return e.name
}

// Dump show symbol values in the scope.
func (e *Env) Dump() {
	e.RLock()
	for k, v := range e.env {
		fmt.Printf("%v = %#v\n", k, v)
	}
	e.RUnlock()
}

// Execute parses and runs source in current scope.
func (e *Env) Execute(src string) (reflect.Value, error) {
	stmts, err := parser.ParseSrc(src)
	if err != nil {
		return NilValue, err
	}
	return Run(stmts, e)
}

func (e *Env) Println(a ...interface{}) (n int, err error) {
	e.RLock()
	defer e.RUnlock()
	return fmt.Fprintln(e.stdout, a...)
}

func (e *Env) Printf(format string, a ...interface{}) (n int, err error) {
	e.RLock()
	defer e.RUnlock()
	return fmt.Fprintf(e.stdout, format, a...)
}

func (e *Env) Print(a ...interface{}) (n int, err error) {
	e.RLock()
	defer e.RUnlock()
	return fmt.Fprint(e.stdout, a...)
}

func (e *Env) StdOut() reflect.Value {
	e.RLock()
	defer e.RUnlock()
	return reflect.ValueOf(e.stdout)
}

func (e *Env) SetStdOut(w io.Writer) {
	e.Lock()
	e.stdout = w
	e.Unlock()
}

func (e *Env) SetSid(s string) error {
	if e.parent == nil {
		e.sid = s
		return e.Define("ГлобальныйИдентификаторСессии", s)
	} else {
		return e.parent.SetSid(s)
	}
}

func (e *Env) GetSid() string {
	if e.parent == nil {
		return e.sid
	} else {
		return e.parent.GetSid()
	}
}
