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

	fv, realFV := calcFV(investmentAmount, expectedARR, expectedInflationRate, years);
	
	fmt.Print("Please enter Investment Amount: ")
	fmt.Scanln(&investmentAmount)
	

	fmt.Println("Future Value: ", round(fv, 2))
	fmt.Println("Real Future Value: ", round(realFV, 2))
}

func round(num float64, decimalPlaces int) float64 {
	x := math.Pow(10, float64(decimalPlaces))
	return math.Round(num*x) / x
}

func calcFV(amt float64, arr float64, inflationRate float64, period float64) (float64, float64) {
	fv := (amt) * math.Pow(1+(arr)/100, (period))
	realFV := fv / math.Pow(1+inflationRate/100, period)

	return fv, realFV;
}