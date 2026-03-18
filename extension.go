package cnc

// Codec marshals and unmarshals extension values.
type Codec interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

// Extension is an ASN.1-inspired pluggable field.
// If Critical is true and the receiver has no registered handler for ID,
// it must abort processing of the enclosing message.
// If Critical is false, unknown extensions are silently ignored.
type Extension struct {
	ID       string `json:"id"`
	Critical bool   `json:"critical"`
	Value    []byte `json:"value"` // opaque; decoded by registered handler
}

type Extensions []Extension
