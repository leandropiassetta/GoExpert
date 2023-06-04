package main

import (
	"fmt"

	"github.com/leandropiassetta/goexpert/7-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{}, &models.Category{})

	// create category
	category := models.Category{Name: "Kitchen"}
	db.Create(&category)

	// create product
	db.Create(&models.Product{
		Name:       "Knife",
		Price:      10.5,
		CategoryID: 2,
	})

	var categories []models.Category

	// db.Model(&models.Category{}).Preload("Products") -> join the table category with the table product and the preload is used to get the data of the table product in the table category
	err = db.Model(&models.Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// print the products of each category (1 category has many products)
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name, category.Name)
		}
	}
}
