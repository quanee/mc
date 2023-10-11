package main

// Token 类型
type TokenType int

const (
	Integer TokenType = iota
	Plus
	Minus
	Multiply
	Divide
)

// Token 结构
type Token struct {
	Type  TokenType
	Value string
}
