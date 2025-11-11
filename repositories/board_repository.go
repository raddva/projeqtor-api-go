package repositories

import (
	"github.com/raddva/projeqtor-api-go/config"
	"github.com/raddva/projeqtor-api-go/models"
)

type BoardRepository interface {
	Create(board *models.Board) error
}

type boardRepository struct {
}

func NewBoardRepository() BoardRepository {
	return &boardRepository{}
}

func (r *boardRepository) Create(board *models.Board) error {
	return config.DB.Create(board).Error
}