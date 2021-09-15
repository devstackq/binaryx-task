package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devstackq/binaryx/models"
)

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
	err = h.Services.InitBalance(&models.Account{UUID: u.UUID, Balance: 100, CurrencyId: 1})
	if err != nil {
		log.Println(err, "init balance:")
	}
	err = h.Services.InitBalance(&models.Account{UUID: u.UUID, Balance: 100, CurrencyId: 2})
	if err != nil {
		log.Println(err, "init balance:")
	}

	w.WriteHeader(200)

}
