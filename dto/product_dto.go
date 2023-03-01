package dto

import "mime/multipart"

type ProductDTO struct {
	Name          string                  `form:"name" json:"name" binding:"required"`
	Description   string                  `form:"description" json:"description" binding:"required"`
	Price         string                  `form:"price" json:"price" `
	ProductTypeId uint                    `form:"product_type_id" json:"product_type_id" binding:"required"`
	Images        []*multipart.FileHeader `form:"images" json:"images" swaggerignore:"true"`
}
