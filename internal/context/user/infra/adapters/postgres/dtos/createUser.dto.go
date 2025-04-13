package dtos

import (
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/entities"
)

type CreateUserDTO struct {
	UserName  string `bson:"name"`
	UserEmail string `bson:"email"`
}

func (dto *CreateUserDTO) ToDomain() *entities.User {
	return (&entities.User{}).From(&dtos.UserDTO{
		UserName:  dto.UserName,
		UserEmail: dto.UserEmail,
	})
}
