package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type OutputStruct struct {
	TaxRate int       `json:"taxRate"`
	Prices  []float64 `json:"prices"`
}

func ReadPrices() []float64 {
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
	output := make([]OutputStruct, 0, len(taxRates))

	for i, taxRate := range taxRates {
		output = append(output, OutputStruct{
			TaxRate: taxRate,
			Prices:  prices[i],
		})
	}

	data, _ := json.MarshalIndent(output, "", "	")
	os.WriteFile("prices.json", data, 0644)
}
