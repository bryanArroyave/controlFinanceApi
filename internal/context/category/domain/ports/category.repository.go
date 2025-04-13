package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
)

type ICategoryRepository interface {
	CreateCategory(ctx context.Context, category *entities.Category) (int, error)
	CreateSubcategory(ctx context.Context, subcategory *entities.Subcategory) (int, error)
	// UpdateCategory(ctx context.Context, category *entities.Category) error
	// DeleteCategory(ctx context.Context, categoryID string) error
	// GetCategoryByID(ctx context.Context, categoryID string) (*entities.Category, error)
	// GetCategoriesByUser(ctx context.Context, userID string) ([]*entities.Category, error)
}
