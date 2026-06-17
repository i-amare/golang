package models

import (
	"fmt"

	"github.com/i-amare/rest-api/db"
	"github.com/i-amare/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO Users (Email, Password)
	VALUES (?, ?)
	`

	res, err := db.DB.Exec(query, u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id
	if err != nil {
		return err
	}

	fmt.Println("User saved successfully")
	return nil
}

func (u *User) GetUserID() (int64, error) {
	query := `
 	SELECT ID
 	FROM Users
 	WHERE Email = ?
  `

	res := db.DB.QueryRow(query, u.Email)
	var id int64
	err := res.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (u *User) ValidateCredentials() (bool, error) {
	query := `
	SELECT Password
	FROM Users
	WHERE Email = ?
	`

	res := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := res.Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	return utils.ValidatePassword(u.Password, hashedPassword), nil
}

func GetAllUsers() ([]User, error) {
	var usersArr []User

	query := `
	SELECT ID, Email
	FROM Users
	`

	res, err := db.DB.Query(query)
	if err != nil {
		return usersArr, err
	}
	defer res.Close()

	for res.Next() {
		u := User{Password: "********"}
		err = res.Scan(&u.ID, &u.Email)
		if err != nil {
			fmt.Println("Broken here")
			return usersArr, err
		}
		usersArr = append(usersArr, u)
	}

	return usersArr, nil
}
