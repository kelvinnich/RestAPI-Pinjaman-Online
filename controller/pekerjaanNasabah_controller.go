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

type PekerjaanNasabahController interface{
	AddCustomerJobsController(ctx *gin.Context)
	CustomerUpdateJobsController(ctx *gin.Context)
	SearchForCustomerJobsByIdController(ctx *gin.Context)
	DeleteCustomerJobsController(ctx *gin.Context)
}


type pekerjaanNasabahController struct{
	customerJobsService service.PekerjaanNasabahService
	jwtService service.JwtService
}

func NewPekerjaanNasabahController(customerJobs service.PekerjaanNasabahService, jwtService service.JwtService)PekerjaanNasabahController{
	return &pekerjaanNasabahController{
		customerJobsService: customerJobs,
		jwtService: jwtService,
	}
}

func(c *pekerjaanNasabahController) AddCustomerJobsController(ctx *gin.Context){
	var addJobsCustomerDTO dto.CreatePekerjaanNasabahDTO
	err := ctx.ShouldBind(&addJobsCustomerDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}else {
		v, err := c.customerJobsService.AddCustomerJobsService(addJobsCustomerDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to procces request add jobs customer", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		}else {
			response := helper.ResponseOK(true, "OK!", v)
			ctx.JSON(http.StatusOK,response)
		}
	}
}

func(c *pekerjaanNasabahController) CustomerUpdateJobsController(ctx *gin.Context){
	var updateJobsCustomerDTO dto.UpdatePekerjaanNasabahDTO
	err := ctx.ShouldBind(&updateJobsCustomerDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	id,err := strconv.ParseInt(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request update customer jobs", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	updateJobsCustomerDTO.Id = int(id)
	update,_ := c.customerJobsService.CustomersJobsUpdateService(updateJobsCustomerDTO)
	response := helper.ResponseOK(true, "OK!", update)
	ctx.JSON(http.StatusOK,response)
}

func(c *pekerjaanNasabahController)SearchForCustomerJobsByIdController(ctx *gin.Context){
	id,err := strconv.ParseInt(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	jobs,err := c.customerJobsService.SearchForCustomerJobsByIdService(int(id))
	if (jobs == &model.Master_Jobs_customers{}){
		response := helper.ErrorResponse("failed to procces data id not found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else {
		response := helper.ResponseOK(true, "OK!", jobs)
		ctx.JSON(http.StatusOK, response)
	}
}

func(c *pekerjaanNasabahController)DeleteCustomerJobsController(ctx *gin.Context){
	var deleteJobsNasabah model.Master_Jobs_customers
	id,err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess request delete jobs customer", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	deleteJobsNasabah.ID = int(id)
	deletes := c.customerJobsService.DeleteCustomerJobsService(deleteJobsNasabah.ID)
	response := helper.ResponseOK(true, "OK!", deletes)
	ctx.JSON(http.StatusOK, response)
}