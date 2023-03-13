package main

import (
	"github/go_ecommerce/docs"
	intializer "github/go_ecommerce/intializers"
	"github/go_ecommerce/models"
	"github/go_ecommerce/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	intializer.LoadEnv()
	intializer.ConnectionDB()
	intializer.DB.AutoMigrate(&models.Product{}, &models.ProductImage{}, &models.ProductType{}, &models.User{}, &models.Address{}, &models.ProductCategory{})
}

//	@title		Tag Service API
//	@version	1.0

//	@schemes	https

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	r := gin.Default()
	r.Static("/assets", "../assets")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = os.Getenv("IP_ADDRESS")
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Authorization"}
	r.Use(cors.New(config))
	routes.MainRoutes(r, intializer.DB)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":4000")

}
