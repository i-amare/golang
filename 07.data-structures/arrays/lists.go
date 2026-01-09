package main

import (
	"fmt"
)

func main() {
	var productNames [4]string
	prices := [4]float32{9.9, 8, 3, 6.87}

	productNames[1] = "Nook"
	featuredPrices := prices[1:3]
	highlightedPrices := featuredPrices[1:]

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])

	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
}
