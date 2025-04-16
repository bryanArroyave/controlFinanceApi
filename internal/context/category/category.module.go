package category

import (
	savecategory "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveCategory"
	savesubcategory "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveSubcategory"
	"go.uber.org/fx"
)

var (
	UsecasesModule = fx.Module("usecases", fx.Provide(
		savecategory.NewSaveCategoryUsecase,
		savesubcategory.NewSaveSubcategoryUsecase,
	))
)
