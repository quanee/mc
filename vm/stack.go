package vm

import (
	"fmt"
	"mc/token"
)

type Stack struct {
	data []float64
}

func (s *Stack) Push(value float64) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() float64 {
	if len(s.data) == 0 {
		return 0.0 // 错误处理
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value
}

func evaluateRPN(operands []float64, operators []token.TokenType) float64 {
	stack := Stack{}
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
