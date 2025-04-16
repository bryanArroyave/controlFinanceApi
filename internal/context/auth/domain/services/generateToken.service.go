package services

import (
	"fmt"
	"os"
	"time"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/auth/dtos"
	"github.com/dgrijalva/jwt-go"
)

type privateClaims struct {
	*dtos.Claims
	jwt.StandardClaims
}

type TokenService struct {
	jwtSecret   string
	serviceName string
}

func NewTokenService() *TokenService {
	return &TokenService{
		jwtSecret:   os.Getenv("JWT_SECRET_KEY"),
		serviceName: os.Getenv("SERVICE_NAME"),
	}
}

func (s *TokenService) GenerateToken(claims *dtos.Claims) (string, time.Time, error) {

	expiredAt := time.Now().AddDate(0, 3, 0)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &privateClaims{
		Claims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			Issuer:    s.serviceName,
		},
	})
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", expiredAt, err
	}

	return tokenString, expiredAt, nil
}

func (s *TokenService) ParseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, &privateClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *TokenService) ValidateToken(tokenString string) (*dtos.Claims, error) {
	token, err := s.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*privateClaims); ok && token.Valid {
		return claims.Claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func (s *TokenService) ExtractClaims(tokenString string) (*dtos.Claims, error) {
	token, err := s.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*privateClaims); ok && token.Valid {
		return claims.Claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
