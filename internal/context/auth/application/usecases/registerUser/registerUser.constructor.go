package registeruser

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
)

type RegisterUserUsecase struct {
	userRepository domainports.IUserRepository
}

func NewRegisterUserUsecase(
	userRepository domainports.IUserRepository,
) applicationports.IRegisterUser {
	return &RegisterUserUsecase{
		userRepository: userRepository,
	}
}
