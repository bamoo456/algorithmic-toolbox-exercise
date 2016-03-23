package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// read 1st number string from std input
	fmt.Print("Enter 1st number: ")
	text1, _ := reader.ReadString('\n')
	fmt.Print("Enter 2nd number: ")
	// read 2nd number string from std input
	text2, _ := reader.ReadString('\n')
	// trim the last char ('\n')
	text1 = text1[:len(text1)-1]
	text2 = text2[:len(text2)-1]
	fmt.Println("Your 1st number is ", text1, "2nd number is ", text2)
	// covert string to float
	num1, _ := strconv.ParseFloat(text1, 64)
	num2, _ := strconv.ParseFloat(text2, 64)
	sum := num1 + num2
	fmt.Println("The sum of this two numbers are: ", sum)

}
