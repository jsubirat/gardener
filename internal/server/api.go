package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jsubirat/gardener/internal/types/measure"

	"github.com/gorilla/mux"
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

	measurements := &measure.Measurement{}
	measurements.measurement = "temperature"
	fields := &measure.Fields{}
	fields.temperature = 32.0
	fields.humidity = 75.0
	measurements.fields = fields

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurements)
}
