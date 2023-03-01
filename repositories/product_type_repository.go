package repositories

import (
	"fmt"
	"github/go_ecommerce/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductTypeRepository interface {
	CreateProductType(*models.ProductType) error
	GetAllProductType() ([]models.ProductType, error)
}

type productTypeRepositoryDB struct {
	db *gorm.DB
}

func NewProductTypeRepositoryDB(db *gorm.DB) ProductTypeRepository {
	return productTypeRepositoryDB{db: db}
}

func (r productTypeRepositoryDB) CreateProductType(product *models.ProductType) error {
	err := r.db.Create(product).Error
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (r productTypeRepositoryDB) GetAllProductType() ([]models.ProductType, error) {
	productType := []models.ProductType{}
	err := r.db.Preload(clause.Associations).Find(&productType).Error
	if err != nil {
		return nil, err
	}
	return productType, nil
}
