package main

import (
	"algorithmic-toolbox-exercise/week2/exercise-3/libs"
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
	} else if nth > 45 {
		fmt.Println("nth is not supported to calculate !")
	} else {
		var fibNum int64 = libs.GetNthFibonacci(nth)
		fmt.Printf("The %d th fibonacci number is %d !\n", nth, fibNum)
	}

}
