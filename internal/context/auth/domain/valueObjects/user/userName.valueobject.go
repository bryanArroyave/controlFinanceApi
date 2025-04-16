package user

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type UserName struct {
	*valueobjects.StringValueObject
}

func NewUserName(value string) *UserName {
	return &UserName{
		valueobjects.NewStringValueObject(value).MinLength(3).MaxLength(50),
	}
}
