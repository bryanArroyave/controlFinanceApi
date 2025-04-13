package getuser

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/entities"
)

func (a *GetUserUsecase) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	return a.userService.GetUserByID(ctx, userID)
}
