package main

import (
	"algorithmic-toolbox-exercise/week3/exercise-9/libs"
	"algorithmic-toolbox-exercise/week3/exercise-9/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("setup the number N of items and the capacity W of a knapsack")

	if inputs := getInput(); len(inputs) != 2 {
		fmt.Println("wrong input params")
	} else {
		itemNum, _ := strconv.Atoi(inputs[0])
		knapsackSize, _ := strconv.Atoi(inputs[1])
		items := []models.Item{}
		fmt.Println("give me the item weight w and value v")
		for i := 0; i < itemNum; i++ {
			itemParams := getInput()
			value, _ := strconv.Atoi(itemParams[0])
			weight, _ := strconv.Atoi(itemParams[1])
			items = append(items, models.Item{float32(weight), float32(value)})
		}
		maxKnapsackValue := libs.GetMaximumKnapsackValue(items, float32(knapsackSize))
		fmt.Println("The maximum knapsack value is ", maxKnapsackValue)
	}
}

func getInput() []string {
	reader := bufio.NewReader(os.Stdin)
	inputs, _ := reader.ReadString('\n')
	inputs = inputs[:len(inputs)-1]
	return strings.Split(inputs, " ")
}
