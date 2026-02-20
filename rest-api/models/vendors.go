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

	fmt.Println("Vendor saved: ", v)
	return nil
}

func GetVendor(id int64) (Vendor, error) {
	query := `
	SELECT * FROM Vendor
	WHERE id = ?
	`

	var v Vendor
	res := db.DB.QueryRow(query, id)
	res.Scan(&v.ID, &v.Name, &v.Description)

	return v, nil
}

func GetAllVendors() ([]Vendor, error) {
	query := `
	SELECT * FROM Vendor
	`

	res, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var vendorArr []Vendor

	for res.Next() {
		var v Vendor
		res.Scan(&v.ID, &v.Name, &v.Description)
		vendorArr = append(vendorArr, v)
	}

	return vendorArr, nil
}

func UpdateVendor(v Vendor) (any, error) {
	query := `
	UPDATE Vendor
	SET Name = ?, Description = ? 
	WHERE id = ?
	`

	res, err := db.DB.Exec(query, v.Name, v.Description, v.ID)
	if err != nil {
		fmt.Println(err)
	}

	return res, err
}
