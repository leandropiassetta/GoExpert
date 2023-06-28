package models

import "gorm.io/gorm"

type Category struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:category_name"`
	Products []Product
}

// many products belongs to one category
type Product struct {
	ID           int          `gorm:"primaryKey"`
	Name         string       `gorm:"column:product_name"`
	Price        float64      `gorm:"column:product_price"`
	CategoryID   int          `gorm:"column:category_id"`
	Category     Category     `gorm:"foreignKey:CategoryID"`
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID"` // this is the relationship
	gorm.Model                // this is used to add the fields created_at, updated_at and deleted_at
}

// has one relationship -> 1 product has a serial number and this number is unique for each product and dont repeat in another product (1:1)
type SerialNumber struct {
	ID        int    `gorm:"primaryKey"` // this is the primary key of the table serial_number
	Number    string `gorm:"column:serial_number"`
	ProductID int    `gorm:"column:product_id"` // this is the foreign key of the table serial_number
}
