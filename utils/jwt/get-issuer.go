package jwt_util

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	SecretKey = "qwe123"
)

func GetIssuer(cookie string) (issuer string, rest_error error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {

		return "", err
	}

	issuer, issuerErr := token.Claims.GetIssuer()

	if issuerErr != nil {
		return "", issuerErr
	}

	return issuer, nil
}
