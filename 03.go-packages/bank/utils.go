package main

import (
	"fmt"
	"os"
	"strconv"
)

const FILENAME = "account.txt"

func writeBalance(accBalance float64) {
	balanceText := fmt.Sprint(accBalance)
	os.WriteFile(FILENAME, []byte(balanceText), 0644)
}

func readBalance() float64 {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		return 0
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0
	}
	return balance
}
