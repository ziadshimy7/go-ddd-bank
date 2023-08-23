package jwt_util

import (
	"strconv"

	domain "github.com/go-ddd-bank/domain/model"
	errors "github.com/go-ddd-bank/utils"
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

func CreateJWT(result *domain.User, expireTime *jwt.NumericDate) (string, *errors.Errors) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{Issuer: strconv.Itoa(int(result.ID)), ExpiresAt: expireTime})
	token, tokenErr := claims.SignedString([]byte(SecretKey))
	if tokenErr != nil {
		err := errors.NewInternalServerError("login failed")
		return "", err
	}

	return token, nil
}
