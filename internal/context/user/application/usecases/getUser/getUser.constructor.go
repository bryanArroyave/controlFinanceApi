package getuser

import (
	usecaseports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/application/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/ports"
)

type GetUserUsecase struct {
	userService ports.IUserService
}

func New(
	userService ports.IUserService,
) usecaseports.IGetUserPort {
	return &GetUserUsecase{
		userService: userService,
	}
}
