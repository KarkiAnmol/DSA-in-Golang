package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TestInsertionSort verifies the functionality of the InsertionSort function with several test cases.
func TestInsertionSort(t *testing.T) {
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
			got := InsertionSort(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.expected)
			} else {
				fmt.Printf("Test case %d '%s'.............PASSED ✅\n", i, tt.name)
			}
		})
	}
}
