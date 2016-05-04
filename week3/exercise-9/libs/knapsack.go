package libs

import (
	"algorithmic-toolbox-exercise/week3/exercise-9/models"
)

func GetMaximumKnapsackValue(items []models.Item, weightLimit float32) float32 {
	var value float32
	currentWeight := weightLimit
	sortedItems := getSortedItems(items)
	for i := len(sortedItems) - 1; i > -1; i-- {
		if sortedItems[i].Weight < currentWeight {
			currentWeight -= sortedItems[i].Weight
			value += sortedItems[i].Value
		} else {
			value += (currentWeight * sortedItems[i].Value / sortedItems[i].Weight)
			break
		}
	}
	return value
}

func getSortedItems(items []models.Item) []models.Item {
	temp := models.Item{}
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			val1 := items[i].Value / items[i].Weight
			val2 := items[j].Value / items[j].Weight
			if val2 < val1 {
				temp = items[i]
				items[i] = items[j]
				items[j] = temp
			}
		}
	}
	return items
}
