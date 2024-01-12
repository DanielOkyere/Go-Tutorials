package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	rates := getUserInput("Rates: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, rates)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - rate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
