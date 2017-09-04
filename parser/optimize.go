package parser

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/bincode"
)

// TODO: переделать на универсальную рефлексию перебора полей структур

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
			r, err := EvalBinOp(e.Operator, r1, r2, none)
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
			r, err := EvalUnOp(e.Operator, v, none)
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
			if ToBool(v) {
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
		r = InvokeConst(e.Value, none)
	case *ast.NumberExpr:
		r, err = InvokeNumber(e.Lit, none)
		if err != nil {
			// ошибки пропускаем
			return none
		}
	case *ast.StringExpr:
		r = reflect.ValueOf(e.Lit)
	}
	return
}

func InvokeConst(v string, defval reflect.Value) reflect.Value {
	switch strings.ToLower(v) {
	case "истина":
		return reflect.ValueOf(true)
	case "ложь":
		return reflect.ValueOf(false)
	case "null":
		return reflect.ValueOf(bincode.NullVar)
	}
	return defval
}

func InvokeNumber(lit string, defval reflect.Value) (reflect.Value, error) {
	if strings.Contains(lit, ".") || strings.Contains(lit, "e") {
		v, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return defval, err
		}
		return reflect.ValueOf(float64(v)), nil
	}
	var i int64
	var err error
	if strings.HasPrefix(lit, "0x") {
		i, err = strconv.ParseInt(lit[2:], 16, 64)
	} else {
		i, err = strconv.ParseInt(lit, 10, 64)
	}
	if err != nil {
		return defval, err
	}
	return reflect.ValueOf(i), nil
}

func EvalUnOp(op string, v, defval reflect.Value) (reflect.Value, error) {
	switch op {
	case "-":
		if v.Kind() == reflect.Float64 {
			return reflect.ValueOf(-v.Float()), nil
		}
		return reflect.ValueOf(-v.Int()), nil
	case "^":
		return reflect.ValueOf(^ToInt64(v)), nil
	case "!":
		return reflect.ValueOf(!ToBool(v)), nil
	default:
		return defval, fmt.Errorf("Неизвестный оператор")
	}
}

func ToBool(v reflect.Value) bool {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0.0
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		vlow := strings.ToLower(v.String())
		if vlow == "true" || vlow == "истина" {
			return true
		}
		if ToInt64(v) != 0 {
			return true
		}
	}
	return false
}

func ToInt64(v reflect.Value) int64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.String:
		s := v.String()
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return int64(i)
		}
	}
	return 0
}

func ToString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	if !v.IsValid() {
		return "Неопределено"
	}
	if v.Kind() == reflect.Bool {
		if v.Bool() {
			return "Истина"
		} else {
			return "Ложь"
		}
	}
	return fmt.Sprint(v.Interface())
}

func ToFloat64(v reflect.Value) float64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	}
	return 0.0
}

func IsNil(v reflect.Value) bool {
	if !v.IsValid() || v.Kind().String() == "unsafe.Pointer" {
		return true
	}
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
		return true
	}
	return false
}

func IsNum(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// equal returns true when lhsV and rhsV is same value.
func Equal(lhsV, rhsV reflect.Value) bool {
	lhsIsNil, rhsIsNil := IsNil(lhsV), IsNil(rhsV)
	if lhsIsNil && rhsIsNil {
		return true
	}
	if (!lhsIsNil && rhsIsNil) || (lhsIsNil && !rhsIsNil) {
		return false
	}
	if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
		lhsV = lhsV.Elem()
	}
	if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
		rhsV = rhsV.Elem()
	}
	if !lhsV.IsValid() || !rhsV.IsValid() {
		return true
	}
	if IsNum(lhsV) && IsNum(rhsV) {
		if rhsV.Type().ConvertibleTo(lhsV.Type()) {
			rhsV = rhsV.Convert(lhsV.Type())
		}
	}
	if lhsV.CanInterface() && rhsV.CanInterface() {
		return reflect.DeepEqual(lhsV.Interface(), rhsV.Interface())
	}
	return reflect.DeepEqual(lhsV, rhsV)
}

func EvalBinOp(op string, lhsV, rhsV, defval reflect.Value) (reflect.Value, error) {
	switch op {

	// TODO: расширить возможные варианты

	case "+":
		if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
			return reflect.ValueOf(ToString(lhsV) + ToString(rhsV)), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() != reflect.Array && rhsV.Kind() != reflect.Slice) {
			return reflect.Append(lhsV, rhsV), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() == reflect.Array || rhsV.Kind() == reflect.Slice) {
			return reflect.AppendSlice(lhsV, rhsV), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(ToFloat64(lhsV) + ToFloat64(rhsV)), nil
		}
		return reflect.ValueOf(ToInt64(lhsV) + ToInt64(rhsV)), nil
	case "-":
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(ToFloat64(lhsV) - ToFloat64(rhsV)), nil
		}
		return reflect.ValueOf(ToInt64(lhsV) - ToInt64(rhsV)), nil
	case "*":
		if lhsV.Kind() == reflect.String && (rhsV.Kind() == reflect.Int || rhsV.Kind() == reflect.Int32 || rhsV.Kind() == reflect.Int64) {
			return reflect.ValueOf(strings.Repeat(ToString(lhsV), int(ToInt64(rhsV)))), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(ToFloat64(lhsV) * ToFloat64(rhsV)), nil
		}
		return reflect.ValueOf(ToInt64(lhsV) * ToInt64(rhsV)), nil
	case "/":
		return reflect.ValueOf(ToFloat64(lhsV) / ToFloat64(rhsV)), nil
	case "%":
		return reflect.ValueOf(ToInt64(lhsV) % ToInt64(rhsV)), nil
	case "==":
		return reflect.ValueOf(Equal(lhsV, rhsV)), nil
	case "!=":
		return reflect.ValueOf(Equal(lhsV, rhsV) == false), nil
	case ">":
		return reflect.ValueOf(ToFloat64(lhsV) > ToFloat64(rhsV)), nil
	case ">=":
		return reflect.ValueOf(ToFloat64(lhsV) >= ToFloat64(rhsV)), nil
	case "<":
		return reflect.ValueOf(ToFloat64(lhsV) < ToFloat64(rhsV)), nil
	case "<=":
		return reflect.ValueOf(ToFloat64(lhsV) <= ToFloat64(rhsV)), nil
	case "|":
		return reflect.ValueOf(ToInt64(lhsV) | ToInt64(rhsV)), nil
	case "||":
		if ToBool(lhsV) {
			return lhsV, nil
		}
		return rhsV, nil
	case "&":
		return reflect.ValueOf(ToInt64(lhsV) & ToInt64(rhsV)), nil
	case "&&":
		if ToBool(lhsV) {
			return rhsV, nil
		}
		return lhsV, nil
	case "**":
		if lhsV.Kind() == reflect.Float64 {
			return reflect.ValueOf(math.Pow(ToFloat64(lhsV), ToFloat64(rhsV))), nil
		}
		return reflect.ValueOf(int64(math.Pow(ToFloat64(lhsV), ToFloat64(rhsV)))), nil
	case ">>":
		return reflect.ValueOf(ToInt64(lhsV) >> uint64(ToInt64(rhsV))), nil
	case "<<":
		return reflect.ValueOf(ToInt64(lhsV) << uint64(ToInt64(rhsV))), nil
	default:
		return defval, fmt.Errorf("Неизвестный оператор")
	}
}
