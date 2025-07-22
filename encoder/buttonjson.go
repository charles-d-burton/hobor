//go:build excludetinygo

package encoder

import (
	"encoding/json"
	"errors"
)

func (b *Button) MarshalJSON() ([]byte, error) {
	b.Platform = "button"
	if b.CommandTopic == "" {
		return nil, errors.New("command topic is empty")
	}
	return json.Marshal(b)
}
