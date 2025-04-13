package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetUser(c echo.Context) error {

	user, err := h.getUserUsecase.GetUser(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	resp := Response{
		Data: user.ToPrimitives(),
	}
	return c.JSON(http.StatusOK, resp)
}
