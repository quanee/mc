package token

import "fmt"

// Token type
type TokenType int

const (
	Number TokenType = iota
	Variable
	String
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
		return "[Toke] Number"
	case Variable:
		return "[Toke] Variable"
	case Plus:
		return "[Toke] Plus[+]"
	case Minus:
		return "[Toke] Minus[-]"
	case Multiply:
		return "[Toke] Multiply[*]"
	case Divide:
		return "[Toke] Divide[/]"
	case Lparen:
		return "[Toke] Lparen[(]"
	case Rparen:
		return "[Toke] Rparen[)]"
	default:
		return fmt.Sprintf("unknow %d", typ)
	}
}
