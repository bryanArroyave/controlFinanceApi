package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtener el token de la cabecera Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Missing authorization header",
			})
		}

		// Verificar que el token sea válido
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Verificar que el método de firma sea correcto
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			// Retornar la clave secreta para la validación
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		// Si hubo un error al analizar el token
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Invalid or expired token",
			})
		}

		// Verificar si los claims son válidos
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			// Puedes guardar la información del usuario en el contexto de Echo para usarla después
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)
			return next(c)
		}

		// Si el token no es válido
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid token",
		})
	}
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(username, email string) (string, error) {
	// Crear los claims del JWT
	claims := Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // El token expira en 24 horas
			Issuer:    "your-app",                            // Emisor del token
		},
	}

	// Crear el token con la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY"))) // Usa una clave secreta
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
