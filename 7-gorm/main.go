package main

import (
	"fmt"

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
	dsn := "root:root@tcp(localhost:3305)/goexpert"

	// open the connection with the database

	// gorm.Open() -> open the connection with the database
	// mysql.Open() -> define the driver of the database
	// driver is  what is used to connect with the database (mysql, postgres, sqlite, etc).
	// &gorm.Config{} -> define the configuration of the connection with the database

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate() -> create the table in the database with the struct Product and the tags of the struct Product are used to create the table in the database with the name of the column and the primary key of the table and the auto increment of the primary key of the table in the database and the type of the column in the database (varchar, int, float, etc) and the size of the column in the database (varchar(255), int(11), etc) and the default value of the column in the database (default 0, default 'name', etc) and the not null of the column in the database (not null, null, etc)
	db.AutoMigrate(&Product{})

	// Create() -> create the record in the database
	// &Product{} -> define the record that will be created in the database

	db.Create(&Product{
		Name:  "Television",
		Price: 2000.10,
	})

	// create batch of records in the database with the same struct
	products := []Product{
		{Name: "Notebook", Price: 3000.10},
		{Name: "Smartphone", Price: 1000.10},
		{Name: "Tablet", Price: 1500.10},
	}

	db.Create(&products)

	// select one
	var product Product

	// db.First(&product, 2)
	// fmt.Println(product)

	db.First(&product, "product_name = ?", "Smartphone")
	fmt.Println(product)

	// select all
	var products2 []Product
	db.Find(&products2)

	for _, product := range products2 {
		fmt.Println(product)
	}
}
