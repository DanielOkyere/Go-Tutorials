package main

import (
	"fmt"
	"math"
)

// Function to check if a number is a perfect square
func isPerfectSquare(n int) bool {
	sqrtN := int(math.Sqrt(float64(n)))
	return sqrtN*sqrtN == n
}

// Function to check if a number is a Fibonacci number
func isFibonacci(num int) string {
	// A number is a Fibonacci number if and only if one of (5*num*num + 4) or (5*num*num - 4) is a perfect square
	if isPerfectSquare(5*num*num+4) || isPerfectSquare(5*num*num-4) {
		return "yes"
	}
	return "no"
}

func MathChallenge(num int) string {
	result:= isFibonacci(num)
	return result
}

func main() {
	// Example usage:
	fmt.Println(MathChallenge(54))
}

