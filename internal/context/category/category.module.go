package category

import (
	addcategory "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/addCategory"
	addsubcategory "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/addSubcategory"
	"go.uber.org/fx"
)

var (
	UsecasesModule = fx.Module("usecases", fx.Provide(
		addcategory.NewAddCategoryUsecase,
		addsubcategory.NewAddSubcategoryUsecase,
	))
)
