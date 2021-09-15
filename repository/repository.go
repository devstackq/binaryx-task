package repository

import (
	"database/sql"

	"github.com/devstackq/binaryx/models"
)

//use cases, interface - for relaition layer
type User interface {
	CreateUser(*models.User) error
}

type Wallet interface {
	AddCurrency(string, float64) error
	InitBalance(*models.Account) error
}

type Repository struct {
	User
	Wallet
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:   NewUserRepository(db),
		Wallet: NewWalletRepository(db),
	}
}
