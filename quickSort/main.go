package quick

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Define the size of the sample array and create it with random integers.
	sampleSize := 10
	sample := make([]int, sampleSize)
	for i := range sample {
		sample[i] = rand.Intn(100) // Fill the array with random numbers between 0 and 99.
	}

	fmt.Println("Original array:", sample)

	// Sort the array using Quick Sort.
	QuickSort(sample)

	// Print the sorted array.
	fmt.Println("Sorted array by Quick Sort:", sample)
}
