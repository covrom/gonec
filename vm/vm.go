package vm

import (
	"errors"
	"io"

	"fmt"

	"github.com/covrom/gonec/gonecparser/ast"
	"github.com/covrom/gonec/gonecparser/variant"
	"github.com/covrom/gonec/gonecparser/token"
	"math/big"
	"time"
)

var (
	InitError = errors.New("Отсутствует исполняемый код модуля")
)

// виртуальная машина

type VirtMachine struct {
	af          *ast.File
	w           io.Writer
	fset        *token.FileSet
	mainContext *variant.Context
}

func NewVM(af *ast.File, w io.Writer, fst *token.FileSet) *VirtMachine {
	return &VirtMachine{
		af:          af,
		w:           w,
		fset:        fst,
		mainContext: nil,
	}
}

func (v *VirtMachine) Run() error {
	for _, ur := range v.af.Unresolved {
		fmt.Printf("Не назначен объект у идентификатора %v\n", ur.Name)
	}
	v.mainContext = variant.NewMainContext()
	defer v.mainContext.Destroy()

	ini := v.af.Scope.Lookup("__init__")
	if ini == nil {
		return InitError
	}
	
	return ast.Inspect(ini.Decl.(*ast.FuncDecl).Body, v.walker)
}

func (v *VirtMachine) walker(n ast.Node) (bool,error) {

	// ast.Print(v.fset,n)

	switch x := n.(type) {
	case *ast.BlockStmt:
	case *ast.AssignStmt:
		//перебираем левые переменные и присваиваем им соответствующие правые
		//с учетом взаимообмена - для этого сначала вычисляем все правые части
		if len(x.Lhs)!=len(x.Rhs){
			return false,fmt.Errorf("Количество переменных не сопадает с количеством выражений присваивания %v",x.Pos)
		}
		vars:=[]*variant.Variant{}
		for nr:=range x.Rhs{
			vars = append(vars,v.evaluate(x.Rhs[nr],v.mainContext))
		}
		for nr:=range x.Lhs{
			vars = append(vars,v.evaluate(x.Rhs[nr],v.mainContext))
		}
		
		ast.Print(v.fset,x.Lhs)
		ast.Print(v.fset,x.Rhs)
		//выражение разобрано - углубляться в дерево не требуется
		return false,nil
	}
	return true,nil
}

func (v *VirtMachine) evaluate(ex ast.Expr, ctx *variant.Context) *variant.Variant {
	nv:=variant.NewVariant()
	switch e:=ex.(type){
		case *ast.BadExpr:
		case *ast.BasicLit:
			switch e.Kind{
				//простые типы значений
				case token.STRING:
					nv.SetString(e.Value)
				case token.NUM:
					f:=big.NewFloat(0.0)
					if f,ok:=f.SetString(e.Value);ok{
						nv.SetNum(*f)
					}
				case token.DATE:
					if dt,err:=time.Parse(time.RFC3339,e.Value);err==nil{
						nv.SetDate(dt)
					}
				case token.UNDEF:
					nv.SetUNDEF()
				case token.TRUE:
					nv.SetBool(true)
				case token.FALSE:
					nv.SetBool(false)				
				case token.NULL:
					nv.SetNULL()
			}
		case *ast.Ident:
			nv = ctx.GetVar(e.Obj.Name)

	}
	return nv
}