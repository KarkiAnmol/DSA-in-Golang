package quick

// QuickSort sorts an array using the Quick Sort algorithm.
// It takes a slice of integers as input and modifies it in place.
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

// quickSort is a recursive helper function that performs the quick sort.
func quickSort(arr []int, low, high int) []int {
	if low < high {
		// Partition the array and get the partition index.
		pi := partition(arr, low, high)

		// Recursively sort the elements before and after the partition.
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
	return arr
}

// partition reorders the elements of the array such that elements less than the pivot are to the left,
// and elements greater than the pivot are to the right.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choosing the last element as the pivot.
	i := low - 1       // Index of the smaller element.

	for j := low; j < high; j++ {
		// If the current element is smaller than or equal to the pivot.
		if arr[j] <= pivot {
			i++                             // Increment index of the smaller element.
			arr[i], arr[j] = arr[j], arr[i] // Swap elements.
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap the pivot element with the element at i+1.
	return i + 1                              // Return the partition index.
}
