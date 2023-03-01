package dto

type RegisterDTO struct {
	Username   string `json:"username" binding:"required" `
	Password   string `json:"password" `
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    AddressDTO
}

type LoginDTO struct {
	Username string ` json:"username" binding:"required" example:"test1"`
	Password string ` json:"password" binding:"required" example:"12345"`
}

type AddressDTO struct {
	Address_line1 string `json:"address_line1"`
	Address_line2 string `json:"address_line2"`
	City          string `json:"city"`
	District      string `json:"district"`
	Sub_district  string `json:"sub_district"`
	Postcode      string `json:"postcode"`
}
