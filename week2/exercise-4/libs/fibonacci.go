package libs

func GetLastDigitOnFibonacci(nth int) int {
	var fibLastDigit = make([]int, nth, nth)
	fibLastDigit[0] = 1
	fibLastDigit[1] = 1
	for i := 2; i < nth; i++ {
		fibLastDigit[i] = (fibLastDigit[i-2] + fibLastDigit[i-1]) % 10
	}
	return fibLastDigit[nth-1]
}
