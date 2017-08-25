// Package parser implements parser for gonec.
package parser

import (
	"errors"
	"fmt"
	"unicode"

	"strings"

	"github.com/covrom/gonec/ast"
	"github.com/covrom/gonec/bincode"
)

const (
	EOF = -1   // End of file.
	EOL = '\n' // End of line.
)

// Error provides a convenient interface for handling runtime error.
// It can be Error inteface with type cast which can call Pos().
type Error struct {
	Message  string
	Pos      ast.Position
	Filename string
	Fatal    bool
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

// Scanner stores informations for lexer.
type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
	canequal bool
	typecast bool
	castType string
}

// opName is correction of operation names.
var opName = map[string]int{
	"функция":           FUNC,
	"возврат":           RETURN,
	"перем":             VAR,
	"вызватьисключение": THROW,
	"если":              IF,
	"для":               FOR,
	"прервать":          BREAK,
	"продолжить":        CONTINUE,
	"из":                IN,
	"иначе":             ELSE,
	// "создать":           NEW,
	"истина":            TRUE,
	"ложь":              FALSE,
	"неопределено":      NIL,
	"модуль":            MODULE,
	"попытка":           TRY,
	"исключение":        CATCH,
	// "окончательно":      FINALLY,
	"выбор":  SWITCH,
	"когда":  CASE,
	"другое": DEFAULT,
	"старт":  GO,
	"канал":  CHAN,
	"новый":  MAKE,

	"или":          OROR,
	"и":            ANDAND,
	"не":           int('!'),
	"конеццикла":   int('}'),
	"конецесли":    int('}'),
	"конецфункции": int('}'),
	"конецпопытки": int('}'),
	"конецвыбора":  int('}'),
	"тогда":        int('{'),
	"цикл":         int('{'),
	"null":         NULL,
	"каждого":      EACH,
	"по":           TO,
	"пока":         WHILE,
	"иначеесли":    ELSIF,

	"строка":     TYPECAST,
	"число":      TYPECAST,
	"булево":     TYPECAST,
	"целоечисло": TYPECAST,
	"массив":     TYPECAST,
	"структура":  TYPECAST,
}

var opCanEqual = map[int]bool{
	RETURN: true,
	THROW:  true,
	IF:     true,
	// FOR:      true,
	IN:       true,
	// NEW:      true,
	TRUE:     true,
	FALSE:    true,
	NIL:      true,
	GO:       true,
	CHAN:     true,
	MAKE:     true,
	OROR:     true,
	ANDAND:   true,
	int('!'): true,
	NULL:     true,
	// EACH:     true,
	TO:    true,
	WHILE: true,
	ELSIF: true,
}

// Init resets code to scan.
func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

