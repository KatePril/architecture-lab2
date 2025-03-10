package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", ".txt file with expression")
	outputFile      = flag.String("o", "", ".txt file with output")
)

func provideOutput(outputDest string) (io.Writer, error) {
	var output io.Writer
	if outputDest == "" {
		output = os.Stdout
	} else {
		file, err := os.Create(outputDest)
		if err != nil {
			return nil, fmt.Errorf("error creating output file: %w", err)
		}
		output = file
	}

	return output, nil
}

func provideInput(inputExpression, inputFile string) (io.Reader, error) {
	var input io.Reader

	if inputExpression != "" {
		input = strings.NewReader(inputExpression)
		return input, nil
	}

	if fileExists(inputFile) {
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, fmt.Errorf("error opening input file: %w", err)
		}
		input = file
	} else {
		input = strings.NewReader(inputFile)
	}
	return input, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	flag.Parse()
	if *inputExpression == "" && *inputFile == "" {
		fmt.Println("Please provide either expression using -e flag or file with expression using -f flag")
	}
	if *inputExpression != "" && *inputFile != "" {
		fmt.Println("You can't provide both -e and -f at the same time")
	}

	input, _ := provideInput(*inputExpression, *inputFile)
	output, _ := provideOutput(*outputFile)

	handler := &lab2.ComputeHandler{Input: input, Output: output}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
