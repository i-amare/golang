package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedARR = 10
	var investmentTimeFrame = 10

	var fv = float64(investmentAmount) * math.Pow(1+float64(expectedARR)/100, float64(investmentTimeFrame))

	fmt.Print(fv)
}
