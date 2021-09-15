package handler

import (
	"log"
	"net/http"

	service "github.com/devstackq/binaryx/service"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

//data save in db psql, tables : user, wallet, coins,
//backend- start - docker container run..
//use interface - poluymorphism -> each wallet - 1 method
//use gorutine
//use middleware - auth user
//handlers, enpoint1 /signup, /signin, /wallet, /transfer, /transaction

func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", h.Signup) //post:fields, fname, lname, email, age
	log.Println("created routers")
	return mux
}
