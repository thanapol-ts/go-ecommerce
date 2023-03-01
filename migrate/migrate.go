package main

import (
	intializer "github/go_ecommerce/intializers"
	"github/go_ecommerce/models"
)

func init() {
	// intializer.LoadEnv()
	intializer.ConnectionDB()
}

func main() {
	intializer.DB.AutoMigrate(&models.Product{}, &models.ProductImage{}, &models.ProductType{}, &models.User{}, &models.Address{}, &models.ProductCategory{})
}
