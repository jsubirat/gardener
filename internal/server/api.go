package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsubirat/gardener/internal/sensors"
)

type api struct {
	router http.Handler
}

// Server type contains a Router with all the available routes properly wired
type Server interface {
	Router() http.Handler
}

// NewServer returns an instance of a Server, properly initialized
func NewServer() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/sensors/dht-22", a.fetchDHT22Sensor).Methods(http.MethodGet)
	r.HandleFunc("/sensors/moisture", a.fetchMoistureSensor).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchDHT22Sensor(w http.ResponseWriter, r *http.Request) {
	log.Println("Called /sensors/dht-22")

	measurement, err := sensors.MeasureDHT22()
	if err != nil {
		log.Println("Error while getting DHT22 measurement")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurement)
}

func (a *api) fetchMoistureSensor(w http.ResponseWriter, r *http.Request) {
	log.Println("Called /sensors/moisture")

	measurement, err := sensors.MeasureMoisture(0)
	if err != nil {
		log.Println("Erro while getting moisture measurement")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measurement)
}
