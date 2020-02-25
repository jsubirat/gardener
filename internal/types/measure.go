package types

type Fields struct {
	temperature float32 `json:"temperature"`
	humidity    float32 `json:"humidity"`
}

type Measurement struct {
	measurement string `json:"measurement"`
	fields      Fields `json:"fields"`
}
