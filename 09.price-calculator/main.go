package main

import (
	"fmt"
	"math"

	"example.com/price-calculator/utils"
)

var output = make([][]float64, 5)

func main() {
	taxRates := []int{0, 10, 14, 20, 24}
	prices := utils.ReadPrice()

	for _, val := range prices {
		calcPostTaxPrices(val, &taxRates)
	}

	// utils.OutputPrices(taxRates, output)
	for i, line := range output {
		fmt.Println("Tax Rate: ", taxRates[i])
		fmt.Println(line)
		fmt.Println("")
	}
}

func calcPostTaxPrices(val float64, taxRates *[]int) []float64 {
	postTaxPrices := make([]float64, 0, len(*taxRates))

	for idx, taxRate := range *taxRates {
		newPrice := val * (1 + (float64(taxRate) / 100))
		newPrice = math.Round(newPrice*100) / 100

		output[idx] = append(output[idx], newPrice)
		postTaxPrices = append(postTaxPrices, newPrice)
	}
	
	return postTaxPrices
}
