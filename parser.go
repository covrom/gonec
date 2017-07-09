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
	switch p.curToken.category {
	case defIdentifier:
		if p.peekToken.category == defOperator {
			switch p.peekToken.toktype {
			case oEq:
				//присвоение без []
				return p.parseLetStatement()
			case oLBr:
				//вызов метода
			case oLSqBr:
				//присвоение [элементу массива]

			default:
				return nil, errors.New("Неизвестная операция над идентификатором переменной")
			}
		} else {

		}
	case defEOF:
		return nil, nil
	}
	return nil, fmt.Errorf("Неизвестная синтаксическая конструкция %s", p.curToken.literal)
}

func (p *parser) parseLetStatement() (stmt *pLetStatement, err error) {
	stmt = &pLetStatement{tok: p.curToken}
	stmt.name = &pIdentifier{tok: p.curToken, val: p.curToken.literal}
	if p.peekToken.category&catAssignable == 0 || p.peekToken.toktype == oLabelStart {
		return nil, errors.New("Недопустимый аргумент присваивания")
	}
	//выражением считаем все, что идет до конца файла или строки (;)
	for p.curToken.category&catEndExpression == 0 {
		p.nextToken()
	}

	return stmt, nil
}
