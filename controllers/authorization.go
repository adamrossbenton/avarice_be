package controllers

import (
	"avarice/dto"
	"avarice/svc"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService	service.LoginService
	jwtService		service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService:	loginService,
		jwtService:		jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Username, true)
	}
	return ""
}