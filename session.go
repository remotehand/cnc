package types

// Well-known session template names. Agents may also advertise custom names.
const (
	SessionTemplateShell        = "std.shell"
	SessionTemplateFileTransfer = "std.file-transfer"
	SessionTemplateDesktop      = "std.desktop"
	SessionTemplateCI           = "std.ci"
	SessionTemplateAgentTool    = "std.agent-tool"
)

type SessionPolicy struct {
	Concurrent int `json:"concurrent,omitempty"` // max concurrent sessions; 0 = unlimited
}

// SessionTemplate describes a session type an agent can host.
// Name may be a well-known constant or a custom string.
// The central sends the template name in SessionRequest; the agent
// looks up the matching SessionTemplate from its CapabilityReport.
type SessionTemplate struct {
	Name            string        `json:"name"`
	Policy          SessionPolicy `json:"policy,omitempty"`
	InitialChannels []ChannelSpec `json:"initial_channels,omitempty"` // opened automatically on session start
	AllowedChannels []ChannelSpec `json:"allowed_channels,omitempty"` // may be requested by central after start
	Extensions      Extensions    `json:"extensions,omitempty"`
}

type SessionState string

const (
	SessionPending SessionState = "pending"
	SessionActive  SessionState = "active"
	SessionClosing SessionState = "closing"
	SessionClosed  SessionState = "closed"
)

// SessionRequest is sent central → agent.
// TargetAgentID nil = fleet-routed (central picks agent).
// Non-nil = pinned to that agent UUID (e.g. CI resume on same machine).
type SessionRequest struct {
	Template      string     `json:"template"` // template name
	TargetAgentID *string    `json:"target_agent_id,omitempty"`
	Labels        KV         `json:"labels,omitempty"`
	Extensions    Extensions `json:"extensions,omitempty"`
}

// Session is the central-authoritative record; distributed to both central and agent.
type Session struct {
	ID              string       `json:"id"`       // UUID
	Template        string       `json:"template"` // template name used
	AgentID         string       `json:"agent_id"`
	State           SessionState `json:"state"`
	CreatedAt       int64        `json:"created_at"` // Unix seconds
	DefaultChannels []Channel    `json:"default_channels,omitempty"`
	Extensions      Extensions   `json:"extensions,omitempty"`
}
