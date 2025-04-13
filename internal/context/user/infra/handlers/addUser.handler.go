package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	AccountID string `json:"account_id"`
	Total     int64  `json:"total"`
}

type Response struct {
	Data any `json:"data"`
}

func (h *handler) RegisterPayment(c echo.Context) error {
	req := Request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	err = h.addUserUsecase.AddUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	resp := Response{
		Data: "success",
	}
	return c.JSON(http.StatusOK, resp)
}
