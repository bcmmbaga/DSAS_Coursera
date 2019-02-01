package toolbox

// In this problem, you will implement the binary search algorithm that allows searching
// very efficiently (even huge) lists, provided that the list is sorted.

func binarySearchRecursive(arr []int, low, high, key int) int {
	if high >= low {
		mid := low + (high-low)/2
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			return binarySearchRecursive(arr, low, mid-1, key)
		} else {
			return binarySearchRecursive(arr, mid+1, high, key)
		}
	}
	return -1
}

func binarySearch(arr []int, key int) int {
	low, high := 0, len(arr)-1
	for high <= low {
		mid := low + (high-low)/2
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
