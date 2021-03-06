package sensors

import (
	"errors"
	"fmt"
	"time"

	mcp "github.com/jdevelop/golang-rpi-extras/mcp3008"
	types "github.com/jsubirat/gardener/internal/types"
)

// MeasureMoisture returns the moisture of the soil
func MeasureMoisture(channel int) (measurement *types.MoistureMeasurement, e error) {
	go measure100Times()

	mcp, err := mcp.NewMCP3008(0, 0, mcp.Mode0, 500000)
	if err != nil {
		return nil, errors.New("Error while connecting to the MCP3008")
	}

	moisture, err := mcp.Measure(channel)
	if err != nil {
		return nil, errors.New("Error while obtaining measure from the MCP3008")
	}

	measurement = &types.MoistureMeasurement{
		Measurement: "moisture-sensor",
		Fields: &types.MoistureFields{
			Moisture: moisture,
		},
	}

	return measurement, nil
}

func measure100Times() {
	mcp, err := mcp.NewMCP3008(0, 0, mcp.Mode0, 500000)
	if err != nil {
		fmt.Println("Error while connecting to the MCP3008")
	}

	for i := 0; i < 100; i++ {
		moisture, err := mcp.Measure(0)
		if err != nil {
			fmt.Println("Error while obtaining measure from the MCP3008")
		} else {
			fmt.Printf("Moisture: %v\n", moisture)
		}
		time.Sleep(2 * 1000000000)
	}
}