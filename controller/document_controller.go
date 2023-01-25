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

type DocumentNasabahController interface{
	UploadDocumentController(ctx *gin.Context)
	UpdateDocumentController(ctx *gin.Context)
	FindDocumentByIdController(ctx *gin.Context)
	DeleteDocumentController(ctx *gin.Context)
}

type documentNasabahController struct{
	documentService service.DocumentService
	jwtService service.JwtService
}

func NewDocumentController(ds service.DocumentService, js service.JwtService)DocumentNasabahController{
	return &documentNasabahController{
		documentService: ds,
		jwtService: js,
	}
}

func(c *documentNasabahController)UploadDocumentController(ctx *gin.Context){
	var documentDto dto.CreateDocumentNasabahDTO
	err := ctx.ShouldBind(&documentDto)
	if err != nil {
		response := helper.ErrorResponse("failed to procces request upload document", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	v := c.documentService.UploadDocument(documentDto)
	response := helper.ResponseOK(true, "OK!", v)
	ctx.JSON(http.StatusOK, response)
}

func(c *documentNasabahController)UpdateDocumentController(ctx *gin.Context){
	var documentDto dto.UpdateDocumentNasabahDTO
	err := ctx.ShouldBind(&documentDto)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess request update document", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id,err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	documentDto.Id = id
	updateDocument := c.documentService.UpdateDocument(documentDto)
	response := helper.ResponseOK(true, "OK!", updateDocument)
	ctx.JSON(http.StatusOK, response)
}

func(c *documentNasabahController)FindDocumentByIdController(ctx *gin.Context){
	id,err := strconv.ParseUint(ctx.Param("id"),0,0)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess find docs by id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	
	 docs, err := c.documentService.GetDocumentById(id)
	if (*docs == model.Master_Document_Customer{}){
		response := helper.ErrorResponse("failed to procces data id not found", "this data is not the same", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else {
		response := helper.ResponseOK(true, "OK!", docs)
		ctx.JSON(http.StatusOK, response)
	}
}

func(c *documentNasabahController)DeleteDocumentController(ctx *gin.Context){
	var docs model.Master_Document_Customer
	id,err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to proccess request delete docs", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	docs.Id = id
	deletes := c.documentService.DeleteDocument(docs.Id)
	response := helper.ResponseOK(true, "OK!", deletes)
	ctx.JSON(http.StatusOK,response)
}