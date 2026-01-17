package main

import (
	"fmt"
	"math"

	"example.com/price-calculator/utils"
)

func main() {
	taxRates := []int{0, 10, 14, 20, 24}
	prices := utils.ReadPrice()

	output := make([][]float64, 5)
	for i := range output {
		output[i] = make([]float64, 0, len(prices))
	}

	for _, val := range prices {
		postTaxPrices := make([]float64, 0, len(taxRates))

		for idx, taxRate := range taxRates {
			newPrice := val * (1 + (float64(taxRate) / 100))
			newPrice = math.Round(newPrice*100) / 100

			output[idx] = append(output[idx], newPrice)
			postTaxPrices = append(postTaxPrices, newPrice)
		}
	}

	for i, val := range output {
		fmt.Println("Tax Rate: ", taxRates[i])
		fmt.Println(val)
		fmt.Println("")
	}
}
