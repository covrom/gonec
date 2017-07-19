package variant

import "sync"

type Context struct {
	sync.RWMutex
	Name      string
	Parent    *Context
	Vars      map[string]Variant
	interrupt *bool //сигнал останова интерпретатора
}

func NewMainContext() *Context {
	b := false
	return &Context{
		Vars:      make(map[string]Variant),
		Parent:    nil,
		interrupt: &b,
	}
}

func (ctx *Context) Destroy(){

}