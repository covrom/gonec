package variant

import "sync"
import "strings"

type Context struct {
	sync.RWMutex
	Name      string
	Parent    *Context
	Vars      map[string]*Variant
	interrupt *bool //сигнал останова интерпретатора
}

func NewMainContext() *Context {
	b := false
	return &Context{
		Vars:      make(map[string]*Variant),
		Parent:    nil,
		interrupt: &b,
	}
}

func (ctx *Context) Destroy() {
	ctx.Lock()
	defer ctx.Unlock()
	ctx.Vars = nil
	ctx.Parent = nil
	*ctx.interrupt = true
}

func (ctx *Context) NewContext(name string) *Context {
	b := false
	return &Context{
		Vars:      make(map[string]*Variant),
		Parent:    ctx,
		interrupt: &b,
		Name:      name,
	}
}

func (ctx *Context) findVar(name string) *Variant {
	if v, ok := ctx.Vars[strings.ToLower(name)]; ctx.Parent != nil && !ok {
		return ctx.Parent.findVar(name)
	} else {
		return v
	}
}

func (ctx *Context) GetVar(name string) *Variant {
	ctx.RLock()
	defer ctx.RUnlock()

	v := ctx.findVar(name)
	//если переменной еще нет, а так же ее нет в родительских контекстах, создаем ее в текущем контексте
	if v == nil {
		v = NewVariant()
		ctx.SetVar(name, v)
	}
	return v
}

func (ctx *Context) SetVar(name string, val *Variant) {
	ctx.Lock()
	defer ctx.Unlock()
	ctx.Vars[strings.ToLower(name)] = val
}
