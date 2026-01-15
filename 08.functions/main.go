package main

import (
	"fmt"
)

type transformFn func(number int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	double := multiply(2)
	triple := multiply(3)

	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(*doubledNumbers)
	fmt.Println(*tripledNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) *[]int {
	result := make([]int, 0, len(*numbers))

	for _, val := range *numbers {
		result = append(result, transform(val))
	}

	return &result
}

func multiply(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
