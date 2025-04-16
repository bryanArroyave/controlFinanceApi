package entities

import (
	"errors"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/valueObjects/subcategory"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/utils"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

type Subcategory struct {
	SubcategoryID *valueobjects.ID
	CategoryID    *valueobjects.ID
	Name          *subcategory.SubcategoryName
	Color         *subcategory.SubcategoryColor
	Budget        *subcategory.SubcategoryBudget
	mapped        *dtos.SubcategoryDTO
}

func NewSubcategory(
	categoryID *valueobjects.ID,
	name *subcategory.SubcategoryName,
	color *subcategory.SubcategoryColor,
	budget *subcategory.SubcategoryBudget,
) (*Subcategory, error) {
	if name == nil || color == nil || budget == nil {
		return nil, errors.New("invalid subcategory")
	}

	c := &Subcategory{
		Name:       name,
		Color:      color,
		Budget:     budget,
		mapped:     &dtos.SubcategoryDTO{},
		CategoryID: categoryID,
	}

	err := c.validate(true)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Subcategory) SetID(id *valueobjects.ID) {
	c.SubcategoryID = id
}

func (c *Subcategory) ToPrimitives() *dtos.SubcategoryDTO {
	return c.mapped
}

func (c *Subcategory) validate(validateError bool) error {
	errorList := make([]error, 0)

	name, err := c.Name.Value()
	utils.VerifyError(err, &errorList, validateError)

	color, err := c.Color.Value()
	utils.VerifyError(err, &errorList, validateError)

	categoryID, err := c.CategoryID.Value()
	utils.VerifyError(err, &errorList, validateError)

	budget, err := c.Budget.Value()
	utils.VerifyError(err, &errorList, validateError)

	c.mapped.CategoryID = uint(categoryID)
	c.mapped.Name = name
	c.mapped.Color = color
	c.mapped.Budget = budget

	if len(errorList) > 0 {
		return errors.Join(errorList...)
	}

	return nil
}
