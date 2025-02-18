package controllers

import (
	errCommon "field-service/common/error"
	"field-service/common/response"
	"field-service/domain/dto"
	services "field-service/services/field"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldController struct {
	fieldService services.IFieldService
}

type IFieldController interface {
	GetAllWithPagination(ctx *gin.Context)
	GetAllWithoutPagination(ctx *gin.Context)
	GetByUUID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewFieldController(fieldService services.IFieldService) *FieldController {
	return &FieldController{
		fieldService: fieldService,
	}
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
			Code:     http.StatusBadRequest,
			Error:    err,
			Messsage: &errorMessage,
			Data:     errorResponse,
			Gin:      ctx,
		})
		return
	}

	return
	// result, err := f.GetAllWithPagination(ctx, &params)
	// if err != nil {
	// 	response.HttpResponse(response.ParamHTTPResp{
	// 		Code:  http.StatusInternalServerError,
	// 		Error: err,
	// 		Gin:   ctx,
	// 	})
	// 	return
	// }
	// response.HttpResponse(response.ParamHTTPResp{
	// 	Code:  http.StatusOK,
	// 	Error: nil,
	// 	Data:  result,
	// 	Gin:   ctx,
	// })

}
