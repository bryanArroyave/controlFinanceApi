package handlers

import (
	"fmt"

	infraports "github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	saveCategoryUsecase    ports.ISaveCategory
	saveSubcategoryUsecase ports.ISaveSubcategory
}

type Result struct {
	fx.Out

	Handler infraports.IHttpHandler `group:"handlers"`
}

func New(
	saveCategoryUsecase ports.ISaveCategory,
	saveSubcategoryUsecase ports.ISaveSubcategory,
) Result {
	return Result{
		Handler: &handler{
			saveCategoryUsecase:    saveCategoryUsecase,
			saveSubcategoryUsecase: saveSubcategoryUsecase,
		},
	}
}

func (h *handler) RegisterRoutes(publicGroup *echo.Group, privateGroup *echo.Group) {

	userGroup := privateGroup.Group("/category")

	userGroup.POST("", h.SaveCategory)
	userGroup.POST("/:category_id/subcategory", h.SaveSubcategory)

	userGroup.GET("", func(c echo.Context) error {

		email := c.Get("email").(string)
		userId := c.Get("userId").(int)

		return c.String(200, "Hello, "+email+"! Your user ID is "+fmt.Sprint(userId))
	})

}

var CategoryModule = fx.Module("category", fx.Provide(
	New,
))
