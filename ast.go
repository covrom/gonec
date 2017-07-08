package gonec

import (
	"bytes"
)

//узел - основная единица AST
type pNode interface {
	//представление токена в исходном коде, после обработки лексером
	tLit() string
	//преобразование в строку
	String() string
}

//команды
type pStmt interface {
	pNode
	stmtNode()
}

//выражения
type pExpr interface {
	pNode
	exprNode()
}

//определения
type pDecl interface {
	pNode
	declNode()
}

//главная нода программы
type pProgram struct {
	stmts []pStmt
}

func (p *pProgram) tLit() string {
	if len(p.stmts) > 0 {
		//литерал первого токена в программе
		return p.stmts[0].tLit()
	}
	return ""
}

func (p *pProgram) String() string {
	//все представления всех команд программы - листинг AST дерева
	var out bytes.Buffer
	for _, s := range p.stmts {
		out.WriteString(s.String())
	}
	return out.String()
}

//присвоение
type pLetStatement struct {
	tok  token
	name *pIdentifier
	val  pExpr
}

func (ls *pLetStatement) stmtNode()    {}
func (ls *pLetStatement) tLit() string { return ls.tok.literal }
func (ls *pLetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.name.String())
	out.WriteString(" = ")
	if ls.val != nil {
		out.WriteString(ls.val.String())
	}
	out.WriteString(";")
	return out.String()
}

//идентификатор переменной - всегда строка
type pIdentifier struct {
	tok token
	val string
}

func (i *pIdentifier) exprNode()      {}
func (i *pIdentifier) tLit() string   { return i.tok.literal }
func (i *pIdentifier) String() string { return i.val }
