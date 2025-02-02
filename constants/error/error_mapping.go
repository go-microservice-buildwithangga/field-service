package error

import (
	errField "field-service/constants/error/field"
	errFieldSchedule "field-service/constants/error/fieldschedule"
)

func ErrMapping(err error) bool {
	allErrors := append(append(GeneralError[:], errField.FieldErrors[:]...), errFieldSchedule.FieldScheduleErrors[:]...)
	for _, item := range allErrors {
		if item.Error() == err.Error() {
			return true
		}
	}
	return false

}
