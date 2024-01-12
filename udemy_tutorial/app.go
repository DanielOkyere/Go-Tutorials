package main

import (
	"fmt"
	"math"
)

func CalculateInvestment () {
	const inflationRate float64 = 2.5
	var investmentRate float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentRate)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentRate * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Println("Future Value: ",futureRealValue)
	fmt.Printf("Future Value %.1f\n FutureRealValue %.1f\n", futureValue, futureRealValue)
	
	// fmt.Println(futureValue)
}