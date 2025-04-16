package registeruser

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
)

type RegisterUserUsecase struct {
	userRepository       domainports.IUserRepository
	userEventsRepository domainports.IUserEventsRepository
}

func NewRegisterUserUsecase(
	userRepository domainports.IUserRepository,
	userEventsRepository domainports.IUserEventsRepository,
) applicationports.IRegisterUser {
	return &RegisterUserUsecase{
		userRepository:       userRepository,
		userEventsRepository: userEventsRepository,
	}
}
