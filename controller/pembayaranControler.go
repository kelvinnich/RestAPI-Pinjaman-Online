package controller

import (
	"net/http"
	"pinjaman-online/dto"
	"pinjaman-online/helper"
	"pinjaman-online/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PembayaranController interface{
	PembayaranPinjamanController(ctx *gin.Context)
	UpdatePembayaranController(ctx *gin.Context)
	ListPembayaranByStatusController(ctx *gin.Context)
}

type pembayaranController struct{
	pembayaranService service.PembayaranService
	jwtService service.JwtService
}

func NewPembayaranController(ps service.PembayaranService, js service.JwtService) PembayaranController{
	return &pembayaranController{
		pembayaranService: ps,
		jwtService: js,
	}
}

func(c *pembayaranController)PembayaranPinjamanController(ctx *gin.Context){
	var dtoPayment dto.CreatePembayaranDTO
	err := ctx.ShouldBind(&dtoPayment) 
	if err != nil {
		response := helper.ErrorResponse("failed to procces request pembayaran ", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	payment, err := c.pembayaranService.PembayaranPinjamanService(dtoPayment)
	if err != nil {
		response := helper.ErrorResponse("failed to create payment", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		
	}
	response := helper.ResponseOK(true, "OK!", payment)
	ctx.JSON(http.StatusOK, response)
}


func (c *pembayaranController) UpdatePembayaranController(ctx *gin.Context){
	var updateDTO dto.UpdatePembayaranDTO
	err := ctx.ShouldBind(&updateDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id,err := strconv.ParseInt(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	updateDTO.Id = int(id)
	update,err := c.pembayaranService.UpdatePembayaranService(updateDTO)
	response := helper.ResponseOK(true, "OK!", update)
	ctx.JSON(http.StatusOK, response)
}


func (c *pembayaranController) ListPembayaranByStatusController(ctx *gin.Context) {
	status := ctx.Param("status")
	pembayarans, err := c.pembayaranService.ListPembayaranByStatusService(status)
	if err != nil {
		response := helper.ErrorResponse( "Error fetching pembayarans", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseOK(true, "OK", pembayarans)
	ctx.JSON(http.StatusOK, response)
}

