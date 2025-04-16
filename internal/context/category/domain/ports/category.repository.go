package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
)

type ICategoryRepository interface {
	SaveCategory(ctx context.Context, userID int, category *entities.Category) (int, error)
	SaveSubcategory(ctx context.Context, subcategory *entities.Subcategory) (int, error)
}
