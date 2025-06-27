package main

import (
	"log"

	"github.com/dipan-ck/atomdb.git/internal/server"
)

func main() {

	//redis-cli -p 6300

	err := server.StartServer("0.0.0.0:6300")

	if err != nil {
		log.Fatal(err)
	}

}
