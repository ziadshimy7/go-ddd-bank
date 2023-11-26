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
)

type UserHandler struct {
	us services.UserServiceHandlers
}

func NewUserHandler(us services.UserServiceHandlers) *UserHandler {
	return &UserHandler{us: us}
}

// RegisterUser             godoc
// @Summary Registers a user and returns the user info with password
// @ID register-user
// @name RegisterUser
// @Produce json
// @Param firstName query string true "first name of the user"
// @Param lastName query string true "last name of the user"
// @Param phone query string true "phone number of the user (must start with a + and country code eg. +7)"
// @Param email query string true "user's email (must be a valid email)"
// @Param password query string true "user's password (must be a strong password, containing an uppercase, lowercase and symbol)"
// @Success 200 {object} dto.UserDTO
// @Failure 401 {object} errors.Errors
// @Tags Auth: Register User
//
// @Router /api/auth/register [post]
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

// Login             godoc
// @Summary Log in the user if the username and password are found in the db
// @ID login
// @name Login
// @Produce json
// @Param email query string true "user's email"
// @Param password query string true "user's password"
// @Success 200 {object} dto.UserDTO
// @Header 200 {string} Set-Cookie "jwt=token; Expires=expires; HttpOnly" true
// @Failure 401 {object} errors.Errors
// @Tags Auth: Login User
//
// @Router /api/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	user := domain.NewUser(&domain.User{})

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
	token, tokenErr := jwt_util.CreateJWT(result, expireTime)

	if tokenErr != nil {
		err := errors.NewInternalServerError("login failed")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, result)
}

// GetUser             godoc
// @Summary Get a single user by jwt token (if passed in headers)
// @ID get-user
// @name GetUser
// @Produce json
//
// @Param 		 Cookie header string  false "jwt"     default(token=xxx)
// @Success 200 {object} dto.UserDTO
// @Failure 401 {object} errors.Errors
// @Tags Auth: Get User
//
//	@Security		ApiKeyAuth
//
// @Router /api/auth/user [get]
func (uhandler *UserHandler) GetUser(c *gin.Context) {
	var user = &domain.User{}

	cookie, err := c.Cookie("jwt")

	if err != nil {
		getErr := errors.NewBadRequestError("Couldn't retrieve cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	issuer, issuerErr := jwt_util.GetIssuer(cookie)

	if issuerErr != nil {
		restErr := errors.NewInternalServerError("error parsing issuer")
		c.JSON(restErr.Status, restErr)
		return
	}

	issuerInt, _ := strconv.ParseInt(issuer, 10, 64)

	user.ID = issuerInt

	userResponse, getErr := uhandler.us.GetUserByID(user)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, userResponse)

}

func (uhandler *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
