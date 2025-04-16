package savesubcategory

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveSubcategory/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/entities"
	subcategoryvalueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/valueObjects/subcategory"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

func (usecase *SaveSubcategoryUsecase) SaveSubcategory(ctx context.Context, userID, categoryID int, params *dtos.SaveSubcategoryParam) (int, error) {

	category, err := usecase.categoryService.GetCategoryByID(ctx, categoryID)

	if err != nil {
		return 0, err
	}

	if category == nil {
		return 0, nil
	}

	subcategory, err := entities.NewSubcategory(
		valueobjects.NewID(int(categoryID)),
		subcategoryvalueobjects.NewSubcategoryName(params.Name),
		subcategoryvalueobjects.NewSubcategoryColor(params.Color),
		subcategoryvalueobjects.NewSubcategoryBudget(params.Budget),
	)

	if err != nil {
		return 0, err
	}
	subcategory.SetID(valueobjects.NewID(params.ID))
	id, err := usecase.categoryRepository.SaveSubcategory(ctx, subcategory)

	if err != nil {
		return 0, err
	}

	return id, nil
}
