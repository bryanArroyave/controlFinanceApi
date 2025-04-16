package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/services"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/handlers/utils"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("missingToken", fmt.Errorf("authorization header is missing"), nil))
		}

		tokenService := services.NewTokenService()

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := tokenService.ParseToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("invalidToken", err, nil))
		}

		if claims, err := tokenService.ExtractClaims(tokenString); err == nil && token.Valid {
			c.Set("email", claims.Email)
			c.Set("userId", claims.UserID)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("invalidToken", err, nil))
	}
}

func SessionCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Check if the session cookie is present
		sessionCookie, err := c.Cookie("control_finance_session_token")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("unauthorized", fmt.Errorf("don't have an active session"), nil))
		}

		tokenService := services.NewTokenService()

		tokenString := sessionCookie.Value
		token, err := tokenService.ParseToken(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("invalidToken", err, nil))
		}

		if claims, err := tokenService.ExtractClaims(tokenString); err == nil && token.Valid {
			c.Set("email", claims.Email)
			c.Set("userId", claims.UserID)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("invalidToken", err, nil))
	}
}
