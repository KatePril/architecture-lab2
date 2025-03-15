package lab2

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	input := strings.NewReader("+ 5 * - 4 2 ^ 3 2")
	output := new(bytes.Buffer)

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	assert.Nil(t, err)
	assert.Equal(t, "5 + 4 - 2 * 3 ^ 2", output.String())
}

func TestCompute_WrongSyntax(t *testing.T) {
	input := strings.NewReader("5 + 4 - 2 * 3 ^ 2")
	output := new(bytes.Buffer)

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	assert.NotNil(t, err)
}

func TestCompute_EmptyString(t *testing.T) {
	input := strings.NewReader("")
	output := new(bytes.Buffer)

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()

	fmt.Println(err)
	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("input error"))
}
