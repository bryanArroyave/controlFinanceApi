package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/dtos"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *dtos.UserDTO) (string, error)
}
