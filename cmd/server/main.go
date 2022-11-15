package main

import (
	"log"

	"github.com/mirroris/proglog/inter"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
