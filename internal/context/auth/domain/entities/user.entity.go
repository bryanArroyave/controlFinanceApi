package entities

import (
	"errors"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/valueObjects/user"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/utils"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/domain/valueObjects"
)

type User struct {
	ID       *valueobjects.ID
	Name     *user.UserName
	Email    *user.UserEmail
	Img      *user.UserImg
	Password *user.UserPassword
	Provider *user.UserProvider
	mapped   *dtos.UserDTO
}

func NewUser(
	name *user.UserName,
	email *user.UserEmail,
	img *user.UserImg,
	password *user.UserPassword,
	provider *user.UserProvider,
) (*User, error) {
	if name == nil || email == nil {
		return nil, errors.New("invalid user")
	}

	c := &User{
		Name:     name,
		Email:    email,
		Password: password,
		Provider: provider,
		Img:      img,
		mapped:   &dtos.UserDTO{},
	}

	err := c.validate(true)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *User) SetID(id *valueobjects.ID) {
	if id == nil {
		return
	}
	c.mapped.ID = uint(id.InecureValue())
	c.ID = id
}

func (c *User) ToPrimitives() *dtos.UserDTO {
	return c.mapped
}

func (c *User) validate(validateError bool) error {
	errorList := make([]error, 0)

	name, err := c.Name.Value()
	utils.VerifyError(err, &errorList, validateError)
	c.mapped.Name = name

	email, err := c.Email.Value()
	utils.VerifyError(err, &errorList, validateError)
	c.mapped.Email = email

	password, err := c.Password.Value()
	utils.VerifyError(err, &errorList, validateError)
	c.mapped.Password = password

	provider, err := c.Provider.Value()
	utils.VerifyError(err, &errorList, validateError)
	c.mapped.Provider = provider

	img, err := c.Img.Value()
	utils.VerifyError(err, &errorList, validateError)
	c.mapped.Img = img

	if len(errorList) > 0 {
		return errors.Join(errorList...)
	}

	return nil
}
