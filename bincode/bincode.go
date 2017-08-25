package bincode

import (
	"strings"

	"github.com/covrom/gonec/ast"
)

///////////////////////////////////////////////////////////////
// компиляция в байткод
///////////////////////////////////////////////////////////////

func BinaryCode(inast []ast.Stmt, reg int, lid *int) (outast BinCode) {
	for _, st := range inast {
		// перебираем все подвыражения и команды, и выстраиваем их в линию
		// если в команде есть выражение - определяем новый id регистра, присваиваем ему выражение, а в команду передаем id этого регистра
		switch s := st.(type) {
		case *ast.ExprStmt:
			outast = append(outast, addBinExpr(s.Expr, reg, lid)...)

		}
	}
	return
}

func appendBin(bins BinCode, b BinStmt, e ast.Expr) BinCode {
	b.SetPosition(e.Position())
	return append(bins, b)
}

func addBinExpr(expr ast.Expr, reg int, lid *int) (bins BinCode) {
	switch e := expr.(type) {
	case *ast.NativeExpr:
		// добавляем команду загрузки значения
		bins = appendBin(bins,
			&BinLOAD{
				Reg: reg, // основной регистр
				Val: e.Value.Interface(),
			}, e)
	case *ast.NumberExpr:
		// команда на загрузку строки в регистр и ее преобразование в число, в регистр
		bins = appendBin(bins,
			&BinLOAD{
				Reg: reg,
				Val: e.Lit,
			}, e)

		bins = appendBin(bins,
			&BinCASTNUM{
				Reg: reg,
			}, e)
	case *ast.StringExpr:
		bins = appendBin(bins,
			&BinLOAD{
				Reg: reg,
				Val: e.Lit,
			}, e)
	case *ast.ConstExpr:
		b := BinLOAD{
			Reg: reg,
		}
		switch strings.ToLower(e.Value) {
		case "истина":
			b.Val = true
		case "ложь":
			b.Val = false
		case "null":
			b.Val = ast.NullVar

		default:
			b.Val = nil

		}
		bins = appendBin(bins, &b, e)
	case *ast.ArrayExpr:
		// создание слайса
		bins = appendBin(bins,
			&BinMAKESLICE{
				Reg: reg,
				Len: len(e.Exprs),
				Cap: len(e.Exprs),
			}, e)

		for i, ee := range e.Exprs {
			// каждое выражение сохраняем в следующем по номеру регистре (относительно регистра слайса)
			bins = append(bins, addBinExpr(ee, reg+1, lid)...)
			bins = appendBin(bins,
				&BinSETIDX{
					Reg:    reg,
					Index:  i,
					ValReg: reg + 1,
				}, ee)
		}
	case *ast.MapExpr:
		// создание мапы
		bins = appendBin(bins,
			&BinMAKEMAP{
				Reg: reg,
				Len: len(e.MapExpr),
			}, e)

		for k, ee := range e.MapExpr {
			bins = append(bins, addBinExpr(ee, reg+1, lid)...)
			bins = appendBin(bins,
				&BinSETKEY{
					Reg:    reg,
					Key:    k,
					ValReg: reg + 1,
				}, ee)
		}
	case *ast.IdentExpr:
		bins = appendBin(bins,
			&BinGET{
				Reg:    reg,
				Id:     e.Id,
				Dotted: strings.Contains(e.Lit, "."),
			}, e)
	case *ast.UnaryExpr:
		bins = append(bins, addBinExpr(e.Expr, reg, lid)...)
		bins = appendBin(bins,
			&BinUNARY{
				Reg: reg,
				Op:  rune(e.Operator[0]),
			}, e)
	case *ast.AddrExpr:
		bins = append(bins, addBinExpr(e.Expr, reg, lid)...)
		bins = appendBin(bins,
			&BinADDR{
				Reg: reg,
			}, e)
	case *ast.DerefExpr:
		bins = append(bins, addBinExpr(e.Expr, reg, lid)...)
		bins = appendBin(bins,
			&BinUNREF{
				Reg: reg,
			}, e)
	case *ast.ParenExpr:
		bins = append(bins, addBinExpr(e.SubExpr, reg, lid)...)
	case *ast.BinOpExpr:
		oper := OperMap[e.Operator]
		// сначала вычисляем левую часть
		bins = append(bins, addBinExpr(e.Lhs, reg, lid)...)
		switch oper {
		case LOR:
			*lid++
			lab := *lid
			// вставляем проверку на истину слева и возвращаем ее, не вычисляя правую часть, иначе возвращаем правую часть
			bins = appendBin(bins,
				&BinJTRUE{
					Reg:    reg,
					JumpTo: lab,
				}, e)
			bins = append(bins, addBinExpr(e.Rhs, reg, lid)...)
			bins = appendBin(bins,
				&BinLABEL{
					Label: lab,
				}, e)
		case LAND:
			*lid++
			lab := *lid
			// вставляем проверку на ложь слева и возвращаем ее, не вычисляя правую часть, иначе возвращаем правую часть
			bins = appendBin(bins,
				&BinJFALSE{
					Reg:    reg,
					JumpTo: lab,
				}, e)
			bins = append(bins, addBinExpr(e.Rhs, reg, lid)...)
			bins = appendBin(bins,
				&BinLABEL{
					Label: lab,
				}, e)
		default:
			bins = append(bins, addBinExpr(e.Rhs, reg+1, lid)...)
			bins = appendBin(bins,
				&BinOPER{
					RegL: reg,
					RegR: reg + 1,
					Op:   oper,
				}, e)
		}
	case *ast.TernaryOpExpr:
		bins = append(bins, addBinExpr(e.Expr, reg, lid)...)
		*lid++
		lab := *lid
		bins = appendBin(bins,
			&BinJFALSE{
				Reg:    reg,
				JumpTo: lab,
			}, e)
		// если истина - берем левое выражение
		bins = append(bins, addBinExpr(e.Lhs, reg, lid)...)
		// прыгаем в конец
		*lid++
		lend := *lid
		bins = appendBin(bins,
			&BinJMP{
				JumpTo: lend,
			}, e)

		// правое выражение
		bins = appendBin(bins,
			&BinLABEL{
				Label: lab,
			}, e)
		bins = append(bins, addBinExpr(e.Rhs, reg, lid)...)
		bins = appendBin(bins,
			&BinLABEL{
				Label: lend,
			}, e)

	case *ast.CallExpr:
		// помещаем аргументы
		// либо в серию регистров, начиная с reg, если их <=7
		// либо в массив аргументов в reg
		if len(e.SubExprs) <= 7 {
			for i := 0; i < len(e.SubExprs); i++ {
				bins = append(bins, addBinExpr(e.SubExprs[i], reg+i, lid)...)
			}
		} else {
			bins = appendBin(bins,
				&BinMAKESLICE{
					Reg: reg,
					Len: len(e.SubExprs),
					Cap: len(e.SubExprs),
				}, e)

			for i, ee := range e.SubExprs {
				// каждое выражение сохраняем в следующем по номеру регистре (относительно регистра слайса)
				bins = append(bins, addBinExpr(ee, reg+1, lid)...)
				bins = appendBin(bins,
					&BinSETIDX{
						Reg:    reg,
						Index:  i,
						ValReg: reg + 1,
					}, ee)
			}
		}
		bins = appendBin(bins,
			&BinCALL{
				Name:    e.Name,
				NumArgs: len(e.SubExprs),
				RegArgs: reg,
				VarArg:  e.VarArg,
				Go:      e.Go,
			}, e)

	case *ast.AnonCallExpr:

	case *ast.MemberExpr:

	case *ast.ItemExpr:

	case *ast.SliceExpr:

	case *ast.FuncExpr:

	case *ast.LetExpr:

	case *ast.LetsExpr:

	case *ast.AssocExpr:

	case *ast.NewExpr:

	case *ast.ChanExpr:

	case *ast.TypeCast:

	case *ast.MakeExpr:

	case *ast.MakeChanExpr:

	case *ast.MakeArrayExpr:

	}

	return
}
