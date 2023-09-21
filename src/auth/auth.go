package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"
)

var (
	signingMethod = jwt.SigningMethodHS256
	secretKey     = helper.LoadEnv("SECRET_KEY", "test")
)

func SecretKey() string {
	return secretKey
}

func SigningMethod() *jwt.SigningMethodHMAC {
	return signingMethod
}

func GenerateAccessToken(data *model.Auth) (string, error) {
	expTime := time.Now().Add(3 * 24 * time.Hour)
	return generateToken(data, expTime)
}

func generateToken(data *model.Auth, exp time.Time) (string, error) {
	data.StandardClaims = jwt.StandardClaims{
		ExpiresAt: exp.Unix(),
	}

	token := jwt.NewWithClaims(SigningMethod(), data)

	tokenString, err := token.SignedString([]byte(SecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
