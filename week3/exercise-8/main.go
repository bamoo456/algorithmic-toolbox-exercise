package main

import (
	"algorithmic-toolbox-exercise/week3/exercise-8/libs"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input the inputMoney fibonacci number you want to know")
	text1, _ := reader.ReadString('\n')
	text1 = text1[:len(text1)-1]
	inputMoney, err := strconv.Atoi(text1)

	if err != nil {
		fmt.Println("Your input value can not be identified!")
		fmt.Println(err)
	} else if inputMoney > 1000 || inputMoney < 0 {
		fmt.Println("inputMoney is not supported to calculate !")
	} else {
		coinNum := libs.GetMinumCoinNum(inputMoney)
		fmt.Println("The minimum coin numbers are:", coinNum)
	}

}
