package user

import (
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/enums"
	valueobjects "github.com/bryanArroyave/golang-utils/valueObjects"
)

type UserProvider struct {
	*valueobjects.StringValueObject
}

func NewUserProvider(value string) *UserProvider {
	return &UserProvider{
		valueobjects.NewStringValueObject(value).Include(enums.GetAllProvidersStrings()...).Optional(),
	}
}
