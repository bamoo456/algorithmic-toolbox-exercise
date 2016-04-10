package libs

func GetNthFibonacci(nth int) int64 {
	var fibNums = make([]int64, nth, nth)
	fibNums[0] = 1
	fibNums[1] = 1
	for i := 2; i < nth; i++ {
		fibNums[i] = fibNums[i-2] + fibNums[i-1]
	}
	return fibNums[nth-1]
}
