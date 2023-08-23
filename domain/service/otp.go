package services

import (
	"github.com/go-ddd-bank/domain/model"
	repo "github.com/go-ddd-bank/domain/repository"
	errors "github.com/go-ddd-bank/utils"
)

type OTPService struct {
	r repo.OTPRepository
}

func NewOTPService(r repo.OTPRepository) *OTPService {
	return &OTPService{r: r}
}

func (os *OTPService) UpdateOTP(otp *domain.OTP) *errors.Errors {
	if err := os.r.UpdateOTP(otp); err != nil {
		return err
	}

	return nil
}
