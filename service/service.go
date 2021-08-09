package service

import (
	"errors"
	"math/rand"

	// "../entity"
	// "../repository"
	"github.com/vipindasvg/golang-rest-api/entity"
	"github.com/vipindasvg/golang-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type Service struct{}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func NewPostService() PostService {
	return &Service{}
}

func (*Service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*Service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*Service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
