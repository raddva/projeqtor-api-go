package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/raddva/projeqtor-api-go/config"
	"github.com/raddva/projeqtor-api-go/controllers"
	"github.com/raddva/projeqtor-api-go/database/seed"
	"github.com/raddva/projeqtor-api-go/repositories"
	"github.com/raddva/projeqtor-api-go/routes"
	"github.com/raddva/projeqtor-api-go/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))

}