package registeruser

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/registerUser/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/entities"
	uservalueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/valueObjects/user"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

func (usecase *RegisterUserUsecase) RegisterUser(ctx context.Context, params *dtos.RegisterUserParam) (int, error) {

	user, err := entities.NewUser(
		uservalueobjects.NewUserName(params.Name),
		uservalueobjects.NewUserEmail(params.Email),
		uservalueobjects.NewUserImg(params.Img),
		uservalueobjects.NewUserPassword(params.Password),
		uservalueobjects.NewUserProvider(params.Provider),
	)

	if err != nil {
		return 0, err
	}
	user.SetID(valueobjects.NewID(params.ID))
	id, err := usecase.userRepository.CreateUser(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}
