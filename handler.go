package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with Input io.Reader and Output io.Writer.
// Its Compute() method should read the expression from Input and write the computed result to the Output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Input)
	if !scanner.Scan() {
		return fmt.Errorf("Input Error")
	}

	expression := scanner.Text()
	transformedExpression, err1 := PrefixToInfix(expression)
	if err1 != nil {
		return err1
	}

	_, err2 := fmt.Fprint(ch.Output, transformedExpression)
	return err2
}
