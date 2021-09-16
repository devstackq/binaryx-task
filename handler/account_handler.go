package handler

import (
	"log"
	"net/http"
)

//счет
type Customer interface {
	GetAccount()
	UpdateAccount()
	TransferMoney()
}

func (h *Handler) getEmailJwt(r *http.Request) (string, error) {

	values, err := r.Cookie("jwt_token")
	if err != nil {
		log.Println(err, "err jwt token")
		return "", err
	}
	log.Println(values)
	return "ss", nil
}
func (h *Handler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	//get acccounts by email , get id, gey by id - all wallets
	// /get email, by jwt ?
	e, err := h.getEmailJwt(r)
	log.Println(e, err)
}

//metdos realize, each currency?
