package repository

import (
	"database/sql"

	"github.com/devstackq/binaryx/models"
)

//use cases, interface - for relaition layer
type User interface {
	GetUserPassword(string) (string, error)
	CreateUser(*models.User) error
}

type Wallet interface {
	AddCurrency(string, float64) error
	InitBalance(*models.Account) error
	GetAccountsByEmail(string) ([]models.Account, error)
	CheckWallet(*models.Account) (*models.Account, error)
	Transfer(*models.Account, *models.Account) error
	GetUUIDByEmail(string) (string, error)
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
