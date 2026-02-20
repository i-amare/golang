package models

import (
	"fmt"

	"github.com/i-amare/rest-api/db"
)

type MenuItem struct {
	ID          string
	Name        string
	Description string
	Price       float32
}

type Vendor struct {
	ID          int64
	Name        string `binding:"required"`
	Description string
	Menu        []MenuItem
}

var vendors = []Vendor{}

func (v Vendor) Save() error {
	query := `
	INSERT INTO Vendor(Name, Description)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(v.Name, v.Description)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	v.ID = id

	vendors = append(vendors, v)
	fmt.Println("Vendor saved: ", id)

	return nil
}

func GetAllVendors() []Vendor {
	return vendors
}
