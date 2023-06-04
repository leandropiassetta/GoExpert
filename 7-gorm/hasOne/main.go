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

	db.AutoMigrate(&models.Product{}, &models.Category{}, &models.SerialNumber{})

	// create category
	category := models.Category{Name: "Electronic"}
	db.Create(&category)

	// create product
	db.Create(&models.Product{
		Name:       "smartphone",
		Price:      2000.10,
		CategoryID: 1,
	})

	// create serial number
	db.Create(&models.SerialNumber{
		Number:    "123456789",
		ProductID: 1,
	})

	var products []models.Product

	// db.Preload("Category") -> join the table category with the table product and the preload is used to get the data of the table category in the table product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Price, product.Category.Name, product.SerialNumber.Number)
	}
}
