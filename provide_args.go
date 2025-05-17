package lab2

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ProvideOutput(outputDest string) (io.Writer, func() error, error) {
	if outputDest == "" {
		return os.Stdout, nil, nil
	}
	file, err := os.Create(outputDest)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating output file: %w", err)
	}

	return file, file.Close, nil
}

func ProvideInput(inputExpression, inputFile string) (io.Reader, func() error, error) {
	if inputExpression != "" {
		return strings.NewReader(inputExpression), nil, nil
	}

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening input file: %w", err)
	}

	return file, file.Close, nil
}
