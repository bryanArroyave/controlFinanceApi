package adapters

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
)

func (s *CategoryAdapter) SaveCategory(ctx context.Context, category *entities.Category) (int, error) {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return 0, err
	}

	categoryPrimitive := category.ToPrimitives()

	categoryModel := &models.Category{
		ID:     categoryPrimitive.ID,
		Name:   categoryPrimitive.Name,
		Color:  categoryPrimitive.Color,
		Type:   categoryPrimitive.Type,
		Budget: float64(categoryPrimitive.Budget),
	}

	result := conn.Save(categoryModel)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(categoryModel.ID), nil
}

func (s *CategoryAdapter) SaveSubcategory(ctx context.Context, category *entities.Subcategory) (int, error) {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return 0, err
	}

	subCategoryPrimitive := category.ToPrimitives()

	subCategoryModel := &models.Subcategory{
		ID:         subCategoryPrimitive.ID,
		Name:       subCategoryPrimitive.Name,
		Color:      subCategoryPrimitive.Color,
		Budget:     float64(subCategoryPrimitive.Budget),
		CategoryID: subCategoryPrimitive.CategoryID,
	}

	result := conn.Save(subCategoryModel)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(subCategoryModel.ID), nil
}

func (s *CategoryAdapter) DeleteCategory(ctx context.Context, categoryID string) error {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return err
	}

	result := conn.Delete(&models.Category{}, categoryID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
