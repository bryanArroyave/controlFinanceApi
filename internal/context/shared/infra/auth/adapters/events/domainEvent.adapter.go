package events

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	messagebroker "github.com/bryanArroyave/golang-utils/events/adapter/messageBroker"
)

type DomainEventAdapter struct {
	publisher message.Publisher
}

func NewDomainEventAdapter(
	publisher message.Publisher,
) *DomainEventAdapter {
	return &DomainEventAdapter{
		publisher: publisher,
	}
}

func (d *DomainEventAdapter) Publish(eventName string, payload any) error {
	msg, err := messagebroker.NewBrokerMessage(payload)
	if err != nil {
		return err
	}
	middleware.SetCorrelationID(watermill.NewUUID(), msg.GetMessage())

	if err := d.publisher.Publish(eventName, msg.GetMessage()); err != nil {
		return err
	}

	return nil

}
