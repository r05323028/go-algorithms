package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: -3}
	root.Right = &TreeNode{Val: 9}

	actual := LevelOrder(root)
	expected := []int{0, -3, 9}
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Errorf("Actual: %v != Expected: %v", actual, expected)
	}
}

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
