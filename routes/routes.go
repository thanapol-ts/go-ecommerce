package routes

import (
	"github/go_ecommerce/controllers"
	"github/go_ecommerce/middleware"
	"github/go_ecommerce/repositories"
	"github/go_ecommerce/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRoutes(r *gin.Engine, db *gorm.DB) {
	//product controller
	productRepositoryDB := repositories.NewProductRepositoryDB(db)
	productService := services.NewProductService(productRepositoryDB)
	productController := controllers.NewProductController(productService)
	//producttype controller
	productTypeRepositoryDB := repositories.NewProductTypeRepositoryDB(db)
	productTypeService := services.NewProductTypeService(productTypeRepositoryDB)
	productTypeController := controllers.NewProductTypeController(productTypeService)
	//productcategory controller
	productCategoryRepositoryDB := repositories.NewProductCategoryRepositoryDB(db)
	productCategoryService := services.NewProductCategoryService(productCategoryRepositoryDB)
	productCategoryController := controllers.NewProductCategoryController(productCategoryService)

	authenRepositoryDB := repositories.NewAuthenRepositoryDB(db)
	authenService := services.NewAuthenService(authenRepositoryDB)
	authenController := controllers.NewAuthenController(authenService)

	// Create routes group.
	route := r.Group("/api/v1")

	// Routes for GET method:
	route.GET("/product", middleware.JWTAuthMiddleware(), productController.GetProducts)
	route.GET("/product/:id", middleware.JWTAuthMiddleware(), productController.GetProductById)
	route.POST("/product/create", middleware.JWTAuthMiddleware(), productController.CreateProduct)
	route.PATCH("/product/update/:id", middleware.JWTAuthMiddleware(), productController.UpdateProduct)

	route.GET("/product-type", middleware.JWTAuthMiddleware(), productTypeController.GetAllProductType)
	route.POST("/product-type/create", middleware.JWTAuthMiddleware(), productTypeController.CreateProductType)

	route.GET("/product-category", middleware.JWTAuthMiddleware(), productCategoryController.GetAllProductCategory)
	route.POST("/product-category/create", middleware.JWTAuthMiddleware(), productCategoryController.CreateProductCategory)

	route.POST("/auth/login", authenController.Login)
	route.POST("/auth/signup", authenController.Singup)

}
