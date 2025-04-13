package ports

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/entities"
)

type IGetUserPort interface {
	GetUser(ctx context.Context, userID string) (*entities.User, error)
}
