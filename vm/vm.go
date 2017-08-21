package vm

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/parser"
)

var (
	NilValue   = reflect.ValueOf((*interface{})(nil))
	NilType    = reflect.TypeOf((*interface{})(nil))
	TrueValue  = reflect.ValueOf(true)
	FalseValue = reflect.ValueOf(false)
)

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	Message string
	Pos     ast.Position
}

var (
	BreakError     = errors.New("Неверное применение оператора Прервать")
	ContinueError  = errors.New("Неверное применение оператора Продолжить")
	ReturnError    = errors.New("Неверное применение оператора Возврат")
	InterruptError = errors.New("Выполнение прервано")
)

// NewStringError makes error interface with message.
func NewStringError(pos ast.Pos, err string) error {
	if pos == nil {
		return &Error{Message: err, Pos: ast.Position{1, 1}}
	}
	return &Error{Message: err, Pos: pos.Position()}
}

// NewErrorf makes error interface with message.
func NewErrorf(pos ast.Pos, format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...), Pos: pos.Position()}
}

// NewError makes error interface with message.
// This doesn't overwrite last error.
func NewError(pos ast.Pos, err error) error {
	if err == nil {
		return nil
	}
	if err == BreakError || err == ContinueError || err == ReturnError {
		return err
	}
	if pe, ok := err.(*parser.Error); ok {
		return pe
	}
	if ee, ok := err.(*Error); ok {
		return ee
	}
	return &Error{Message: err.Error(), Pos: pos.Position()}
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

// Func is function interface to reflect functions internaly.
type Func func(args ...reflect.Value) (reflect.Value, error)

func (f Func) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

func ToFunc(f Func) reflect.Value {
	return reflect.ValueOf(f)
}

// Run executes statements in the specified environment.
func Run(stmts []ast.Stmt, env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error
	for _, stmt := range stmts {
		if _, ok := stmt.(*ast.BreakStmt); ok {
			return NilValue, BreakError
		}
		if _, ok := stmt.(*ast.ContinueStmt); ok {
			return NilValue, ContinueError
		}
		rv, err = RunSingleStmt(stmt, env)
		if err != nil {
			return rv, err
		}
		if _, ok := stmt.(*ast.ReturnStmt); ok {
			return reflect.ValueOf(rv), ReturnError
		}
	}
	return rv, nil
}

// Interrupts the execution of any running statements in the specified environment.
//
// Note that the execution is not instantly aborted: after a call to Interrupt,
// the current running statement will finish, but the next statement will not run,
// and instead will return a NilValue and an InterruptError.
func Interrupt(env *Env) {
	env.Lock()
	*(env.interrupt) = true
	env.Unlock()
}

// RunSingleStmt executes one statement in the specified environment.
func RunSingleStmt(stmt ast.Stmt, env *Env) (reflect.Value, error) {
	env.Lock()
	if *(env.interrupt) {
		*(env.interrupt) = false
		env.Unlock()

		return NilValue, InterruptError
	}
	env.Unlock()

	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		return rv, nil
	case *ast.VarStmt:
		rv := NilValue
		var err error
		rvs := []reflect.Value{}
		for _, expr := range stmt.Exprs {
			rv, err = invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(expr, err)
			}
			rvs = append(rvs, rv)
		}
		result := []interface{}{}
		for i, name := range stmt.Names {
			if i < len(rvs) {
				env.Define(name, rvs[i])
				result = append(result, rvs[i].Interface())
			}
		}
		return reflect.ValueOf(result), nil
	case *ast.LetsStmt:
		rv := NilValue
		var err error
		vs := []interface{}{}
		for _, rhs := range stmt.Rhss {
			rv, err = invokeExpr(rhs, env)
			if err != nil {
				return rv, NewError(rhs, err)
			}
			if rv == NilValue {
				vs = append(vs, nil)
			} else if rv.IsValid() && rv.CanInterface() {
				vs = append(vs, rv.Interface())
			} else {
				vs = append(vs, nil)
			}
		}
		rvs := reflect.ValueOf(vs)
		if len(stmt.Lhss) > 1 && rvs.Len() == 1 {
			item := rvs.Index(0)
			if item.Kind() == reflect.Interface {
				item = item.Elem()
			}
			if item.Kind() == reflect.Slice {
				rvs = item
			}
		}
		for i, lhs := range stmt.Lhss {
			if i >= rvs.Len() {
				break
			}
			v := rvs.Index(i)
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			_, err = invokeLetExpr(lhs, v, env)
			if err != nil {
				return rvs, NewError(lhs, err)
			}
		}
		if rvs.Len() == 1 {
			return rvs.Index(0), nil
		}
		return rvs, nil
	case *ast.IfStmt:
		// Если Тогда ИначеЕсли Иначе КонецЕсли
		rv, err := invokeExpr(stmt.If, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if ast.ToBool(rv) {
			// Then
			newenv := env //.NewEnv()
			//defer newenv.Destroy()
			rv, err = Run(stmt.Then, newenv)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		done := false
		if len(stmt.ElseIf) > 0 {
			for _, stmt := range stmt.ElseIf {
				stmt_if := stmt.(*ast.IfStmt)
				// ElseIf
				rv, err = invokeExpr(stmt_if.If, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				if !ast.ToBool(rv) {
					continue
				}
				// ElseIf Then
				done = true
				rv, err = Run(stmt_if.Then, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				break
			}
		}
		if !done && len(stmt.Else) > 0 {
			// Else
			newenv := env //.NewEnv()
			//defer newenv.Destroy()
			rv, err = Run(stmt.Else, newenv)
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	case *ast.TryStmt:
		newenv := env //.NewEnv()
		//defer newenv.Destroy()
		_, err := Run(stmt.Try, newenv)
		if err != nil {
			// Catch
			cenv := env //.NewEnv()
			//defer cenv.Destroy()
			// if stmt.Var != "" {
			cenv.Define(ast.UniqueNames.Set("описаниеошибки"), reflect.ValueOf(err))
			// }
			_, e1 := Run(stmt.Catch, cenv)
			if e1 != nil {
				err = NewError(stmt.Catch[0], e1)
			} else {
				err = nil
			}
		}
		// if len(stmt.Finally) > 0 {
		// 	// Finally
		// 	fenv := env.NewEnv()
		// 	defer fenv.Destroy()
		// 	_, e2 := Run(stmt.Finally, newenv)
		// 	if e2 != nil {
		// 		err = NewError(stmt.Finally[0], e2)
		// 	}
		// }
		return NilValue, NewError(stmt, err)
	case *ast.LoopStmt:
		newenv := env.NewEnv()
		defer newenv.Destroy()
		for {
			if stmt.Expr != nil {
				ev, ee := invokeExpr(stmt.Expr, newenv)
				if ee != nil {
					return ev, ee
				}
				if !ast.ToBool(ev) {
					break
				}
			}

			rv, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				if err == ReturnError {
					return rv, err
				}
				return rv, NewError(stmt, err)
			}
		}
		return NilValue, nil
	case *ast.ForStmt:
		val, ee := invokeExpr(stmt.Value, env)
		if ee != nil {
			return val, ee
		}
		if val.Kind() == reflect.Interface {
			val = val.Elem()
		}
		if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
			//отключено создание новой области видимости внутри цикла
			newenv := env //.NewEnv()
			//defer newenv.Destroy()

			for i := 0; i < val.Len(); i++ {
				iv := val.Index(i)
				if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
					iv = iv.Elem()
				}
				newenv.Define(stmt.Var, iv)
				rv, err := Run(stmt.Stmts, newenv)
				if err != nil {
					if err == BreakError {
						err = nil
						break
					}
					if err == ContinueError {
						err = nil
						continue
					}
					if err == ReturnError {
						return rv, err
					}
					return rv, NewError(stmt, err)
				}
			}
			return NilValue, nil
		} else if val.Kind() == reflect.Chan {
			newenv := env //.NewEnv()
			//defer newenv.Destroy()

			for {
				iv, ok := val.Recv()
				if !ok {
					break
				}
				if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
					iv = iv.Elem()
				}
				newenv.Define(stmt.Var, iv)
				rv, err := Run(stmt.Stmts, newenv)
				if err != nil {
					if err == BreakError {
						err = nil
						break
					}
					if err == ContinueError {
						err = nil
						continue
					}
					if err == ReturnError {
						return rv, err
					}
					return rv, NewError(stmt, err)
				}
			}
			return NilValue, nil
		} else {
			return NilValue, NewStringError(stmt, "Цикл может выполняться только для переменных-коллекций или каналов")
		}
	case *ast.NumForStmt:
		newenv := env //.NewEnv()
		//defer newenv.Destroy()
		fv1, err := invokeExpr(stmt.Expr1, newenv)
		if err != nil {
			return NilValue, err
		}
		if !ast.IsNum(fv1) {
			return NilValue, NewStringError(stmt, "Цикл должен иметь числовой итератор")
		}

		fv2, err := invokeExpr(stmt.Expr2, newenv)
		if err != nil {
			return NilValue, err
		}
		if !ast.IsNum(fv2) {
			return NilValue, NewStringError(stmt, "Цикл должен иметь числовое целевое значение")
		}
		fvi1, fvi2 := ast.ToInt64(fv1), ast.ToInt64(fv2)
		fviadd := int64(1)
		if fvi1 > fvi2 {
			fviadd = int64(-1) // если конечное значение меньше первого, идем в обратном порядке
		}

		for fvi := fvi1; true; fvi = fvi + fviadd {

			newenv.Define(stmt.Name, fvi) //изменение итератора в теле цикла никак не влияет на стабильность перебора

			rv, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				if err == ReturnError {
					return rv, err
				}
				return rv, NewError(stmt, err)
			}

			if fvi == fvi2 { //когда дошли до равенства - проходим итерацию и прекращаем цикл
				break
			}
		}
		return NilValue, nil
	// case *ast.CForStmt:
	// 	newenv := env //.NewEnv()
	// 	//defer newenv.Destroy()
	// 	_, err := invokeExpr(stmt.Expr1, newenv)
	// 	if err != nil {
	// 		return NilValue, err
	// 	}
	// 	for {
	// 		fb, err := invokeExpr(stmt.Expr2, newenv)
	// 		if err != nil {
	// 			return NilValue, err
	// 		}
	// 		if !ast.ToBool(fb) {
	// 			break
	// 		}

	// 		rv, err := Run(stmt.Stmts, newenv)
	// 		if err != nil {
	// 			if err == BreakError {
	// 				err = nil
	// 				break
	// 			}
	// 			if err == ContinueError {
	// 				err = nil
	// 				continue
	// 			}
	// 			if err == ReturnError {
	// 				return rv, err
	// 			}
	// 			return rv, NewError(stmt, err)
	// 		}
	// 		_, err = invokeExpr(stmt.Expr3, newenv)
	// 		if err != nil {
	// 			return NilValue, err
	// 		}
	// 	}
	// 	return NilValue, nil
	case *ast.ReturnStmt:
		rvs := []interface{}{}
		switch len(stmt.Exprs) {
		case 0:
			return NilValue, nil
		case 1:
			rv, err := invokeExpr(stmt.Exprs[0], env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		for _, expr := range stmt.Exprs {
			rv, err := invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if ast.IsNil(rv) {
				rvs = append(rvs, nil)
			} else if rv.IsValid() {
				rvs = append(rvs, rv.Interface())
			} else {
				rvs = append(rvs, nil)
			}
		}
		return reflect.ValueOf(rvs), nil
	case *ast.ThrowStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if !rv.IsValid() {
			return NilValue, NewError(stmt, err)
		}
		return rv, NewStringError(stmt, fmt.Sprint(rv.Interface()))
	case *ast.ModuleStmt:
		newenv := env.NewEnv()
		//newenv.SetName(stmt.Name)
		rv, err := Run(stmt.Stmts, newenv)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		env.DefineGlobal(stmt.Name, reflect.ValueOf(newenv))
		return rv, nil
	case *ast.SwitchStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		done := false
		var default_stmt *ast.DefaultStmt
		for _, ss := range stmt.Cases {
			if ssd, ok := ss.(*ast.DefaultStmt); ok {
				default_stmt = ssd
				continue
			}
			case_stmt := ss.(*ast.CaseStmt)
			cv, err := invokeExpr(case_stmt.Expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if !ast.Equal(rv, cv) {
				continue
			}
			rv, err = Run(case_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			done = true
			break
		}
		if !done && default_stmt != nil {
			rv, err = Run(default_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	case *ast.SelectStmt:
		// основные варианты - если нет секции "другое", то обходятся в цикле, пока хотя бы один не сработает
		// если есть секция "другое" - выполняется она и делается выход из цикла
	startslct:
		done := false
		var rv reflect.Value
		var err error
		var default_stmt *ast.DefaultStmt
		for _, ss := range stmt.Cases {
			err = nil
			rv = NilValue
			if ssd, ok := ss.(*ast.DefaultStmt); ok {
				default_stmt = ssd
				continue
			}
			case_stmt := ss.(*ast.CaseStmt)

			switch e := case_stmt.Expr.(type) {
			case *ast.ChanExpr:
				rhs, err := invokeExpr(e.Rhs, env)
				if err != nil {
					return NilValue, NewError(case_stmt.Expr, err)
				}

				if e.Lhs == nil {
					if rhs.Kind() == reflect.Chan {
						// есть только правая часть - это чтение из канала без сохранения в переменной слева
						var ok bool
						rv, ok = rhs.TryRecv()
						if !ok {
							//не прочитано из канала - идем дальше
							continue
						}
					}
				} else {
					lhs, err := invokeExpr(e.Lhs, env)
					if err != nil {
						return NilValue, NewError(case_stmt.Expr, err)
					}
					if lhs.Kind() == reflect.Chan {
						// слева - канал, тогда пишем в него выражение справа
						ok := lhs.TrySend(rhs)
						if !ok {
							// не отправлено в канал
							continue
						}
					} else if rhs.Kind() == reflect.Chan {
						// слева - выражение, а справа - канал, тогда выражение - это переменная для присваивания
						var ok bool
						rv, ok = rhs.TryRecv()
						if !ok {
							continue
						}
						rv, err = invokeLetExpr(e.Lhs, rv, env)
					}
				}
			default:
				return NilValue, NewStringError(case_stmt.Expr, "При выборе вариантов из каналов допустимы только выражения с каналами")
			}

			if err != nil {
				return rv, NewError(stmt, err)
			}

			rv, err = Run(case_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			done = true
			break
		}
		if !done {
			if default_stmt != nil {
				rv, err = Run(default_stmt.Stmts, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
			} else {
				// если нет секции "другое", возвращаемся к выбору из каналов
				goto startslct
			}
		}
		return rv, nil
	default:
		return NilValue, NewStringError(stmt, "неизвестная конструкция")
	}
}

func invokeLetExpr(expr ast.Expr, rv reflect.Value, env *Env) (reflect.Value, error) {
	switch lhs := expr.(type) {
	case *ast.IdentExpr:
		if env.Set(lhs.Id, rv) != nil {
			if strings.Contains(lhs.Lit, ".") {
				return NilValue, NewErrorf(expr, "Имя неопределено '%s'", lhs.Lit)
			}
			env.Define(lhs.Id, rv)
		}
		return rv, nil
	case *ast.MemberExpr:
		v, err := invokeExpr(lhs.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}

		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice {
			v = v.Index(0)
		}

		if !v.IsValid() {
			return NilValue, NewStringError(expr, "Поле недоступно")
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Struct:
			v = ast.FieldByNameCI(v, lhs.Name)
			if !v.CanSet() {
				return NilValue, NewStringError(expr, "Значение не может быть изменено")
			}
			v.Set(rv)
		case reflect.Map:
			v.SetMapIndex(reflect.ValueOf(ast.UniqueNames.Get(lhs.Name)), rv)
		default:
			if !v.CanSet() {
				return NilValue, NewStringError(expr, "Значение не может быть изменено")
			}
			v.Set(rv)
		}
		return v, nil
	case *ast.ItemExpr:
		v, err := invokeExpr(lhs.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		i, err := invokeExpr(lhs.Index, env)
		if err != nil {
			return i, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Индекс должен быть целым числом")
			}
			ii := int(i.Int())
			if ii < 0 {
				ii += v.Len()
			}
			if ii < 0 || ii >= v.Len() {
				return NilValue, NewStringError(expr, "Индекс за пределами границ")
			}

			vv := v.Index(ii)
			if !vv.CanSet() {
				return NilValue, NewStringError(expr, "Значение не может быть изменено")
			}
			vv.Set(rv)
			return rv, nil
		}
		if v.Kind() == reflect.String {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Индекс должен быть целым числом")
			}
			rvs := []rune(rv.String())
			if len(rvs) != 1 {
				return NilValue, NewStringError(expr, "Длина строки должна быть ровно один символ")
			}
			r := []rune(v.String())
			vlen := len(r)
			ii := int(i.Int())
			if ii < 0 {
				ii += vlen
			}
			if ii < 0 || ii >= vlen {
				return NilValue, NewStringError(expr, "Индекс за пределами границ")
			}
			// заменяем руну
			r[ii] = rvs[0]
			return reflect.ValueOf(string(r)), nil
		}
		if v.Kind() == reflect.Map {
			if i.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Ключ должен быть строкой")
			}
			v.SetMapIndex(i, rv)
			return rv, nil
		}
		return v, NewStringError(expr, "Неверная операция")
	case *ast.SliceExpr:
		v, err := invokeExpr(lhs.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		rb, err := invokeExpr(lhs.Begin, env)
		if err != nil {
			return rb, NewError(expr, err)
		}
		re, err := invokeExpr(lhs.End, env)
		if err != nil {
			return re, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			vv, err := ast.SliceAt(v, rb, re, NilValue)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			if !vv.CanSet() {
				return NilValue, NewStringError(expr, "Диапазон не может быть изменен")
			}
			vv.Set(rv)
			return rv, nil
		}
		if v.Kind() == reflect.String {

			r, ii, ij, err := ast.StringToRuneSliceAt(v, rb, re)
			if err != nil {
				return NilValue, NewError(expr, err)
			}

			rvs := []rune(rv.String())
			if len(rvs) != len(r) {
				return NilValue, NewStringError(expr, "Длина строки должна быть равна длине диапазона")
			}

			// заменяем руны
			copy(r[ii:ij], rvs)

			return reflect.ValueOf(string(r)), nil
		}

		return v, NewStringError(expr, "Неверная операция")
	}
	return NilValue, NewStringError(expr, "Неверная операция")
}

// invokeExpr evaluates one expression.
func invokeExpr(expr ast.Expr, env *Env) (reflect.Value, error) {
	switch e := expr.(type) {
	case *ast.NativeExpr:
		return e.Value, nil
	case *ast.NumberExpr:
		i, err := ast.InvokeNumber(e.Lit, NilValue)
		if err != nil {
			return NilValue, NewError(expr, err)
		}
		return i, nil
	case *ast.IdentExpr:
		return env.Get(e.Id)
	case *ast.StringExpr:
		return reflect.ValueOf(e.Lit), nil
	case *ast.ArrayExpr:
		a := make([]interface{}, len(e.Exprs))
		for i, expr := range e.Exprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return arg, NewError(expr, err)
			}
			a[i] = arg.Interface()
		}
		return reflect.ValueOf(a), nil
	case *ast.MapExpr:
		m := make(map[string]interface{})
		for k, expr := range e.MapExpr {
			v, err := invokeExpr(expr, env)
			if err != nil {
				return v, NewError(expr, err)
			}
			m[k] = v.Interface()
		}
		return reflect.ValueOf(m), nil
	case *ast.DerefExpr:
		v := NilValue
		var err error
		switch ee := e.Expr.(type) {
		case *ast.IdentExpr:
			v, err = env.Get(ee.Id)
			if err != nil {
				return v, err
			}
		case *ast.MemberExpr:
			v, err := invokeExpr(ee.Expr, env)
			if err != nil {
				return v, NewError(expr, err)
			}
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if v.Kind() == reflect.Slice {
				v = v.Index(0)
			}
			if v.IsValid() && v.CanInterface() {
				if vme, ok := v.Interface().(*Env); ok {
					m, err := vme.Get(ee.Name)
					if !m.IsValid() || err != nil {
						return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
					}
					return m, nil
				}
			}

			//m := v.MethodByName(ast.UniqueNames.Get(ee.Name))
			m := ast.MethodByNameCI(v, ee.Name)

			if !m.IsValid() {
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					// m = v.FieldByName(ast.UniqueNames.Get(ee.Name))
					m = ast.FieldByNameCI(v, ee.Name)
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
					}
				} else if v.Kind() == reflect.Map {
					m = v.MapIndex(reflect.ValueOf(ast.UniqueNames.Get(ee.Name)))
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
					}
				} else {
					return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
				}
				v = m
			} else {
				v = m
			}
		default:
			return NilValue, NewStringError(expr, "Неверная операция для значения")
		}
		if v.Kind() != reflect.Ptr {
			return NilValue, NewStringError(expr, "Невозможно извлечь значение ссылки")
		}
		return v.Elem(), nil
	case *ast.AddrExpr:
		v := NilValue
		var err error
		switch ee := e.Expr.(type) {
		case *ast.IdentExpr:
			v, err = env.Get(ee.Id)
			if err != nil {
				return v, err
			}
		case *ast.MemberExpr:
			v, err := invokeExpr(ee.Expr, env)
			if err != nil {
				return v, NewError(expr, err)
			}
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if v.Kind() == reflect.Slice {
				v = v.Index(0)
			}
			if v.IsValid() && v.CanInterface() {
				if vme, ok := v.Interface().(*Env); ok {
					m, err := vme.Get(ee.Name)
					if !m.IsValid() || err != nil {
						return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
					}
					return m, nil
				}
			}

			// m := v.MethodByName(ast.UniqueNames.Get(ee.Name))
			m := ast.MethodByNameCI(v, ee.Name)

			if !m.IsValid() {
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					// m = v.FieldByName(ast.UniqueNames.Get(ee.Name))
					m = ast.FieldByNameCI(v, ee.Name)
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
					}
				} else if v.Kind() == reflect.Map {
					m = v.MapIndex(reflect.ValueOf(ast.UniqueNames.Get(ee.Name)))
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
					}
				} else {
					return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", ee.Name))
				}
				v = m
			} else {
				v = m
			}
		default:
			return NilValue, NewStringError(expr, "Неверная операция над значением")
		}
		if !v.CanAddr() {
			i := v.Interface()
			return reflect.ValueOf(&i), nil
		}
		return v.Addr(), nil
	case *ast.UnaryExpr:
		v, err := invokeExpr(e.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		r, err := ast.EvalUnOp(e.Operator, v, NilValue)
		return r, NewError(expr, err)
	case *ast.ParenExpr:
		v, err := invokeExpr(e.SubExpr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		return v, nil
	case *ast.FuncExpr:
		f := reflect.ValueOf(func(expr *ast.FuncExpr, env *Env) Func {
			return func(args ...reflect.Value) (reflect.Value, error) {
				if !expr.VarArg {
					if len(args) != len(expr.Args) {
						return NilValue, NewStringError(expr, "Неверное количество аргументов")
					}
				}
				newenv := env.NewEnv()
				if expr.VarArg {
					newenv.Define(expr.Args[0], reflect.ValueOf(args))
				} else {
					for i, arg := range expr.Args {
						newenv.Define(arg, args[i])
					}
				}
				rr, err := Run(expr.Stmts, newenv)
				if err == ReturnError {
					err = nil
					rr = rr.Interface().(reflect.Value)
				}
				return rr, err
			}
		}(e, env))
		env.Define(e.Name, f)
		return f, nil
	case *ast.MemberExpr:
		v, err := invokeExpr(e.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice {
			v = v.Index(0)
		}
		if v.IsValid() && v.CanInterface() {
			if vme, ok := v.Interface().(*Env); ok {
				m, err := vme.Get(e.Name)
				if !m.IsValid() || err != nil {
					return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", e.Name))
				}
				return m, nil
			}
		}

		m := ast.MethodByNameCI(v, e.Name)

		if !m.IsValid() {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				m = ast.FieldByNameCI(v, e.Name)
				if !m.IsValid() {
					return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", e.Name))
				}
			} else if v.Kind() == reflect.Map {
				m = v.MapIndex(reflect.ValueOf(ast.UniqueNames.Get(e.Name)))
				if !m.IsValid() {
					return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", e.Name))
				}
			} else {
				return NilValue, NewStringError(expr, fmt.Sprintf("Неверная операция '%s'", e.Name))
			}
		}
		return m, nil
	case *ast.ItemExpr:
		v, err := invokeExpr(e.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		i, err := invokeExpr(e.Index, env)
		if err != nil {
			return i, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Индекс массива должен быть целым числом")
			}
			ii := int(i.Int())
			if ii < 0 {
				ii += v.Len()
			}
			if ii < 0 || ii >= v.Len() {
				return NilValue, NewStringError(expr, "Индекс за пределами границ")
			}
			return v.Index(ii), nil
		}
		if v.Kind() == reflect.Map {
			if i.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Ключ структуры должен быть строкой")
			}
			return v.MapIndex(i), nil
		}
		if v.Kind() == reflect.String {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Индекс массива должен быть целым числом")
			}
			rs := []rune(v.Interface().(string))
			ii := int(i.Int())
			if ii < 0 {
				ii += len(rs)
			}
			if ii < 0 || ii >= len(rs) {
				return NilValue, NewStringError(expr, "Индекс за пределами границ")
			}
			return reflect.ValueOf(string(rs[ii])), nil
		}
		return v, NewStringError(expr, "Неверная операция")
	case *ast.SliceExpr:
		v, err := invokeExpr(e.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		rb, err := invokeExpr(e.Begin, env)
		if err != nil {
			return rb, NewError(expr, err)
		}
		re, err := invokeExpr(e.End, env)
		if err != nil {
			return re, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			rv, err := ast.SliceAt(v, rb, re, NilValue)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			return rv, nil
		}
		if v.Kind() == reflect.String {
			rv, err := ast.StringAt(v, rb, re, NilValue)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			return rv, nil
		}
		return v, NewStringError(expr, "Неверная операция")
	case *ast.AssocExpr:
		switch e.Operator {
		case "++":
			if alhs, ok := e.Lhs.(*ast.IdentExpr); ok {
				v, err := env.Get(alhs.Id)
				if err != nil {
					return v, err
				}
				if v.Kind() == reflect.Float64 {
					v = reflect.ValueOf(ast.ToFloat64(v) + 1.0)
				} else {
					v = reflect.ValueOf(ast.ToInt64(v) + 1)
				}
				if env.Set(alhs.Id, v) != nil {
					env.Define(alhs.Id, v)
				}
				return v, nil
			}
		case "--":
			if alhs, ok := e.Lhs.(*ast.IdentExpr); ok {
				v, err := env.Get(alhs.Id)
				if err != nil {
					return v, err
				}
				if v.Kind() == reflect.Float64 {
					v = reflect.ValueOf(ast.ToFloat64(v) - 1.0)
				} else {
					v = reflect.ValueOf(ast.ToInt64(v) - 1)
				}
				if env.Set(alhs.Id, v) != nil {
					env.Define(alhs.Id, v)
				}
				return v, nil
			}
		}

		v, err := invokeExpr(&ast.BinOpExpr{Lhs: e.Lhs, Operator: e.Operator[0:1], Rhs: e.Rhs}, env)
		if err != nil {
			return v, err
		}

		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		return invokeLetExpr(e.Lhs, v, env)
	case *ast.LetExpr:
		rv, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return rv, NewError(e, err)
		}
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		return invokeLetExpr(e.Lhs, rv, env)
	case *ast.LetsExpr:
		rv := NilValue
		var err error
		vs := []interface{}{}
		for _, rhs := range e.Rhss {
			rv, err = invokeExpr(rhs, env)
			if err != nil {
				return rv, NewError(rhs, err)
			}
			if rv == NilValue {
				vs = append(vs, nil)
			} else if rv.IsValid() && rv.CanInterface() {
				vs = append(vs, rv.Interface())
			} else {
				vs = append(vs, nil)
			}
		}
		rvs := reflect.ValueOf(vs)
		for i, lhs := range e.Lhss {
			if i >= rvs.Len() {
				break
			}
			v := rvs.Index(i)
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			_, err = invokeLetExpr(lhs, v, env)
			if err != nil {
				return rvs, NewError(lhs, err)
			}
		}
		return rvs, nil
	case *ast.NewExpr:
		rt, err := env.Type(e.Type)
		if err != nil {
			return NilValue, NewError(expr, err)
		}
		return reflect.New(rt), nil
	case *ast.BinOpExpr:
		lhsV := NilValue
		rhsV := NilValue
		var err error

		lhsV, err = invokeExpr(e.Lhs, env)
		if err != nil {
			return lhsV, NewError(expr, err)
		}
		if lhsV.Kind() == reflect.Interface {
			lhsV = lhsV.Elem()
		}
		if e.Rhs != nil {
			rhsV, err = invokeExpr(e.Rhs, env)
			if err != nil {
				return rhsV, NewError(expr, err)
			}
			if rhsV.Kind() == reflect.Interface {
				rhsV = rhsV.Elem()
			}
		}
		r, err := ast.EvalBinOp(e.Operator, lhsV, rhsV, NilValue)
		return r, NewError(expr, err)
	case *ast.ConstExpr:
		return ast.InvokeConst(e.Value, NilValue), nil
	case *ast.AnonCallExpr:
		f, err := invokeExpr(e.Expr, env)
		if err != nil {
			return f, NewError(expr, err)
		}
		if f.Kind() == reflect.Interface {
			f = f.Elem()
		}
		if f.Kind() != reflect.Func {
			return f, NewStringError(expr, "Неизвестная функция")
		}
		return invokeExpr(&ast.CallExpr{Func: f, SubExprs: e.SubExprs, VarArg: e.VarArg, Go: e.Go}, env)
	case *ast.CallExpr:
		f := NilValue

		if e.Func != nil {
			f = e.Func.(reflect.Value)
		} else {
			var err error
			ff, err := env.Get(e.Name)
			if err != nil {
				return f, err
			}
			f = ff
		}
		_, isReflect := f.Interface().(Func)

		args := []reflect.Value{}
		l := len(e.SubExprs)
		for i, expr := range e.SubExprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return arg, NewError(expr, err)
			}

			if i < f.Type().NumIn() {
				if !f.Type().IsVariadic() {
					it := f.Type().In(i)
					if arg.Kind().String() == "unsafe.Pointer" {
						arg = reflect.New(it).Elem()
					}
					if arg.Kind() != it.Kind() && arg.IsValid() && arg.Type().ConvertibleTo(it) {
						arg = arg.Convert(it)
					} else if arg.Kind() == reflect.Func {
						if _, isFunc := arg.Interface().(Func); isFunc {
							rfunc := arg
							arg = reflect.MakeFunc(it, func(args []reflect.Value) []reflect.Value {
								for i := range args {
									args[i] = reflect.ValueOf(args[i])
								}
								if e.Go {
									go func() {
										rfunc.Call(args)
									}()
									return []reflect.Value{}
								}
								var rets []reflect.Value
								for _, v := range rfunc.Call(args)[:it.NumOut()] {
									rets = append(rets, v.Interface().(reflect.Value))
								}
								return rets
							})
						}
					} else if !arg.IsValid() {
						arg = reflect.Zero(it)
					}
				}
			}
			if !arg.IsValid() {
				arg = NilValue
			}

			if !isReflect {
				if e.VarArg && i == l-1 {
					for j := 0; j < arg.Len(); j++ {
						args = append(args, arg.Index(j).Elem())
					}
				} else {
					args = append(args, arg)
				}
			} else {
				if arg.Kind() == reflect.Interface {
					arg = arg.Elem()
				}
				if e.VarArg && i == l-1 {
					for j := 0; j < arg.Len(); j++ {
						args = append(args, reflect.ValueOf(arg.Index(j).Elem()))
					}
				} else {
					args = append(args, reflect.ValueOf(arg))
				}
			}
		}
		ret := NilValue
		var err error
		fnc := func() {
			defer func() {
				if os.Getenv("GONEC_DEBUG") == "" {
					if ex := recover(); ex != nil {
						if e, ok := ex.(error); ok {
							err = e
						} else {
							err = errors.New(fmt.Sprint(ex))
						}
					}
				}
			}()
			if f.Kind() == reflect.Interface {
				f = f.Elem()
			}
			rets := f.Call(args)
			if isReflect {
				ev := rets[1].Interface()
				if ev != nil {
					err = ev.(error)
				}
				ret = rets[0].Interface().(reflect.Value)
			} else {
				for i, expr := range e.SubExprs {
					if ae, ok := expr.(*ast.AddrExpr); ok {
						if id, ok := ae.Expr.(*ast.IdentExpr); ok {
							invokeLetExpr(id, args[i].Elem().Elem(), env)
						}
					}
				}
				if f.Type().NumOut() == 1 {
					ret = rets[0]
				} else {
					var result []interface{}
					for _, r := range rets {
						result = append(result, r.Interface())
					}
					ret = reflect.ValueOf(result)
				}
			}
		}
		if e.Go {
			go fnc()
			return NilValue, nil
		}
		fnc()
		if err != nil {
			return ret, NewError(expr, err)
		}
		return ret, nil
	case *ast.TernaryOpExpr:
		rv, err := invokeExpr(e.Expr, env)
		if err != nil {
			return rv, NewError(expr, err)
		}
		if ast.ToBool(rv) {
			lhsV, err := invokeExpr(e.Lhs, env)
			if err != nil {
				return lhsV, NewError(expr, err)
			}
			return lhsV, nil
		}
		rhsV, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return rhsV, NewError(expr, err)
		}
		return rhsV, nil
	case *ast.TypeCast:
		// приведение типов, включая приведение типов в массиве как новый типизированный массив
		// убрать из стандартной библиотеки функции преобразования
		eType := e.Type
		if e.TypeExpr != nil {
			// создаем по имени типа
			ev, err := invokeExpr(e.TypeExpr, env)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			if ev.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Имя типа должно быть строкой")
			}
			eType = ast.UniqueNames.Set(ast.ToString(ev))
		}
		nt, err := env.Type(eType)
		if err != nil {
			return NilValue, err
		}
		rv, err := invokeExpr(e.CastExpr, env)
		if err != nil {
			return rv, NewError(expr, err)
		}

		return ast.TypeCastConvert(rv, nt, false, NilValue)

	case *ast.MakeExpr:
		eType := e.Type
		if e.TypeExpr != nil {
			// создаем по имени типа
			ev, err := invokeExpr(e.TypeExpr, env)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			if ev.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Имя типа должно быть строкой")
			}
			eType = ast.UniqueNames.Set(ast.ToString(ev))
		}
		rt, err := env.Type(eType)
		if err != nil {
			return NilValue, NewError(expr, err)
		}
		if rt.Kind() == reflect.Map {
			return reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt), nil
		}
		return reflect.Zero(rt), nil
	case *ast.MakeChanExpr:

		var size int
		if e.SizeExpr != nil {
			rv, err := invokeExpr(e.SizeExpr, env)
			if err != nil {
				return NilValue, err
			}
			size = int(ast.ToInt64(rv))
		}
		return func() (refval reflect.Value, err error) {
			defer func() {
				if os.Getenv("GONEC_DEBUG") == "" {
					if ex := recover(); ex != nil {
						if e, ok := ex.(error); ok {
							err = e
						} else {
							err = errors.New(fmt.Sprint(ex))
						}
					}
				}
			}()
			return reflect.ValueOf(make(chan interface{}, size)), nil
		}()
	case *ast.MakeArrayExpr:

		var alen int
		if e.LenExpr != nil {
			rv, err := invokeExpr(e.LenExpr, env)
			if err != nil {
				return NilValue, err
			}
			alen = int(ast.ToInt64(rv))
		}
		var acap int
		if e.CapExpr != nil {
			rv, err := invokeExpr(e.CapExpr, env)
			if err != nil {
				return NilValue, err
			}
			acap = int(ast.ToInt64(rv))
		} else {
			acap = alen
		}
		return func() (refval reflect.Value, err error) {
			defer func() {
				if os.Getenv("GONEC_DEBUG") == "" {
					if ex := recover(); ex != nil {
						if e, ok := ex.(error); ok {
							err = e
						} else {
							err = errors.New(fmt.Sprint(ex))
						}
					}
				}
			}()
			return reflect.ValueOf(make([]interface{}, alen, acap)), nil
		}()
	case *ast.ChanExpr:
		rhs, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return NilValue, NewError(expr, err)
		}

		if e.Lhs == nil {
			if rhs.Kind() == reflect.Chan {
				rv, _ := rhs.Recv()
				return rv, nil
			}
		} else {
			lhs, err := invokeExpr(e.Lhs, env)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			if lhs.Kind() == reflect.Chan {
				lhs.Send(rhs)
				return NilValue, nil
			} else if rhs.Kind() == reflect.Chan {
				rv, ok := rhs.Recv()
				if !ok {
					return NilValue, NewErrorf(expr, "Ошибка работы с каналом")
				}
				return invokeLetExpr(e.Lhs, rv, env)
			}
		}
		return NilValue, NewStringError(expr, "Неверная операция с каналом")
	default:
		return NilValue, NewStringError(expr, "Неизвестное выражение")
	}
}
