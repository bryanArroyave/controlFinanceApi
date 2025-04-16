package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/dtos"
)

type IGetUserByEmail interface {
	GetUserByEmail(ctx context.Context, email string) (*dtos.UserDTO, error)
}
