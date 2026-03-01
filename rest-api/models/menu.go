package models

import (
	"fmt"

	"github.com/i-amare/rest-api/db"
)

type MenuItem struct {
	ID       int64
	Name     string
	Price    float64
	VendorID int64
}

func (m *MenuItem) Save() error {
	query := `
	INSERT INTO MenuItems (Name, Price, VendorID)
	VALUES (?, ?, ?)
	`

	res, err := db.DB.Exec(query, m.Name, m.Price, m.VendorID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	fmt.Println("ID: ", id)
	m.ID = id
	if err != nil {
		return err
	}

	_, err = GetVendor(m.VendorID)
	if err != nil {
		return err
	}

	fmt.Println("Menu item saved successfully")
	return nil
}
