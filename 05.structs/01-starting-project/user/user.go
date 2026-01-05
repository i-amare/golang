package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	dateCreated time.Time
}

func NewUser(firstName string, lastName string, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Invalid Input")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		birthdate:   birthdate,
		dateCreated: time.Now(),
	}, nil
}

func (user User) OutputUserData() {
	fmt.Println("{")
	fmt.Println("firstName: ", user.firstName)
	fmt.Println("lastName: ", user.lastName)
	fmt.Println("birthdate: ", user.birthdate)
	fmt.Println("dateCreated: ", user.dateCreated)
	fmt.Println("}")
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
	fmt.Println("Cleared User Name")
}
