package main

import (
	"github.com/leandropiassetta/goexpert/7-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	dsn := "root:root@tcp(localhost:3305)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{}, &models.Category{})

	// initialize transaction
	tx := db.Begin()

	var category models.Category

	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error

	if err != nil {
		panic(err)
	}

	category.Name = "Garden"
	tx.Debug().Save(&category)
	tx.Commit()
}
