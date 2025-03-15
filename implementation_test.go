package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfixSimple1(t *testing.T) {
	res, err := PrefixToInfix("- 1 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "1 - 2", res)
	}
}

func TestPrefixToInfixSimple2(t *testing.T) {
	res, err := PrefixToInfix("^ 1 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "1 ^ 3", res)
	}
}

func TestPrefixToInfixSimple3(t *testing.T) {
	res, err := PrefixToInfix("+ - 4 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 - 5 + 8", res)
	}
}

func TestPrefixToInfixSimple4(t *testing.T) {
	res, err := PrefixToInfix("- 4 + 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 - (5 + 8)", res)
	}
}

func TestPrefixToInfixSimple5(t *testing.T) {
	res, err := PrefixToInfix("* - 4 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "(4 - 5) * 8", res)
	}
}

func TestPrefixToInfixSimple6(t *testing.T) {
	res, err := PrefixToInfix("- * 4 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 * 5 - 8", res)
	}
}

func TestPrefixToInfixSimple7(t *testing.T) {
	res, err := PrefixToInfix("- 4 * 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 - 5 * 8", res)
	}
}

func TestPrefixToInfixSimple8(t *testing.T) {
	res, err := PrefixToInfix("* 4 - 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 * (5 - 8)", res)
	}
}

func TestPrefixToInfixSimple9(t *testing.T) {
	res, err := PrefixToInfix("- 4 / 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 - 5 / 8", res)
	}
}

func TestPrefixToInfixSimple10(t *testing.T) {
	res, err := PrefixToInfix("/ 4 - 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 / (5 - 8)", res)
	}
}

func TestPrefixToInfixSimple11(t *testing.T) {
	res, err := PrefixToInfix("^ / 4 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "(4 / 5) ^ 8", res)
	}
}

func TestPrefixToInfixSimple12(t *testing.T) {
	res, err := PrefixToInfix("^ 4 / 5 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 ^ (5 / 8)", res)
	}
}

func TestPrefixToInfix(t *testing.T) {
	res, err := PrefixToInfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "5 + (4 - 2) * 3", res)
	}
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 1 2")
	fmt.Println(res)

	// Output:
	// 1 + 2
}
