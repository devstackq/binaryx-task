package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/devstackq/binaryx/models"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}

var JwtSecret = []byte("secret_key")

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {

	u := models.User{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
	}
	err = h.Services.CreateUser(&u)
	//then -> nextHandler, wallet - btc, eth set by uiud
	if err != nil {
		log.Println(err, ": signup err")
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	//set default currency and put money
	err = h.Services.InitBalance(&models.Account{UUID: u.UUID, Balance: 100, CurrencyId: 1, Email: u.Email})
	if err != nil {
		log.Println(err, "init balance:1")
	}
	err = h.Services.InitBalance(&models.Account{UUID: u.UUID, Balance: 100, CurrencyId: 2, Email: u.Email})
	if err != nil {
		log.Println(err, "init balance:2")
	}
	w.WriteHeader(200)
}

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {

	u := models.User{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err, "read err")
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err, "unmarshal lerr")
	}

	err = h.Services.Signin(&u)
	if err != nil {
		log.Println(err, "signin err")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//if req data ok -> set jwt token -> redirect profile page..
	expirationTime := time.Now().Add(15 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//if ok - set jwt client side
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	log.Println("set jwt,", tokenString)
	w.WriteHeader(http.StatusOK)
}
