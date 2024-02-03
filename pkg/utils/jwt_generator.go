package utils

import (
	"errors"
	"os"
	"time"

	"github.com/auth_app/app/models"
	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken(data *models.User) (*string, error) {
	if data == nil {
		return nil, errors.New("user data is nil")
	}

	exp := GetTimeNow().Add(time.Minute * time.Duration(15)).Unix()
	claims := jwt.MapClaims{
		"user_id": data.ID,
		"exp":     exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
