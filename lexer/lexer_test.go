package lexer

import (
	"testing"

	"github.com/srtkkou/beryl/token"
)

func TestNextToken(t *testing.T) {
	input := "html\n" +
		"  body\n" +
		"    div\n"

	tests := []struct {
		exType    token.TokenType
		exLiteral string
	}{
		{token.Tag, "html"},
		{token.LF, "\n"},
		{token.Indent, "  "},
		{token.Tag, "body"},
		{token.LF, "\n"},
		{token.Indent, "  "},
		{token.Indent, "  "},
		{token.Tag, "div"},
		{token.LF, "\n"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.exType {
			t.Fatalf("tests[%d] - token type error. ex=%q, got=%q",
				i, tt.exType, tok.Type)
		}
		if tok.Literal != tt.exLiteral {
			t.Fatalf("tests[%d] - literal error. ex=%q, got=%q",
				i, tt.exLiteral, tok.Literal)
		}
	}
}
