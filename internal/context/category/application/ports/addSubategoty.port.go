package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/addSubcategory/dtos"
)

type ISaveSubcategory interface {
	SaveSubcategory(ctx context.Context, categoryID int, params *dtos.SaveSubcategoryParam) (int, error)
}
