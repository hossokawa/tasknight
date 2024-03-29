package auth

import (
	"fmt"
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
}

func GenerateJWT(id string, username string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      id,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Failed to create token")
	}

	return tokenStr, nil
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func GetTokenFromCookie(c echo.Context) string {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return err.Error()
	}
	tokenStr := cookie.Value
	return tokenStr
}

func GetUserId(c echo.Context) string {
	token := GetTokenFromCookie(c)
	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(token, claims, nil)
	return claims["sub"].(string)
}
