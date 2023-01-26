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

type PinjamanController interface{
	CreatePinjamanController(ctx *gin.Context)
	UpdatePinjamanController(ctx *gin.Context)
	SearchPinjamanByIdController(ctx *gin.Context)
	DeletePinjamanController(ctx *gin.Context)
	UpdateStatusApprovalPinjamanController(ctx *gin.Context)
}

type pinjamanController struct {
	pinjamanService service.PinjamanService
	jwtService service.JwtService
}

func NewPinajamanController(pj service.PinjamanService, js service.JwtService)PinjamanController{
	return &pinjamanController{
		pinjamanService: pj,
		jwtService: js,
	}
}

func(c *pinjamanController)CreatePinjamanController(ctx *gin.Context){
	var createPinjamanDTO dto.CreatePinjamanDTO
	err := ctx.ShouldBind(&createPinjamanDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}else {
		v,err := c.pinjamanService.CreatePinjamanService(createPinjamanDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to procces request create pinjaman", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		}else {
			response := helper.ResponseOK(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func(c *pinjamanController)UpdatePinjamanController(ctx *gin.Context){
	var updatePinjamanDTO dto.UpdatePinjamanDTO
	err := ctx.ShouldBind(&updatePinjamanDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}else {
		id,err := strconv.ParseUint(ctx.Param("id"),0,0)
		if err != nil {
			response := helper.ErrorResponse("failed to procces parse id", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
			return
		}

		updatePinjamanDTO.Id = id
		update,err := c.pinjamanService.UpdatePinjamanService(updatePinjamanDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to procces id not found", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		}else {
			response := helper.ResponseOK(true, "OK!", update)
			ctx.JSON(http.StatusOK,response)
		}
	}
}

func(c *pinjamanController)SearchPinjamanByIdController(ctx *gin.Context){
	id,err := strconv.ParseUint(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to procces parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	pinjaman, err := c.pinjamanService.SearchPinjamanByIdService(id)
	if (*pinjaman == model.Master_Loan{}){
		response := helper.ErrorResponse("failed to procces data id not found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else {
		response := helper.ResponseOK(true, "OK!", pinjaman)
		ctx.JSON(http.StatusOK, response)
	}
}

func(c *pinjamanController)DeletePinjamanController(ctx *gin.Context){
	var deletePinjaman model.Master_Loan
	id,err := strconv.ParseUint(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess request parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	deletePinjaman.Id = id
	deletes := c.pinjamanService.DeletePinjamanService(deletePinjaman.Id)
	response := helper.ResponseOK(true, "OK!", deletes)
	ctx.JSON(http.StatusOK, response)
}

func (c *pinjamanController) UpdateStatusApprovalPinjamanController(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess request parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	pinjaman, err := c.pinjamanService.UpdateLoanStatus(id)
	if err != nil {
		response := helper.ErrorResponse("failed to update loan approval status", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.ResponseOK(true, "OK!", pinjaman)
		ctx.JSON(http.StatusOK, response)
	}
}
