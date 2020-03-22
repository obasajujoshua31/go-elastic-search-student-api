package main

import (
	"go-elastic-search-student-api/server"
	"log"
)

func main() {
	err := server.StartServer()

	if err != nil {
		log.Fatal(err)
	}
}
