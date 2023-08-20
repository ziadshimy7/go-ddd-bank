package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ddd-bank/infrastructure/api"
)

type AccountRoutes struct {
	ahandler *api.AccountHandler
}

func InitiateAccountRoutes(ahandler *api.AccountHandler) *AccountRoutes {
	return &AccountRoutes{ahandler: ahandler}
}

func (h *AccountRoutes) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/account", h.ahandler.GetAccount)
}
