package cnc

type ChannelType string

const (
	ChannelControl ChannelType = "std.control"
	ChannelTTY     ChannelType = "std.tty"
)

// ChannelDisposition controls whether the agent forwards or drops this channel.
type ChannelDisposition string

const (
	ChannelForward ChannelDisposition = "fwd"
	ChannelDiscard ChannelDisposition = "dsc"
)

type ChannelDirection string

const (
	ChannelBidirectional ChannelDirection = "BIDI"
	ChannelAgentToServer ChannelDirection = "MISO"
	ChannelServerToAgent ChannelDirection = "MOSI"
)

// ChannelSpec describes a channel in a SessionTemplate (no ID or SessionID).
type ChannelSpec struct {
	Name           string             `json:"name"` // type-specific name, e.g. stdin/stdout/stderr
	Type           ChannelType        `json:"type"`
	Direction      ChannelDirection   `json:"direction"`
	Disposition    ChannelDisposition `json:"disposition"`
	BufferSizeHint int                `json:"buffer_size_hint,omitempty"` // bytes; 0 = default
	Extensions     Extensions         `json:"extensions,omitempty"`
}

type Channel struct {
	ChannelSpec

	ID        string `json:"id"` // UUID
	SessionID string `json:"session_id"`
}

// ChannelUpdate reconfigures a live channel. Nil pointer fields = no change.
type ChannelUpdate struct {
	ChannelID   string              `json:"channel_id"`
	Disposition *ChannelDisposition `json:"disposition,omitempty"`
	Extensions  Extensions          `json:"extensions,omitempty"`
}
