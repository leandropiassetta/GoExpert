package database

import (
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) CreateProduct(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	// if page and limit are 0, it means that the user didn't send these parameters, so we will return all the products
	// if page and limit are not 0, we will return the products paginated and sorted by the created_at column (asc or desc)
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}

// create findall with

func (p *Product) FindProductByID(id string) (*entity.Product, error) {
	var product entity.Product

	err := p.DB.First(&product, "id = ?", id).Error

	// if the product is not found, gorm returns an error
	// &product is the pointer to the product variable that we want to fill with the data from the database
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	// check if the product exists in the database before updating it (if it doesn't exist, gorm returns an error)
	// if the product exists, gorm will update it
	_, err := p.FindProductByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	// check if the product exists in the database before deleting it (if it doesn't exist, gorm returns an error)
	// if the product exists, gorm will delete it
	product, err := p.FindProductByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
