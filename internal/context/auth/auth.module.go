package auth

import (
	getuserbyemail "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/getUserByEmail"
	loginuser "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/loginUser"
	registeruser "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/registerUser"
	"go.uber.org/fx"
)

var (
	UsecasesModule = fx.Module("usecases", fx.Provide(
		registeruser.NewRegisterUserUsecase,
		loginuser.NewLoginUserUsecase,
		getuserbyemail.NewGetUserByEmailUsecase,
	))
)
