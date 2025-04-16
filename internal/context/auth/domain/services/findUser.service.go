package services

import (
	"context"
	"fmt"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
)

type UserFinder struct {
	userService ports.IUserService
}

func NewUserFinder(userService ports.IUserService) *UserFinder {
	return &UserFinder{
		userService: userService,
	}
}

func (finder *UserFinder) FindUserByEmail(ctx context.Context, email string) (*dtos.UserDTO, error) {
	user, err := finder.userService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		// TODO: Crear error personalizado
		return nil, fmt.Errorf("user not found")
	}

	return user.ToPrimitives(), nil
}
