package entities

import (
	"errors"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/dtos"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/valueObjects"
)

type User struct {
	UserName  *valueobjects.UserName
	UserEmail *valueobjects.UsereEmail
	mapped    *dtos.UserDTO
}

func NewUser(
	userName *valueobjects.UserName,
	userEmail *valueobjects.UsereEmail,
) (*User, error) {

	if userName == nil || userEmail == nil {
		return nil, errors.New("invalid User")
	}

	j := &User{
		UserName:  userName,
		UserEmail: userEmail,
		mapped:    &dtos.UserDTO{},
	}

	err := j.validate(true)
	if err != nil {
		return nil, err
	}
	return j, nil

}

func (j *User) ToPrimitives() *dtos.UserDTO {
	return j.mapped
}

func (j *User) validate(validateError bool) error {

	errorList := make([]error, 0)

	userName, err := j.UserName.Value()
	verifyError(err, &errorList, validateError)

	userEmail, err := j.UserEmail.Value()
	verifyError(err, &errorList, validateError)

	j.mapped.UserName = userName
	j.mapped.UserEmail = userEmail

	if len(errorList) > 0 {
		return errors.Join(errorList...)
	}

	return nil
}

func verifyError(err error, errorList *[]error, validateError bool) {
	if err != nil && validateError {
		*errorList = append(*errorList, err)
	}
}

func (j *User) From(dto *dtos.UserDTO) *User {

	nu := &User{
		UserName:  valueobjects.NewUserName(dto.UserName),
		UserEmail: valueobjects.NewUsereEmail(dto.UserEmail),
		mapped:    &dtos.UserDTO{},
	}

	_ = nu.validate(false)

	return nu

}
