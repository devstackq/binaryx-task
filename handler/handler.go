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

func (h *Handler) InitRouter() *http.ServeMux {

	routes := h.createRoutes()

	log.Println("created routers")
	mux := http.NewServeMux()
	//add middleware each auth route
	for _, route := range routes {
		if route.IsAuth {
			// route.Handler = h.IsCookieValid(route.Handler)
		}
		mux.HandleFunc(route.Path, route.Handler)
	}
	return mux
}

func (h *Handler) createRoutes() []Route {

	return []Route{
		{
			Path:    "/signup",
			Handler: h.Signup,
			IsAuth:  false,
		},
		// {
		// 	Path:    "/account",
		// 	Handler: h.GetAccount,
		// 	IsAuth:  true,
		// },
	}
}
