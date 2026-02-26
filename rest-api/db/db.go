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
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS "Users" (
	"ID"	INTEGER NOT NULL,
	"Email"	TEXT NOT NULL UNIQUE,
	"Password"	TEXT NOT NULL,
	PRIMARY KEY("ID" AUTOINCREMENT)
	)
	`

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
	CREATE TABLE IF NOT EXISTS "Vendors" (
	"ID"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	"Description"	TEXT NOT NULL,
	"OwnerID"	INTEGER,
	PRIMARY KEY("ID" AUTOINCREMENT),
	FOREIGN KEY("OwnerID") REFERENCES "Users"("ID")
)
	`

	var err error
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println("Could not create users table")
	}

	_, err = DB.Exec(createMenuItemsTable)
	if err != nil {
		fmt.Println("Could not create menu items table")
	}
	_, err = DB.Exec(createVendorTable)
	if err != nil {
		fmt.Println("Could not create vendor table")
	}
}
