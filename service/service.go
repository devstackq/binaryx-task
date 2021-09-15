package service

import (
	"github.com/devstackq/binaryx/models"
	"github.com/devstackq/binaryx/repository"
)

//business logic
type User interface {
	CreateUser(*models.User) error
}
type Wallet interface {
	InitBalance(*models.Account) error
	AddCurrency(string, float64) error
}

type Service struct {
	User
	Wallet
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User:   NewUserService(r.User),
		Wallet: NewWalletService(r.Wallet),
	}
}
