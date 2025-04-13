package addcategory

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/ports"
)

type AddCategoryUsecase struct {
	categoryRepository domainports.ICategoryRepository
}

func NewAddCategoryUsecase(
	categoryRepository domainports.ICategoryRepository,
) applicationports.IAddCategory {
	return &AddCategoryUsecase{
		categoryRepository: categoryRepository,
	}
}
