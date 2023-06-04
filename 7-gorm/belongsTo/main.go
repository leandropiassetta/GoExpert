package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"column:category_name"`
	gorm.Model
}

// many products belongs to one category
type Product struct {
	ID         int      `gorm:"primaryKey"`
	Name       string   `gorm:"column:product_name"`
	Price      float64  `gorm:"column:product_price"`
	CategoryID int      `gorm:"column:category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// // create category
	category := Category{Name: "Electronic"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      5000.10,
		CategoryID: 1,
	})

	var products []Product

	// db.Preload("Category") -> join the table category with the table product and the preload is used to get the data of the table category in the table product
	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Price, product.Category.Name)
	}
}
