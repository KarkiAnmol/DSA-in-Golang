package merge

// MergeSort sorts an array using the merge sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])  // Recursively sort the first half
	right := MergeSort(arr[mid:]) // Recursively sort the second half

	return merge(left, right) // Merge the two halves
}

// merge combines two sorted sub-arrays into a single sorted array
func merge(left, right []int) []int {
	var result []int
	l, r := 0, 0

	// Traverse both arrays and insert smaller value from left or right into result
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Collect remaining elements of left, if any
	result = append(result, left[l:]...)

	// Collect remaining elements of right, if any
	result = append(result, right[r:]...)

	return result
}
