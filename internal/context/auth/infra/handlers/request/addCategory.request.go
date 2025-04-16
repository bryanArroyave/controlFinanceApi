package request

import "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/registerUser/dtos"

type RegisterUserRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *RegisterUserRequest) MapToUsecaseParam() *dtos.RegisterUserParam {
	return &dtos.RegisterUserParam{
		ID:    r.ID,
		Name:  r.Name,
		Email: r.Email,
	}
}
