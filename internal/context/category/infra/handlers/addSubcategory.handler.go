package handlers

import (
	"net/http"
	"strconv"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/handlers/request"
	"github.com/labstack/echo/v4"
)

func (h *handler) AddSubcategory(c echo.Context) error {
	req := request.SaveSubcategoryRequest{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	categoryIDStr := c.Param("category_id")

	categoryID, err := strconv.Atoi(categoryIDStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"data": map[string]string{
				"error": "invalid category id",
			},
		})
	}

	id, err := h.addSubcategoryUsecase.SaveSubcategory(c.Request().Context(), categoryID, req.MapToUsecaseParam())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"data": map[string]string{
				"error": err.Error(),
			},
		})
	}

	resp := Response{
		Data: map[string]any{
			"subcategory_id": id,
		},
	}
	return c.JSON(http.StatusOK, resp)
}
