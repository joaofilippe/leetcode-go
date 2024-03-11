package longest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		}, {
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			[]string{"reflower", "flow", "flight"},
			"",
		},
	}

	for _, test := range tests {
		got := longestCommonPrefix(test.input)
		assert.Equal(t, test.expected, got)
	}
}
