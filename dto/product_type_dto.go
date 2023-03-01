package dto

import "mime/multipart"

type ProductTypeDTO struct {
	Name              string                `form:"name" json:"name" binding:"required"`
	Description       string                `form:"description" json:"description" binding:"required"`
	ProductCategoryId uint                  `form:"product_category_id" json:"product_category_id" binding:"required"`
	Images            *multipart.FileHeader `form:"images" json:"images" swaggerignore:"true"`
}
