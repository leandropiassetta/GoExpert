package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// the tags are used to define the name of the column in the database and the primary key of the table in the database and the auto increment of the primary key
// the tags are used by the gorm to create the table in the database
// the gorm is a library util to work with the database in go
type Product struct {
	ID    int     `gorm:"primaryKey"`
	Name  string  `gorm:"column:product_name"`
	Price float64 `gorm:"column:product_price"`
}

func main() {
	// my connection with the database
	// dsn is a string that contains the connection information
	// dsn -> data source name
	dsn := "root:root@tcp(localhost:3306)/goexpert"

	// open the connection with the database

	// gorm.Open() -> open the connection with the database
	// mysql.Open() -> define the driver of the database
	// driver is  what is used to connect with the database (mysql, postgres, sqlite, etc).
	// &gorm.Config{} -> define the configuration of the connection with the database

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
