package dtos

type CategoryDTO struct {
	ID     uint
	Name   string
	Icon   string
	Color  string
	Type   string
	Budget int
}

type SubcategoryDTO struct {
	ID         uint
	Name       string
	Icon       string
	Color      string
	Budget     int
	CategoryID uint
}

type TransactionDTO struct {
	Date            string
	CategoryID      string
	SubcategoryID   *string
	AccountFromID   string
	AccountToID     *string
	Note            string
	Description     string
	Type            string
	Amount          float64
	PeriodicityType string
	InstallmentInfo *InstallmentInfoDTO
	RecurringInfo   *RecurringInfoDTO
}

type InstallmentInfoDTO struct {
	TotalInstallments int
	Commission        float64
}

type RecurringInfoDTO struct {
	Periodicity string
}
