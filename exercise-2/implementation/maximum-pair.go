package impl

import "strconv"

// find the maximum pair string
func MaximumPair(sequences []string) (int64, int64) {
	var max1, max2 int64
	for _, numString := range sequences {
		num, _ := strconv.ParseInt(numString, 10, 32)
		if num > max1 {
			if max1 > max2 {
				max2 = max1
			}
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}
	return max2, max1
}
