package slice

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

func TestDeleteStringSlice(t *testing.T) {
	var tests = []struct {
		inputSlice       []string
		inputDeleteIndex int
		expected         []string
	}{
		{[]string{"0", "1", "2"}, 0, []string{"1", "2"}},
		{[]string{"0", "1", "2"}, 1, []string{"0", "2"}},
		{[]string{"0", "1", "2"}, 2, []string{"0", "1"}},
	}

	for _, test := range tests {
		assert.Equal(t, DeleteStringSlice(test.inputSlice, test.inputDeleteIndex), test.expected)
	}
}

func TestDeleteStringSlicePanic(t *testing.T) {
	var tests = []struct {
		inputSlice       []string
		inputDeleteIndex int
		expected         []string
	}{
		{[]string{"0", "1", "2"}, 3, []string{"0", "2"}},
	}

	for _, test := range tests {
		defer func() { recover() }()
		_ = DeleteStringSlice(test.inputSlice, test.inputDeleteIndex)
		t.Errorf("did not panic")
	}
}

func TestStringSliceEqule(t *testing.T) {
	var tests = []struct {
		inputSliceA []string
		inputSliceB []string
		expected    bool
	}{
		{[]string{}, []string{}, true},
		{[]string{}, []string{""}, false},
		{[]string{}, []string(nil), false},
		{[]string{"0", "1", "2"}, []string{"0", "1", "2"}, true},
		{[]string{"0", "1", "2"}, []string{"0", "1", "3"}, false},
		{[]string{"0", "1", "2", "中文"}, []string{"0", "1", "2", "中文"}, true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, StringSliceEqule(test.inputSliceA, test.inputSliceB))
	}
}
