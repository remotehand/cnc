package json

import (
	"encoding/json"

	"github.com/remotehand/cnc"
)

// Codec implements cnc.Codec using standard library JSON.
type Codec struct{}

var _ cnc.Codec = Codec{}

func (Codec) Marshal(v any) ([]byte, error)     { return json.Marshal(v) }
func (Codec) Unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }
