package main

import (
	"log"

	"github.com/juanefec/go-url-shortner/server"
)

func main() {
	s := server.NewServer()
	log.Println("Started listening.")
	log.Fatalln(s.ListenAndServe())
}
