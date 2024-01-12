package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() float64{
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func main() {
	var accountBalance float64 = readBalanceFromFile()

	fmt.Println("Welcome to Daniel Okyere's Bank!")
	fmt.Println("What would you like to do?")
	for {

		fmt.Println("Select an option to proceed")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Make an investment")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Your chosen option \n")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		} else if choice == 2 {
			var deposit float64
			fmt.Print("Your deposit: \n")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount must be greater than 0")
				continue
			}
			accountBalance += deposit
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new balance is %.2f", accountBalance)
		} else if choice == 3 {
			var amtWithdrawal float64
			fmt.Print("How much you want to withdraw: \n")
			fmt.Scan(&amtWithdrawal)
			if amtWithdrawal <= 0 {
				fmt.Println("Invalid amount must be greater than 0")
				continue
			}
			if amtWithdrawal > accountBalance {
				fmt.Println("You cannot withdraw more than you have")
				continue
			}
			accountBalance -= amtWithdrawal
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new balance is %.2f", accountBalance)
		} else if choice == 4 {
			fmt.Print("Invenstment is not ready yet....")
		} else {
			fmt.Println("Good bye, auf wiedersen!!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank")

}