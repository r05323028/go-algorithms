package sorts

import (
	"reflect"
	"testing"
)

func runTestCases(t *testing.T, sortFunc func([]int) []int) {
	testCases := []struct {
		input    []int
		expected []int
		name     string
	}{
		{
			input:    []int{3, 2, 1, 4},
			expected: []int{1, 2, 3, 4},
			name:     "Simple",
		},
		{
			input:    []int{1, 4, 2, 3, 2},
			expected: []int{1, 2, 2, 3, 4},
			name:     "Has Duplicates",
		},
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice",
		},
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Only One Element",
		},
		{
			input:    []int{1, 0, -1},
			expected: []int{-1, 0, 1},
			name:     "Has Negatives",
		},
	}

	for _, test := range testCases {
		actual := sortFunc(test.input)
		sorted := reflect.DeepEqual(actual, test.expected)
		if !sorted {
			t.Errorf("%v: actual=%v, expected=%v", test.name, actual, test.expected)
		}
	}
}

func TestMergeSort(t *testing.T) {
	runTestCases(t, MergeSort)
}

func TestInsertionSort(t *testing.T) {
	runTestCases(t, InsertionSort)
}

func TestBubbleSort(t *testing.T) {
	runTestCases(t, BubbleSort)
}
