//go:build excludetinygo

package encoder

import (
	"encoding/json"
	"errors"
)

func (sensor *Sensor) MarshalJSON() ([]byte, error) {
	err := sensor.validateAvailability()
	if err != nil {
		return nil, err
	}
	if sensor.UniqueID == "" {
		return nil, errors.New("unique_id-undefined")
	}
	if sensor.Name == "" {
		return nil, errors.New("name-undefined")
	}
	return json.Marshal(sensor)
}

func (t *Temperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

func (h *Humidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(h)
}

func (th *TempAndHumidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(th)
}
