package controller

import (
	"log"
	"net/http"
	"pinjaman-online/dto"
	"pinjaman-online/helper"
	"pinjaman-online/model"
	"pinjaman-online/service"
	"strconv"

	"github.com/gin-gonic/gin"
)


type HistoryPembayaranController interface{
	GetAllHistoryPembayaranController(ctx *gin.Context)
	GetHistoryPembayaranByIdController(ctx *gin.Context)
	UpdateHistoryPembayaranController(ctx *gin.Context)
	DeleteHistoryPembayaranController(ctx *gin.Context)
}

type historyPembayaranController struct {
	hps service.HistoryPembayaranService
	js service.JwtService
}


func NewHistoryPembayaranController(historyPembayaranService service.HistoryPembayaranService, jwtService service.JwtService) HistoryPembayaranController{
	return &historyPembayaranController{
		hps: historyPembayaranService,
		js: jwtService,
	}
}

func(c *historyPembayaranController) GetAllHistoryPembayaranController(ctx *gin.Context){
		history, err := c.hps.GetAllHistoryPembayaran()
		if err != nil {
			log.Printf("Error history controller %v", err)
			response := helper.ErrorResponse("faileld to procces request get all data", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseOK(true, "OK!", history)
		ctx.JSON(http.StatusOK, response)

}

func(c *historyPembayaranController) GetHistoryPembayaranByIdController(ctx *gin.Context){
	id,err := strconv.ParseUint(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to procces parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	history,err := c.hps.GetHistoryPembayaranByID(id)
	if err != nil {
		response := helper.ErrorResponse("failed to procces get by id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", history)
	ctx.JSON(http.StatusOK, response)
}

func(c *historyPembayaranController)UpdateHistoryPembayaranController(ctx *gin.Context){
	var updateDto dto.UpdateHistoryPembayaranDTO
	err := ctx.ShouldBind(&updateDto)
	if err != nil {
		response := helper.ErrorResponse("failed to procces ", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id,err := strconv.ParseUint(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to procces parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	updateDto.Id = id
	updateHistoryPembyaran,err := c.hps.UpdateHistoryPembayaranService(updateDto)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request update", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", updateHistoryPembyaran)
	ctx.JSON(http.StatusOK,response)
}

func (c *historyPembayaranController) DeleteHistoryPembayaranController(ctx *gin.Context){
	var deleteHistoryPembayaran model.Master_Payment_History
	id,err := strconv.ParseUint(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to procces parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	
	}

	deleteHistoryPembayaran.Id =  id
	delete := c.hps.DeleteHistoryPembayaranService(deleteHistoryPembayaran.Id)
	response := helper.ResponseOK(true, "OK!",delete )
	ctx.JSON(http.StatusOK, response)
}