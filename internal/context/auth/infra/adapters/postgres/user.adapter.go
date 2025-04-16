package adapters

import (
	"github.com/bryanArroyave/golang-utils/gorm/ports"
)

type UserAdapter struct {
	dbManager ports.IDBManager
}

func NewUserAdapter(dbManager ports.IDBManager) *UserAdapter {
	return &UserAdapter{
		dbManager: dbManager,
	}
}
