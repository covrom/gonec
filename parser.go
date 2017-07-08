package gonec

import (
	"io"
)

type parser struct {
	l      []token //lexer
	errors []string

	peekpos   int   //позиция следующего токена
	curToken  token //текущий токен
	peekToken token //следующий токен

	//prefixParseFns map[token.TokenType]prefixParseFn
	//infixParseFns  map[token.TokenType]infixParseFn
}

func newParser(l []token) *parser {
	p := &parser{l: l}
	// читаем curToken и peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *parser) nextToken() {
	p.curToken = p.peekToken
	if p.peekpos < len(p.l) {
		p.peekToken = p.l[p.peekpos]
		p.peekpos++
	} else {
		p.peekToken = token{}
	}
}

// Parse формирует AST - абстрактное дерево исполнения
func (i *interpreter) Parse(tokens []token, w io.Writer) (err error) {

	//парсер
	

	return nil
}
