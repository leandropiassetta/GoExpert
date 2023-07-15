package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// many to many relationship is the form more complex of relationship between tables in the database.
// many to many relationship is used when one record of a table can be related with many records of another table and one record of another table can be related with many records of the first table.

// in this cases is necessary create a table that will be the relationship between the two tables.

// table intermediary (relationship) -> table that will be the relationship between the two tables

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // intermediary table
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"` // intermediary table
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// create category

	category1 := Category{Name: "Kitchen"}
	category2 := Category{Name: "Electronics"}

	db.Create(&category1)
	db.Create(&category2)

	// create product

	db.Create(&Product{Name: "Microwave", Price: 100.00, Categories: []Category{category1, category2}})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
}
