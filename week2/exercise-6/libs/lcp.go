package libs

func GetLCP(a, b int64) int64 {
	var gcd int64
	if a > b {
		gcd = getGCD(a, b)
	} else {
		gcd = getGCD(b, a)
	}
	return a * b / gcd
}

func getGCD(a, b int64) int64 {
	if b == 0 {
		return a
	}
	if b == 1 {
		return 1
	}
	return getGCD(b, a%b)
}
