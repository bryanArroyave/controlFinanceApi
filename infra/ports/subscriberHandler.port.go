package ports

import (
	"github.com/ThreeDotsLabs/watermill/message"
	routerbroker "github.com/bryanArroyave/golang-utils/events/adapter/routerBroker"
)

type ISubscriberHandler interface {
	RegisterRoutes(router *routerbroker.Router, subscriber message.Subscriber)
}
