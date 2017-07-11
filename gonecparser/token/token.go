// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

import "strconv"
import "strings"

// Token is the set of lexical tokens of the Go programming language.
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT // main
	// INT    // 12345
	NUM // 123.45
	// IMAG   // 123.45i
	DATE   // 'a'
	STRING // "abc"

	literal_end

	operator_beg
	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =

	NEQ // !=
	LEQ // <=
	GEQ // >=
	// DEFINE   // :=
	// ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	LABEL 	  // ~

	operator_end

	keyword_beg
	// Keywords
	BREAK
	CASE
	CHAN
	// CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	// FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	// INTERFACE
	MAP
	PACKAGE
	// RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR

	EXPORT
	THEN
	ELSIF
	ENDIF
	EACH
	IN
	TO
	WHILE
	DO
	ENDDO
	PROC
	ENDFUNC
	ENDPROC
	TRY
	ENDTRY
	EXCEPT
	RAISE
	NEW
	NULL
	UNDEF
	LAND  // &&
	LOR   // ||
	NOT    // !

	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT: "IDENT",
	NUM:   "число",
	// FLOAT:  "FLOAT",
	// IMAG:   "IMAG",
	DATE:   "дата",
	STRING: "строка",
	NULL:   "null",
	UNDEF:  "неопределено",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND:  "и",
	LOR:   "или",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "=",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "не",

	NEQ: "<>",
	LEQ: "<=",
	GEQ: ">=",
	// DEFINE:   ":=",
	// ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
	LABEL:     "~",

	BREAK: "прервать",
	CASE:  "когда",
	CHAN:  "канал",
	// CONST:    "const",
	CONTINUE: "продолжить",

	DEFAULT: "другой",
	DEFER:   "позже",
	ELSE:    "иначе",
	// FALLTHROUGH: "fallthrough",
	FOR: "для",

	FUNC:   "функция",
	GO:     "поток",
	GOTO:   "перейти",
	IF:     "если",
	IMPORT: "импорт",

	// INTERFACE: "interface",
	MAP:     "соответствие",
	PACKAGE: "пакет",
	// RANGE:     "range",
	RETURN: "возврат",

	SELECT: "переключить",
	STRUCT: "структура",
	SWITCH: "выбор",
	TYPE:   "тип",
	VAR:    "перем",

	EXPORT:  "экспорт",
	THEN:    "тогда",
	ELSIF:   "иначеесли",
	ENDIF:   "конецесли",
	EACH:    "каждого",
	IN:      "из",
	TO:      "по",
	WHILE:   "пока",
	DO:      "цикл",
	ENDDO:   "конеццикла",
	PROC:    "процедура",
	ENDFUNC: "конецфункции",
	ENDPROC: "конецпроцедуры",
	TRY:     "попытка",
	ENDTRY:  "конецпопытки",
	EXCEPT:  "исключение",
	RAISE:   "вызватьисключение",
	NEW:     "новый",
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (op Token) Precedence() int {
	switch op {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
		return 5
	}
	return LowestPrec
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
//
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[strings.ToLower(ident)]; is_keyword {
		return tok
	}
	return IDENT
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok Token) IsLiteral() bool { return literal_beg < tok && tok < literal_end }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok Token) IsOperator() bool { return operator_beg < tok && tok < operator_end }

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func (tok Token) IsKeyword() bool { return keyword_beg < tok && tok < keyword_end }
