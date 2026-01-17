package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadPrice() []float32 {
	prices := make([]float32, 0, 50)

	file, err := os.Open("store/prices.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return prices
}