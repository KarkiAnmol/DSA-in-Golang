package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize the random number generator with the current time as a seed
	rand.Seed(time.Now().UnixNano())

	// Define the size of the sample array and create it with random integers
	sampleSize := 10
	sample := make([]int, sampleSize)
	for i := range sample {
		sample[i] = rand.Intn(100) // Fill the array with random numbers between 0 and 99
	}

	// Print the original, unsorted array
	fmt.Println("Original array:", sample)

	// Measure the time taken by the Merge Sort algorithm to sort the array
	start := time.Now()           // Record the start time
	sorted := MergeSort(sample)   // Sort the array using Merge Sort
	duration := time.Since(start) // Calculate the duration by subtracting start from the current time

	// Print the sorted array and the time taken to sort
	fmt.Println("Sorted array by Merge Sort:", sorted)
	fmt.Printf("Execution time for Merge Sort: %v\n", duration)
}
