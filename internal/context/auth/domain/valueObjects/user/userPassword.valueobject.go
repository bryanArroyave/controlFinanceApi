package user

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type UserPassword struct {
	*valueobjects.StringValueObject
}

func NewUserPassword(value string) *UserPassword {
	return &UserPassword{
		valueobjects.NewStringValueObject(value).MinLength(5).Pattern("^[a-zA-Z0-9]*$").MaxLength(20).Optional(),
	}
}
