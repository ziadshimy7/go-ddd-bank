package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/go-ddd-bank/domain/service"
	errors "github.com/go-ddd-bank/utils"
	jwt_util "github.com/go-ddd-bank/utils/jwt"
	"github.com/golang-jwt/jwt/v5"
)

type AccountHandler struct {
	as *services.AccountService
}

func NewAccountHandler(as *services.AccountService) *AccountHandler {
	return &AccountHandler{as: as}
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	cookie, err := c.Cookie("jwt")

	if err != nil {
		getErr := errors.NewInternalServerError("Couldn't retrieve cookie")
		c.JSON(getErr.Status, getErr)
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(jwt_util.SecretKey), nil
	})

	if err != nil {
		restErr := errors.NewInternalServerError("error parsing cookie")

		c.JSON(restErr.Status, restErr)
		return
	}

	issuer, issuerErr := token.Claims.GetIssuer()

	if issuerErr != nil {
		restErr := errors.NewInternalServerError("error parsing issuer")
		c.JSON(restErr.Status, restErr)
		return
	}

	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

	result, getErr := h.as.GetAccountByID(issuerInt)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
