package repositories

import (
	"fmt"
	"github/go_ecommerce/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductCategoryRepository interface {
	CreateProductCategory(*models.ProductCategory) error
	GetAllProductCategory() ([]models.ProductCategory, error)
}

type productCategoryRepositoryDB struct {
	db *gorm.DB
}

func NewProductCategoryRepositoryDB(db *gorm.DB) ProductCategoryRepository {
	return productCategoryRepositoryDB{db: db}
}

func (r productCategoryRepositoryDB) CreateProductCategory(product *models.ProductCategory) error {
	err := r.db.Create(product).Error
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (r productCategoryRepositoryDB) GetAllProductCategory() ([]models.ProductCategory, error) {
	productCategory := []models.ProductCategory{}
	err := r.db.Preload(clause.Associations).Find(&productCategory).Error
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}
