package handlers

import (
	"net/http"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/handlers/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (h *handler) Login(c echo.Context) error {
	switch c.QueryParam("provider") {
	case "github":
		url := h.githuboauth2Config.AuthCodeURL(h.oauth2StateString, oauth2.AccessTypeOffline)
		return c.Redirect(http.StatusFound, url)
	case "google":
		url := h.googleOauth2Config.AuthCodeURL(h.oauth2StateString, oauth2.AccessTypeOffline)
		return c.Redirect(http.StatusFound, url)
	}
	// TODO: implementación con correo y contraseña

	return c.JSON(http.StatusBadRequest, utils.BuildSuccessResponse("invalid provider", nil))
}
