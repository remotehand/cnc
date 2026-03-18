package cbor

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/remotehand/cnc"
)

var (
	enc cbor.EncMode
	dec cbor.DecMode
)

func init() {
	var err error
	enc, err = cbor.CanonicalEncOptions().EncMode()
	if err != nil {
		panic("cbor enc mode: " + err.Error())
	}
	dec, err = cbor.DecOptions{}.DecMode()
	if err != nil {
		panic("cbor dec mode: " + err.Error())
	}
}

// Codec implements cnc.Codec using CBOR with canonical encoding.
type Codec struct{}

var _ cnc.Codec = Codec{}

func (Codec) Marshal(v any) ([]byte, error)     { return enc.Marshal(v) }
func (Codec) Unmarshal(data []byte, v any) error { return dec.Unmarshal(data, v) }
