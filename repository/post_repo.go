package repository

import (
	// "../entity"
	"github.com/vipindasvg/golang-rest-api/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

