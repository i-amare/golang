package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedARR float64 = 10
	var investmentTimeFrame float64 = 10

	var fv = (investmentAmount) * math.Pow(1+(expectedARR)/100, (investmentTimeFrame))

	fmt.Print(fv)
}
