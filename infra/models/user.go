package models

import "time"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	IDNumber string `gorm:"column:identification_number"`
	Email    string
	Phone    string

	CreatedAt time.Time
	UpdatedAt time.Time

	Transactions []*Transaction `gorm:"foreignKey:UserID"`
	Accounts     []*Account     `gorm:"foreignKey:UserID"`
	Categories   []*Category    `gorm:"foreignKey:UserID"`
}

type Account struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint
	User           User
	Type           string // cash, debit, credit, savings, investment, loan
	Name           string
	InitialBalance float64
	Description    string
	BillingDate    *time.Time // For credit cards
	PaymentDueDate *time.Time
	Fee            *float64

	CreatedAt time.Time
	UpdatedAt time.Time

	TransactionsFrom []*Transaction `gorm:"foreignKey:AccountFromID"` // Movimientos que salen
	TransactionsTo   []*Transaction `gorm:"foreignKey:AccountToID"`   // Movimientos que entran (transferencia)
}

type Category struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Color  string
	Type   string // income, expense
	Budget float64
	UserID *uint `gorm:"default:null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	User          *User
	Subcategories []*Subcategory `gorm:"foreignKey:CategoryID"`
	Transactions  []*Transaction `gorm:"foreignKey:CategoryID"`
}

type Subcategory struct {
	ID         uint `gorm:"primaryKey"`
	CategoryID uint
	Name       string
	Color      string
	Budget     float64

	CreatedAt time.Time
	UpdatedAt time.Time

	Category     *Category
	Transactions []*Transaction `gorm:"foreignKey:SubcategoryID"`
}

type Transaction struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uint
	AccountFromID   *uint
	AccountToID     *uint
	CategoryID      *uint
	SubcategoryID   *uint
	Amount          float64
	Note            string
	Description     string
	Type            string  // income, expense, transfer
	FrequencyType   *string // installment, recurring, null
	Installments    *int    // if installment
	InstallmentFee  *float64
	RecurringPeriod *string // daily, weekly, monthly, etc.
	Date            time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

	User        *User
	AccountFrom *Account `gorm:"foreignKey:AccountFromID"`
	AccountTo   *Account `gorm:"foreignKey:AccountToID"`
	Category    *Category
	Subcategory *Subcategory
}
