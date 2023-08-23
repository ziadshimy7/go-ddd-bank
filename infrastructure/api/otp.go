package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domain "github.com/go-ddd-bank/domain/model"
	services "github.com/go-ddd-bank/domain/service"
	errors "github.com/go-ddd-bank/utils"
	jwt_util "github.com/go-ddd-bank/utils/jwt"
	"github.com/pquerna/otp/totp"
)

type OTPHandler struct {
	os *services.OTPService
}

func NewOTPHandler(os *services.OTPService) *OTPHandler {
	return &OTPHandler{os: os}
}

func (otphandler *OTPHandler) GetOTP(c *gin.Context) {
	otp := &domain.OTP{}
	user := &domain.User{}

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
	fmt.Println(user.Email)

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: user.Email,
		SecretSize:  15,
	})

	if err != nil {
		err := errors.NewBadRequestError("Couldn't generate Secret Key for one time password, Please try again later !")
		c.JSON(err.Status, err)
		return
	}

	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

	otp.User_id = issuerInt
	otp.Otp_secret = key.Secret()
	otp.Otp_auth_url = key.URL()

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

// func (otphandler *OTPHandler) VerifyOTP(c *gin.Context) {
// 	var token *domain.OTPTokenRequest
// 	var user *domain.User

// 	cookie, _ := c.Cookie("jwt")

// 	_, issuerErr := jwt_util.GetIssuer(cookie)

// 	if issuerErr != nil {
// 		restErr := errors.NewBadRequestError("Couldn't get user information")
// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&token); err != nil {
// 		tokenErr := errors.NewBadRequestError(err.Error())
// 		c.JSON(tokenErr.Status, tokenErr)
// 	}

// 	// get user service and get the user by using it
// 	getErr := otphandler.os.GetUserByID(user)

// 	if getErr != nil {
// 		restErr := errors.NewBadRequestError("Couldn't get user information")
// 		c.JSON(http.StatusBadRequest, restErr)
// 	}

// 	isValid := totp.Validate(token.Token, user.Email)

// 	if !isValid {
// 		restErr := errors.NewBadRequestError("Token isn't valid")

// 		c.JSON(http.StatusBadRequest, restErr)
// 		return
// 	}

// 	// dataToUpdate := domain.User{
// 	// 	Otp_enabled:  true,
// 	// 	Otp_verified: true,
// 	// }

// 	userResponse := gin.H{
// 		"id":          user.ID,
// 		"name":        user.FirstName,
// 		"email":       user.Email,
// 		"otp_enabled": token.Otp_enabled,
// 	}

// 	c.JSON(http.StatusOK, gin.H{"otp_verified": true, "user": userResponse})

// }
