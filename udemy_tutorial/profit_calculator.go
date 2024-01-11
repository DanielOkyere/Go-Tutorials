package main

import (
	"fmt"
)

func main () {
	var revenue , expenses float64
	const tax_rate float64 = 0.020

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	var earnings_before_tax float64 = revenue - expenses
	
	var earnings_after_tax float64 = revenue - (earnings_before_tax * tax_rate)
	
	var ratio float64 =  earnings_before_tax / earnings_after_tax

	fmt.Println(earnings_after_tax)
	fmt.Println(earnings_before_tax)
	fmt.Println(ratio)
}