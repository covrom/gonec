package gonec

import (
	"errors"
	"io"
	"strings"

	"github.com/covrom/gonec/gonecscan"
)

func (i *interpreter) Lexer(r io.Reader, w io.Writer) (tokens []token, err error) {
	
	//лексический анализ
	
	var s gonecscan.Scanner

	s.Error = func(s *gonecscan.Scanner, msg string) {
		err = errors.New(msg)
	}

	s.Init(r)

	var tok rune

	for tok != gonecscan.EOF {
		tok = s.Scan()
		if err != nil {
			return
		}

		nt := token{literal: s.TokenText()}
		ntlit := strings.ToLower(nt.literal)
		var ok bool
		switch tok {
		case gonecscan.Ident:
			nt.toktype, ok = keywordMap[ntlit]
			if !ok {
				nt.category = defIdentifier
			} else {
				nt.category = defKeyword
			}
		case gonecscan.String:
			// строки возвращаются без переносов и комментариев
			nt.category = defValueString
		case gonecscan.Int:
			nt.category = defValueInt
		case gonecscan.Float:
			nt.category = defValueFloat
		case gonecscan.Date:
			nt.category = defValueDate
		default:
			nt.toktype, ok = operMap[ntlit]
			if !ok {
				nt.toktype, ok = delimMap[ntlit]
				if !ok {
					nt.toktype, ok = pointMap[ntlit]
					if !ok {
						nt.category = defUnknown
					} else {
						nt.category = defPoint
					}
				} else {
					nt.category = defDelimiter
				}
			} else {
				nt.category = defOperator
			}
		}

		tokens = append(tokens, nt)

	}

	return nil, nil
}
