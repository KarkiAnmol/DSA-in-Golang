package main

import (
	"fmt"
)

func main() {
	sample := []int{12, 11, 13, 5, 6}
	fmt.Println("Original array:", sample)
	sorted := InsertionSort(sample)
	fmt.Println("Sorted array:", sorted)
}
