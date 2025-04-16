package subscribers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	infraports "github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	routerbroker "github.com/bryanArroyave/golang-utils/events/adapter/routerBroker"
	"go.uber.org/fx"
)

type handler struct {
	saveCategoryUsecase ports.ISaveCategory
}

type Result struct {
	fx.Out

	Handler infraports.ISubscriberHandler `group:"subscriber_handlers"`
}

func New(
	saveCategoryUsecase ports.ISaveCategory,
) Result {
	return Result{
		Handler: &handler{
			saveCategoryUsecase: saveCategoryUsecase,
		},
	}
}

func (h *handler) RegisterRoutes(router *routerbroker.Router, subscriber message.Subscriber) {
	router.AddHandler("on_user_registered", "user.registered", subscriber, h.OnUserRegister)
}

var CategorySubscriberModule = fx.Module("categorySubscriber", fx.Provide(
	New,
))
