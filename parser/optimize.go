package parser

import (
	"reflect"

	"github.com/covrom/gonec/ast"
)

func constFolding(inast []ast.Stmt) []ast.Stmt {
	for i, st := range inast {
		switch s := st.(type) {
		case *ast.ExprStmt:
			inast[i].(*ast.ExprStmt).Expr = simplifyExprFolding(s.Expr)
		case *ast.VarStmt:
			for i2, e2 := range s.Exprs {
				inast[i].(*ast.VarStmt).Exprs[i2] = simplifyExprFolding(e2)
			}
		case *ast.LetsStmt:
			for i2, e2 := range s.Lhss {
				inast[i].(*ast.LetsStmt).Lhss[i2] = simplifyExprFolding(e2)
			}
			for i2, e2 := range s.Rhss {
				inast[i].(*ast.LetsStmt).Rhss[i2] = simplifyExprFolding(e2)
			}
		case *ast.IfStmt:
			inast[i].(*ast.IfStmt).If = simplifyExprFolding(s.If)
			inast[i].(*ast.IfStmt).Then = constFolding(s.Then)
			inast[i].(*ast.IfStmt).Else = constFolding(s.Else)
			inast[i].(*ast.IfStmt).ElseIf = constFolding(s.ElseIf)
		case *ast.TryStmt:
			inast[i].(*ast.TryStmt).Try = constFolding(s.Try)
			inast[i].(*ast.TryStmt).Catch = constFolding(s.Catch)
		case *ast.LoopStmt:
			inast[i].(*ast.LoopStmt).Expr = simplifyExprFolding(s.Expr)
			inast[i].(*ast.LoopStmt).Stmts = constFolding(s.Stmts)
		case *ast.ForStmt:
			inast[i].(*ast.ForStmt).Value = simplifyExprFolding(s.Value)
			inast[i].(*ast.ForStmt).Stmts = constFolding(s.Stmts)
		case *ast.NumForStmt:
			inast[i].(*ast.NumForStmt).Expr1 = simplifyExprFolding(s.Expr1)
			inast[i].(*ast.NumForStmt).Expr2 = simplifyExprFolding(s.Expr2)
			inast[i].(*ast.NumForStmt).Stmts = constFolding(s.Stmts)
		case *ast.ReturnStmt:
			for i2, e2 := range s.Exprs {
				inast[i].(*ast.ReturnStmt).Exprs[i2] = simplifyExprFolding(e2)
			}
		case *ast.ThrowStmt:
			inast[i].(*ast.ThrowStmt).Expr = simplifyExprFolding(s.Expr)
		case *ast.ModuleStmt:
			inast[i].(*ast.ModuleStmt).Stmts = constFolding(s.Stmts)
		case *ast.SwitchStmt:
			inast[i].(*ast.SwitchStmt).Expr = simplifyExprFolding(s.Expr)
			inast[i].(*ast.SwitchStmt).Cases = constFolding(s.Cases)
		case *ast.SelectStmt:
			inast[i].(*ast.SelectStmt).Cases = constFolding(s.Cases)
		case *ast.CaseStmt:
			inast[i].(*ast.CaseStmt).Expr = simplifyExprFolding(s.Expr)
			inast[i].(*ast.CaseStmt).Stmts = constFolding(s.Stmts)
		case *ast.DefaultStmt:
			inast[i].(*ast.DefaultStmt).Stmts = constFolding(s.Stmts)
		}
	}
	return inast
}

