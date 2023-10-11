package main

import (
	"fmt"
)

func main() {
	expression := "3 + 4 * 2 - 6 / 2"
	parser := &Parser{tokens: make([]Token, 0)}
	if err := parser.parseExpression(expression); err != nil {
		fmt.Println(err)
		return
	}

	ast, err := parser.buildAST()
	if err != nil {
		fmt.Println("build AST error:", err)
		return
	}

	result := evaluateASTWithPriority(ast)
	fmt.Printf("expression result: %f\n", result)
}
