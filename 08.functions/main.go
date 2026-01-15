package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	double := multiply(2)
	triple := multiply(3)

	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(*doubledNumbers)
	fmt.Println(*tripledNumbers)
}

func transformNumbers(numbers *[]int, transform func(number int) int) *[]int {
	doubledNumbers := make([]int, 0, len(*numbers))

	for _, val := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(val))
	}

	return &doubledNumbers
}

func multiply(factor int) func(number int) int {

	f := func(number int) int {
		return number * factor
	}

	return f
}
