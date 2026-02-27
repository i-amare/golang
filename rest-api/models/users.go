package models

import (
	"fmt"

	"github.com/i-amare/rest-api/db"
)

type User struct {
	ID       int64
	Email    string
	Password string
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
