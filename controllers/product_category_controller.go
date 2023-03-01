package controllers

import (
	"fmt"
	"github/go_ecommerce/dto"
	"github/go_ecommerce/response"
	"github/go_ecommerce/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productCategoryController struct {
	productCategoryService services.ProductCategoryService
}

func NewProductCategoryController(productCategoryService services.ProductCategoryService) productCategoryController {
	return productCategoryController{productCategoryService: productCategoryService}
}

// CreateProduct creates a new product.
//
//	@Summary		Create a new product Category.
//	@Description	Create a new product Category.
//	@Tags			product-Category
//	@Accept			multipart/form-data
//	@Produce		json
//
//	@Param			ProductCategory	formData	dto.ProductCategoryDTO	true	"Product Category name"
//	@Param			images			formData	file					true	"Product Category images"
//	@Security		Bearer
//	@Router			/product-category/create [post]
func (pc *productCategoryController) CreateProductCategory(ctx *gin.Context) {
	var dto dto.ProductCategoryDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pc.productCategoryService.CreateProductCategory(&dto, ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := response.Response{
		Result:  true,
		Status:  http.StatusOK,
		Message: "create product success",
	}
	ctx.JSON(http.StatusOK, res)
}

// GetProductTags		godoc
//
//	@Summary		get products Category
//	@Description	Get data all from table product Category
//	@Produce		application/json
//	@Tags			product-Category
//	@Success		200	{object}	response.Response{}
//	@Router			/product-category [get]
//
//	@Security		Bearer
func (pc *productCategoryController) GetAllProductCategory(ctx *gin.Context) {
	products, err := pc.productCategoryService.GetAllProductCategory()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := response.Response{
		Result:  true,
		Status:  http.StatusOK,
		Message: "get products success",
		Data:    products,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, res)
}
