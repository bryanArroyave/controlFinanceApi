package migrations

import (
	"github.com/bryanArroyave/eventsplit/back/user-service/infra/models"
	"github.com/bryanArroyave/golang-utils/gorm/ports"
)

func StartMigrations(saasDBManager ports.IDBManager) error {
	conn, err := saasDBManager.GetConnection()
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Category{},
		&models.Subcategory{},
		&models.Transaction{},
	)
	if err != nil {
		return err
	}

	return nil
}
