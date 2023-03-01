package repositories

import (
	"github/go_ecommerce/models"

	"gorm.io/gorm"
)

type AuthenRepository interface {
	Singup(*models.User) error
	Login(*models.User) (*models.User, error)
}

type authenRepositoryDB struct {
	db *gorm.DB
}

func NewAuthenRepositoryDB(db *gorm.DB) AuthenRepository {
	return authenRepositoryDB{db: db}
}

func (r authenRepositoryDB) Singup(user *models.User) error {

	err := r.db.Create(user).Error

	if err != nil {
		return err
	}
	return nil
}

func (r authenRepositoryDB) Login(user *models.User) (*models.User, error) {

	errCreate := r.db.First(user, "username = ?", user.Username).Error
	if errCreate != nil {
		return nil, errCreate
	}
	return user, nil
}
