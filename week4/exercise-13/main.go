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
	resultChannels := make([]chan int, len(targetArray), len(targetArray))
	for i := 0; i < len(resultArray); i++ {
		resultChannels[i] = make(chan int)
		go goBinarySearch(sortedArray, 0, len(resultArray)-1, targetArray[i], resultChannels[i])
	}
	for i := 0; i < len(resultArray); i++ {
		resultArray[i] = <-resultChannels[i]
	}
	fmt.Println("The Search Result is ", resultArray)
}

func goBinarySearch(array []int, low int, high int, key int, channel chan int) {
	channel <- libs.BinarySearch(array, low, high, key)
	close(channel)
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
