package json

import (
	"encoding/json"

	types "github.com/remotehand/cnc"
)

// Codec implements types.Codec using standard library JSON.
type Codec struct{}

var _ types.Codec = Codec{}

func (Codec) Marshal(v any) ([]byte, error)     { return json.Marshal(v) }
func (Codec) Unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }
