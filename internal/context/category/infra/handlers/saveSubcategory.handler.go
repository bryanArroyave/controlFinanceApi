package handlers

import (
	"net/http"
	"strconv"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/handlers/request"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/handlers/utils"
	"github.com/labstack/echo/v4"
)

func (h *handler) SaveSubcategory(c echo.Context) error {
	req := request.SaveSubcategoryRequest{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.BuildErrorResponse("invalidRequest", err, nil))
	}

	categoryIDStr := c.Param("category_id")

	categoryID, err := strconv.Atoi(categoryIDStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("invalidCategoryId", err, nil))
	}

	id, err := h.saveSubcategoryUsecase.SaveSubcategory(c.Request().Context(), c.Get("userId").(int), categoryID, req.MapToUsecaseParam())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("unexpectedError", err, nil))
	}

	resp := Response{
		Data: map[string]any{
			"subcategory_id": id,
		},
	}
	return c.JSON(http.StatusOK, resp)
}
