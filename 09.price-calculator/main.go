package main

import (
	"fmt"

	"example.com/price-calculator/utils"
)

func main() {
	prices := utils.ReadPrice()
	fmt.Println(prices)
}
