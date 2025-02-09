package movezeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{[]int{1, 0, 2, 0, 3, 0, 4, 0, 5}, []int{1, 2, 3, 4, 5, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		moveZeroes(test.input)
		assert.Equal(t, test.expected, test.input)
	}
}
