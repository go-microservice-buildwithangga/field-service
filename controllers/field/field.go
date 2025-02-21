package controllers

import (
	errCommon "field-service/common/error"
	"field-service/common/response"
	"field-service/domain/dto"
	"field-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type FieldController struct {
	service services.IServiceRegistry
}

type IFieldController interface {
	GetAllWithPagination(ctx *gin.Context)
	GetAllWithoutPagination(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)

	Update(*gin.Context)

	Delete(*gin.Context)
}

func NewFieldController(service services.IServiceRegistry) *FieldController {
	return &FieldController{
		service: service,
	}
}

func (f *FieldController) Delete(ctx *gin.Context) {
	err := f.service.GetField().Delete(ctx, ctx.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Gin:  ctx,
	})
}

func (f *FieldController) Update(ctx *gin.Context) {
	var request dto.UpdateFieldRequest
	err := ctx.ShouldBindWith(&request, binding.FormMultipart)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		dataErr := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Data:    dataErr,
			Message: &errMessage,
			Error:   err,
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetField().Update(ctx, ctx.Param("uuid"), &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})

}

func (f *FieldController) Create(ctx *gin.Context) {
	var request dto.FieldRequest
	err := ctx.ShouldBindWith(&request, binding.FormMultipart)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Error:   err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     ctx,
		})

		return
	}

	result, err := f.service.GetField().Create(ctx, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Data: result,
		Gin:  ctx,
	})

}

func (f *FieldController) GetByUUID(ctx *gin.Context) {
	result, err := f.service.GetField().GetByUUID(ctx, ctx.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})

}

func (f *FieldController) GetAllWithoutPagination(ctx *gin.Context) {
	result, err := f.service.GetField().GetAllWithoutPagination(ctx)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})

}
func (f *FieldController) GetAllWithPagination(ctx *gin.Context) {
	var params dto.FieldRequestParam
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		errorMessage := http.StatusText(http.StatusBadRequest)
		errorResponse := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Error:   err,
			Message: &errorMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetField().GetAllWithPagination(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})

}
