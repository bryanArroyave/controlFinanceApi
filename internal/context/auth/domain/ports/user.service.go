package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/entities"
)

type IUserService interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}
