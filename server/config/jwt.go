package config

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	strava "github.com/zerefwayne/rooots/server/models/strava"
)

var mySigningKey = []byte("AllYourBase")

func GenerateJwtToken(oauthResponse *strava.ExchangeTokenResponseBody) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_token":  oauthResponse.AccessToken,
		"refresh_token": oauthResponse.RefreshToken,
		"exp":           oauthResponse.ExpiresAt,
	})

	return token.SignedString(mySigningKey)
}

type MyCustomClaims struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	jwt.RegisteredClaims
}

func ValidateJwtToken(tokenStr string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return true, nil
	}

	return false, fmt.Errorf("invalid token string: %s", tokenStr)
}
