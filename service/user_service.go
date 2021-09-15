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

func (us *UserService) CreateUser(user models.User) error {

	u := models.User{}
	if len(u.Password) < 8 {
		return errors.New("password length must be more 8")
	}
	if len(u.FirstName) < 1 {
		return errors.New("first name field not empty")
	}
	if len(u.LastName) < 1 {
		return errors.New("last name field not empty")
	}

	if !validDomain(u.Email) {
		return errors.New("invalid email domain")
	}

	hash, err := HashPassword(u.Password)
	if err != nil {
		return errors.New("hash pwd error")
		log.Println(err)
	}
	u.Password = hash

	uuid := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		log.Println(err)
		return errors.New("uuid error")
	}

	u.UUID = uuid

	us.repository.CreateUser(user)

	return nil
}
