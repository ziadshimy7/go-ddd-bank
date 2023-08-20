package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ddd-bank/infrastructure/api"
)

type UserAccountRoutes struct {
	uahandler *api.UserAccountHandler
}

func InitiateUserAccountRoutes(uahandler *api.UserAccountHandler) *UserAccountRoutes {
	return &UserAccountRoutes{uahandler: uahandler}
}

func (h *UserAccountRoutes) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/user/account", h.uahandler.GetUserAccountDetails)
}
