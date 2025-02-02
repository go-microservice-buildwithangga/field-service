package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSqlError            = errors.New("database server failed to execute query")
	ErrTooManyRequest      = errors.New("too many request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
)

var GeneralError = []error{
	ErrInternalServerError,
	ErrSqlError,
	ErrTooManyRequest,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
