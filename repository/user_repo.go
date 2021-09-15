package repository

import (
	"database/sql"

	"github.com/devstackq/binaryx/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(u models.User) error {
	sqlStmt := `
	INSERT INTO users (firstname, lastname, email, password, uuid)
	VALUES ($1, $2, $3, $4)`
	_, err := ur.db.Exec(sqlStmt, u.FirstName, u.LastName, u.Email, u.Password, u.UUID)
	if err != nil {
		return err
	}
	return nil
}
