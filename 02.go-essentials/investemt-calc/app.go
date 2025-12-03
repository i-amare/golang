package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := 1000.0
	expectedARR := 10.0
	years := 10.0
	expectedInflationRate := 2.5

	fmt.Print("Please enter Investment Amount: ")
	fmt.Scanln(&investmentAmount)
	
	fv := (investmentAmount) * math.Pow(1+(expectedARR)/100, (years))
	realFV := fv / math.Pow(1+expectedInflationRate/100, years)

	fmt.Println("Future Value: ", round(fv, 2))
	fmt.Println("Real Future Value: ", round(realFV, 2))
}

func round(num float64, decimalPlaces int) float64 {
	x := math.Pow(10, float64(decimalPlaces))
	return math.Round(num*x) / x
}
