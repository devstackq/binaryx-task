package handler

import (
	"log"
	"net/http"

	service "github.com/devstackq/binaryx/service"
)

type Handler struct {
	Services *service.Service
}
type Route struct {
	Path    string
	Handler http.HandlerFunc
	IsAuth  bool
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

//middleware, check valid token, token expires 15min
func (h *Handler) validJwtToken(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("jwt_token")
		if err != nil {
			log.Println(err, "err jwt token")
			return
		}
		f.ServeHTTP(w, r)
	}
}

func (h *Handler) InitRouter() *http.ServeMux {

	routes := h.createRoutes()
	mux := http.NewServeMux()
	for _, route := range routes {
		if route.IsAuth {
			route.Handler = h.validJwtToken(route.Handler)
		}
		mux.HandleFunc(route.Path, route.Handler)
	}
	return mux
}

func (h *Handler) createRoutes() []Route {

	return []Route{
		//firstname, lastname, email, password
		{
			Path:    "/signup",
			Handler: h.Signup,
			IsAuth:  false,
		},
		//email, password
		{
			Path:    "/signin",
			Handler: h.Signin,
			IsAuth:  false,
		},
		//get query(authorized user)
		{
			Path:    "/wallets",
			Handler: h.GetAccounts,
			IsAuth:  true,
		},
		//recepient(email), amount, currencyid
		{
			Path:    "/transfer",
			Handler: h.TransferMoney,
			IsAuth:  true,
		},
	}
}
