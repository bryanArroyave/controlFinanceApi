package utils

func VerifyError(err error, errorList *[]error, validateError bool) {
	if err != nil && validateError {
		*errorList = append(*errorList, err)
	}
}
