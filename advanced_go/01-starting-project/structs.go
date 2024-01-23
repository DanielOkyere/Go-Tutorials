package main

import (
	"fmt"
	"time"
	"errors"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	return &user {
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user 
	appUser, error := newUser(userFirstName, userLastName, userBirthdate)
	// ... do something awesome with that gathered data!

	if error != nil {
		return 
	}
	appUser.outPutUserDetails()
	appUser.clearUserName()
	appUser.outPutUserDetails()

	// fmt.Println(firstName, lastName, birthdate)
}

func (u *user) outPutUserDetails(){
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
