//go:build excludetinygo

package encoder

import (
	"encoding/json"
	"errors"
)

func (bss *BinarySensorState) MarshalJSON() ([]byte, error) {
	switch bss.State {
	case "ON", "OFF":
		return json.Marshal(bss)
	default:
		return nil, errors.New("state is not ON or OFF")
	}
}
