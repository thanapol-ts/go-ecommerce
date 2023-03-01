package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Password   string
	First_name string
	Last_name  string
	Email      string `gorm:"unique"`
	Phone      string
	Address    Address `gorm:"foreignKey:User_id"`
}

func (b *User) TableName() string {
	return "user"
}

type Address struct {
	ID            uint `gorm:"primarkey"`
	Address_line1 string
	Address_line2 string
	City          string
	District      string
	Sub_district  string
	Postcode      string
	User_id       uint
}

func (b *Address) TableName() string {
	return "address"
}
