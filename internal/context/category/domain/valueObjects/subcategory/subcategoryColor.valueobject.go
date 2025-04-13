package subcategory

import valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"

type SubcategoryColor struct {
	*valueobjects.StringValueObject
}

func NewSubcategoryColor(value string) *SubcategoryColor {
	return &SubcategoryColor{
		valueobjects.NewStringValueObject(value).MinLength(3).MaxLength(50),
	}
}
