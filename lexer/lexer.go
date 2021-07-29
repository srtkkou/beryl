package lexer

import (
	"github.com/srtkkou/garnet/token"
)

type Lexer struct {
	input   string
	pos     int
	nextPos int
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}
	l.pos = l.nextPos
	l.nextPos++
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	}
	return l.input[l.nextPos]
}

func (l *Lexer) NextToken() (t token.Token) {
	switch l.ch {
	case ' ':
		t = newToken(token.Space, l.ch)
	case '\n':
		t = newToken(token.LF, l.ch)
	case '#':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{
				Type:    token.PrintComment,
				Literal: literal,
			}
		} else {
			t = newToken(token.Comment, l.ch)
		}
	case '.':
		t = newToken(token.Period, l.ch)
	case '|':
		t = newToken(token.Block, l.ch)
	case '[':
		t = newToken(token.LBracket, l.ch)
	case ']':
		t = newToken(token.RBracket, l.ch)
	case '"':
		t = newToken(token.Quote, l.ch)
	case '=':
		t = newToken(token.Assign, l.ch)
	case '-':
		t = newToken(token.Code, l.ch)
	case '/':
		t = newToken(token.SingleTagMark, l.ch)
		//	case '\t': TODO: ERROR
		//		t = newToken(token.Tab, l.ch)
	case 0:
		t = newToken(token.EOF, 'e')
		t.Literal = ""
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.Tag
			return t
		} else {
			t = newToken(token.Illegal, l.ch)
		}
	}
	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) readIdentifier() string {
	firstPos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[firstPos:l.pos]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
