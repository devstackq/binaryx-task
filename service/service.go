package service

import (
	"github.com/devstackq/binaryx/models"
	"github.com/devstackq/binaryx/repository"
)

//business logic
type User interface {
	CreateUser(models.User) error
}

type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
	}
}
