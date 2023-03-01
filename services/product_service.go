package services

import (
	"errors"
	"fmt"
	"github/go_ecommerce/dto"
	"github/go_ecommerce/errs"
	"github/go_ecommerce/logs"
	"github/go_ecommerce/models"
	"github/go_ecommerce/repositories"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(*dto.ProductDTO, *gin.Context) error
	GetAllProduct() ([]models.Product, error)
	GetProductById(int) (*models.Product, error)
	UpdateProduct(int, *dto.ProductDTO, *gin.Context) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return productService{productRepo: productRepo}
}

func (s productService) CreateProduct(productDTO *dto.ProductDTO, ctx *gin.Context) error {
	fmt.Println("images", productDTO.Images)

	images := []models.ProductImage{}
	for _, file := range productDTO.Images {
		fileType := strings.Split(file.Filename, ".")[1]
		fileName := uuid.New().String() + "." + fileType

		if err := ctx.SaveUploadedFile(file, "../assets/"+fileName); err != nil {
			return err
		}

		images = append(images, models.ProductImage{
			FileName:   fileName,
			FileUrl:    "http://localhost:4000/assets/" + fileName,
			FileFormat: fileType,
		})
	}

	err := s.productRepo.CreateProduct(&models.Product{
		Name:          productDTO.Name,
		Description:   productDTO.Description,
		Price:         productDTO.Price,
		ProductTypeID: productDTO.ProductTypeId,
		Product_Image: images,
	})

	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotfoundError("product not found")
		}
		return err
	}
	return nil
}

func (s productService) GetAllProduct() ([]models.Product, error) {
	products, err := s.productRepo.GetAllProduct()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return products, nil
}

func (s productService) GetProductById(id int) (*models.Product, error) {
	product, err := s.productRepo.GetProductById(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotfoundError("product not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return product, nil
}

func (s productService) UpdateProduct(id int, productDTO *dto.ProductDTO, ctx *gin.Context) error {
	product, err := s.productRepo.GetProductById(id)
	if err != nil {
		return err
	}
	product.Name = productDTO.Name
	product.Description = productDTO.Description
	product.Price = productDTO.Price
	product.ProductTypeID = productDTO.ProductTypeId

	if len(productDTO.Images) > 0 {
		for _, file := range product.Product_Image {
			err := os.Remove("../assets/" + file.FileName)
			if err != nil {
				return err
			}
			errDeleteImage := s.productRepo.DeleteProductImage(int(file.ID))
			if errDeleteImage != nil {
				return errDeleteImage
			}
		}
		images := []models.ProductImage{}
		for _, file := range productDTO.Images {
			fileType := strings.Split(file.Filename, ".")[1]
			fileName := uuid.New().String() + "." + fileType

			if err := ctx.SaveUploadedFile(file, "../assets/"+fileName); err != nil {
				return err
			}

			images = append(images, models.ProductImage{
				FileName:   fileName,
				FileUrl:    "http://localhost:4000/assets/" + fileName,
				FileFormat: fileType,
			})
		}
		product.Product_Image = images
	}
	errUpdate := s.productRepo.UpdateProduct(product)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}
