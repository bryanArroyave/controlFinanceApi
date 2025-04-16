package getuserbyemail

import (
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/ports"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
)

type GetUserByEmailUsecase struct {
	userService domainports.IUserService
}

func NewGetUserByEmailUsecase(
	userService domainports.IUserService,
) applicationports.IGetUserByEmail {
	return &GetUserByEmailUsecase{
		userService: userService,
	}
}
