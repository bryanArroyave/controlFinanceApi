package valueobjects

import (
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

type UserName struct {
	*valueobjects.StringValueObject
}

func NewUserName(value string) *UserName {
	return &UserName{
		valueobjects.NewStringValueObject(value).MinLength(5).MaxLength(30).Pattern("^[a-zA-Z][a-zA-Z0-9]*$"),
	}
}
