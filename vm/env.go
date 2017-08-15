package vm

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/parser"
)

// Env provides interface to run VM. This mean function scope and blocked-scope.
// If stack goes to blocked-scope, it will make new Env.
type Env struct {
	sync.RWMutex
	name      string
	env       map[int]reflect.Value
	typ       map[int]reflect.Type
	parent    *Env
	interrupt *bool
	stdout    io.Writer
	sid       string
}

// NewEnv creates new global scope.
func NewEnv() *Env {
	b := false

	return &Env{
		env:       make(map[int]reflect.Value),
		typ:       make(map[int]reflect.Type),
		parent:    nil,
		interrupt: &b,
		stdout:    os.Stdout,
	}
}

// NewEnv creates new child scope.
func (e *Env) NewEnv() *Env {
	return &Env{
		env:       make(map[int]reflect.Value),
		typ:       make(map[int]reflect.Type),
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
		env:       make(map[int]reflect.Value),
		typ:       make(map[int]reflect.Type),
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
	//ni := strings.ToLower(n)
	m := &Env{
		env:    make(map[int]reflect.Value),
		typ:    make(map[int]reflect.Type),
		parent: e,
		name:   n, //ni,
		stdout: e.stdout,
	}
	e.Define(ast.UniqueNames.Set(n), m)
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
func (e *Env) Addr(k int) (reflect.Value, error) {

	for ee := e; ee != nil; ee = ee.parent {
		ee.RLock()
		defer ee.RUnlock()
		if v, ok := ee.env[k]; ok {
			return v.Addr(), nil
		}
	}
	return NilValue, fmt.Errorf("Имя неопределено '%s'", ast.UniqueNames.Get(k))
}

// TypeName определяет имя типа по типу значения
func (e *Env) TypeName(t reflect.Type) int {

	for ee := e; ee != nil; ee = ee.parent {
		ee.RLock()
		defer ee.RUnlock()
		for k, v := range ee.typ {
			if v == t {
				return k
			}
		}
	}
	return ast.UniqueNames.Set(t.String())
}

// Type returns type which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Type(k int) (reflect.Type, error) {

	for ee := e; ee != nil; ee = ee.parent {
		ee.RLock()
		defer ee.RUnlock()
		if v, ok := ee.typ[k]; ok {
			return v, nil
		}
	}
	return NilType, fmt.Errorf("Тип неопределен '%s'", ast.UniqueNames.Get(k))
}

// Get returns value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Get(k int) (reflect.Value, error) {

	for ee := e; ee != nil; ee = ee.parent {
		ee.RLock()
		defer ee.RUnlock()
		if v, ok := ee.env[k]; ok {
			return v, nil
		}
	}
	return NilValue, fmt.Errorf("Имя неопределено '%s'", ast.UniqueNames.Get(k))
}

// Set modifies value which specified as symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Set(k int, v interface{}) error {

	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}

	for ee := e; ee != nil; ee = ee.parent {
		ee.Lock()
		defer ee.Unlock()
		if _, ok := ee.env[k]; ok {
			ee.env[k] = val
			return nil
		}
	}
	return fmt.Errorf("Имя неопределено '%s'", ast.UniqueNames.Get(k))
}

// DefineGlobal defines symbol in global scope.
func (e *Env) DefineGlobal(k int, v interface{}) error {
	for ee := e; ee != nil; ee = ee.parent {
		if ee.parent == nil {
			return ee.Define(k, v)
		}
	}
	return fmt.Errorf("Отсутствует глобальный контекст!")
}

// DefineType defines type which specifis symbol in global scope.
func (e *Env) DefineType(k int, t interface{}) error {
	for ee := e; ee != nil; ee = ee.parent {
		if ee.parent == nil {
			typ, ok := t.(reflect.Type)
			if !ok {
				typ = reflect.TypeOf(t)
			}
			ee.Lock()
			ee.typ[k] = typ
			ee.Unlock()
			return nil
		}
	}
	return fmt.Errorf("Отсутствует глобальный контекст!")
}

func (e *Env) DefineTypeS(k string, t interface{}) error {
	return e.DefineType(ast.UniqueNames.Set(k), t)
}

// Define defines symbol in current scope.
func (e *Env) Define(k int, v interface{}) error {
	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}

	e.Lock()
	e.env[k] = val
	e.Unlock()

	return nil
}

func (e *Env) DefineS(k string, v interface{}) error {
	return e.Define(ast.UniqueNames.Set(k), v)
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
		e.Printf("%d %s = %#v\n", k, ast.UniqueNames.Get(k), v)
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
	// e.RLock()
	// defer e.RUnlock()
	return fmt.Fprintln(e.stdout, a...)
}

func (e *Env) Printf(format string, a ...interface{}) (n int, err error) {
	// e.RLock()
	// defer e.RUnlock()
	return fmt.Fprintf(e.stdout, format, a...)
}

func (e *Env) Sprintf(format string, a ...interface{}) string {
	// e.RLock()
	// defer e.RUnlock()
	return fmt.Sprintf(format, a...)
}

func (e *Env) Print(a ...interface{}) (n int, err error) {
	// e.RLock()
	// defer e.RUnlock()
	return fmt.Fprint(e.stdout, a...)
}

func (e *Env) StdOut() reflect.Value {
	// e.RLock()
	// defer e.RUnlock()
	return reflect.ValueOf(e.stdout)
}

func (e *Env) SetStdOut(w io.Writer) {
	// e.Lock()
	//пренебрегаем возможными коллизиями при установке потока вывода, т.к. это совсем редкая операция
	e.stdout = w
	// e.Unlock()
}

func (e *Env) SetSid(s string) error {
	for ee := e; ee != nil; ee = ee.parent {
		if ee.parent == nil {
			ee.sid = s
			return ee.Define(ast.UniqueNames.Set("ГлобальныйИдентификаторСессии"), s)
		}
	}
	return fmt.Errorf("Отсутствует глобальный контекст!")
}

func (e *Env) GetSid() string {
	for ee := e; ee != nil; ee = ee.parent {
		if ee.parent == nil {
			// пренебрегаем возможными коллизиями, т.к. изменение номера сессии - это совсем редкая операция
			return ee.sid
		}
	}
	return ""
}
