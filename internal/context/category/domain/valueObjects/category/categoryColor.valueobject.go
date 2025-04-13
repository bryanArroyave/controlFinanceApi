package category

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type CategoryColor struct {
	*valueobjects.StringValueObject
}

func NewCategoryColor(value string) *CategoryColor {
	return &CategoryColor{
		valueobjects.NewStringValueObject(value).MinLength(3).MaxLength(50),
	}
}
