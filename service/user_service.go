package service

import (
	"errors"
	"log"
	"net/mail"

	"github.com/devstackq/binaryx/models"
	"github.com/devstackq/binaryx/repository"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
}

//handle request
func validDomain(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (us *UserService) CreateUser(user *models.User) error {

	if len(user.Password) < 8 {
		return errors.New("password length must be more 8")
	}
	if len(user.FirstName) < 1 {
		return errors.New("first name field not empty")
	}
	if len(user.LastName) < 1 {
		return errors.New("last name field not empty")
	}

	if !validDomain(user.Email) {
		return errors.New("invalid email domain")
	}
	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return errors.New("hash pwd error")
	}
	user.Password = hash

	uuid := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		log.Println(err)
		return errors.New("uuid error")
	}

	user.UUID = uuid

	us.repository.CreateUser(user)

	return nil
}
