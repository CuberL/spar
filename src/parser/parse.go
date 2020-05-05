package parser

import (
	"io"

	"github.com/CuberL/spar/src/lexer"
	"github.com/CuberL/spar/src/types"
)

// Parse returns parsed Spanner DDL statements.
func Parse(r io.Reader) (*types.DDStatements, error) {
	impl := lexer.NewLexerImpl(r, &keywordTokenizer{})
	l := newLexerWrapper(impl)

	yyParse(l)

	if l.err != nil {
		return nil, l.err
	} else {
		return &l.result, nil
	}
}
