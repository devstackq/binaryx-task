package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"

	"github.com/devstackq/bynaryx/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//handle request
func validDomain(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
func Signup(r *http.Request) (statusCode int, data map[string]interface{}){
var errMsg string
	u := models.User{}
b, err := ioutil.ReadAll(r.Body)
if err != nil {
	log.Println(err)
}
err = json.Unmarshal(b, &u)
if err != nil {
	log.Println(err)
}

if len(u.Password) <8 {
	errMsg = "password length must be more 8"
}
if len(u.FirstName)  < 1  {
	errMsg = "first name field not empty"
}
if len(u.LastName)  < 1  {
	errMsg = "last name field not empty"
}

if !validDomain(u.Email) {
	errMsg = "invalid email domain"
}

hash, err := HashPassword(u.Password) 
if err != nil {
	errMsg = "uuid error"
	log.Println(err)
}
u.Password = hash

uuid := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		log.Println(err)
		errMsg = "uuid error"
	}

u.UUID =  uuid

//  repository.Signup(user)


//signup - sql


//then -> nextHandler, wallet - btc, eth set by uiud




//valid email. password
if errMsg != "" {
return 400,	map[string]interface{}{
		"error": errMsg,
	}
}	
	return 200, nil
}