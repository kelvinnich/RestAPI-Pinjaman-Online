package controller

import (
	"fmt"
	"net/http"
	"pinjaman-online/dto"
	"pinjaman-online/helper"
	"pinjaman-online/service"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


type NasabahController interface {
	UpdateNasabahController(context *gin.Context)
	ProfileNasabahController(context *gin.Context)
}

type nasabahController struct {
	nasabahService service.NasabahServic
	jwtService      service.JwtService
}

func NewNasabahController(nasabahservice service.NasabahServic, jwtService service.JwtService) NasabahController {
	return &nasabahController{
		nasabahService: nasabahservice,
		jwtService:      jwtService,
	}
}

func (c *nasabahController) UpdateNasabahController(context *gin.Context) {
	var customerUpdateDTO dto.UpdateNasabahDTO
	err := context.ShouldBind(&customerUpdateDTO)
	if err != nil {
		response := helper.ErrorResponse("Fail to process request", err.Error(), helper.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateTokenService(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["customer_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	customerUpdateDTO.Id = id
	customer := c.nasabahService.UpdateNasabah(customerUpdateDTO)
	response := helper.ResponseOK(true, "OK!", customer)
	context.JSON(http.StatusOK, response)
}

func (c *nasabahController) ProfileNasabahController(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateTokenService(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	customer := c.nasabahService.ProfileNasabah(id)
	res := helper.ResponseOK(true, "OK!", customer)
	context.JSON(http.StatusOK, res)
}