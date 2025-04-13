package adapters

import (
	"context"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/entities"
	domainports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/ports"
	valueobjects "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/user/domain/valueObjects"
	"github.com/bryanArroyave/golang-utils/gorm/ports"
)

type UserService struct {
	dbManager ports.IDBManager
}

func NewUserService(dbManager ports.IDBManager) domainports.IUserService {
	return &UserService{
		dbManager: dbManager,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*entities.User, error) {
	conn, err := s.dbManager.GetConnection()
	if err != nil {
		return nil, err
	}

	var userModel models.User
	result := conn.First(&userModel, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	userName := valueobjects.NewUserName(userModel.Name)
	userEmail := valueobjects.NewUsereEmail(userModel.Email)

	return entities.NewUser(userName, userEmail)
}
