package subcategory

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type SubcategoryName struct {
	*valueobjects.StringValueObject
}

func NewSubcategoryName(value string) *SubcategoryName {
	return &SubcategoryName{
		valueobjects.NewStringValueObject(value).MinLength(3).MaxLength(50),
	}
}
