package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/entities"
)

type IUserService interface {
	GetUserByID(ctx context.Context, userID string) (*entities.User, error)
}
