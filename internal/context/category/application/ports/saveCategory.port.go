package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveCategory/dtos"
)

type ISaveCategory interface {
	SaveCategory(ctx context.Context, userID int, params *dtos.SaveCategoryParam) (int, error)
}
