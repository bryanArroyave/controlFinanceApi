package request

import "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveSubcategory/dtos"

type SaveSubcategoryRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Budget int    `json:"budget"`
}

func (r *SaveSubcategoryRequest) MapToUsecaseParam() *dtos.SaveSubcategoryParam {
	return &dtos.SaveSubcategoryParam{
		ID:     r.ID,
		Name:   r.Name,
		Color:  r.Color,
		Budget: r.Budget,
	}
}
