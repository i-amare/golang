package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	dateCreated time.Time
}

func newUser(firstName string, lastName string, birthdate string) *User {
	return &User{
		firstName:   firstName,
		lastName:    lastName,
		birthdate:   birthdate,
		dateCreated: time.Now(),
	}
}

func (user User) outputUserData() {
	fmt.Println("{")
	fmt.Println("firstName: ", user.firstName)
	fmt.Println("lastName: ", user.lastName)
	fmt.Println("birthdate: ", user.birthdate)
	fmt.Println("dateCreated: ", user.dateCreated)
	fmt.Println("}")
}

func (user *User) clearUserName() {
	user.firstName = ""
	user.lastName = ""
	fmt.Println("Cleared User Name")
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(firstName, lastName, birthdate);

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
