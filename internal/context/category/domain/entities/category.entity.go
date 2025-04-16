package entities

import (
	"errors"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/dtos"
	category "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/valueObjects/category"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/utils"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

type Category struct {
	ID     *valueobjects.ID
	Name   *category.CategoryName
	Color  *category.CategoryColor
	Type   *category.CategoryType
	Budget *category.CategoryBudget
	mapped *dtos.CategoryDTO
}

func NewCategory(
	name *category.CategoryName,
	color *category.CategoryColor,
	categoryType *category.CategoryType,
	budget *category.CategoryBudget,
) (*Category, error) {
	if name == nil || color == nil || categoryType == nil || budget == nil {
		return nil, errors.New("invalid category")
	}

	c := &Category{
		Name:   name,
		Color:  color,
		Type:   categoryType,
		Budget: budget,
		mapped: &dtos.CategoryDTO{},
	}

	err := c.validate(true)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Category) SetID(id *valueobjects.ID) {
	if id == nil {
		return
	}
	c.mapped.ID = uint(id.InecureValue())
	c.ID = id
}

func (c *Category) ToPrimitives() *dtos.CategoryDTO {
	return c.mapped
}

func (c *Category) validate(validateError bool) error {
	errorList := make([]error, 0)

	name, err := c.Name.Value()
	utils.VerifyError(err, &errorList, validateError)

	color, err := c.Color.Value()
	utils.VerifyError(err, &errorList, validateError)

	categoryType, err := c.Type.Value()
	utils.VerifyError(err, &errorList, validateError)

	budget, err := c.Budget.Value()
	utils.VerifyError(err, &errorList, validateError)

	c.mapped.Name = name
	c.mapped.Color = color
	c.mapped.Type = categoryType
	c.mapped.Budget = budget

	if len(errorList) > 0 {
		return errors.Join(errorList...)
	}

	return nil
}
