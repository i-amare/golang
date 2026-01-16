package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sum := sumNumbers(&numbers)

	fmt.Println(sum)
}

func sumNumbers(numbers *[]int) int {
	sum := 0

	for _, val := range *numbers {
		sum += val
	}

	return sum
}
