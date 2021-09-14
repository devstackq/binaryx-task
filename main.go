package main

import (
	"log"
	"net/http"

	"github.com/devstackq/bynaryx/handlers"
	"github.com/devstackq/bynaryx/repository"
)

//data save in db psql, tables : user, wallet, coins,
//backend- start - docker container run..
//use interface - poluymorphism -> each wallet - 1 method
//use gorutine
//use middleware - auth user
//handlers, enpoint1 /signup, /signin, /wallet, /transfer, /transaction

type Test struct {
	Id int `json:"id"`
Name string `json:"name"`
}
// var Db *sql.DB

func main(){

	db, err := repository.CreateDB()

	//chain interface relation between layer -> repos->services->handlers
	//outer layer connect -> inner - with interfaces, then realize interfaces
	// repos := repository.NewRepository(db)
	// services := service.NewService(repos)
	// handler := handler.NewHandler(services)

	// s := &Server{
	// 	http: &http.Server{
	// 		Addr:         ":8000",
	// 		Handler:      handlers.NewRouter(),
	// 		WriteTimeout: 10 * time.Second,
	// 		ReadTimeout:  10 * time.Second,
	// 	},
	// }
	// return s.http.ListenAndServe()


	r := handlers.NewRouter()

	r.POST("/signup", handlers.Signup)

	//preapre route, path handlker -> add routes struct -> start server with handlers

	log.Println("server start 8000")
	err := http.ListenAndServe(":8000", r)
if err != nil {
	log.Println(err)
}
}
