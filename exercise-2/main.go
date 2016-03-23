package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// find the maximum pair string
func maximumPair(base string, pairs []string) (int64, int64) {
	var maxNum int64
	baseNum, _ := strconv.ParseInt(base, 10, 32)
	for _, numString := range pairs {
		num, _ := strconv.ParseInt(numString, 10, 32)
		if num > maxNum {
			if maxNum > baseNum {
				baseNum = maxNum
			}
			maxNum = num
		}
	}
	return baseNum, maxNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your base number: ")
	baseNum, _ := reader.ReadString('\n')
	baseNum = baseNum[:len(baseNum)-1]

	fmt.Println("The baseNum is ", baseNum)

	fmt.Println("please enter your pair numbers: ")
	pairStrings, _ := reader.ReadString('\n')
	pairStrings = pairStrings[:len(pairStrings)-1]

	pairNums := strings.Split(pairStrings, " ")

	fmt.Println("The pair strings are", pairNums)

	v1, v2 := maximumPair(baseNum, pairNums)
	fmt.Println("The multiple result of maximum pair is :", v1*v2)

}
