package seed

import (
	"log"

	"github.com/google/uuid"
	"github.com/raddva/projeqtor-api-go/config"
	"github.com/raddva/projeqtor-api-go/models"
	"github.com/raddva/projeqtor-api-go/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("Secret101%")

	admin := models.User{
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed too seed admin", err)
	} else {
		log.Println("Admin seeded")
	}
}
