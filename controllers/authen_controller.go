package controllers

import (
	"github/go_ecommerce/dto"
	"github/go_ecommerce/response"
	"github/go_ecommerce/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authenController struct {
	authenService services.AuthenService
}

func NewAuthenController(authenService services.AuthenService) authenController {
	return authenController{authenService: authenService}
}

//	@Create			Signup
//	@Router			/auth/signup [post]
//	@Tags			User
//	@Description	create user
//	@Param			User	body	dto.RegisterDTO	true	"User"
func (ac *authenController) Singup(ctx *gin.Context) {
	var user dto.RegisterDTO
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.NewResponseError(false, ctx, http.StatusBadRequest, err.Error())
		return
	}
	err := ac.authenService.Signup(&user)
	if err != nil {
		response.NewResponseError(true, ctx, http.StatusConflict, err.Error())
		return
	}

	response.NewResponseSuccess(true, ctx, http.StatusOK, "signup success")

}

//	@Create			login
//	@Router			/auth/login [post]
//	@Tags			User
//	@Description	login
//	@Param			User	body	dto.LoginDTO	true	"User"
func (ac *authenController) Login(ctx *gin.Context) {
	var user dto.LoginDTO
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.NewResponseError(false, ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := ac.authenService.Login(&user)
	if err != nil {
		response.NewResponseError(true, ctx, http.StatusUnauthorized, "login failed")
		return
	}
	response.NewResponseSuccessWithData(true, ctx, http.StatusOK, "login success", token)
	return
}
