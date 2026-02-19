package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createMenuItemsTable := `
	CREATE TABLE IF NOT EXISTS "MenuItems" (
	"ItemID"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT NOT NULL,
	"Price"	REAL NOT NULL,
	"VendorID"	INTEGER,
	PRIMARY KEY("ItemID" AUTOINCREMENT),
	FOREIGN KEY("VendorID") REFERENCES "Vendor"("ID")
	)
	`
	createVendorTable := `
	CREATE TABLE IF NOT EXISTS "Vendor" (
	"ID"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT NOT NULL,
	"Description"	TEXT,
	PRIMARY KEY("ID" AUTOINCREMENT)
	)
	`

	var err error
	_, err = DB.Exec(createMenuItemsTable)
	if err != nil {
		fmt.Println("Could not menu items table")
	}
	_, err = DB.Exec(createVendorTable)
	if err != nil {
		fmt.Println("Could not create vendor table")
	}
}
