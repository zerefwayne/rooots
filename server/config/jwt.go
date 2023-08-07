package config

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zerefwayne/rooots/server/models"
)

var mySigningKey = []byte("AllYourBase")

func GenerateJwtToken(oauthResponse *models.ExchangeTokenResponseBody) (string, error) {
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

	log.Printf("DECONSTRUCTED TOKEN: %+v\n", token)

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		log.Printf("VALID TOKEN. CLAIMS: %+v\n", claims)
		return true, nil
	}

	return false, fmt.Errorf("invalid token string: %s", tokenStr)
}
