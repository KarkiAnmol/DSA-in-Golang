package selection

import (
	"fmt"
	"reflect"
	"testing"
)

// TestSelectionSort verifies the functionality of the SelectionSort function with several test cases.
// It defines multiple scenarios to ensure that SelectionSort behaves as expected across a variety of inputs.
func TestSelectionSort(t *testing.T) {
	// Define a slice of test cases
	tests := []struct {
		name     string // A descriptive name for the test case to help identify what it is testing
		input    []int  // The input array to sort
		expected []int  // The expected output array after sorting
	}{
		{"sorted input", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse order", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"empty input", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"random order", []int{3, 1, 4, 1, 5}, []int{1, 1, 3, 4, 5}},
	}

	// Iterate through each test case
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectionSort(tt.input) // Perform the sort on the input
			// Check if the output from the sort matches the expected output
			if !reflect.DeepEqual(got, tt.expected) {
				// If it doesn't match, report an error
				t.Errorf("Test case %d '%s'.............FAILED ❌\nExpected: %v\nActual:   %v\n", i, tt.name, tt.expected, got)
				// Uncommented error message for more standard formatting
				// t.Errorf("SelectionSort() = %v, want %v", got, tt.expected)
			} else {
				// If it does match, print a success message
				fmt.Printf("Test case %d '%s'.............PASSED ✅\n", i, tt.name)
			}
		})
	}
}
