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

	// User 
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Board
	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NeewBoardService(boardRepo, userRepo, boardMemberRepo)
	boardController := controllers.NewBoardController(boardService)

	routes.Setup(app, userController, boardController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))

}