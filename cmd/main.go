package main

import (
	"log"

	"github.com/devstackq/binaryx/server"
)

func main() {
	s := server.NewServer()
	err := s.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
