package api

import (
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ddd-bank/domain/dto"
	domain "github.com/go-ddd-bank/domain/model"
	services "github.com/go-ddd-bank/domain/service"
	errors "github.com/go-ddd-bank/utils"
	jwt_util "github.com/go-ddd-bank/utils/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func mapStructFields(source interface{}, destination interface{}) {
	sourceValue := reflect.ValueOf(source).Elem()
	destinationValue := reflect.ValueOf(destination).Elem()
	sourceType := sourceValue.Type()

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceFieldName := sourceType.Field(i).Name
		destinationField := destinationValue.FieldByName(sourceFieldName)

		if destinationField.IsValid() && destinationField.CanSet() {
			sourceFieldValue := sourceValue.Field(i)
			destinationField.Set(sourceFieldValue)
		}
	}
}

type UserHandler struct {
	us *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user domain.User
	userResponse := new(dto.UserDTO)

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
	mapStructFields(result, userResponse)

	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) Login(c *gin.Context) {
	var user *domain.User
	userResponse := new(dto.UserDTO)

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

	mapStructFields(result, userResponse)
	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, userResponse)
}

func (uhandler *UserHandler) GetUser(c *gin.Context) {
	var user = &domain.User{}
	userResponse := new(dto.UserDTO)

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
	mapStructFields(user, userResponse)

	c.JSON(http.StatusOK, user)

}

func (uhandler *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
