package dtos

type SaveSubcategoryParam struct {
	ID         int
	CategoryID int
	Name       string
	Color      string
	Type       string
	Budget     int
}
