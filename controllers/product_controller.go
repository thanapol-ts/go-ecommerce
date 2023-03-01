package controllers

import (
	"fmt"
	"github/go_ecommerce/dto"
	"github/go_ecommerce/errs"
	"github/go_ecommerce/logs"
	"github/go_ecommerce/response"
	"github/go_ecommerce/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) productController {
	return productController{productService: productService}
}

// CreateProduct creates a new product.
//
//	@Summary		Create a new product.
//	@Description	Create a new product.
//	@Tags			product
//	@Accept			multipart/form-data
//	@Produce		json
//
//	@Param			Product	formData	dto.ProductDTO	true	"Product name"
//	@Param			images	formData	[]file			true	"Product images"
//
//	@Router			/product/create [post]
//
//	@Security		Bearer
func (pc *productController) CreateProduct(ctx *gin.Context) {
	var dto dto.ProductDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		appErr, ok := err.(errs.AppError)
		if ok {
			logs.Error(appErr)
			response.NewResponseError(false, ctx, appErr.Code, appErr.Message)
			return
		}
	}
	err := pc.productService.CreateProduct(&dto, ctx)
	if err != nil {
		appErr, ok := err.(errs.AppError)
		if ok {
			logs.Error(appErr)
			response.NewResponseError(false, ctx, appErr.Code, appErr.Message)
			return
		} else {
			response.NewResponseError(false, ctx, 400, err.Error())
			return
		}
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
//	@Tags			product
//	@Success		200	{object}	response.Response{}
//	@Router			/product [get]
//
//	@Security		Bearer
func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productService.GetAllProduct()

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

// GetProductTags		godoc
//
//	@Summary		get product by id
//	@Description	Get data by id from table product
//	@Produce		application/json
//	@Tags			product
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	response.Response{}
//	@Router			/product/:id [get]
//
//	@Security		Bearer
func (pc *productController) GetProductById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	products, err := pc.productService.GetProductById(id)

	if err != nil {

		appErr, ok := err.(errs.AppError)
		if ok {
			logs.Error(appErr)
			response.NewResponseError(false, ctx, appErr.Code, appErr.Message)
			return
		}

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

// UpdateProductTags		godoc
//
//	@Summary		update product
//	@Description	update peoduct
//	@Accept			multipart/form-data
//	@Produce		json
//	@Tags			product
//	@Param			id		path		int				true	"Product ID"
//	@Param			Product	formData	dto.ProductDTO	true	"Product name"
//	@Param			images	formData	[]file			false	"Product images"
//	@Success		200		{object}	response.Response{}
//	@Router			/product/update/:id [PATCH]
//
//	@Security		Bearer
func (pc *productController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var dto dto.ProductDTO
	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := pc.productService.UpdateProduct(id, &dto, ctx)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := response.Response{
		Result:  true,
		Status:  http.StatusOK,
		Message: "get products success",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, res)
}
