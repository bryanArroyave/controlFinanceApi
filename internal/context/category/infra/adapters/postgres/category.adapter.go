package adapters

import (
	"github.com/bryanArroyave/golang-utils/gorm/ports"
)

type CategoryAdapter struct {
	dbManager ports.IDBManager
}

func NewCategoryAdapter(dbManager ports.IDBManager) *CategoryAdapter {
	return &CategoryAdapter{
		dbManager: dbManager,
	}
}
