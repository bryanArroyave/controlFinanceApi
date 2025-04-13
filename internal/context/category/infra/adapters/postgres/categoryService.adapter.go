package adapters

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/valueObjects/category"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

func (s *CategoryAdapter) GetCategoryByID(ctx context.Context, categoryID int) (*entities.Category, error) {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return nil, err
	}

	var categoryModel *models.Category
	result := conn.First(&categoryModel, categoryID)
	if result.Error != nil {
		return nil, result.Error
	}

	c, err := entities.NewCategory(
		category.NewCategoryName(categoryModel.Name),
		category.NewCategoryColor(categoryModel.Color),
		category.NewCategoryType(categoryModel.Type),
		category.NewCategoryBudget(int(categoryModel.Budget)),
	)

	c.SetID(valueobjects.NewID(int(categoryModel.ID)))

	if err != nil {
		return nil, err
	}

	return c, nil
}
