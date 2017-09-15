package ast

import (
	"reflect"

	"github.com/covrom/gonec/bincode/binstmt"
	"github.com/covrom/gonec/builtins"
	"github.com/covrom/gonec/env"
	"github.com/covrom/gonec/pos"
)

// Stmt provides all of interfaces for statement.
type Stmt interface {
	pos.Pos
	stmt()
	Simplify()
	BinTo(*binstmt.BinStmts, int, *int)
}

// StmtImpl provide commonly implementations for Stmt..
type StmtImpl struct {
	pos.PosImpl // StmtImpl provide Pos() function.
}

// stmt provide restraint interface.
func (x *StmtImpl) stmt() {}

type Stmts []Stmt

func (x Stmts) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	for _, st := range x {
		st.BinTo(bins, reg, lid)
	}
}

// NoneStmt используется для пропуска блоков кода, например, Else
type NoneStmt struct {
	StmtImpl
}

func (x *NoneStmt) Simplify()                                       {}
func (s *NoneStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {}

// ExprStmt provide expression statement.
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

func (x *ExprStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
}

func (s *ExprStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	s.Expr.BinTo(bins, reg, lid, true)
	// *bins = append(*bins, addBinExpr(s.Expr, reg, lid, true)...)
}

// IfStmt provide "if/else" statement.
type IfStmt struct {
	StmtImpl
	If     Expr
	Then   Stmts
	ElseIf Stmts // This is array of IfStmt
	Else   Stmts
}

func (x *IfStmt) Simplify() {
	x.If = x.If.Simplify()
	for _, st := range x.Then {
		st.Simplify()
	}
	for _, st := range x.ElseIf {
		st.Simplify()
	}
	for _, st := range x.Else {
		st.Simplify()
	}
}

func (s *IfStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	*lid++
	lend := *lid

	// Если
	s.If.BinTo(bins, reg, lid, false)

	*lid++
	lf := *lid

	bins.Append(binstmt.NewBinJFALSE(reg, lf, s))

	// Тогда
	s.Then.BinTo(bins, reg, lid)

	bins.Append(binstmt.NewBinJMP(lend, s))

	// ИначеЕсли
	bins.Append(binstmt.NewBinLABEL(lf, s))

	for _, elif := range s.ElseIf {
		stmtif := elif.(*IfStmt)

		stmtif.If.BinTo(bins, reg, lid, false)

		// если ложь, то перейдем на следующее условие
		*lid++
		li := *lid

		bins.Append(binstmt.NewBinJFALSE(reg, li, stmtif))

		stmtif.Then.BinTo(bins, reg, lid)

		bins.Append(binstmt.NewBinJMP(lend, stmtif))

		bins.Append(binstmt.NewBinLABEL(li, stmtif))
	}

	// Иначе
	if len(s.Else) > 0 {
		s.Else.BinTo(bins, reg, lid)
	}
	// КонецЕсли
	bins.Append(binstmt.NewBinLABEL(lend, s))

	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))
}

// TryStmt provide "try/catch/finally" statement.
type TryStmt struct {
	StmtImpl
	Try Stmts
	// Var     string
	Catch Stmts
	// Finally Stmts
}

func (x *TryStmt) Simplify() {
	for _, st := range x.Try {
		st.Simplify()
	}
	for _, st := range x.Catch {
		st.Simplify()
	}
}

func (s *TryStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	*lid++
	lend := *lid
	*lid++
	li := *lid
	// эта инструкция сообщает, в каком регистре будет отслеживаться ошибка выполнения кода до блока CATCH
	// по-умолчанию, ошибка в регистрах не отслеживается, а передается по уровням исполнения вирт. машины
	bins.Append(binstmt.NewBinTRY(reg, li, s))

	s.Try.BinTo(bins, reg+1, lid) // чтобы не затереть регистр с ошибкой, увеличиваем номер

	// сюда переходим, если в блоке выше возникла ошибка
	bins.Append(binstmt.NewBinLABEL(li, s))

	// CATCH работает как JFALSE, и определяет функцию ОписаниеОшибки()
	bins.Append(binstmt.NewBinCATCH(reg, lend, s))

	// тело обработки ошибки
	s.Catch.BinTo(bins, reg, lid) // регистр с ошибкой больше не нужен, текст определен функцией

	bins.Append(binstmt.NewBinLABEL(lend, s))
	// КонецПопытки

	// снимаем со стека состояние обработки ошибок, чтобы последующий код не был включен в текущую обработку
	bins.Append(binstmt.NewBinPOPTRY(li, s))

	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))
}

