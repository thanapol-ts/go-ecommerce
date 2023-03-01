package controllers

import (
	"fmt"
	"github/go_ecommerce/dto"
	"github/go_ecommerce/response"
	"github/go_ecommerce/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productTypeController struct {
	productTypeService services.ProductTypeService
}

func NewProductTypeController(productTypeService services.ProductTypeService) productTypeController {
	return productTypeController{productTypeService: productTypeService}
}

// CreateProduct creates a new product.
//
//	@Summary		Create a new product.
//	@Description	Create a new product.
//	@Tags			product-type
//	@Accept			multipart/form-data
//	@Produce		json
//
//	@Param			ProductType	formData	dto.ProductTypeDTO	true	"Product Type name"
//	@Param			images		formData	file				true	"Product Type images"
//	@Security		Bearer
//	@Router			/product-type/create [post]
func (pc *productTypeController) CreateProductType(ctx *gin.Context) {
	dto := dto.ProductTypeDTO{}
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pc.productTypeService.CreateProductType(&dto, ctx)
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
//	@Summary		get products
//	@Description	Get data all from table product
//	@Produce		application/json
//	@Tags			product-type
//	@Success		200	{object}	response.Response{}
//	@Router			/product-type [get]
//
//	@Security		Bearer
func (pc *productTypeController) GetAllProductType(ctx *gin.Context) {
	products, err := pc.productTypeService.GetAllProductType()

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