// Scan analyses token, and decide identify or literals.
func (s *Scanner) Scan() (tok int, lit string, pos ast.Position, err error) {
	if s.typecast {
		//вставляем название типа
		s.typecast = false
		tok = IDENT
		lit = s.castType
		pos = s.pos()
		err = nil
		return
	}
retry:
	s.skipBlank()
	pos = s.pos()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
		}
		lowlit := strings.ToLower(lit)
		if name, ok := opName[lowlit]; ok {
			tok = name
			_, s.canequal = opCanEqual[tok]
			if tok == TYPECAST {
				s.typecast = true
				s.castType = lowlit
			}
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		tok = NUMBER
		lit, err = s.scanNumber()
		if err != nil {
			return
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString('"')
		if err != nil {
			return
		}
	case ch == '\'':
		tok = STRING
		lit, err = s.scanString('\'')
		if err != nil {
			return
		}
	case ch == '`':
		tok = STRING
		lit, err = s.scanRawString()
		if err != nil {
			return
		}
	default:
		switch ch {
		case EOF:
			tok = EOF
		case '#':
			for !isEOL(s.peek()) {
				s.next()
			}
			goto retry
		case '!':
			s.next()
			switch s.peek() {
			case '=':
				tok = NEQ
				lit = "!="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			switch s.peek() {
			case '=':
				tok = EQEQ
				lit = "=="
			default:
				s.back()
				if s.canequal {
					tok = EQEQ
					lit = "=="
				} else {
					tok = int(ch)
					lit = string(ch)
					s.canequal = true
				}
			}
		case '+':
			s.next()
			switch s.peek() {
			case '+':
				tok = PLUSPLUS
				lit = "++"
			case '=':
				tok = PLUSEQ
				lit = "+="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '-':
			s.next()
			switch s.peek() {
			case '-':
				tok = MINUSMINUS
				lit = "--"
			case '=':
				tok = MINUSEQ
				lit = "-="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '*':
			s.next()
			switch s.peek() {
			case '*':
				tok = POW
				lit = "**"
			case '=':
				tok = MULEQ
				lit = "*="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '/':
			s.next()
			switch s.peek() {
			case '/':
				for !isEOL(s.peek()) {
					s.next()
				}
				goto retry
			case '=':
				tok = DIVEQ
				lit = "/="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			switch s.peek() {
			case '=':
				tok = GE
				lit = ">="
			case '>':
				tok = SHIFTRIGHT
				lit = ">>"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '<':
			s.next()
			switch s.peek() {
			case '-':
				tok = OPCHAN
				lit = "<-"
			case '=':
				tok = LE
				lit = "<="
			case '<':
				tok = SHIFTLEFT
				lit = "<<"
			case '>':
				tok = NEQ
				lit = "!="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '|':
			s.next()
			switch s.peek() {
			case '|':
				tok = OROR
				lit = "||"
			case '=':
				tok = OREQ
				lit = "|="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '&':
			s.next()
			switch s.peek() {
			case '&':
				tok = ANDAND
				lit = "&&"
			case '=':
				tok = ANDEQ
				lit = "&="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '.':
			s.next()
			if s.peek() == '.' {
				s.next()
				if s.peek() == '.' {
					tok = VARARG
				} else {
					err = fmt.Errorf(`syntax error "%s"`, "..")
					return
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '\n':
			tok = int(ch)
			lit = string(ch)
			//первое равенство в строке - это будет присваивание
			//нельзя переносить знак равенства в проверке на равенство в выражениях с присваиванием
			//можно поставить знак равенства в конце строки и только потом перенести строку
			s.canequal = false
		case ';':
			//смена оператора - меняем признак возможности сравнения
			tok = int(ch)
			lit = string(ch)
			s.canequal = false
		case '(':
			tok = int(ch)
			lit = string(ch)
		case ')', ']':
			tok = int(ch)
			lit = string(ch)
		case '{', '}':
			tok = int(ch)
			lit = string(ch)
			s.canequal = false
		case '?':
			s.next()
			switch s.peek() {
			case '(':
				tok = TERNARY
				lit = "?"
				s.canequal = true //присваивания внутри тернарного оператора не бывает
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case ':', '%', ',', '^':
			tok = int(ch)
			lit = string(ch)
		case '[':
			s.next()
			if s.peek() == ']' {
				s.next()
				if s.peek() == '(' {
					s.back()
					tok = ARRAYLIT
					lit = "[]"
				} else {
					s.back()
					s.back()
					tok = int(ch)
					lit = string(ch)
				}
			} else {
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		default:
			err = fmt.Errorf(`syntax error "%s"`, string(ch))
			tok = int(ch)
			lit = string(ch)
			return
		}
		s.next()
	}
	return
}

// isLetter returns true if the rune is a letter for identity.
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isDigit returns true if the rune is a number.
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isHex returns true if the rune is a hex digits.
func isHex(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

// isEOL returns true if the rune is at end-of-line or end-of-file.
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

// isBlank returns true if the rune is empty character..
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

// peek returns current rune in the code.
func (s *Scanner) peek() rune {
	if s.reachEOF() {
		return EOF
	}
	return s.src[s.offset]
}

// next moves offset to next.
func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

// current returns the current offset.
func (s *Scanner) current() int {
	return s.offset
}

// offset sets the offset value.
func (s *Scanner) set(o int) {
	s.offset = o
}

// back moves back offset once to top.
func (s *Scanner) back() {
	s.offset--
}

// reachEOF returns true if offset is at end-of-file.
func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

// pos returns the position of current.
func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

// skipBlank moves position into non-black character.
func (s *Scanner) skipBlank() {
	for isBlank(s.peek()) {
		s.next()
	}
}

// scanIdentifier returns identifier begining at current position.
func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

// scanNumber returns number begining at current position.
func (s *Scanner) scanNumber() (string, error) {
	var ret []rune
	ch := s.peek()
	ret = append(ret, ch)
	s.next()
	if ch == '0' && s.peek() == 'x' {
		ret = append(ret, s.peek())
		s.next()
		for isHex(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	} else {
		for isDigit(s.peek()) || s.peek() == '.' {
			ret = append(ret, s.peek())
			s.next()
		}
		if s.peek() == 'e' {
			ret = append(ret, s.peek())
			s.next()
			if isDigit(s.peek()) || s.peek() == '+' || s.peek() == '-' {
				ret = append(ret, s.peek())
				s.next()
				for isDigit(s.peek()) || s.peek() == '.' {
					ret = append(ret, s.peek())
					s.next()
				}
			}
			for isDigit(s.peek()) || s.peek() == '.' {
				ret = append(ret, s.peek())
				s.next()
			}
		}
		if isLetter(s.peek()) {
			return "", errors.New("identifier starts immediately after numeric literal")
		}
	}
	return string(ret), nil
}

// scanRawString returns raw-string starting at current position.
func (s *Scanner) scanRawString() (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("unexpected EOF")
			break
		}
		if s.peek() == '`' {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

// scanString returns string starting at current position.
// This handles backslash escaping.
func (s *Scanner) scanString(l rune) (string, error) {
	var ret []rune
eos:
	for {
		s.next()
		switch s.peek() {
		case EOL:
			return "", errors.New("unexpected EOL")
		case EOF:
			return "", errors.New("unexpected EOF")
		case l:
			s.next()
			break eos
		case '\\':
			s.next()
			switch s.peek() {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			}
			ret = append(ret, s.peek())
			continue
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

// Lexer provides inteface to parse codes.
type Lexer struct {
	s     *Scanner
	lit   string
	pos   ast.Position
	e     error
	stmts []ast.Stmt
}

// Lex scans the token and literals.
func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if err != nil {
		l.e = &Error{Message: fmt.Sprintf("%s", err.Error()), Pos: pos, Fatal: true}
	}
	lval.tok = ast.Token{Tok: tok, Lit: lit}
	lval.tok.SetPosition(pos)
	l.lit = lit
	l.pos = pos
	return tok
}

// Error sets parse error.
func (l *Lexer) Error(msg string) {
	l.e = &Error{Message: msg, Pos: l.pos, Fatal: false}
}

// Parser provides way to parse the code using Scanner.
func Parse(s *Scanner) ([]ast.Stmt, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.stmts, l.e
}

func EnableErrorVerbose() {
	yyErrorVerbose = true
}

// ParserSrc provides way to parse the code from source.
func ParseSrc(src string) ([]ast.Stmt, error) {
	scanner := &Scanner{
		src: []rune(src),
	}
	prs, err := Parse(scanner)
	if err != nil {
		return prs, err
	}
	// оптимизируем дерево AST
	// свертка констант и нативные значения
	prs = constFolding(prs)
	// fmt.Printf("%#v\n", prs[0].(*ast.LetsStmt).Rhss[0].(*ast.SliceExpr))
	lid := 0
	bin := bincode.BinaryCode(prs, 0, &lid)
	fmt.Printf("%v\n", bin)

	return prs, err
}
