package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidationResponse(err error) (validationResponse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			resHandlerErr := handleValidationError(err)
			validationResponse = append(validationResponse, resHandlerErr)
		}
	}
	return validationResponse
}

func handleValidationError(err validator.FieldError) ValidationResponse {
	switch err.Tag() {
	case "required":
		return ValidationResponse{
			Field:   err.Field(),
			Message: fmt.Sprintf("Field %s is required", err.Field()),
		}
	case "email":
		return ValidationResponse{
			Field:   err.Field(),
			Message: fmt.Sprintf("Field %s must be a valid email", err.Field()),
		}
	default:
		return handleDefaultValidationError(err)
	}
}

func handleDefaultValidationError(err validator.FieldError) ValidationResponse {
	ErrValidator, ok := ErrValidator[err.Tag()]
	if !ok {
		count := strings.Count(ErrValidator, "%s")
		if count == 1 {
			return ValidationResponse{
				Field:   err.Field(),
				Message: fmt.Sprintf(ErrValidator, err.Field()),
			}
		}
		return ValidationResponse{
			Field:   err.Field(),
			Message: fmt.Sprintf(ErrValidator, err.Field(), err.Param()),
		}
	}
	return ValidationResponse{
		Field:   err.Field(),
		Message: fmt.Sprintf("Something went wrong with field %s", err.Field()),
	}
}

func WrapError(err error) error {
	logrus.Errorf("error:  %v", err)
	return err

}
