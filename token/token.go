package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "Illegal"
	EOF     = "EOF"
	Space   = " "
	Tab     = "\t"
	//Indent       = "  "
	LF           = "\n"
	Comment      = "#"
	PrintComment = "#+"
	Period       = "."
	Block        = "|"
	LBracket     = "["
	RBracket     = "]"
	Quote        = `"`
	Assign       = "="
	Code         = "-"
	//PrintCode     = "+"
	SingleTagMark = "/"
	Tag           = "Tag"
)
