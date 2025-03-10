package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
	"os"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", ".txt file with expression")
	outputFile      = flag.String("o", "", ".txt file with output")
)

func main() {
	flag.Parse()
	if *inputExpression == "" && *inputFile == "" {
		fmt.Println("Please provide either expression using -e flag or file with expression using -f flag")
	}
	if *inputExpression != "" && *inputFile != "" {
		fmt.Println("You can't provide both -e and -f at the same time")
	}

	input, _ := lab2.ProvideInput(*inputExpression, *inputFile)
	output, _ := lab2.ProvideOutput(*outputFile)

	handler := &lab2.ComputeHandler{Input: input, Output: output}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
