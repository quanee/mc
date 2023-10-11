package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	tokens []Token
}

func NewParser() Parser {
	return Parser{tokens: make([]Token, 0)}
}

// 解析表达式为 Token 列表
func (p *Parser) parseExpression(expression string) error {
	expression = strings.ReplaceAll(expression, " ", "") // 移除空格

	for i := 0; i < len(expression); i++ {
		char := expression[i]
		switch {
		case char >= '0' && char <= '9':
			// parse int
			start := i
			for i < len(expression) && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.') {
				i++
			}
			p.tokens = append(p.tokens, Token{Integer, expression[start:i]})
			i--
		case char == '+':
			p.tokens = append(p.tokens, Token{Plus, "+"})
		case char == '-':
			p.tokens = append(p.tokens, Token{Minus, "-"})
		case char == '*':
			p.tokens = append(p.tokens, Token{Multiply, "*"})
		case char == '/':
			p.tokens = append(p.tokens, Token{Divide, "/"})
		}
	}

	return nil
}

// build AST
func (p *Parser) buildAST() (*Node, error) {
	return p.parseExpressionRecursive()
}

func (p *Parser) parseExpressionRecursive() (*Node, error) {
	if len(p.tokens) == 0 {
		return nil, nil
	}

	node, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for len(p.tokens) > 0 {
		operator, err := p.parseBinaryOperator()
		if err != nil {
			return nil, err
		}

		rightNode, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		node = &Node{
			Type:     BinaryOperation,
			Left:     node,
			Operator: operator.Type,
			Right:    rightNode,
		}
	}

	return node, nil
}

func (p *Parser) parseTerm() (*Node, error) {
	if len(p.tokens) == 0 {
		return nil, nil
	}

	token := p.tokens[0]
	p.tokens = p.tokens[1:]

	if token.Type == Integer {
		value, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return nil, err
		}
		return &Node{Type: Number, Value: value}, nil
	}

	return nil, fmt.Errorf("invalid express")
}

func (p *Parser) parseBinaryOperator() (Token, error) {
	if len(p.tokens) == 0 {
		return Token{}, fmt.Errorf("unexcepted EOF")
	}

	token := p.tokens[0]
	p.tokens = p.tokens[1:]

	if token.Type == Plus || token.Type == Minus || token.Type == Multiply || token.Type == Divide {
		return token, nil
	}

	return Token{}, fmt.Errorf("unexcepted operator: %v", token.Type)
}
