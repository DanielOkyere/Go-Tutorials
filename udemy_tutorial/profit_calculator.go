package main

import (
	"fmt"
	"errors"
	"os"
)

// Goals
// 1) Validate user input
// => Show error message & exit if invalid input is provided
// - No negative numbers 
// - Not 0
// 2) Store calculated results into a file
const fileName = "data_store.txt"

func writeValuesToFile(revenue, expenses, rates float64){
	revenueText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", revenue, expenses, rates)
	os.WriteFile(fileName, []byte(revenueText), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Print(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Print(err)
		return
	}
	rates, err := getUserInput("Rates: ")
	if err != nil {
		fmt.Print(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, rates)
	writeValuesToFile(ebt, profit, rates)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0.0, errors.New("Invalid amout entered")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - rate/100)
	ratio := 0.0

	if profit <= 0 {
		return ebt, profit, ratio
	} else {
		ratio = ebt / profit
	}

	return ebt, profit, ratio
}
