package ast

import (
	"mc/token"
	"mc/vm"
)

// AST 节点类型
type NodeType int

const (
	Number NodeType = iota
	BinaryOperation
)

// AST 节点
type Node struct {
	Type NodeType

	// 如果节点类型为 Number
	Value float64

	// 如果节点类型为 BinaryOperation
	Left     *Node
	Operator token.TokenType
	Right    *Node
}

// add operator priority
func (ast *Node) ConvertToStack() []*Node {
	if ast == nil {
		return []*Node{}
	}

	if ast.Type == Number {
		return []*Node{ast}
	}

	leftRPN := ast.Left.ConvertToStack()
	rightRPN := ast.Right.ConvertToStack()

	rightRPN = append(rightRPN, ast)

	return append(leftRPN, rightRPN...)

}

func EvaluateRPN(operands []float64, operators []token.TokenType) float64 {
	stack := vm.Stack{}

	stack.Push(operands[0])

	return stack.Pop()
}
