package main

import "fmt"

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
	Operator TokenType
	Right    *Node
}

// 执行计算
func evaluateAST(node *Node) float64 {
	if node == nil {
		return 0.0
	}

	if node.Type == Number {
		return node.Value
	}

	leftValue := evaluateAST(node.Left)
	rightValue := evaluateAST(node.Right)

	switch node.Operator {
	case Plus:
		return leftValue + rightValue
	case Minus:
		return leftValue - rightValue
	case Multiply:
		return leftValue * rightValue
	case Divide:
		if rightValue == 0.0 {
			// 处理除零错误
			fmt.Println("除零错误")
			return 0.0
		}
		return leftValue / rightValue
	default:
		// 处理未知运算符
		fmt.Println("未知运算符")
		return 0.0
	}
}

// 添加运算符优先级
func evaluateASTWithPriority(node *Node) float64 {
	if node == nil {
		return 0.0
	}

	if node.Type == Number {
		fmt.Println(node.Value)
		return node.Value
	}

	//fmt.Println(node.Operator)
	// 首先处理乘法和除法
	if node.Operator == Multiply || node.Operator == Divide {
		leftValue := evaluateASTWithPriority(node.Left)
		rightValue := evaluateASTWithPriority(node.Right)

		switch node.Operator {
		case Multiply:
			return leftValue * rightValue
		case Divide:
			if rightValue == 0.0 {
				// 处理除零错误
				fmt.Println("除零错误")
				return 0.0
			}
			return leftValue / rightValue
		}
	}

	// 处理加法和减法
	if node.Operator == Plus || node.Operator == Minus {
		leftValue := evaluateASTWithPriority(node.Left)
		rightValue := evaluateASTWithPriority(node.Right)

		switch node.Operator {
		case Plus:
			return leftValue + rightValue
		case Minus:
			return leftValue - rightValue
		}
	}

	// 处理未知运算符
	fmt.Println("未知运算符")
	return 0.0
}
