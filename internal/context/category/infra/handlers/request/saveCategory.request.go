package request

import "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/usecases/saveCategory/dtos"

type SaveCategoryRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Type   string `json:"type"`
	Budget int    `json:"budget"`
}

func (r *SaveCategoryRequest) MapToUsecaseParam() *dtos.SaveCategoryParam {
	return &dtos.SaveCategoryParam{
		ID:     r.ID,
		Name:   r.Name,
		Color:  r.Color,
		Type:   r.Type,
		Budget: r.Budget,
	}
}
