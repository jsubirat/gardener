package types

type Fields struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

type Measurement struct {
	Measurement string  `json:"measurement"`
	Fields      *Fields `json:"fields"`
}
