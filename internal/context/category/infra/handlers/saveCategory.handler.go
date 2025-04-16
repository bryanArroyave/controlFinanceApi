package handlers

import (
	"net/http"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/handlers/request"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Data any `json:"data"`
}

func (h *handler) SaveCategory(c echo.Context) error {
	req := request.SaveCategoryRequest{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	id, err := h.saveCategoryUsecase.SaveCategory(c.Request().Context(), c.Get("userId").(int), req.MapToUsecaseParam())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"data": map[string]string{
				"error": err.Error(),
			},
		})
	}

	resp := Response{
		Data: map[string]any{
			"category_id": id,
		},
	}
	return c.JSON(http.StatusOK, resp)
}
