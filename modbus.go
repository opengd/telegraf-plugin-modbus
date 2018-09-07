package modbus

import (
	"fmt"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/filter"
	"github.com/influxdata/telegraf/plugins/inputs"

	"github.com/goburrow/modbus"
)

type ModbusData struct {

}

func NewModbusData() *ModbusData {
	return &ModbusData{

	}
}

func (_ *ModbusData) Description() string {
	return "Read Modbus data from RTU"
}

var sampleConfig = ""

func (_ *ModbusData) SampleConfig() string {
	return sampleConfig
}

func (s *ModbusData) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("modbus", func() telegraf.Input {
		return &ModbusData{

		}
	})
} 
