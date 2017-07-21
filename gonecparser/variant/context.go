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

func (ctx *Context) GetVar(name string) *Variant {
	ctx.RLock()
	defer ctx.RUnlock()

	if v, ok := ctx.Vars[strings.ToLower(name)]; ctx.Parent != nil && !ok {
		return ctx.Parent.GetVar(name)
	} else {
		return v
	}

}

func (ctx *Context) SetVar(name string, val *Variant) {
	ctx.Lock()
	defer ctx.Unlock()
	ctx.Vars[strings.ToLower(name)] = val
}
