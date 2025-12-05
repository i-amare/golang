package main

import "fmt"

func main() {
	var revenue float64 = inputFloat64("Please input revenue: ")
	var expenses float64 = inputFloat64("Please input expenses: ")
	var taxRate float64 = inputFloat64("Please input tax rate: ")

	preTaxProfit := revenue - expenses
	taxExpense := taxRate * preTaxProfit
	postTaxProfit := preTaxProfit - taxExpense
	ratio := postTaxProfit / preTaxProfit

	fmt.Println("Earnings before tax: ", preTaxProfit)
	fmt.Printf("Tax Expense: (%v)\n", taxExpense)
	fmt.Println("Profit for period: ", postTaxProfit)
	fmt.Println("Ratio: ", ratio)
}

func inputFloat64(msg string) float64 {
	var value float64

	fmt.Print(msg)
	fmt.Scanln(&value)

	return value
}
