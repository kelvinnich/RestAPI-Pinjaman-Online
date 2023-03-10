package controller

import (
	"net/http"
	"pinjaman-online/dto"
	"pinjaman-online/helper"
	"pinjaman-online/model"
	"pinjaman-online/service"
	"strconv"

	"github.com/gin-gonic/gin"
)


type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	
	authService service.AuthenticationService
	jwtService service.JwtService
}

func NewAuthController(authService service.AuthenticationService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginNasabahDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
	response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if authResult == nil {
	response := helper.ErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	return
	}
	v := authResult.(model.Master_Customer)
	generatedToken := c.jwtService.GenerateTokenService(strconv.FormatUint(v.Id, 10))
	v.Token = generatedToken
	response := helper.ResponseOK(true, "OK!", v)
	ctx.JSON(http.StatusOK, response)
	}

	
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterNasabahDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) || !c.authService.IsDuplicateNIk(registerDTO.NoKtp){
		response := helper.ErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObject{})
		ctx.JSON(http.StatusConflict, response)
		
	} else {
		createdCustomer := c.authService.CreateNasabah(registerDTO)
		token := c.jwtService.GenerateTokenService(strconv.FormatUint(createdCustomer.Id, 10))
		createdCustomer.Token = token
		response := helper.ResponseOK(true, "OK!", createdCustomer)
		ctx.JSON(http.StatusCreated, response)
	} 
}