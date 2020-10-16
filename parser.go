package querystring

//go:generate goyacc -o querystring.y.go querystring.y

import (
	"fmt"
	"strings"
)

var debugLexer bool

// Parse querystring and return Condition
func Parse(query string) (rq Condition, err error) {
	if query == "" {
		return nil, nil
	}

	lex := newLexerWrapper(newConditionStringLex(strings.NewReader(query)))
	doParse(lex)

	if len(lex.errs) > 0 {
		return nil, fmt.Errorf(strings.Join(lex.errs, "\n"))
	}
	return lex.query, nil
}

func doParse(lex *lexerWrapper) {
	defer func() {
		r := recover()
		if r != nil {
			lex.errs = append(lex.errs, fmt.Sprintf("parse error: %v", r))
		}
	}()

	yyParse(lex)
}

const (
	queryShould = iota
	queryMust
	queryMustNot
)

type lexerWrapper struct {
	lex   yyLexer
	errs  []string
	query Condition
}

func newLexerWrapper(lex yyLexer) *lexerWrapper {
	return &lexerWrapper{
		lex:   lex,
		query: nil,
	}
}

func (l *lexerWrapper) Lex(lval *yySymType) int {
	return l.lex.Lex(lval)
}

func (l *lexerWrapper) Error(s string) {
	l.errs = append(l.errs, s)
}

func init() {
	yyErrorVerbose = true
}
