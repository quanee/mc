package token

import "fmt"

// Token type
type TokenType int

const (
	Number TokenType = iota
	Variable
	Plus
	Minus
	Multiply
	Divide
	Lparen // (
	Rparen // )
)

// Token define
type Token struct {
	Type  TokenType
	Value string
}

func (typ TokenType) String() string {
	switch typ {
	case Number:
		return "[Toke Type] Number"
	case Variable:
		return "[Toke Type] Variable"
	case Plus:
		return "[Toke Type] Plus[+]"
	case Minus:
		return "[Toke Type] Minus[-]"
	case Multiply:
		return "[Toke Type] Multiply[*]"
	case Divide:
		return "[Toke Type] Divide[/]"
	case Lparen:
		return "[Toke Type] Lparen[(]"
	case Rparen:
		return "[Toke Type] Divide"
	default:
		return fmt.Sprintf("unknow %d", typ)
	}
}
