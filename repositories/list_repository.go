package repositories

import (
	"github.com/raddva/projeqtor-api-go/config"
	"github.com/raddva/projeqtor-api-go/models"
)

type ListRepository interface {
}

type listRepository struct {
}

func NewListRepository() ListRepository {
	return &listRepository{}
}

func (r *listRepository) Create(list *models.List) error {
	return config.DB.Create(list).Error
}