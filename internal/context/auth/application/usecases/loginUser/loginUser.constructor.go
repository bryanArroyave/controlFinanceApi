package loginuser

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
)

type LoginUserUsecase struct {
	userRepository domainports.IUserRepository
}

func NewLoginUserUsecase(
	userRepository domainports.IUserRepository,
) applicationports.ILoginUser {
	return &LoginUserUsecase{
		userRepository: userRepository,
	}
}
