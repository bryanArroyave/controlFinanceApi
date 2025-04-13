package handlers

import (
	"github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	usecasesports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/application/ports"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	addUserUsecase usecasesports.IAddUserPort
	getUserUsecase usecasesports.IGetUserPort
}

type Result struct {
	fx.Out

	Handler ports.IHttpHandler `group:"handlers"`
}

func New(
	addUserUsecase usecasesports.IAddUserPort,
	getUSerUsecase usecasesports.IGetUserPort,
) Result {
	return Result{
		Handler: &handler{
			addUserUsecase: addUserUsecase,
			getUserUsecase: getUSerUsecase,
		},
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {

	userGroup := e.Group("/users")

	userGroup.POST("", h.RegisterPayment)
	userGroup.GET("", h.RegisterPayment)
	userGroup.GET("/:id", h.GetUser)

}

var UserModule = fx.Module("user", fx.Provide(
	New,
))
