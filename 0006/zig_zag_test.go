package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ZigZag(t *testing.T) {
	tests := []struct {
		Input   string
		Output  string
		NumRows int
	}{
		{
			"PAYPALISHIRING",
			"PAHNAPLSIIGYIR",
			3,
		},
		{
			"PAYPALISHIRING",
			"PINALSIGYAHRPI",
			4,
		},
	}

	for _, test := range tests {
		str := convert(test.Input, test.NumRows)

		assert.Equal(t, test.Output, str, "should be equals")
	}
}
