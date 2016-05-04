package libs

func BinarySearch(a []int, low, high, key int) int {
	// do not find matched position
	if high < low {
		return -1
	}
	mid := low + (high-low)/2
	if a[mid] == key {
		return mid
	} else if key < a[mid] {
		return BinarySearch(a, low, mid-1, key)
	} else {
		return BinarySearch(a, mid+1, high, key)
	}
}
