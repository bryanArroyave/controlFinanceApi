package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/registerUser/dtos"
)

type IRegisterUser interface {
	RegisterUser(ctx context.Context, params *dtos.RegisterUserParam) (int, error)
}
