package adapters

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/entities"
)

func (s *UserAdapter) CreateUser(ctx context.Context, user *entities.User) (int, error) {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return 0, err
	}

	userPrimitive := user.ToPrimitives()

	userModel := &models.User{
		ID:       userPrimitive.ID,
		Name:     userPrimitive.Name,
		Email:    userPrimitive.Email,
		Provider: userPrimitive.Provider,
		Img:      userPrimitive.Img,
	}

	result := conn.Save(userModel)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(userModel.ID), nil
}
