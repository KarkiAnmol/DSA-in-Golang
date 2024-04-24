package main

import (
	"fmt"
)

// Insertion Sort sorts an array using the insertion sort algorithm
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Move elements of arr[0..i-1], that are greater than key, to one position ahead
		// of their current position
		for j > -1 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	sample := []int{12, 11, 13, 5, 2}
	fmt.Println("Original array:", sample)
	sorted := InsertionSort(sample)
	fmt.Println("Sorted array:", sorted)
}
