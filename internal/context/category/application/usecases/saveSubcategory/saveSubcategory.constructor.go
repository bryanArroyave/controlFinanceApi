package savesubcategory

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/ports"
)

type SaveSubcategoryUsecase struct {
	categoryService    domainports.ICategoryService
	categoryRepository domainports.ICategoryRepository
}

func NewSaveSubcategoryUsecase(
	categoryService domainports.ICategoryService,
	categoryRepository domainports.ICategoryRepository,
) applicationports.ISaveSubcategory {
	return &SaveSubcategoryUsecase{
		categoryService:    categoryService,
		categoryRepository: categoryRepository,
	}
}
