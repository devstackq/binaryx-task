package repository

import (
	"database/sql"

	"github.com/devstackq/binaryx/models"
)

//use cases, interface - for relaition layer
type User interface {
	CreateUser(models.User) error
}

type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
