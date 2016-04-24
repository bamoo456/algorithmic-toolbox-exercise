package libs

import "fmt"
import "math/big"

func GetNthFibonacciMod(nth *big.Int, m int64) int64 {
	lastNums := [2]*big.Int{big.NewInt(0), big.NewInt(1)}
	reminders := []*big.Int{big.NewInt(0), big.NewInt(1)}
	result := big.NewInt(0)
	var maxReminders int64 = 10000
	var i int64
	var idx int
	var cycle int
	var targetNth = big.NewInt(0).Add(nth, big.NewInt(1))

	// if nth is less than 10, then directly compute the reminder
	if targetNth.Int64() < 10 {
		return computeFibNum(targetNth.Int64()) % m
	}
	// to restrict max iteration on for-loop
	maxIter := big.NewInt(maxReminders)

	for i = 2; big.NewInt(i).Cmp(big.NewInt(0).Sub(maxIter, big.NewInt(1))) == -1; i++ {
		result = big.NewInt(0).Add(lastNums[0], lastNums[1])
		lastNums[0] = lastNums[1]
		lastNums[1] = result
		rem := big.NewInt(0).Mod(result, big.NewInt(m))
		reminders = append(reminders, rem)
		if reminders[idx].Cmp(rem) == 0 {
			idx++
		} else {
			idx = 0
		}
		if idx > 3 {
			cycle = len(reminders) - idx
			break
		}
	}
	if cycle == 0 {
		fmt.Println("Do not find the reminders cycle", cycle)
		return 0
	}
	var rIdx int64
	if cycle > 0 {
		rIdx = targetNth.Mod(targetNth, big.NewInt(int64(cycle))).Int64()
	} else {
		rIdx = targetNth.Int64()
	}
	return reminders[rIdx-1].Int64()
}

func computeFibNum(n int64) int64 {
	var result, i int64
	lastNums := [2]int64{0, 1}
	if n == 0 || n == 1 {
		return lastNums[n]
	}
	for i = 2; i < n-1; i++ {
		result = lastNums[0] + lastNums[1]
		lastNums[0] = lastNums[1]
		lastNums[1] = result
	}
	return result
}
