package subcategory

import (
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

// TODO: Crear float value object
type SubcategoryBudget struct {
	*valueobjects.IntValueObject
}

func NewSubcategoryBudget(value int) *SubcategoryBudget {
	return &SubcategoryBudget{
		valueobjects.NewIntValueObject(value).Min(0),
	}
}
