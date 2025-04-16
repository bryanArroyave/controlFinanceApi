package getuserbyemail

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/dtos"
)

func (usecase *GetUserByEmailUsecase) GetUserByEmail(ctx context.Context, email string) (*dtos.UserDTO, error) {

	user, err := usecase.userService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {

		return nil, nil
	}

	return user.ToPrimitives(), nil
}
