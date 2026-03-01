package models

import (
	"fmt"

	"github.com/i-amare/rest-api/db"
)

type Vendor struct {
	ID          int64
	Name        string `binding:"required"`
	Description string
	Menu        []MenuItem
}

func (v *Vendor) Save() error {
	query := `
	INSERT INTO Vendors (Name, Description)
	VALUES (?, ?)
	`

	res, err := db.DB.Exec(query, v.Name, v.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	v.ID = id

	fmt.Println("Vendor saved: ", v)
	return nil
}

func GetVendor(id int64) (Vendor, error) {
	query := `
	SELECT ID, Name, Description
	FROM Vendors
	WHERE ID = ?
	`

	var v Vendor
	res := db.DB.QueryRow(query, id)
	err := res.Scan(&v.ID, &v.Name, &v.Description)
	if err != nil {
		return Vendor{}, err
	}

	query = `
	SELECT ItemID, Name, Price 
	FROM MenuItems
	WHERE VendorID = ?
	`

	menu := []MenuItem{}
	row, err := db.DB.Query(query, id)
	if err != nil {
		return v, nil
	}
	defer row.Close()

	for row.Next() {
		var m MenuItem
		if err := row.Scan(&m.ID, &m.Name, &m.Price); err != nil {
			fmt.Println("MenuItem: ", m)
			fmt.Println(err.Error())
			fmt.Println("s")
			continue
		}
		menu = append(menu, m)
	}
	v.Menu = menu

	return v, nil
}

func GetAllVendors() ([]Vendor, error) {
	query := `
	SELECT ID, Name, Description
	FROM Vendors
	`

	res, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var vendorArr []Vendor

	query = `
	SELECT ItemID, Name, Price 
	FROM MenuItems
	WHERE VendorID = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return []Vendor{}, err
	}

	for res.Next() {
		var v Vendor
		if err := res.Scan(&v.ID, &v.Name, &v.Description); err != nil {
			return nil, err
		}

		menu := []MenuItem{}
		row, err := stmt.Query(v.ID)
		if err != nil {
			continue
		}
		defer row.Close()

		for row.Next() {
			var m MenuItem
			if err := row.Scan(&m.ID, &m.Name, &m.Price); err != nil {
				fmt.Println("MenuItem: ", m)
				fmt.Println(err.Error())
				fmt.Println("s")
				continue
			}
			menu = append(menu, m)
		}
		v.Menu = menu

		vendorArr = append(vendorArr, v)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return vendorArr, nil
}

func UpdateVendor(v Vendor) (any, error) {
	query := `
	UPDATE Vendors
	SET Name = ?, Description = ? 
	WHERE id = ?
	`

	res, err := db.DB.Exec(query, v.Name, v.Description, v.ID)
	if err != nil {
		fmt.Println(err)
	}

	return res, err
}

func DeleteVendor(id int64) error {
	query := `
	DELETE FROM Vendors
	WHERE id = ?
 	`

	_, err := db.DB.Exec(query, id)
	if err != nil {
		fmt.Println(id)
		return err
	}

	return nil
}
