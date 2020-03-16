package sensors

import (
	"errors"
	"fmt"

	dht "github.com/d2r2/go-dht"
	types "github.com/jsubirat/gardener/internal/types"
)

const dht22GpioPort = 17
const numReadRetries = 10

// MeasureDHT22 reads the temperature and humidity using the DHT22 sensor
func MeasureDHT22() (measurement *types.DHT22Measurement, err error) {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT22, dht22GpioPort, false, numReadRetries)
	if err != nil {
		return nil, errors.New("Error while reading DHT22 sensor")
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %.2f*C, Humidity = %.2f%% (retried %d times)\n", temperature, humidity, retried)

	measurement = &types.DHT22Measurement{
		Measurement: "sensor-dht22",
		Fields: &types.DHT22Fields{
			Temperature: temperature,
			Humidity:    humidity,
		},
	}

	return measurement, nil
}
