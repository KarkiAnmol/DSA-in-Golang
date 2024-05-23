package quick

import (
	"fmt"
	"reflect"
	"testing"
)

// TestQuickSort verifies the functionality of the QuickSort function with several test cases.
func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"sorted input", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse order", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"empty input", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"random order", []int{3, 1, 4, 1, 5}, []int{1, 1, 3, 4, 5}},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Test case %d '%s'.............FAILED ❌\nExpected: %v\nActual:   %v\n", i, tt.name, tt.expected, tt.input)
				// t.Errorf("QuickSort() = %v, want %v", tt.input, tt.expected)
			} else {
				fmt.Printf("Test case %d '%s'.............PASSED ✅\n", i, tt.name)
			}
		})
	}
}
