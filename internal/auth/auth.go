package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	Id       string
	Username string
	Email    string
	jwt.RegisteredClaims
}

func GenerateJWT(id string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to create token")
	}

	return tokenStr, nil
}

func GetUserId(tokenStr string) string {
	return ""
}
