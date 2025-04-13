package request

import "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/addCategory/dtos"

type AddCategoryRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Type   string `json:"type"`
	Budget int    `json:"budget"`
}

func (r *AddCategoryRequest) MapToUsecaseParam() *dtos.AddCategoryParam {
	return &dtos.AddCategoryParam{
		ID:     r.ID,
		Name:   r.Name,
		Color:  r.Color,
		Type:   r.Type,
		Budget: r.Budget,
	}
}
