package minlength

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input string
	expected int
}

var tests = []Test{
	{"", 0},
	{"ABFCACDB", 2},
	Test{"ACBBD", 5},
}

func Test_minLen(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, minLength(test.input), fmt.Sprintf("Failed for %s", test.input))
	}
}