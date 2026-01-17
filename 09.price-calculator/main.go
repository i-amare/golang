package main

import (
	"fmt"
	"math"

	"example.com/price-calculator/utils"
)

func main() {
	taxRates := []int{0, 10, 14, 20, 24}
	prices := utils.ReadPrice()

	for _, val := range(prices) {
		postTaxPrices := calcPostTaxPrices(val, taxRates)	
		fmt.Println(postTaxPrices)
	}

	fmt.Println(prices)
}

func calcPostTaxPrices(val float64, taxRates []int) []float64 {
		postTaxPrices := make([]float64, 0, len(taxRates))

		for _, taxRate := range(taxRates) {
			newPrice := val * (1 + (float64(taxRate)/100))
			newPrice = math.Round(newPrice * 100) / 100
			postTaxPrices = append(postTaxPrices, newPrice)
		}
		return postTaxPrices
}