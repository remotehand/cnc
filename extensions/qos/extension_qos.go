package qos

import (
	"github.com/remotehand/cnc"
)

// ExtensionQoS is the well-known extension ID for QoS hints.
const ExtensionQoS = "std.qos"

type QoSHint struct {
	DSCP DSCP `json:"dscp,omitempty"`
}

// NewQoSExtension encodes a QoSHint as an Extension using the given Codec.
func NewQoSExtension(hint QoSHint, c types.Codec) (types.Extension, error) {
	b, err := c.Marshal(hint)
	if err != nil {
		return types.Extension{}, err
	}
	return types.Extension{ID: ExtensionQoS, Value: b}, nil
}

// QoSHintFromExtension decodes a QoSHint from an Extension using the given Codec.
func QoSHintFromExtension(ext types.Extension, c types.Codec) (QoSHint, error) {
	var hint QoSHint
	if err := c.Unmarshal(ext.Value, &hint); err != nil {
		return QoSHint{}, err
	}
	return hint, nil
}

// DSCP is a 6-bit Differentiated Services Code Point (0–63).
type DSCP uint8

// DSCP constants per RFC 2474 (Class Selector), RFC 2597 (Assured Forwarding),
// RFC 3246 (Expedited Forwarding), and RFC 8622 (Lower Effort).
const (
	// Default / Best Effort (CS0)
	DSCPDefault DSCP = 0

	// Class Selector codepoints (RFC 2474)
	DSCPCS0 DSCP = 0
	DSCPCS1 DSCP = 8
	DSCPCS2 DSCP = 16
	DSCPCS3 DSCP = 24
	DSCPCS4 DSCP = 32
	DSCPCS5 DSCP = 40
	DSCPCS6 DSCP = 48
	DSCPCS7 DSCP = 56

	// Assured Forwarding class 1 (RFC 2597)
	DSCPAF11 DSCP = 10
	DSCPAF12 DSCP = 12
	DSCPAF13 DSCP = 14

	// Assured Forwarding class 2 (RFC 2597)
	DSCPAF21 DSCP = 18
	DSCPAF22 DSCP = 20
	DSCPAF23 DSCP = 22

	// Assured Forwarding class 3 (RFC 2597)
	DSCPAF31 DSCP = 26
	DSCPAF32 DSCP = 28
	DSCPAF33 DSCP = 30

	// Assured Forwarding class 4 (RFC 2597)
	DSCPAF41 DSCP = 34
	DSCPAF42 DSCP = 36
	DSCPAF43 DSCP = 38

	// Expedited Forwarding (RFC 3246)
	DSCPEF DSCP = 46

	// Lower Effort (RFC 8622)
	DSCPLE DSCP = 1
)
