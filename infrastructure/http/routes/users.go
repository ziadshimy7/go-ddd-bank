package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ddd-bank/infrastructure/api"
)

type UserRoutes struct {
	uhandler *api.UserHandler
}

func InitiateUserRoutes(uhandler *api.UserHandler) *UserRoutes {
	return &UserRoutes{uhandler: uhandler}
}

func (h *UserRoutes) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/auth/register", h.uhandler.RegisterUser)
	r.POST("/api/auth/login", h.uhandler.Login)
	r.GET("/api/auth/user", h.uhandler.GetUser)
	r.GET("/api/auth/logout", h.uhandler.Logout)
}
