package validmountainarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input    []int
	expected bool
}

var tests = []Test{
	{[]int{2, 1}, false},
	{[]int{3, 5, 5}, false},
	{[]int{0, 3, 2, 1}, true},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, false},
	{[]int{0, 2, 3, 4, 5, 2, 1, 0}, true},
}

func Test_validMountainArray(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, validMountainArray(test.input), "Failed for input %v", test.input)
	}
}
