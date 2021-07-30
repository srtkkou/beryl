package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	//FilePath string
	//Line int
	//Column int
}

const (
	Illegal      = "Illegal"
	EOF          = "EOF"
	Space        = " "
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
	VarMark       = "$"
)
