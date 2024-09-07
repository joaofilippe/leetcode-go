package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestInput struct {
	Nums     []int
	Val      int
	Expected int
}

func Test_RemoveElement(t *testing.T) {
	tests := []TestInput{
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			3,
		},
		{
			[]int{3, 2, 2, 3},
			2,
			2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, RemoveElement(test.Nums, test.Val))
	}
}
