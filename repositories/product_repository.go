package repositories

import (
	"fmt"
	"github/go_ecommerce/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	CreateProduct(*models.Product) error
	GetAllProduct() ([]models.Product, error)
	GetProductById(int) (*models.Product, error)
	UpdateProduct(*models.Product) error
	DeleteProductImage(int) error
}

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	return productRepositoryDB{db: db}
}

func (r productRepositoryDB) CreateProduct(product *models.Product) error {

	err := r.db.Create(product).Error
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func (r productRepositoryDB) GetAllProduct() ([]models.Product, error) {

	products := []models.Product{}
	err := r.db.Preload(clause.Associations).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r productRepositoryDB) GetProductById(id int) (*models.Product, error) {

	products := models.Product{}
	err := r.db.Preload(clause.Associations).First(&products, id).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (r productRepositoryDB) UpdateProduct(product *models.Product) error {

	err := r.db.Save(product).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r productRepositoryDB) DeleteProductImage(id int) error {

	err := r.db.Delete(&models.ProductImage{}, id).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
