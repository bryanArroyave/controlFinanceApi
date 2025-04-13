package addsubcategory

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/ports"
)

type AddSubcategoryUsecase struct {
	categoryService    domainports.ICategoryService
	categoryRepository domainports.ICategoryRepository
}

func NewAddSubcategoryUsecase(
	categoryService domainports.ICategoryService,
	categoryRepository domainports.ICategoryRepository,
) applicationports.ISaveSubcategory {
	return &AddSubcategoryUsecase{
		categoryService:    categoryService,
		categoryRepository: categoryRepository,
	}
}
