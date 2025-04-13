package addcategory

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/addCategory/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
	categoryvalueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/valueObjects/category"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

func (usecase *AddCategoryUsecase) AddCategory(ctx context.Context, params *dtos.AddCategoryParam) (int, error) {

	category, err := entities.NewCategory(
		categoryvalueobjects.NewCategoryName(params.Name),
		categoryvalueobjects.NewCategoryColor(params.Color),
		categoryvalueobjects.NewCategoryType(params.Type),
		categoryvalueobjects.NewCategoryBudget(params.Budget),
	)

	if err != nil {
		return 0, err
	}
	category.SetID(valueobjects.NewID(params.ID))
	id, err := usecase.categoryRepository.CreateCategory(ctx, category)

	if err != nil {
		return 0, err
	}

	return id, nil
}
