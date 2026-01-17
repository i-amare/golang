package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadPrice() []float64 {
	prices := make([]float64, 0, 50)

	file, err := os.Open("store/prices.txt")

	if err != nil {
		fmt.Println(err)
		return prices
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		price, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			fmt.Println(err)
			continue
		}
		
		prices = append(prices, price)
	}

	return prices
}

func OutputPrices(taxRates []int, prices [][]float64) {

}