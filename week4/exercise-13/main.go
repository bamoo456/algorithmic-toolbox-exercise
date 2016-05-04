package main

import (
	"algorithmic-toolbox-exercise/week4/exercise-13/libs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("setup the sorted array size and inside elements")
	sortedArray := getFormattedInputs()
	fmt.Println("setup the search target array size and inside elements")
	targetArray := getFormattedInputs()
	resultArray := make([]int, len(targetArray), len(targetArray))
	for i := 0; i < len(resultArray); i++ {
		resultArray[i] = libs.BinarySearch(sortedArray, 0, len(resultArray)-1, targetArray[i])
	}
	fmt.Println("The Search Result is ", resultArray)
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)
	inputs, _ := reader.ReadString('\n')
	inputs = inputs[:len(inputs)-1]
	return strings.Split(inputs, " ")
}

func getFormattedInputs() []int {
	inputs := getInput()
	arraySize, _ := strconv.Atoi(inputs[0])
	if len(inputs)-1 != arraySize {
		fmt.Println("wrong input params")
	}
	formattedInputs := make([]int, arraySize, arraySize)
	for i := 1; i < len(inputs); i++ {
		formattedInputs[i-1], _ = strconv.Atoi(inputs[i])
	}
	return formattedInputs
}
