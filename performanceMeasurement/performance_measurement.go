package main

import (
	insertion "DSA-in-Golang/insertionSort"
	merge "DSA-in-Golang/mergeSort"
	quick "DSA-in-Golang/quickSort"
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// main executes the performance test for sorting algorithms and records the results in a CSV file.
func main() {
	sizes := []int{100, 1000, 10000, 50000}
	file, err := os.Create("performance_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Sort Type", "Array Size", "Execution Time (ms)"})

	for _, size := range sizes {
		array := generateRandomArray(size)

		duration := testSorting(insertion.InsertionSort, array)
		writer.Write([]string{"Insertion Sort", strconv.Itoa(size), formatDuration(duration)})

		duration = testSorting(merge.MergeSort, array)
		writer.Write([]string{"Merge Sort", strconv.Itoa(size), formatDuration(duration)})

		duration = testSorting(quick.QuickSort, array)
		writer.Write([]string{"Quick Sort", strconv.Itoa(size), formatDuration(duration)})
	}
}

// generateRandomArray returns a slice of the specified size filled with random integers.
func generateRandomArray(size int) []int {
	array := make([]int, size)
	for i := range array {
		array[i] = rand.Intn(10000)
	}
	return array
}

// testSorting times the sorting process of the given function on the input array.
func testSorting(sortFunc func([]int) []int, original []int) time.Duration {
	array := make([]int, len(original))
	copy(array, original)
	start := time.Now()
	sortFunc(array)
	return time.Since(start)
}

// formatDuration converts a duration to a string in milliseconds.
func formatDuration(d time.Duration) string {
	return strconv.FormatFloat(d.Seconds()*1000, 'f', 2, 64)
}
