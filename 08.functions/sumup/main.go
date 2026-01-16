package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func main() {
	numbers := []int{1, 2, 3}
	sum1 := sumArray(&numbers)
	sum2 := sum(2.4, 3, 4)

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func sum[T Number](numbers ...T) T {
	var sum T = 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func sumArray[T Number](numbers *[]T) T {
	var sum T = 0

	for _, val := range *numbers {
		sum += val
	}

	return sum
}
