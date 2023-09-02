package main

import (
	"fmt"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	repo "github.com/go-ddd-bank/domain/repository"
	services "github.com/go-ddd-bank/domain/service"
	"github.com/go-ddd-bank/infrastructure/api"
	"github.com/go-ddd-bank/infrastructure/db"
	infrastructure "github.com/go-ddd-bank/infrastructure/http/routes"
	"net/http"
	_ "net/http/pprof"
)

var (
	username       = "ziadshimy7"
	password       = "example-password"
	host           = "127.0.0.1:3306"
	schema         = "users_db"
	dataDriver     = "mysql"
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, host, schema)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	//database ==> Instantiating database connection (mysql)
	dbConn, err := db.NewMySqlConnection(dataDriver, dataSourceName)

	if err != nil {
		panic(err)
	}

	defer dbConn.CloseDbConnection()

	r := gin.Default()
	r.Use(CORSMiddleware())

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	//repositories =>> deals with persistance layer and domain interactions
	userRepo := repo.NewUserRepository(dbConn)
	accountRepo := repo.NewAccountRepository(dbConn)
	userAccountRepo := repo.NewAccountUserViewRepository(dbConn)
	otpRepo := repo.NewOTPRepository(dbConn)

	//services =>> application layer, it orchestrates different domain and repositories actions together
	userService := services.NewUserService(userRepo)
	accountService := services.NewAccountService(accountRepo)
	userAccountService := services.NewUserAccountService(userAccountRepo)
	otpService := services.NewOTPService(otpRepo)

	//handlers =>> http handlers
	userHandler := api.NewUserHandler(userService)
	accountHandler := api.NewAccountHandler(accountService)
	userAccountHandler := api.NewUserAccountHandler(userAccountService)
	otpHandler := api.NewOTPHandler(otpService, userService)

	// routes
	userRoutes := infrastructure.InitiateUserRoutes(userHandler)
	accountRoutes := infrastructure.InitiateAccountRoutes(accountHandler)
	userAccountRoutes := infrastructure.InitiateUserAccountRoutes(userAccountHandler)
	otpRoutes := infrastructure.InitiateOTPRoutes(otpHandler)

	userRoutes.RegisterRoutes(r)
	accountRoutes.RegisterRoutes(r)
	userAccountRoutes.RegisterRoutes(r)
	otpRoutes.RegisterRoutes(r)

	r.Run(":8080")

}
