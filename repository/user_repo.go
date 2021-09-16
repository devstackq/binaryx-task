package repository

import (
	"database/sql"
	"log"

	"github.com/devstackq/binaryx/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserPassword(email string) (password string, err error) {

	sqlStatement := `SELECT password FROM users WHERE email=$1;`
	row := ur.db.QueryRow(sqlStatement, email)
	err = row.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (ur *UserRepository) CreateUser(u *models.User) error {
	sqlStmt := `
	INSERT INTO users (firstname, lastname, email, password)
	VALUES ($1, $2, $3, $4)`
	_, err := ur.db.Exec(sqlStmt, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
