package savecategory

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/ports"
)

type SaveCategoryUsecase struct {
	categoryRepository domainports.ICategoryRepository
}

func NewSaveCategoryUsecase(
	categoryRepository domainports.ICategoryRepository,
) applicationports.ISaveCategory {
	return &SaveCategoryUsecase{
		categoryRepository: categoryRepository,
	}
}
