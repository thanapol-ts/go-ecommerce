package services

import (
	"github/go_ecommerce/dto"
	"github/go_ecommerce/models"
	"github/go_ecommerce/repositories"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductTypeService interface {
	CreateProductType(*dto.ProductTypeDTO, *gin.Context) error
	GetAllProductType() ([]models.ProductType, error)
}

type productTypeService struct {
	productTypeRepo repositories.ProductTypeRepository
}

func NewProductTypeService(productTypeRepo repositories.ProductTypeRepository) ProductTypeService {
	return productTypeService{productTypeRepo: productTypeRepo}
}

func (s productTypeService) CreateProductType(productTypeDTO *dto.ProductTypeDTO, ctx *gin.Context) error {
	fileType := strings.Split(productTypeDTO.Images.Filename, ".")[1]
	fileName := uuid.New().String() + "." + fileType
	if err := ctx.SaveUploadedFile(productTypeDTO.Images, "../assets/"+fileName); err != nil {
		return err
	}

	err := s.productTypeRepo.CreateProductType(&models.ProductType{
		Name:              productTypeDTO.Name,
		Description:       productTypeDTO.Description,
		FileName:          fileName,
		FileUrl:           "http://localhost:4000/assets/" + fileName,
		FileFormat:        fileType,
		Status:            "Active",
		ProductCategoryID: productTypeDTO.ProductCategoryId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s productTypeService) GetAllProductType() ([]models.ProductType, error) {
	products, err := s.productTypeRepo.GetAllProductType()
	if err != nil {
		return nil, err
	}

	return products, nil
}
