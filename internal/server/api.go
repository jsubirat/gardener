package server

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	measurement := &types.Measurement{}
	measurement.Measurement = "sensor-dht22"
	fields := &types.Fields{}
	fields.Temperature = 32.0
	fields.Humidity = 75.0
	measurement.Fields = fields

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurement)
}
