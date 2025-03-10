package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.input)
	if !scanner.Scan() {
		return fmt.Errorf("Input Error")
	}

	expression := scanner.Text()
	transformedExpression, err1 := PrefixToInfix(expression)
	if err1 != nil {
		return err1
	}

	_, err2 := fmt.Fprint(ch.output, transformedExpression)
	return err2
}
