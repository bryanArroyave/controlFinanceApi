package ports

type IUserEventsRepository interface {
	Publish(eventName string, payload any) error
}