// ForStmt provide "for in" expression statement.
type ForStmt struct {
	StmtImpl
	Var   int //string
	Value Expr
	Stmts Stmts
}

func (x *ForStmt) Simplify() {
	x.Value = x.Value.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

func (s *ForStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	// для каждого
	s.Value.BinTo(bins, reg, lid, false)

	*lid++
	lend := *lid
	*lid++
	li := *lid

	regiter := reg + 1
	regval := reg + 2
	regsub := reg + 3
	// инициализируем итератор, параметры цикла и цикл в стеке циклов
	bins.Append(binstmt.NewBinFOREACH(reg, regiter, lend, li, s))

	// очередная итерация
	// сюда же переходим по Продолжить
	bins.Append(binstmt.NewBinLABEL(li, s))

	bins.Append(binstmt.NewBinNEXT(reg, regiter, regval, lend, s))

	// устанавливаем переменную-итератор
	bins.Append(binstmt.NewBinSET(regval, s.Var, s))

	s.Stmts.BinTo(bins, regsub, lid)

	// повторяем итерацию
	bins.Append(binstmt.NewBinJMP(li, s))

	// КонецЦикла
	bins.Append(binstmt.NewBinLABEL(lend, s))

	// снимаем со стека наличие цикла для Прервать и Продолжить
	bins.Append(binstmt.NewBinPOPFOR(li, s))

	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))
}

// NumForStmt name = expr1 to expr2
type NumForStmt struct {
	StmtImpl
	Name  int //string
	Expr1 Expr
	Expr2 Expr
	Stmts Stmts
}

func (x *NumForStmt) Simplify() {
	x.Expr1 = x.Expr1.Simplify()
	x.Expr2 = x.Expr2.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

func (s *NumForStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	// для .. по ..
	regfrom := reg + 1
	regto := reg + 2
	regsub := reg + 3

	s.Expr1.BinTo(bins, regfrom, lid, false)
	s.Expr2.BinTo(bins, regto, lid, false)

	*lid++
	lend := *lid
	*lid++
	li := *lid

	// инициализируем итератор, параметры цикла и цикл в стеке циклов
	bins.Append(binstmt.NewBinFORNUM(reg, regfrom, regto, lend, li, s))

	// очередная итерация
	// сюда же переходим по Продолжить
	bins.Append(binstmt.NewBinLABEL(li, s))

	bins.Append(binstmt.NewBinNEXTNUM(reg, regfrom, regto, lend, s))

	// устанавливаем переменную-итератор
	bins.Append(binstmt.NewBinSET(reg, s.Name, s))

	s.Stmts.BinTo(bins, regsub, lid)
	// повторяем итерацию
	bins.Append(binstmt.NewBinJMP(li, s))

	// КонецЦикла
	bins.Append(binstmt.NewBinLABEL(lend, s))

	// снимаем со стека наличие цикла для Прервать и Продолжить
	bins.Append(binstmt.NewBinPOPFOR(li, s))

	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))

}

// CForStmt provide C-style "for (;;)" expression statement.
// type CForStmt struct {
// 	StmtImpl
// 	Expr1 Expr
// 	Expr2 Expr
// 	Expr3 Expr
// 	Stmts Stmts
// }

// LoopStmt provide "for expr" expression statement.
type LoopStmt struct {
	StmtImpl
	Expr  Expr
	Stmts Stmts
}

