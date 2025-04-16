package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/entities"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, category *entities.User) (int, error)
}
