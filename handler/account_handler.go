package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devstackq/binaryx/models"
	"github.com/dgrijalva/jwt-go"
)

//счет
type Customer interface {
	GetAccount()
	UpdateAccount()
	TransferMoney()
}

func ParseToken(tokenStr string) (*Claims, error) {
	var claims *Claims
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return claims, err
}
func (h *Handler) GetEmailByJwt(w http.ResponseWriter, r *http.Request) (string, error) {

	tokenStr, err := r.Cookie("jwt_token")
	if err != nil {
		log.Println(err, "err jwt token")
		return "", err
	}

	token, err := ParseToken(tokenStr.Value)
	if err != nil {
		return "", err
	}

	return token.Email, nil
}

func (h *Handler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	email, err := h.GetEmailByJwt(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	//get by email - uuid, by uuid get wallets
	seq, err := h.Services.GetAccounts(email)
	if err != nil {
		log.Println(err)
		return
	}
	b, err := json.Marshal(seq)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(seq)
	w.Write(b)
}

func (h *Handler) TransferMoney(w http.ResponseWriter, r *http.Request) {
	// get by email - uuid1, uuid2 -> send transfer, update wallets field - balnce when transfer done
	//get by jwt, email, uuid -> body {who, cash, }
	//check if money norm -> trnasfer
	a := models.Account{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(b, &a)
	if err != nil {
		log.Println(err)
	}

	email, err := h.GetEmailByJwt(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	a.Email = email
	err = h.Services.TransferMoney(&a)
	if err != nil {
		log.Println(err)
		return
	}

}

//metdos realize, each currency?
