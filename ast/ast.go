package ast

import (
	"fmt"

	"mc/token"
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
func EvaluateASTWithPriority(node *Node) float64 {
	if node == nil {
		return 0.0
	}

	if node.Type == Number {
		fmt.Println(node.Value)
		return node.Value
	}

	// first process multiply adn divide
	if node.Operator == token.Multiply || node.Operator == token.Divide {
		leftValue := EvaluateASTWithPriority(node.Left)
		rightValue := EvaluateASTWithPriority(node.Right)

		fmt.Println(node.Operator)
		switch node.Operator {
		case token.Multiply:
			return leftValue * rightValue
		case token.Divide:
			if rightValue == 0.0 {
				// process division by zero error
				fmt.Println("division by zero")
				return 0.0
			}
			return leftValue / rightValue
		}
	}

	// process plus and minus
	if node.Operator == token.Plus || node.Operator == token.Minus {
		leftValue := EvaluateASTWithPriority(node.Left)
		rightValue := EvaluateASTWithPriority(node.Right)

		fmt.Println(node.Operator)
		switch node.Operator {
		case token.Plus:
			return leftValue + rightValue
		case token.Minus:
			return leftValue - rightValue
		}
	}

	// 处理未知运算符
	fmt.Println("未知运算符")
	return 0.0
}
