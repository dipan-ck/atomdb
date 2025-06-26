package main

import (
	"log"

	"github.com/dipan-ck/redis-clone-golang.git/internal/server"
)

func main() {

	err := server.StartServer("0.0.0.0:6300")

	if err != nil {
		log.Fatal(err)
	}

}
