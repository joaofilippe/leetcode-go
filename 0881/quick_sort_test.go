package numrescueboats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{10, -1, 2, 5, 0}, []int{-1, 0, 2, 5, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, quickSort(test.input))
	}
}
