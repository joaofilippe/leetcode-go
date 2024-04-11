package validparentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Input    string
	Expected bool
}

func Test_isValid(t *testing.T) {
	tests := []Test{
		{Input: "()", Expected: true},
		{Input: "()[]{}", Expected: true},
		{Input: "(]", Expected: false},
		{Input: "([)]", Expected: false},
		{Input: "{[]}", Expected: true},
	}

	for _, test := range tests {
		if isValid(test.Input) != test.Expected {
			valid := isValid(test.Input)
			assert.Equal(t, test.Expected, valid, fmt.Sprintf("Input: %s", test.Input))
		}
	}
}
