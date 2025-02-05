package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"field-service/constants"
	errConst "field-service/constants/error"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   *string     `json:"token,omitempty"`
}

type ParamHTTPResp struct {
	Code     int
	Error    error
	Messsage *string
	Gin      *gin.Context
	Data     interface{}
	Token    *string
}

func HttpResponse(param ParamHTTPResp) {

	if param.Error == nil {
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Success,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
			Token:   param.Token,
		})
		return
	}
	message := errConst.ErrInternalServerError.Error()
	if param.Messsage != nil {
		message = *param.Messsage
	} else if param.Error != nil {
		if errConst.ErrMapping(param.Error) {
			message = param.Error.Error()
		}
	}
	param.Gin.JSON(param.Code, Response{
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
		Token:   param.Token,
	})

}
