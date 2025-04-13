package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
)

type ICategoryService interface {
	GetCategoryByID(ctx context.Context, categoryID int) (*entities.Category, error)
	// GetCategoriesByUser(ctx context.Context, userID string) ([]*entities.Category, error)
}
