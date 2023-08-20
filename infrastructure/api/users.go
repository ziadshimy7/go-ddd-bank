package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	domain "github.com/go-ddd-bank/domain/model"
	services "github.com/go-ddd-bank/domain/service"
	errors "github.com/go-ddd-bank/utils"
	jwt_util "github.com/go-ddd-bank/utils/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
)

type UserHandler struct {
	us *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(err.Status, err)
		return
	}

	result, err := h.us.CreateUser(&user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) Login(c *gin.Context) {
	var user *domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(err.Status, err)
		return
	}

	result, err := h.us.GetUserByEmail(user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var expireTime *jwt.NumericDate = jwt.NewNumericDate(time.Now().Add(time.Hour * 72))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{Issuer: strconv.Itoa(int(result.ID)), ExpiresAt: expireTime})
	token, tokenErr := claims.SignedString([]byte(jwt_util.SecretKey))
	if tokenErr != nil {
		err := errors.NewInternalServerError("login failed")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, result)
}

func (uhandler *UserHandler) GetUser(c *gin.Context) {
	var user = &domain.User{}

	cookie, err := c.Cookie("jwt")

	if err != nil {
		getErr := errors.NewInternalServerError("Couldn't retrieve cookie")
		c.JSON(getErr.Status, getErr)
	}

	issuer, issuerErr := jwt_util.GetIssuer(cookie)

	if issuerErr != nil {
		restErr := errors.NewInternalServerError("error parsing issuer")
		c.JSON(restErr.Status, restErr)
		return
	}

	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

	user.ID = issuerInt

	getErr := uhandler.us.GetUserByID(user)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func (uhandler *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (uhandler *UserHandler) GetOTP(c *gin.Context) {
	user := &domain.User{}

	cookie, _ := c.Cookie("jwt")

	issuer, issuerErr := jwt_util.GetIssuer(cookie)

	if issuerErr != nil {
		restErr := errors.NewBadRequestError("Couldn't get user information")

		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	issuerInt, err := strconv.ParseInt(issuer, 10, 64)

	if err != nil {
		restErr := errors.NewBadRequestError("Couldn't get user information")

		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.ID = issuerInt

	getErr := uhandler.us.GetUserByID(user)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: user.Email,
		SecretSize:  15,
	})

	if err != nil {
		panic(err)
	}

	// dataToUpdate := &domain.User{
	// 	Otp_secret:   key.Secret(),
	// 	Otp_auth_url: key.URL(),
	// }

	// ac.DB.Model(&user).Updates(dataToUpdate)

	otpResponse := gin.H{
		"base32":      key.Secret(),
		"otpauth_url": key.URL(),
	}

	c.JSON(http.StatusOK, otpResponse)
}
