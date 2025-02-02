package error

import "errors"

var ErrFieldStatusNotFound = errors.New("field status not found")

var ErrFieldStatusIsExists = errors.New("field status already exists")

var FieldScheduleErrors = []error{
	ErrFieldStatusNotFound,
	ErrFieldStatusIsExists,
}
