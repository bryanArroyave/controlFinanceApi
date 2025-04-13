package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/addCategory/dtos"
)

type IAddCategory interface {
	AddCategory(ctx context.Context, params *dtos.AddCategoryParam) (int, error)
}
