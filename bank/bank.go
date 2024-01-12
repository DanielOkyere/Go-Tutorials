package main

import (
	"fmt"
	"example.com/bank/fileops"
	"example.com/bank/communication"
	"github.com/Pallinder/go-randomdata"
)



const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		//panic("Cant find account store")
	}

	fmt.Println("Welcome to Daniel Okyere's Bank!")
	fmt.Println("Reach us 24/7 on ", randomdata.PhoneNumber())
	fmt.Println("What would you like to do?")
	for {
		communication.PresentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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


