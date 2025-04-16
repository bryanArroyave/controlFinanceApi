package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveSubcategory/dtos"
)

type ISaveSubcategory interface {
	SaveSubcategory(ctx context.Context, userID, categoryID int, params *dtos.SaveSubcategoryParam) (int, error)
}
