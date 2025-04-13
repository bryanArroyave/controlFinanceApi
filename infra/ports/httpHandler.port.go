package ports

import "github.com/labstack/echo/v4"

type IHttpHandler interface {
	RegisterRoutes(e *echo.Echo)
}
