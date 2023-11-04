package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	repo "github.com/go-ddd-bank/domain/repository"
	services "github.com/go-ddd-bank/domain/service"
	"github.com/go-ddd-bank/infrastructure/api"
	"github.com/go-ddd-bank/infrastructure/db"
	"github.com/go-ddd-bank/infrastructure/http/middleware"
	infrastructure "github.com/go-ddd-bank/infrastructure/http/routes"
	"github.com/joho/godotenv"
)

var ()

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataDriver := os.Getenv("DATABASE_DRIVER")
	databaseHost := os.Getenv("DATABASE_HOST")
	if databaseHost == "" {
		databaseHost = "127.0.0.1:3306"
	}

	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseSchema := "users_db"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", databaseUsername, databasePassword, databaseHost, databaseSchema)

	//database ==> Instantiating database connection (mysql)
	dbConn, err := db.NewMySqlConnection(dataDriver, dataSourceName)

	if err != nil {
		panic(err)
	}

	defer dbConn.CloseDbConnection()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

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
