package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func MiddlewareJWT(next echo.HandlerFunc) echo.HandlerFunc {
	secretKey := os.Getenv("SECRET_JWT")
	return func(c echo.Context) error {
		tokenString := extractTokenFromHeader(c.Request())

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
		}

		return next(c)
	}
}

func extractTokenFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	tokenParts := strings.Split(bearerToken, " ")
	if len(tokenParts) == 2 && strings.ToLower(tokenParts[0]) == "bearer" {
		return tokenParts[1]
	}
	return ""
}
