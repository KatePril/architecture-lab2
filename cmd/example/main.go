package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
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

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PrefixToInfix("+ 2 2")
	fmt.Println(res)
}
