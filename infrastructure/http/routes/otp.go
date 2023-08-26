package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ddd-bank/infrastructure/api"
)

type OTPRoutes struct {
	otphandler *api.OTPHandler
}

func InitiateOTPRoutes(otphandler *api.OTPHandler) *OTPRoutes {
	return &OTPRoutes{otphandler: otphandler}
}

func (h *OTPRoutes) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/auth/otp", h.otphandler.GetOTP)
	r.POST("/api/auth/otp/verify", h.otphandler.VerifyOTP)
}
