package main

import (
	"algorithmic-toolbox-exercise/week2/exercise-7/libs"
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input the nth fibonacci number you want to know")
	text1, _ := reader.ReadString('\n')
	text1 = text1[:len(text1)-1]
	nth64, err := strconv.ParseInt(text1, 10, 64)
	nth := big.NewInt(nth64)
	text2, _ := reader.ReadString('\n')
	text2 = text2[:len(text2)-1]
	moder, err := strconv.ParseInt(text2, 10, 64)
	if err != nil {
		fmt.Println("Your input value can not be identified!")
		fmt.Println(err)
	} else {
		var fibMod int64 = libs.GetNthFibonacciMod(nth, moder)
		fmt.Printf("The %d th fibonacci mod %d is %d !\n", nth, moder, fibMod)
	}

}
