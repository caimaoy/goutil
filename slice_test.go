package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	var tests = []struct {
		inputString string
		inputSlice  []string
		expected    bool
	}{
		{"", []string{}, false},
		{"", []string{""}, true},
		{"", []string{"a"}, false},
		{"", []string{"中文", "foo"}, false},
		{"a", []string{}, false},
		{"a", []string{"a"}, true},
		{"foo", []string{"中文", "foo", ""}, true},
		{"中文", []string{"foo", "中文"}, true},
	}

	for _, test := range tests {
		assert.Equal(t, StringInSlice(test.inputString, test.inputSlice), test.expected)
	}
}
