package dto

import "mime/multipart"

type ProductCategoryDTO struct {
	Name        string                `form:"name" json:"name" binding:"required"`
	Description string                `form:"description" json:"description" binding:"required"`
	Images      *multipart.FileHeader `form:"images" json:"images" swaggerignore:"true"`
}
