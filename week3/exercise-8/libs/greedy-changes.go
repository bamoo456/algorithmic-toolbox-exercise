package libs

func GetMinumCoinNum(money int) int {
	var denominations = []int{1, 3, 5}
	currentMoney := money
	result := 0
	for i := len(denominations) - 1; i >= 0; i-- {
		num := getCoinNum(currentMoney, denominations[i])
		currentMoney = currentMoney - (num * denominations[i])
		result += num
	}
	return result
}

func getCoinNum(money, unit int) int {
	if money-unit >= 0 {
		return getCoinNum(money-unit, unit) + 1
	}
	return 0
}
