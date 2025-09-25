package main

import (
	"log"
	"ms-golang-echo/config"
	"ms-golang-echo/internal/handler"
	"ms-golang-echo/internal/migration"
	"ms-golang-echo/internal/repository"
	"ms-golang-echo/internal/usecase"
	"ms-golang-echo/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using system env")
	}
	// init logger
	config.InitLogger()

	// init dependencies
	config.PostgresDB()
	// 2. jalankan migration
	migration.Migrate()

	config.InitRedis()
	config.InitMinio()

	userRepo := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	config.InitValidator(e)

	e.Use(middleware.CORS())

	e.GET("/ping", func(c echo.Context) error {
		return config.Success(c, "pong")
	})

	api := e.Group("/users")
	routes.UserRoutes(api, userHandler)

	// run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
