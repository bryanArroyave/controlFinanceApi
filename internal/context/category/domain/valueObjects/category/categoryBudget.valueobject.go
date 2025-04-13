package category

import (
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

// TODO: Crear float value object
type CategoryBudget struct {
	*valueobjects.IntValueObject
}

func NewCategoryBudget(value int) *CategoryBudget {
	return &CategoryBudget{
		valueobjects.NewIntValueObject(value).Min(0).Optional(),
	}
}
