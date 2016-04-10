package libs

func FindGCD(x, y int32) int32 {
	if x > y {
		return euclidGCD(x, y)
	} else {
		return euclidGCD(y, x)
	}
}

func euclidGCD(a, b int32) int32 {
	if b == 0 {
		return a
	}
	if b == 1 {
		return 1
	}
	return euclidGCD(b, a%b)
}
