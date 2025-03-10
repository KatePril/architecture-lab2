package lab2

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ProvideOutput(outputDest string) (io.Writer, error) {
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

func ProvideInput(inputExpression, inputFile string) (io.Reader, error) {
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
