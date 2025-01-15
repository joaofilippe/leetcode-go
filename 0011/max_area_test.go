package maxarea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	input    []int
	expected int
}

var tests = []Test{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
	{[]int{4, 3, 2, 1, 4}, 16},
	{[]int{1, 2, 1}, 2},
}

func Test_maxArea(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, maxArea(test.input), "Failed for input %v", test.input)
	}
}
