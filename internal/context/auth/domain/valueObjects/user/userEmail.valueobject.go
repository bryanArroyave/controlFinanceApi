package user

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type UserEmail struct {
	*valueobjects.StringValueObject
}

func NewUserEmail(value string) *UserEmail {
	return &UserEmail{
		valueobjects.NewStringValueObject(value).Pattern(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	}
}