func (x *LoopStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

func (s *LoopStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	*lid++
	lend := *lid
	*lid++
	li := *lid
	bins.Append(binstmt.NewBinWHILE(lend, li, s))

	// очередная итерация
	// сюда же переходим по Продолжить
	bins.Append(binstmt.NewBinLABEL(li, s))

	s.Expr.BinTo(bins, reg, lid, false)

	bins.Append(binstmt.NewBinJFALSE(reg, lend, s))

	// тело цикла
	s.Stmts.BinTo(bins, reg+1, lid)

	// повторяем итерацию
	bins.Append(binstmt.NewBinJMP(li, s))

	// КонецЦикла
	bins.Append(binstmt.NewBinLABEL(lend, s))

	// снимаем со стека наличие цикла для Прервать и Продолжить
	bins.Append(binstmt.NewBinPOPFOR(li, s))

	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))

}

// BreakStmt provide "break" expression statement.
type BreakStmt struct {
	StmtImpl
}

func (x *BreakStmt) Simplify() {}

func (s *BreakStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	bins.Append(binstmt.NewBinBREAK(s))
}

// ContinueStmt provide "continue" expression statement.
type ContinueStmt struct {
	StmtImpl
}

func (x *ContinueStmt) Simplify() {}

func (s *ContinueStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	bins.Append(binstmt.NewBinCONTINUE(s))
}

// ForStmt provide "return" expression statement.
type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

func (x *ReturnStmt) Simplify() {
	for i := range x.Exprs {
		x.Exprs[i] = x.Exprs[i].Simplify()
	}
}

func (s *ReturnStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {

	if len(s.Exprs) == 0 {
		bins.Append(binstmt.NewBinLOAD(reg, nil, false, s))
	}
	if len(s.Exprs) == 1 {
		// одиночное значение в reg
		s.Exprs[0].BinTo(bins, reg, lid, false)
	} else {
		// создание слайса в reg
		bins.Append(binstmt.NewBinMAKESLICE(reg, len(s.Exprs), len(s.Exprs), s))

		for i, ee := range s.Exprs {
			ee.BinTo(bins, reg+1, lid, false)
			bins.Append(binstmt.NewBinSETIDX(reg, i, reg+1, ee))
		}
	}
	// в reg имеем значение или структуру возврата
	bins.Append(binstmt.NewBinFREE(reg+1, s))
	bins.Append(binstmt.NewBinRET(reg, s))
}

// ThrowStmt provide "throw" expression statement.
type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

func (x *ThrowStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
}

func (s *ThrowStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	s.Expr.BinTo(bins, reg, lid, false)
	bins.Append(binstmt.NewBinTHROW(reg, s))
}

// ModuleStmt provide "module" expression statement.
type ModuleStmt struct {
	StmtImpl
	Name  int //string
	Stmts Stmts
}

func (x *ModuleStmt) Simplify() {
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

func (s *ModuleStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	if s.Name == env.UniqueNames.Set("_") {
		// добавляем все операторы в текущий контекст
		s.Stmts.BinTo(bins, reg, lid)
	} else {
		bins.Append(binstmt.NewBinMODULE(s.Name, BinaryCode(s.Stmts, 0, lid), s))
	}
}

// SwitchStmt provide switch statement.
type SwitchStmt struct {
	StmtImpl
	Expr  Expr
	Cases Stmts
}

func (x *SwitchStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
	for _, st := range x.Cases {
		st.Simplify()
	}
}

func (s *SwitchStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	s.Expr.BinTo(bins, reg, lid, true)
	// сравниваем с каждым case
	*lid++
	lend := *lid
	var default_stmt *DefaultStmt
	for _, ss := range s.Cases {
		if ssd, ok := ss.(*DefaultStmt); ok {
			default_stmt = ssd
			continue
		}
		*lid++
		li := *lid
		case_stmt := ss.(*CaseStmt)
		case_stmt.Expr.BinTo(bins, reg+1, lid, false)
		bins.Append(binstmt.NewBinEQUAL(reg+2, reg, reg+1, case_stmt))
		bins.Append(binstmt.NewBinJFALSE(reg+2, li, case_stmt))
		case_stmt.Stmts.BinTo(bins, reg, lid)
		bins.Append(binstmt.NewBinJMP(lend, case_stmt))
		bins.Append(binstmt.NewBinLABEL(li, case_stmt))
	}
	if default_stmt != nil {
		default_stmt.Stmts.BinTo(bins, reg, lid)
	}
	bins.Append(binstmt.NewBinLABEL(lend, s))
	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))
}

