package main

import (
	"fmt"
)

func main() {
	prices := []float64{200, 149.99}
	updatedPrices := append(prices, 249.99);
	
	fmt.Println(prices)
	fmt.Println(updatedPrices)
}

// func main() {
// 	var productNames [4]string
// 	prices := [4]float32{9.9, 8, 3, 6.87}

// 	productNames[1] = "Nook"
// 	featuredPrices := prices[1:3]
// 	highlightedPrices := featuredPrices[1:]

// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)

// 	highlightedPrices[0] = 219.78
// 	fmt.Println(prices)

// 	var bigArr [5]string
// 	smallArr := bigArr[:2]
// 	fmt.Println(len(bigArr), cap(bigArr))
// 	fmt.Println(len(smallArr), cap(smallArr))
// }
