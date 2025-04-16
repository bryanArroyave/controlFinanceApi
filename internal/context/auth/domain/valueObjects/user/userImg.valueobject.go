package user

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type UserImg struct {
	*valueobjects.StringValueObject
}

func NewUserImg(value string) *UserImg {
	return &UserImg{
		valueobjects.NewStringValueObject(value),
	}
}
