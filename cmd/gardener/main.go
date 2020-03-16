package main

import (
	"log"
	"net/http"

	"github.com/jsubirat/gardener/internal/server"
)

func main() {
	s := server.NewServer()

	log.Fatal(http.ListenAndServe(":80", s.Router()))
}
