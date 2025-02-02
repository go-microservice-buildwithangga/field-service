package error

func ErrMapping(err error) bool {
	allErrors := GeneralError[:]
	for _, item := range allErrors {
		if item.Error() == err.Error() {
			return true
		}
	}
	return false

}
