package services

import (
	"github/go_ecommerce/dto"
	"github/go_ecommerce/models"
	"github/go_ecommerce/repositories"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductCategoryService interface {
	CreateProductCategory(*dto.ProductCategoryDTO, *gin.Context) error
	GetAllProductCategory() ([]models.ProductCategory, error)
}

type productCategoryService struct {
	productCategoryRepo repositories.ProductCategoryRepository
}

func NewProductCategoryService(productCategoryRepo repositories.ProductCategoryRepository) ProductCategoryService {
	return productCategoryService{productCategoryRepo: productCategoryRepo}
}

func (s productCategoryService) CreateProductCategory(productCategoryDTO *dto.ProductCategoryDTO, ctx *gin.Context) error {
	fileType := strings.Split(productCategoryDTO.Images.Filename, ".")[1]
	fileName := uuid.New().String() + "." + fileType
	if err := ctx.SaveUploadedFile(productCategoryDTO.Images, "../assets/"+fileName); err != nil {
		return err
	}

	err := s.productCategoryRepo.CreateProductCategory(&models.ProductCategory{
		Name:        productCategoryDTO.Name,
		Description: productCategoryDTO.Description,
		FileName:    fileName,
		FileUrl:     "http://localhost:4000/assets/" + fileName,
		FileFormat:  fileType,
		Status:      "Active",
	})

	if err != nil {
		return err
	}

	return nil
}

func (s productCategoryService) GetAllProductCategory() ([]models.ProductCategory, error) {
	products, err := s.productCategoryRepo.GetAllProductCategory()
	if err != nil {
		return nil, err
	}

	return products, nil
}
