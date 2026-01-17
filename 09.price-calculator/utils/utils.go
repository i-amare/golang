package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type p struct {
	TaxRate int       `json:"taxRate"`
	Prices  []float64 `json:"prices"`
}

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
	output := make([]p, 0, len(taxRates))

	for i, taxRate := range taxRates {
		data := p{
			TaxRate: taxRate,
			Prices:  prices[i],
		}
		fmt.Println(data)
		output = append(output, data)
	}

	fmt.Println(output)

	d, _ := json.MarshalIndent(output, "", "	")
	os.WriteFile("prices.json", d, 0644)
}
