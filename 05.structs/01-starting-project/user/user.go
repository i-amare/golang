package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	Birthdate   string
	DateCreated time.Time
}

func NewUser(firstName string, lastName string, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Invalid Input")
	}

	return &User{
		FirstName:   firstName,
		LastName:    lastName,
		Birthdate:   birthdate,
		DateCreated: time.Now(),
	}, nil
}

func (user User) OutputUserData() {
	fmt.Println("{")
	fmt.Println("firstName: ", user.FirstName)
	fmt.Println("lastName: ", user.LastName)
	fmt.Println("birthdate: ", user.Birthdate)
	fmt.Println("dateCreated: ", user.DateCreated)
	fmt.Println("}")
}

func (user *User) ClearUserName() {
	user.FirstName = ""
	user.LastName = ""
	fmt.Println("Cleared User Name")
}
