package category

import (
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

type CategoryType struct {
	*valueobjects.StringValueObject
}

func NewCategoryType(value string) *CategoryType {
	return &CategoryType{
		valueobjects.NewStringValueObject(value).Include("income", "expense"),
	}
}
