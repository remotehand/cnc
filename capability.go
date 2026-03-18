package types

// CapabilityReport is sent by the agent; entirely untrusted by the central.
type CapabilityReport struct {
	ProtocolVersion ProtocolVersion   `json:"protocol_version"`
	Hostname        string            `json:"hostname"`       // non-unique, untrusted
	Tags            KV                `json:"tags,omitempty"` // agent-reported, untrusted
	Capabilities    []SessionTemplate `json:"capabilities"`
	Extensions      Extensions        `json:"extensions,omitempty"`
}
