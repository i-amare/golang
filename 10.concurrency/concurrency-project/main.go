package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	routines := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		routines[i] = make(chan bool)
		go calc(taxRate, routines[i])
	}

	for i, done := range routines {
		fmt.Println(i, ": ", <- done)
	}
}

func calc(taxRate float64, done chan bool) {
	fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
	// cmdm := cmdmanager.New()
	priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
	err := priceJob.Process()

	if err != nil {
		fmt.Println("Could not process job")
		fmt.Println(err)
	}

	done <- true
}
