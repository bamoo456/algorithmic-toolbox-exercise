package impl

// import "strconv"

// find the maximum pair string
func MaximumPair(sequences []int64) (int64, int64) {
	var max1, max2 int64
	for _, num := range sequences {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}
	return max1, max2
}
