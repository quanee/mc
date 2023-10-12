package ast

import (
	"fmt"

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
func (ast *Node) ConvertToStack() ([]float64, []token.TokenType) {
	if ast == nil {
		return nil, nil
	}

	var operands []float64
	var operators []token.TokenType

	if ast.Type == Number {
		operands = append(operands, ast.Value)
	} else if ast.Type == BinaryOperation {
		leftOperands, leftOperators := ast.Left.ConvertToStack()
		rightOperands, rightOperators := ast.Right.ConvertToStack()

		operands = append(operands, leftOperands...)
		operands = append(operands, rightOperands...)
		operators = append(operators, leftOperators...)
		operators = append(operators, rightOperators...)
		operators = append(operators, ast.Operator)
	}

	return operands, operators
}

func evaluateRPN(operands []float64, operators []token.TokenType) float64 {
	stack := vm.Stack{}
	for i := 0; i < len(operators); i++ {
		if operators[i] == token.Plus || operators[i] == token.Minus || operators[i] == token.Multiply || operators[i] == token.Divide {
			rightOperand := stack.Pop()
			leftOperand := stack.Pop()
			result := 0.0

			switch operators[i] {
			case token.Plus:
				result = leftOperand + rightOperand
			case token.Minus:
				result = leftOperand - rightOperand
			case token.Multiply:
				result = leftOperand * rightOperand
			case token.Divide:
				if rightOperand == 0.0 {
					// 处理除零错误
					fmt.Println("除零错误")
					return 0.0
				}
				result = leftOperand / rightOperand
			}

			stack.Push(result)
		} else {
			stack.Push(operands[i])
		}
	}

	if len(operands) > len(operators) {
		stack.Push(operands[len(operands)-1])
	}

	return stack.Pop()
}
