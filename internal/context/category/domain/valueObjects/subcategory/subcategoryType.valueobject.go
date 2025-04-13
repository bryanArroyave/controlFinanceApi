package subcategory

import (
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

type SubcategoryType struct {
	*valueobjects.StringValueObject
}

func NewSubcategoryType(value string) *SubcategoryType {
	return &SubcategoryType{
		valueobjects.NewStringValueObject(value).Include("income", "expense"),
	}
}
