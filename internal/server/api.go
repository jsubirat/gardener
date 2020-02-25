package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
	"github.com/gorilla/mux"
	types "github.com/jsubirat/gardener/internal/types"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/sensors", a.fetchSensors).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchSensors(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Called to get sensors\n")

	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n", temperature, humidity, retried)

	measurement := &types.Measurement{}
	measurement.Measurement = "sensor-dht22"
	fields := &types.Fields{}
	fields.Temperature = temperature
	fields.Humidity = humidity
	measurement.Fields = fields

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurement)
}
