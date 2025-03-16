package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluatePrefixExpression(t *testing.T) {
	cases := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{expression: "+ 2 3", expected: 5, expectErr: false},
		{expression: "- 10 4", expected: 6, expectErr: false},
		{expression: "* 3 4", expected: 12, expectErr: false},
		{expression: "/ 10 2", expected: 5, expectErr: false},
		{expression: "^ 2 3", expected: 8, expectErr: false},
		{expression: "/ 10 0", expected: 0, expectErr: true},
		{expression: "", expected: 0, expectErr: true},
		{expression: "* + 2 3 4", expected: 20, expectErr: false},
		{expression: "+ * 2 3 * 4 5", expected: 26, expectErr: false},
	}

	for _, c := range cases {
		result, err := EvaluatePrefixExpression(c.expression)
		if c.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expected, result)
		}
	}
}

func ExampleEvaluatePrefixExpression() {
	res, _ := EvaluatePrefixExpression("+ 2 3")
	fmt.Println(res)
	// Output:
	// 5
}
