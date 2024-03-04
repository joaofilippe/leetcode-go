package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{123, 321},
		{-123, -321},
		{100, 1},
		{1, 1},
		{1534236469, 0},
	}

	for _, test := range tests {
		got := reverse(test.input)
		assert.Equal(t, test.output, got)
	}
}
