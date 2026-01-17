package main

import (
	"fmt"

	"example.com/price-calculator/utils"
)

func main() {
	taxRates := []int{0, 10, 14, 20, 24}
	prices := utils.ReadPrice()

	for _, val := range(prices) {
		postTaxPrices := make([]float64, 0, len(taxRates))

		for _, taxRate := range(taxRates) {
			newPrice := val * (1 + (float64(taxRate)/100))
			postTaxPrices = append(postTaxPrices, newPrice)
		}
		
		fmt.Println(postTaxPrices)
		fmt.Println("")
	}

	fmt.Println(prices)
}
