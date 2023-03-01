package services

import (
	"fmt"
	"github/go_ecommerce/dto"
	"github/go_ecommerce/errs"
	"github/go_ecommerce/logs"
	"github/go_ecommerce/models"
	"github/go_ecommerce/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginModel struct {
	Access_token string ` json:"access_token"`
}

type AuthenService interface {
	Signup(*dto.RegisterDTO) error
	Login(*dto.LoginDTO) (*LoginModel, error)
}

type authenService struct {
	authenRepo repositories.AuthenRepository
}

func NewAuthenService(authenRepo repositories.AuthenRepository) AuthenService {
	return authenService{authenRepo: authenRepo}
}

func (s authenService) Signup(registerDTO *dto.RegisterDTO) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), 10)
	if err != nil {
		return err
	}

	errSignup := s.authenRepo.Singup(&models.User{
		Username:   registerDTO.Username,
		Password:   string(hash),
		Email:      registerDTO.Email,
		First_name: registerDTO.First_name,
		Last_name:  registerDTO.Last_name,
		Phone:      registerDTO.Phone,
		Address: models.Address{
			Address_line1: registerDTO.Address.Address_line1,
			Address_line2: registerDTO.Address.Address_line2,
			City:          registerDTO.Address.City,
			District:      registerDTO.Address.District,
			Sub_district:  registerDTO.Address.Sub_district,
			Postcode:      registerDTO.Address.Postcode,
		},
	})

	if errSignup != nil {
		logs.Error(errSignup)
		return errSignup
	}
	return nil
}

func (s authenService) Login(loginDTO *dto.LoginDTO) (*LoginModel, error) {

	user, err := s.authenRepo.Login(&models.User{
		Username: loginDTO.Username,
		Password: loginDTO.Password,
	})
	logs.Info(user.Password)
	logs.Info(loginDTO.Password)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password))

	if errPassword != nil {
		fmt.Println("errPassword = ", errPassword)
		return nil, errPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		"user": user.Username,
	})
	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECERT")))
	logs.Info(tokenString)

	if errToken != nil {
		logs.Error(errToken)
		return nil, err
	}

	return &LoginModel{
		Access_token: tokenString,
	}, nil
}
