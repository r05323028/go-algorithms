package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, test := range testCases {
		actual := TwoSum(test.nums, test.target)
		equal := reflect.DeepEqual(actual, test.expected)
		if !equal {
			t.Errorf("Actual: %v != Expected: %v", actual, test.expected)
		}
	}
}
