package parser

import (
	"fmt"
	"strconv"
	"strings"

	"mc/ast"
	"mc/token"
)

type Parser struct {
	tokens []token.Token
}

func NewParser() Parser {
	return Parser{tokens: make([]token.Token, 0)}
}

// parse expression to Token list
func (p *Parser) ParseExpression(expression string) error {
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
			p.tokens = append(p.tokens, token.Token{token.Number, expression[start:i]})
			i--
		case char == '+':
			p.tokens = append(p.tokens, token.Token{token.Plus, "+"})
		case char == '-':
			p.tokens = append(p.tokens, token.Token{token.Minus, "-"})
		case char == '*':
			p.tokens = append(p.tokens, token.Token{token.Multiply, "*"})
		case char == '/':
			p.tokens = append(p.tokens, token.Token{token.Divide, "/"})
		}
	}

	return nil
}

// build AST
func (p *Parser) BuildAST() (*ast.Node, error) {
	return p.parseExpressionRecursive()
}

func (p *Parser) parseExpressionRecursive() (*ast.Node, error) {
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

		node = &ast.Node{
			Type:     ast.BinaryOperation,
			Left:     node,
			Operator: operator.Type,
			Right:    rightNode,
		}
	}

	return node, nil
}

func (p *Parser) parseTerm() (*ast.Node, error) {
	if len(p.tokens) == 0 {
		return nil, nil
	}

	tok := p.tokens[0]
	p.tokens = p.tokens[1:]

	if tok.Type == token.Number {
		value, err := strconv.ParseFloat(tok.Value, 64)
		if err != nil {
			return nil, err
		}
		return &ast.Node{Type: ast.Number, Value: value}, nil
	}

	return nil, fmt.Errorf("invalid express")
}

func (p *Parser) parseBinaryOperator() (token.Token, error) {
	if len(p.tokens) == 0 {
		return token.Token{}, fmt.Errorf("unexcepted EOF")
	}

	tok := p.tokens[0]
	p.tokens = p.tokens[1:]

	if tok.Type == token.Plus || tok.Type == token.Minus || tok.Type == token.Multiply || tok.Type == token.Divide {
		return tok, nil
	}

	return token.Token{}, fmt.Errorf("unexcepted operator: %v", tok.Type)
}
