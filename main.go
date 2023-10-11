package main

import (
	"fmt"

	"mc/ast"
	"mc/parser"
)

func main() {
	expression := "3 + 4 * 2 - 6 / 2"
	parser := parser.NewParser()
	if err := parser.ParseExpression(expression); err != nil {
		fmt.Println(err)
		return
	}

	astree, err := parser.BuildAST()
	if err != nil {
		fmt.Println("build AST error:", err)
		return
	}

	fmt.Printf("%#v", astree)
	result := ast.EvaluateASTWithPriority(astree)
	fmt.Printf("expression result: %f\n", result)
}
