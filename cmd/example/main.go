package main

import (
	"flag"
	"fmt"
	lab2 "github.com/KatePril/architecture-lab2"
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
		os.Exit(1)
	}
	if *inputExpression != "" && *inputFile != "" {
		fmt.Println("You can't provide both -e and -f at the same time")
		os.Exit(1)
	}

	input, inputCloseCallback, _ := lab2.ProvideInput(*inputExpression, *inputFile)
	if input == nil {
		fmt.Println("Please provide a valid expression using -e flag or existing file with expression using -f flag")
		os.Exit(1)
	}
	output, outputCloseCallback, _ := lab2.ProvideOutput(*outputFile)

	handler := &lab2.ComputeHandler{Input: input, Output: output}

	err := handler.Compute()
	if inputCloseCallback != nil {
		finishError(inputCloseCallback())
	}
	if outputCloseCallback != nil {
		finishError(outputCloseCallback())
	}
	finishError(err)
}

func finishError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
