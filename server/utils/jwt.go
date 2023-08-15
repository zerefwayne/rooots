package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte(os.Getenv("JWT_SIGNING_KEY"))

type SessionJwt struct {
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
}

type SessionJwtClaims struct {
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(sessionJwtBody *SessionJwt) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"refresh_token": sessionJwtBody.RefreshToken,
		"user_id":       sessionJwtBody.UserId,
	})

	return token.SignedString(mySigningKey)
}

func ValidateJwtToken(tokenStr string) (bool, *SessionJwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &SessionJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return false, nil, err
	}

	claims, ok := token.Claims.(*SessionJwtClaims)
	if ok && token.Valid {
		return true, claims, nil
	}

	return false, nil, fmt.Errorf("invalid token string: %s", tokenStr)
}
