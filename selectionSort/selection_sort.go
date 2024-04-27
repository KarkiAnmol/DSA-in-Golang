package selection

// SelectionSort sorts an array using the selection sort algorithm.
// This algorithm iterates over the array and for each position,
// finds the smallest element in the unsorted part to swap with the element at the current position.
func SelectionSort(arr []int) []int {
	n := len(arr) // Get the number of elements in the array

	// Loop through each element of the array except the last one
	for i := 0; i < n-1; i++ {
		// Assume the minimum element is the first element of the unsorted part
		minIdx := i

		// Check the rest of the array to find a smaller element
		for j := i + 1; j < n; j++ {
			// Update minIdx if a smaller element is found
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		// Swap the found minimum element with the element at the current position
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr // Return the sorted array
}
