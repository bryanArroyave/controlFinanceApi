package valueobjects

import (
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

type UsereEmail struct {
	*valueobjects.StringValueObject
}

func NewUsereEmail(value string) *UsereEmail {
	return &UsereEmail{
		valueobjects.NewStringValueObject(value).MaxLength(100).Pattern("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"),
	}
}
