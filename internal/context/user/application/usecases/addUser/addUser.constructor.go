package adduser

import (
	usecaseports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/application/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/ports"
)

type AddUserUsecase struct {
	userService ports.IUserService
}

func New(
	userService ports.IUserService,
) usecaseports.IAddUserPort {
	return &AddUserUsecase{
		userService: userService,
	}
}
