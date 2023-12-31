package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	domain "github.com/go-ddd-bank/domain/model"
	services "github.com/go-ddd-bank/domain/service"
	errors "github.com/go-ddd-bank/utils"
	jwt_util "github.com/go-ddd-bank/utils/jwt"
	sendemail "github.com/go-ddd-bank/utils/send-email"
	"github.com/pquerna/otp/totp"
)

type OTPHandler struct {
	os *services.OTPService
	us *services.UserService
}

func NewOTPHandler(os *services.OTPService, us *services.UserService) *OTPHandler {
	return &OTPHandler{os: os, us: us}
}

func (otphandler *OTPHandler) GetOTP(c *gin.Context) {
	otp := &domain.OTP{}
	user := new(domain.User)

	cookie, _ := c.Cookie("jwt")

	issuer, issuerErr := jwt_util.GetIssuer(cookie)
	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(err.Status, err)
		return
	}

	if issuerErr != nil {
		restErr := errors.NewBadRequestError("Couldn't get user information")

		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "reenbank.com",
		AccountName: user.Email,
		SecretSize:  15,
		Period:      60,
	})

	if err != nil {
		otpErr := errors.NewBadRequestError("Couldn't get user information")

		c.JSON(http.StatusBadRequest, otpErr)
		return
	}

	passcode, err := totp.GenerateCode(key.Secret(), time.Now())

	if err != nil {
		err := errors.NewBadRequestError("Couldn't generate Secret Key for one time password, Please try again later !")
		c.JSON(err.Status, err)
		return
	}

	message := fmt.Sprintf("Dear user, your one-time passcode for authentication is: %s. Please use this code to complete your sign in process.", passcode)
	mailErr := sendemail.SendEmail(message, []string{user.Email})

	if mailErr != nil {
		c.JSON(http.StatusInternalServerError, mailErr.Error)
	}

	otp.User_id = issuerInt
	otp.Otp_secret = key.Secret()
	otp.Otp_auth_url = key.URL()
	otp.Passcode = passcode

	updateErr := otphandler.os.UpdateOTP(otp)

	if updateErr != nil {
		c.JSON(http.StatusBadRequest, updateErr)
		return
	}

	otpResponse := gin.H{
		"base32":      key.Secret(),
		"otpauth_url": key.URL(),
	}

	c.JSON(http.StatusOK, otpResponse)
}

func (otphandler *OTPHandler) VerifyOTP(c *gin.Context) {
	var token *domain.OTPTokenRequest
	otp := domain.OTP{}
	user := new(domain.User)

	cookie, _ := c.Cookie("jwt")

	issuer, issuerErr := jwt_util.GetIssuer(cookie)

	if issuerErr != nil {
		restErr := errors.NewBadRequestError("Couldn't get user information")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

	if err := c.ShouldBindJSON(&token); err != nil {
		tokenErr := errors.NewBadRequestError(err.Error())
		c.JSON(tokenErr.Status, tokenErr)
		return
	}

	otp.User_id = issuerInt

	getErr := otphandler.os.GetOTPSecret(&otp)

	if getErr != nil {
		restErr := errors.NewBadRequestError("Couldn't get user information")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	isValid := otp.Validate(token.Token)

	if !isValid {
		restErr := errors.NewBadRequestError("Token isn't valid")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.ID = issuerInt

	_, getErr = otphandler.us.GetUserByID(user)

	if getErr != nil {
		restErr := errors.NewBadRequestError("Couldn't get user information")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	updateErr := otphandler.os.EnableUserOTP(&domain.OTP{User_id: otp.User_id})

	if updateErr != nil {
		restErr := errors.NewBadRequestError("Updating the one time password resulted in an error !")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	userResponse := gin.H{
		"id":           otp.User_id,
		"name":         user.FirstName,
		"email":        user.Email,
		"otp_enabled":  true,
		"otp_verified": true,
	}

	c.JSON(http.StatusOK, userResponse)
}

// func (otphandler *OTPHandler) VerifyOTP(c *gin.Context) {
// 	var token *domain.OTPTokenRequest
// 	otp := domain.OTP{}
// 	user := new(domain.User)

// 	cookie, _ := c.Cookie("jwt")

// 	issuer, issuerErr := jwt_util.GetIssuer(cookie)

// 	if issuerErr != nil {
// 		restErr := errors.NewBadRequestError("Couldn't get user information")
// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}

// 	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

// 	if err := c.ShouldBindJSON(&token); err != nil {
// 		tokenErr := errors.NewBadRequestError(err.Error())
// 		c.JSON(tokenErr.Status, tokenErr)
// 		return
// 	}

// 	otp.User_id = issuerInt

// 	getErr := otphandler.os.GetOTPSecret(&otp)

// 	if getErr != nil {
// 		restErr := errors.NewBadRequestError("Couldn't get user information")
// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}
// 	fmt.Println(otp.Passcode, token.Token)
// 	isValid := totp.Validate(token.Token, otp.Otp_secret)

// 	if !isValid {
// 		restErr := errors.NewBadRequestError("Token isn't valid")
// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}
// 	user.ID = issuerInt

// 	_, getErr = otphandler.us.GetUserByID(user)

// 	if getErr != nil {
// 		restErr := errors.NewBadRequestError("Couldn't get user information")
// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}

// 	updateErr := otphandler.os.EnableUserOTP(&domain.OTP{User_id: otp.User_id})

// 	if updateErr != nil {
// 		restErr := errors.NewBadRequestError("Updating the one time password resulted in an error !")
// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}

// 	userResponse := gin.H{
// 		"id":           otp.User_id,
// 		"name":         user.FirstName,
// 		"email":        user.Email,
// 		"otp_enabled":  true,
// 		"otp_verified": true,
// 	}

// 	c.JSON(http.StatusOK, userResponse)
// }
