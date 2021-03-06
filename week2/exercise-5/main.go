package main

import (
	"algorithmic-toolbox-exercise/week2/exercise-5/libs"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your 1st number:")
	text1, _ := reader.ReadString('\n')
	text1 = text1[:len(text1)-1]
	fmt.Println("Please input your 2nd number:")
	text2, _ := reader.ReadString('\n')
	text2 = text2[:len(text2)-1]
	num1, err1 := strconv.ParseInt(text1, 10, 32)
	num2, err2 := strconv.ParseInt(text2, 10, 32)
	if err1 != nil || err2 != nil {
		fmt.Println("Your input value can not be identified!")
	} else {
		var gcd int32 = libs.FindGCD(int32(num1), int32(num2))
		fmt.Printf("The gcd of %d and %d is %d !\n", int32(num1), int32(num2), gcd)
	}

}
