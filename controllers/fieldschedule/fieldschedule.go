package controllers

import (
	errCommon "field-service/common/error"
	"field-service/common/response"
	"field-service/domain/dto"
	"field-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldScheduleController struct {
	service services.IServiceRegistry
}

// Delete implements IFieldScheduleController.

// GenerateScheduleForOneMonth implements IFieldScheduleController.
func (f *FieldScheduleController) GenerateScheduleForOneMonth(ctx *gin.Context) {
	var params dto.GenerateFieldScheduleFromOneMonthRequest
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
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errData := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errData,
			Error:   err,
			Gin:     ctx,
		})
	}

	err := f.service.GetFieldSchedule().GenerateScheduleForOneMonth(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Gin:   ctx,
			Error: err,
		})
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Gin:  ctx,
	})

}

type IFieldScheduleController interface {
	GetAllWithPagination(*gin.Context)
	GetAllByFieldIDAndDate(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	UpdateStatus(*gin.Context)
	Delete(*gin.Context)
	GenerateScheduleForOneMonth(*gin.Context)
}

func NewFieldScheduleController(service services.IServiceRegistry) IFieldScheduleController {
	return &FieldScheduleController{
		service: service,
	}
}

// Create implements IFieldScheduleController.
func (f *FieldScheduleController) Create(ctx *gin.Context) {
	var request dto.FieldScheduleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errData := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Data:    errData,
			Message: &errMessage,
			Error:   err,
			Gin:     ctx,
		})
		return
	}

	err := f.service.GetFieldSchedule().Create(ctx, &request)
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
		Gin:  ctx,
	})
}

// GetByUUID implements IFieldScheduleController.
func (f *FieldScheduleController) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	result, err := f.service.GetFieldSchedule().GetByUUID(ctx, uuid)
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

// GetAllWithPagination implements IFieldScheduleController.
func (f *FieldScheduleController) GetAllWithPagination(ctx *gin.Context) {
	var params dto.FieldScheduleRequestParam
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
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errData := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Data:    errData,
			Error:   err,
			Message: &errMessage,
			Gin:     ctx,
		})

		return
	}

	result, err := f.service.GetFieldSchedule().GetAllWithPagination(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusInternalServerError,
			Gin:   ctx,
			Error: err,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
	return

}

// GetAllByFieldIDAndDate implements IFieldScheduleController.
func (f *FieldScheduleController) GetAllByFieldIDAndDate(ctx *gin.Context) {
	var params dto.FieldScheduleByFieldIDAndDateRequestParam
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(validate); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errData := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Error:   err,
			Message: &errMessage,
			Data:    errData,
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().GetAllByFieldIDAndDate(ctx, ctx.Param("uuid"), params.Date)
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

	return

}

func (f *FieldScheduleController) Update(ctx *gin.Context) {
	var params dto.UpdateFieldScheduleRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Error:   err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	result, err := f.service.GetFieldSchedule().Update(ctx, ctx.Param("uuid"), &params)
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
		Gin:  ctx,
		Data: result,
	})
}

func (f *FieldScheduleController) Delete(ctx *gin.Context) {
	err := f.service.GetFieldSchedule().Delete(ctx, ctx.Param("uuid"))
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
		Gin:  ctx,
	})
}

// UpdateStatus implements IFieldScheduleController.
func (f *FieldScheduleController) UpdateStatus(ctx *gin.Context) {
	var request dto.UpdateStatusScheduleRquest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code:  http.StatusBadRequest,
			Error: err,
			Gin:   ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errCommon.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Error:   err,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	err = f.service.GetFieldSchedule().UpdateStatus(ctx, &request)
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
		Gin:  ctx,
	})
}
