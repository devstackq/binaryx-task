package main

import (
	"log"

	"github.com/devstackq/binaryx/server"
)

//data save in db psql, tables : user, wallet, coins,
//backend- start - docker container run..
//use interface - poluymorphism -> each wallet - 1 method
//use gorutine
//use middleware - auth user
//handlers, enpoint1 /signup, /signin, /wallet, /transfer, /transaction

//use middleware todo:
//init currency eth, btc - DB todo:
func main() {

	s := server.NewServer()
	err := s.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
