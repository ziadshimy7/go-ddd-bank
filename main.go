package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-ddd-bank/docs"
	repo "github.com/go-ddd-bank/domain/repository"
	services "github.com/go-ddd-bank/domain/service"
	"github.com/go-ddd-bank/infrastructure/api"
	"github.com/go-ddd-bank/infrastructure/db"
	"github.com/go-ddd-bank/infrastructure/http/middleware"
	infrastructure "github.com/go-ddd-bank/infrastructure/http/routes"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go + Gin Domain Driven Design Bank
// @version 1.0
// @description This is a sample bank server

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataDriver := os.Getenv("DATABASE_DRIVER")
	// databaseHost := os.Getenv("DATABASE_HOST")
	// if databaseHost == "" {
	databaseHost := "127.0.0.1:3306"
	// }

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

	//run swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

}