func simplifyExprFolding(expr ast.Expr) ast.Expr {
	none := reflect.Value{}
	switch e := expr.(type) {
	case *ast.BinOpExpr:
		// упрощаем подвыражения
		e.Lhs = simplifyExprFolding(e.Lhs)
		e.Rhs = simplifyExprFolding(e.Rhs)

		r1 := exprAsValue(e.Lhs)
		r2 := exprAsValue(e.Rhs)
		if r1.IsValid() && r2.IsValid() {
			r, err := ast.EvalBinOp(e.Operator, r1, r2, none)
			if err == nil && r.IsValid() {
				// log.Println("Set native value!")
				return &ast.NativeExpr{Value: r}
			}
		}
		return e
	case *ast.LetExpr:
		e.Lhs = simplifyExprFolding(e.Lhs)
		e.Rhs = simplifyExprFolding(e.Rhs)
		return e
	case *ast.UnaryExpr:
		e.Expr = simplifyExprFolding(e.Expr)
		v := exprAsValue(e.Expr)
		if v.IsValid() {
			r, err := ast.EvalUnOp(e.Operator, v, none)
			if err == nil && r.IsValid() {
				return &ast.NativeExpr{Value: r}
			}
		}
		return e
	case *ast.ParenExpr:
		e.SubExpr = simplifyExprFolding(e.SubExpr)
		v := exprAsValue(e.SubExpr)
		if v.IsValid() {
			return &ast.NativeExpr{Value: v}
		}
		return e
	case *ast.TernaryOpExpr:
		e.Expr = simplifyExprFolding(e.Expr)
		e.Lhs = simplifyExprFolding(e.Lhs)
		e.Rhs = simplifyExprFolding(e.Rhs)

		v := exprAsValue(e.Expr)
		r1 := exprAsValue(e.Lhs)
		r2 := exprAsValue(e.Rhs)

		if v.IsValid() && r1.IsValid() && r2.IsValid() && v.Kind() == reflect.Bool {
			if ast.ToBool(v) {
				return &ast.NativeExpr{Value: r1}
			} else {
				return &ast.NativeExpr{Value: r2}
			}
		}
		return e
	case *ast.ArrayExpr:
		a := make([]interface{}, len(e.Exprs))
		waserrors := false
		for i := range e.Exprs {
			e.Exprs[i] = simplifyExprFolding(e.Exprs[i])
			arg := exprAsValue(e.Exprs[i])
			if arg.IsValid() {
				a[i] = arg.Interface()
			} else {
				waserrors = true
			}
		}
		if waserrors {
			return e
		} else {
			return &ast.NativeExpr{Value: reflect.ValueOf(a)}
		}
	case *ast.MapExpr:
		waserrors := false
		m := make(map[string]interface{})
		for k, v := range e.MapExpr {
			vv := simplifyExprFolding(v)
			e.MapExpr[k] = vv
			arg := exprAsValue(vv)
			if arg.IsValid() {
				m[k] = arg.Interface()
			} else {
				waserrors = true
			}
		}
		if waserrors {
			return e
		} else {
			return &ast.NativeExpr{Value: reflect.ValueOf(m)}
		}
	case *ast.CallExpr:
		for i := range e.SubExprs {
			e.SubExprs[i] = simplifyExprFolding(e.SubExprs[i])
		}
		return e
	case *ast.AnonCallExpr:
		e.Expr = simplifyExprFolding(e.Expr)
		for i := range e.SubExprs {
			e.SubExprs[i] = simplifyExprFolding(e.SubExprs[i])
		}
		return e
	case *ast.ItemExpr:
		e.Value = simplifyExprFolding(e.Value)
		e.Index = simplifyExprFolding(e.Index)
		return e
	// case *ast.LetsExpr:
	// 	for i2, e2 := range e.Lhss {
	// 		e.Lhss[i2] = simplifyExprFolding(e2)
	// 	}
	// 	for i2, e2 := range e.Rhss {
	// 		e.Rhss[i2] = simplifyExprFolding(e2)
	// 	}
	// 	return e
	case *ast.AssocExpr:
		e.Lhs = simplifyExprFolding(e.Lhs)
		e.Rhs = simplifyExprFolding(e.Rhs)
		return e
	case *ast.ChanExpr:
		e.Lhs = simplifyExprFolding(e.Lhs)
		e.Rhs = simplifyExprFolding(e.Rhs)
		return e
	case *ast.MakeArrayExpr:
		e.CapExpr = simplifyExprFolding(e.CapExpr)
		e.LenExpr = simplifyExprFolding(e.LenExpr)
		return e
	case *ast.TypeCast:
		e.CastExpr = simplifyExprFolding(e.CastExpr)
		e.TypeExpr = simplifyExprFolding(e.TypeExpr)
		return e
	case *ast.MakeExpr:
		e.TypeExpr = simplifyExprFolding(e.TypeExpr)
		return e
	case *ast.MakeChanExpr:
		e.SizeExpr = simplifyExprFolding(e.SizeExpr)
		return e
	case *ast.PairExpr:
		e.Value = simplifyExprFolding(e.Value)
		return e
	case *ast.AddrExpr:
		e.Expr = simplifyExprFolding(e.Expr)
		return e
	case *ast.DerefExpr:
		e.Expr = simplifyExprFolding(e.Expr)
		return e
	case *ast.MemberExpr:
		e.Expr = simplifyExprFolding(e.Expr)
		return e
	case *ast.SliceExpr:
		e.Value = simplifyExprFolding(e.Value)
		e.Begin = simplifyExprFolding(e.Begin)
		e.End = simplifyExprFolding(e.End)
		return e
	case *ast.FuncExpr:
		e.Stmts = constFolding(e.Stmts)
		return e
	default:
		// одиночные значения - преобразовываем в нативные
		r := exprAsValue(e)
		if r.IsValid() {
			return &ast.NativeExpr{Value: r}
		}
	}
	// если не преобразовали - вернем исходное выражение
	return expr
}

func exprAsValue(expr ast.Expr) (r reflect.Value) {
	none := reflect.Value{}
	var err error
	switch e := expr.(type) {
	case *ast.NativeExpr:
		r = e.Value
	case *ast.ConstExpr:
		r = ast.InvokeConst(e.Value, none)
	case *ast.NumberExpr:
		r, err = ast.InvokeNumber(e.Lit, none)
		if err != nil {
			// ошибки пропускаем
			return none
		}
	case *ast.StringExpr:
		r = reflect.ValueOf(e.Lit)
	}
	return
}