// SelectStmt provide switch statement.
type SelectStmt struct {
	StmtImpl
	Cases Stmts
}

func (x *SelectStmt) Simplify() {
	for _, st := range x.Cases {
		st.Simplify()
	}
}

func (s *SelectStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	*lid++
	lstart := *lid
	bins.Append(binstmt.NewBinLABEL(lstart, s))

	*lid++
	lend := *lid
	var default_stmt *DefaultStmt
	for _, ss := range s.Cases {
		if ssd, ok := ss.(*DefaultStmt); ok {
			default_stmt = ssd
			continue
		}
		*lid++
		li := *lid
		case_stmt := ss.(*CaseStmt)
		e, ok := case_stmt.Expr.(*ChanExpr)
		if !ok {
			panic(binstmt.NewStringError(case_stmt, "При выборе вариантов из каналов допустимы только выражения с каналами"))
		}
		// определяем значение справа
		e.Rhs.BinTo(bins, reg, lid, false)
		if e.Lhs == nil {
			// слева нет значения - это временное чтение из канала без сохранения значения в переменной
			bins.Append(binstmt.NewBinTRYRECV(reg, reg+1, reg+2, reg+3, e.Rhs))
			// если канал закрыт или не получено значение - идем в следующую ветку
			bins.Append(binstmt.NewBinJFALSE(reg+2, li, s))
		} else {
			// значение слева
			e.Lhs.BinTo(bins, reg+1, lid, false)

			// проверяем: слева канал?
			bins.Append(binstmt.NewBinMV(reg+1, reg+3, e))
			bins.Append(binstmt.NewBinISKIND(reg+3, reflect.Chan, e))

			*lid++
			li3 := *lid

			bins.Append(binstmt.NewBinJFALSE(reg+3, li3, e))

			// слева канал - пишем в него правое
			bins.Append(binstmt.NewBinTRYSEND(reg+1, reg, reg+2, e.Lhs))

			*lid++
			li2 := *lid

			// если отправлено значение - выполняем код блока
			bins.Append(binstmt.NewBinJTRUE(reg+2, li2, s))

			// если не отправлено значение - идем в следующую ветку
			// если канал закрыт - будет паника
			bins.Append(binstmt.NewBinJMP(li, s))

			// иначе справа канал, а слева переменная (установим, если прочитали из канала)
			bins.Append(binstmt.NewBinLABEL(li3, s))

			bins.Append(binstmt.NewBinTRYRECV(reg, reg+1, reg+2, reg+3, e.Rhs))

			// если канал закрыт или не получено значение - идем в следующую ветку
			bins.Append(binstmt.NewBinJFALSE(reg+2, li, s))

			// устанавливаем переменную прочитанным значением
			e.Lhs.(CanLetExpr).BinLetTo(bins, reg+1, lid)

			bins.Append(binstmt.NewBinLABEL(li2, s))
		}
		// отправили или прочитали - выполняем ветку кода и выходим из цикла
		case_stmt.Stmts.BinTo(bins, reg, lid)

		// выходим из цикла
		bins.Append(binstmt.NewBinJMP(lend, case_stmt))

		// к следующему case
		bins.Append(binstmt.NewBinLABEL(li, s))
	}
	// если ни одна из веток не сработала - проверяем default
	if default_stmt != nil {
		default_stmt.Stmts.BinTo(bins, reg, lid)
	} else {
		// допускаем обработку других горутин
		bins.Append(binstmt.NewBinGOSHED(s))
		bins.Append(binstmt.NewBinJMP(lstart, s))
	}
	bins.Append(binstmt.NewBinLABEL(lend, s))
	// освобождаем память
	bins.Append(binstmt.NewBinFREE(reg+1, s))
}

// CaseStmt provide switch/case statement.
type CaseStmt struct {
	StmtImpl
	Expr  Expr
	Stmts Stmts
}

