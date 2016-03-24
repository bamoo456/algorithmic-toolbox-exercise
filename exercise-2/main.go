package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// find the maximum pair string
func maximumPair(sequences []string) (int64, int64) {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("please enter your sequence numbers: ")
	sequences, _ := reader.ReadString('\n')
	sequences = sequences[:len(sequences)-1]

	sequenceNums := strings.Split(sequences, " ")

	fmt.Println("The sequence numbers are", sequenceNums)

	v1, v2 := maximumPair(sequenceNums)
	fmt.Println("The multiple result of maximum pair is :", v1*v2)
}
