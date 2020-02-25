package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
	"github.com/jsubirat/gardener/internal/server"
)

func main() {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n", temperature, humidity, retried)

	s := server.New()

	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
