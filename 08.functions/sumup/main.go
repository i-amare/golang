package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func main() {
	numbers := []int{1, 2, 3}
	sum := sumArray(&numbers)

	fmt.Println(sum)
}

func sumArray[T Number](numbers *[]T) T {
	var sum T = 0

	for _, val := range *numbers {
		sum += val
	}

	return sum
}
