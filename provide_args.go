package lab2

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ProvideOutput(outputDest string) (io.Writer, io.Closer, error) {
	if outputDest == "" {
		return os.Stdout, nil, nil
	}
	file, err := os.Create(outputDest)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating output file: %w", err)
	}

	return file, file, nil
}

func ProvideInput(inputExpression, inputFile string) (io.Reader, io.Closer, error) {

	if inputExpression != "" {
		return strings.NewReader(inputExpression), nil, nil
	}

	if fileExists(inputFile) {
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, nil, fmt.Errorf("error opening input file: %w", err)
		}
		return file, file, nil
	}
	return nil, nil, fmt.Errorf("input file not found")

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
