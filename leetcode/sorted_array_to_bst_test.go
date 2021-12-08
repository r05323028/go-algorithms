package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{-10, -3, 0, 5, 9},
			expected: []int{0, -3, 9, -10, 5},
		},
	}
	for _, test := range testCases {
		node := SortedArrayToBST(test.input)
		actual := LevelOrder(node)
		equal := reflect.DeepEqual(actual, test.expected)
		if !equal {
			t.Errorf("Actual: %v != Expected: %v", actual, test.expected)
		}
	}
}