func (x *CaseStmt) Simplify() {
	x.Expr = x.Expr.Simplify()
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

func (s *CaseStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	//ничего не делаем, эти блоки обрабатываются в родительских контекстах
}

// DefaultStmt provide switch/default statement.
type DefaultStmt struct {
	StmtImpl
	Stmts Stmts
}

func (x *DefaultStmt) Simplify() {
	for _, st := range x.Stmts {
		st.Simplify()
	}
}

func (s *DefaultStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	//ничего не делаем, эти блоки обрабатываются в родительских контекстах
}

// LetsStmt provide multiple statement of let.
type LetsStmt struct {
	StmtImpl
	Lhss     []Expr
	Operator string
	Rhss     []Expr
}

func (x *LetsStmt) Simplify() {
	for i := range x.Lhss {
		x.Lhss[i] = x.Lhss[i].Simplify()
	}
	for i := range x.Rhss {
		x.Rhss[i] = x.Rhss[i].Simplify()
	}
}

func (s *LetsStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	// если справа одно выражение - присваиваем его всем левым
	// и если там массив, то по очереди элементы, начиная с 0-го
	// иначе с обеих сторон должно быть одинаковое число выражений, они попарно присваиваются
	if len(s.Rhss) == 1 && len(s.Lhss) > 1 {
		s.Rhss[0].BinTo(bins, reg, lid, false)
		// проверяем на массив
		*lid++
		lend := *lid
		*lid++
		li := *lid
		bins.Append(binstmt.NewBinISSLICE(reg, reg+1, s))
		bins.Append(binstmt.NewBinJFALSE(reg+1, li, s))

		// присваиваем из слайса
		i := 0
		for _, e := range s.Lhss {
			// в рег+1 сохраним очередной элемент
			bins.Append(binstmt.NewBinMV(reg, reg+1, e))
			bins.Append(binstmt.NewBinLOAD(reg+2, core.VMInt(i), false, e))
			bins.Append(binstmt.NewBinGETIDX(reg+1, reg+2, e))
			e.(CanLetExpr).BinLetTo(bins, reg+1, lid)
			i++
		}
		bins.Append(binstmt.NewBinJMP(lend, s))

		// присваиваем одно и то же значение
		bins.Append(binstmt.NewBinLABEL(li, s))
		for _, e := range s.Lhss {
			e.(CanLetExpr).BinLetTo(bins, reg, lid)
		}
		bins.Append(binstmt.NewBinLABEL(lend, s))
	} else {
		if len(s.Lhss) == len(s.Rhss) {
			// сначала все вычисляем в разные регистры, затем все присваиваем
			// так обеспечиваем взаимный обмен
			for i := range s.Rhss {
				s.Rhss[i].BinTo(reg+i, lid, false)
			}
			for i, e := range s.Lhss {
				e.(CanLetExpr).BinLetTo(bins, reg+i, lid)
			}
		} else {
			// ошибка
			panic(binstmt.NewStringError(s, "Количество переменных и значений должно совпадать или значение должно быть одно"))
		}
	}
}

// VarStmt provide statement to let variables in current scope.
type VarStmt struct {
	StmtImpl
	Names []int //string
	Exprs []Expr
}

func (x *VarStmt) Simplify() {
	for i := range x.Exprs {
		x.Exprs[i] = x.Exprs[i].Simplify()
	}
}

func (s *VarStmt) BinTo(bins *binstmt.BinStmts, reg int, lid *int) {
	// если справа одно выражение - присваиваем его всем левым
	// иначе с обеих сторон должно быть одинаковое число выражений, они попарно присваиваются
	if len(s.Exprs) == 1 {
		s.Exprs[0].BinTo(bins, reg, lid, false)
		for _, e := range s.Names {
			bins.Append(binstmt.NewBinSET(reg, e, s))
		}
	} else {
		if len(s.Exprs) == len(s.Names) {
			for i, e := range s.Exprs {
				e.BinTo(bins, reg, lid, false)
				bins.Append(binstmt.NewBinSET(reg, s.Names[i], s))
			}
		} else {
			// ошибка
			panic(binstmt.NewStringError(s, "Количество переменных и значений должно совпадать или значение должно быть одно"))
		}
	}
}
