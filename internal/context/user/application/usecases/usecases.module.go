package usecases

import (
	getuser "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/application/usecases/addUser"
	adduser "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/application/usecases/getUser"
	"go.uber.org/fx"
)

var UsecasesModule = fx.Module("usecases", fx.Provide(
	adduser.New,
	getuser.New,
))
