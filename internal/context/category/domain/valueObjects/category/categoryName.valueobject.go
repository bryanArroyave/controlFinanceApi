package category

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type CategoryName struct {
	*valueobjects.StringValueObject
}

func NewCategoryName(value string) *CategoryName {
	return &CategoryName{
		valueobjects.NewStringValueObject(value).MinLength(3).MaxLength(50),
	}
}
