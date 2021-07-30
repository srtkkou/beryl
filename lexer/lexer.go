package lexer

import (
	"unicode/utf8"

	"github.com/srtkkou/garnet/token"
)

type Lexer struct {
	input   string
	pos     int
	nextPos int
	letter  rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readLetter()
	return l
}

func (l *Lexer) readLetter() {
	if l.nextPos >= len(l.input) {
		l.letter = '0'
	} else {
		runes := []rune(l.input)
		l.letter = runes[l.nextPos]
	}
	l.pos = l.nextPos
	l.nextPos += utf8.RuneLen(l.letter)
}

func (l *Lexer) peekLetter() rune {
	if l.nextPos >= len(l.input) {
		return '0'
	}
	runes := []rune(l.input)
	return runes[l.nextPos]
}

func (l *Lexer) NextToken() (t token.Token) {
	switch l.letter {
	case ' ':
		t = newToken(token.Space, l.letter)
	case '\n':
		t = newToken(token.LF, l.letter)
	case '#':
		if l.peekLetter() == '=' {
			nowLetter := l.letter
			l.readLetter()
			literal := string(nowLetter) + string(l.letter)
			t = token.Token{
				Type:    token.PrintComment,
				Literal: literal,
			}
		} else {
			t = newToken(token.Comment, l.letter)
		}
	case '.':
		t = newToken(token.Period, l.letter)
	case '|':
		t = newToken(token.Block, l.letter)
	case '[':
		t = newToken(token.LBracket, l.letter)
	case ']':
		t = newToken(token.RBracket, l.letter)
	case '"':
		t = newToken(token.Quote, l.letter)
	case '=':
		t = newToken(token.Assign, l.letter)
	case '-':
		t = newToken(token.Code, l.letter)
	case '/':
		t = newToken(token.SingleTagMark, l.letter)
		//	case '\t': TODO: ERROR
		//		t = newToken(token.Tab, l.ch)
	case 0:
		t = newToken(token.EOF, 'e')
		t.Literal = ""
	default:
		if isValidLetter(l.letter) {
			t.Literal = l.readIdentifier()
			t.Type = token.Tag
			return t
		} else {
			t = newToken(token.Illegal, l.letter)
		}
	}
	l.readLetter()
	return t
}

func newToken(tokenType token.TokenType, r rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(r),
	}
}

func (l *Lexer) readIdentifier() string {
	firstPos := l.pos
	for isValidLetter(l.letter) {
		l.readLetter()
	}
	return l.input[firstPos:l.pos]
}

func isValidLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}
