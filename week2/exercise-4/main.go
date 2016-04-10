package main

import (
	"algorithmic-toolbox-exercise/week2/exercise-4/libs"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input the nth fibonacci number you want to know")
	text1, _ := reader.ReadString('\n')
	text1 = text1[:len(text1)-1]
	nth, err := strconv.Atoi(text1)
	if err != nil {
		fmt.Println("Your input value can not be identified!")
		fmt.Println(err)
	} else {
		var fibNum int = libs.GetLastDigitOnFibonacci(nth)
		fmt.Printf("The last digit of %d th fibonacci number is %d !\n", nth, fibNum)
	}

}
