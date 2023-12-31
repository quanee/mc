package main

import (
	"fmt"

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
	optnum := astree.ConvertToStack()
	for _, opt := range optnum {
		fmt.Printf("expression result: %#v\n", opt)
	}

	//fmt.Printf("expression result: %v\n", ast.EvaluateRPN(optnum))
}
