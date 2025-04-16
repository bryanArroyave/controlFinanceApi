package adapters

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/entities"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/valueObjects/user"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

func (s *UserAdapter) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return nil, err
	}

	var userModel *models.User
	result := conn.Where("email = ?", email).Find(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if userModel == nil || userModel.ID == 0 {
		return nil, nil
	}

	c, err := entities.NewUser(
		user.NewUserName(userModel.Name),
		user.NewUserEmail(userModel.Email),
		user.NewUserImg(userModel.Img),
		user.NewUserPassword(userModel.Password),
		user.NewUserProvider(userModel.Provider),
	)

	if err != nil {
		return nil, err
	}
	c.SetID(valueobjects.NewID(int(userModel.ID)))

	return c, nil
}
