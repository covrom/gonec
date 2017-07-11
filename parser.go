package gonec

import (
	"errors"
	"fmt"
	"io"
)

type parser struct {
	l []token //lexer

	peekpos   int   //позиция следующего токена
	curToken  token //текущий токен
	peekToken token //следующий токен

	//prefixParseFns map[token.TokenType]prefixParseFn
	//infixParseFns  map[token.TokenType]infixParseFn
}

func newParser(lex []token) *parser {
	p := &parser{
		l: lex,
	}
	// читаем curToken и peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *parser) nextToken() {
	if p.curToken.category&defEOF != 0 {
		return
	}
	p.curToken = p.peekToken
	if p.peekpos < len(p.l) {
		p.peekToken = p.l[p.peekpos]
		p.peekpos++
	} else {
		p.peekToken = token{}
	}
}

func (p *parser) curTokenCatIs(c uint) bool {
	return p.curToken.category == c
}
func (p *parser) curTokenTypeIs(t rune) bool {
	return p.curToken.toktype == t
}
func (p *parser) peekTokenCatIs(c uint) bool {
	return p.peekToken.category == c
}
func (p *parser) peekTokenTypeIs(t rune) bool {
	return p.peekToken.toktype == t
}
func (p *parser) curTokenIs(c uint, t rune) bool {
	return p.curTokenCatIs(c) && p.curTokenTypeIs(t)
}
func (p *parser) peekTokenIs(c uint, t rune) bool {
	return p.peekTokenCatIs(c) && p.peekTokenTypeIs(t)
}

func (p *parser) expectCatPeek(c uint) bool {
	if p.peekTokenCatIs(c) {
		p.nextToken()
		return true
	}
	return false
}

func (p *parser) parseProgram() (prog pProgram, err error) {
	prog = pProgram{stmts: []pStmt{}, errors: []string{}}
	var stmt pStmt
	for p.curToken.category&defEOF == 0 {
		stmt, err = p.parseStatement()
		if err != nil {
			prog.errors = append(prog.errors, fmt.Sprintf("Ошибка %d : %d : %v", p.curToken.srcline, p.curToken.srccol, err))
			err = errors.New("Обнаружены ошибки синтаксиса")
		} else {
			prog.stmts = append(prog.stmts, stmt)
		}
		p.nextToken()
	}
	return
}

// Parse формирует AST - абстрактное дерево исполнения
func (i *interpreter) Parse(tokens []token, w io.Writer) (pProgram, error) {
	//парсер
	return newParser(tokens).parseProgram()
}

func (p *parser) parseStatement() (stmt pStmt, err error) {

	//если в начале не идет ключевая конструкция "для" и т.п., то
	//собираем оператор и смотрим, есть ли в нем присваивание
	//если есть, будет два операнда
	//если нет, то есть только правый операнд

	switch p.curToken.category {
	case defIdentifier:
		if p.peekTokenCatIs(defOperator) {
			switch p.peekToken.toktype {
			case oLBr:
				//вызов метода
			case oLSqBr:
				//присвоение [элементу массива]
			case oPoint:
				//точка

			default:
				return nil, errors.New("Неизвестная операция над идентификатором переменной")
			}
		} else if p.expectCatPeek(defAssignator) {
			//присвоение без []
			return p.parseLetStatement()
		}
	case defPoint:
		
	case defEOF:
		return nil, nil
	}
	return nil, fmt.Errorf("Неизвестная синтаксическая конструкция %s", p.curToken.literal)
}

func (p *parser) parseLetStatement() (stmt *pLetStatement, err error) {
	stmt = &pLetStatement{
		tok: p.curToken,
	}
	stmt.name = &pIdentifier{tok: p.curToken, val: p.curToken.literal}
	if p.peekToken.category&catAssignable == 0 || p.peekToken.toktype == oLabelStart {
		return nil, errors.New("Недопустимый аргумент присваивания")
	}
	//выражением считаем все, что идет до конца файла или строки (;)

	//stmt.val = p.parseExpression(LOWEST)

	for p.curToken.category&catEndExpression == 0 {
		p.nextToken()
	}

	return stmt, nil
}
