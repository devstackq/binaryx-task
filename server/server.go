package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devstackq/binaryx/handler"
	"github.com/devstackq/binaryx/repository"
	"github.com/devstackq/binaryx/service"
)

type Server struct {
	http *http.Server
}

func NewServer() *Server {
	db, err := repository.CreateDB()
	if err != nil {
		log.Println(err, "err create tables")
	}

	//chain interface relation between layer -> repos->services->handlers
	//outer layer connect -> inner - with interfaces, then realize interfaces
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	//create firstly currincies, btc & eth
	err = services.Wallet.AddCurrency("BTC", 40000.0)
	//check if !exist field ->insert new currency
	if err != nil {
		log.Println(err)
	}
	err = services.Wallet.AddCurrency("ETH", 1500.0)
	if err != nil {
		log.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}
	//custom server, set port, cusom routers, read/write timeout
	s := &Server{
		http: &http.Server{
			Addr:         port,
			Handler:      handler.InitRouter(),
			WriteTimeout: 6 * time.Second,
			ReadTimeout:  6 * time.Second,
		},
	}
	return s
}

func (s *Server) Run() error {
	log.Println("start server in port: ", s.http.Addr)
	return s.http.ListenAndServe()
}
