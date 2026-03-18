package types

// RegisterRequest is sent by the agent on startup.
// AgentID is empty on first registration; the central assigns a UUID.
// On reconnect, carry the previously assigned UUID. Auth is via Extensions.
type RegisterRequest struct {
	AgentID    string           `json:"agent_id,omitempty"`
	Capability CapabilityReport `json:"capability"`
	Extensions Extensions       `json:"extensions,omitempty"`
}

type AgentConfig struct {
	HeartbeatInterval int `json:"heartbeat_interval"` // seconds
}

type RegisterResponse struct {
	AgentID     string      `json:"agent_id"`               // confirmed or newly assigned UUID
	TrustedTags KV          `json:"trusted_tags,omitempty"` // central-assigned, authoritative
	Config      AgentConfig `json:"config"`
	Extensions  Extensions  `json:"extensions,omitempty"`
}
