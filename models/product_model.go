package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Description   string
	Price         string
	Product_Image []ProductImage `gorm:"foreignKey:ProductID"`
	ProductType   ProductType
	ProductTypeID uint
}

type ProductImage struct {
	ID         uint `gorm:"primaryKey"`
	FileName   string
	FileUrl    string
	FileFormat string
	ProductID  uint
}

type ProductType struct {
	gorm.Model
	Name              string
	Description       string
	FileName          string
	FileUrl           string
	FileFormat        string
	Status            string
	ProductCategory   ProductCategory
	ProductCategoryID uint
}

type ProductCategory struct {
	gorm.Model
	Name        string
	Description string
	FileName    string
	FileUrl     string
	FileFormat  string
	Status      string
}

func (b *Product) TableName() string {
	return "product"
}

func (b *ProductImage) TableName() string {
	return "productImage"
}

func (b *ProductType) TableName() string {
	return "productType"
}
