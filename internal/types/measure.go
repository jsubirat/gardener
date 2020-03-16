package types

// DHT22Fields contains the metrics that can be captured by the DHT22 sensor
type DHT22Fields struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

// DHT22Measurement is a measurement returned by the DHT22 sensor
type DHT22Measurement struct {
	Measurement string       `json:"measurement"`
	Fields      *DHT22Fields `json:"fields"`
}

// MoistureFields contains the metrics that can be captured by the moisture sensor
type MoistureFields struct {
	Moisture int `json:"moisture"`
}

// MoistureMeasurement is a measurement returned by the moisture sensor
type MoistureMeasurement struct {
	Measurement string          `json:"measurement"`
	Fields      *MoistureFields `json:"fields"`
}
