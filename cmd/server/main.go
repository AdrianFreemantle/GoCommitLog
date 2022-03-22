package main

import (
	"log"

	"github.com/AdrianFreemantle/LogService/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
