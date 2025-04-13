package seeds

import (
	utilsports "github.com/bryanArroyave/golang-utils/gorm/ports"
)

type ControlFinancialSeed struct {
	dbManager utilsports.IDBManager
}

func NewControlFinancialSeed(dbManager utilsports.IDBManager) *ControlFinancialSeed {
	return &ControlFinancialSeed{dbManager: dbManager}

}
func (c *ControlFinancialSeed) Exec() {
	conn, err := c.dbManager.GetConnection()
	if err != nil {
		panic(err)
	}

	seedCategories(conn)
}
