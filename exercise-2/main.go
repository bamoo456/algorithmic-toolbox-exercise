package main

import (
	"algorithmic-toolbox-exercise/exercise-2/implementation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("please enter your sequence numbers: ")
	sequences, _ := reader.ReadString('\n')
	sequences = sequences[:len(sequences)-1]

	sequenceNums := strings.Split(sequences, " ")

	fmt.Println("The sequence numbers are", sequenceNums)

	v1, v2 := impl.MaximumPair(sequenceNums)
	fmt.Println("The multiple result of maximum pair is :", v1*v2)
}
