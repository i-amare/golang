package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	
	doubledNumbers := doubleNumbers(&numbers)

	fmt.Println(doubledNumbers)
}

func doubleNumbers(numbers *[]int) *[]int {
	doubledNumbers := make([]int, 0, len(*numbers))

	for _, val := range(*numbers) {
		doubledNumbers = append(doubledNumbers, double(val))
	}

	return &doubledNumbers;
}

func double(number int) int {
	return  number * 2
}