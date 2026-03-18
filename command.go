package cnc

type CommandType string

const (
	CommandExec   CommandType = "exec"   // run to completion
	CommandSpawn  CommandType = "spawn"  // persistent process
	CommandSignal CommandType = "signal" // send signal to running process
	CommandResize CommandType = "resize" // resize PTY
)

type CommandState string

const (
	CommandPending  CommandState = "pending"
	CommandRunning  CommandState = "running"
	CommandExited   CommandState = "exited"
	CommandKilled   CommandState = "killed"
	CommandTimedOut CommandState = "timed-out"
)

type PTYConfig struct {
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
	Term string `json:"term,omitempty"` // e.g. "xterm-256color"
}

// ExecPayload is used by CommandExec and CommandSpawn.
type ExecPayload struct {
	Argv    []string          `json:"argv"`
	Env     map[string]string `json:"env,omitempty"`
	Cwd     string            `json:"cwd,omitempty"`
	Timeout int               `json:"timeout,omitempty"` // seconds; 0 = none
	PTY     *PTYConfig        `json:"pty,omitempty"`
}

type SignalPayload struct {
	Signal string `json:"signal"` // e.g. "SIGTERM"
}

type ResizePayload struct {
	Cols int `json:"cols"`
	Rows int `json:"rows"`
}

// Command carries a typed payload pre-serialized with the same codec as the outer envelope.
// The receiver decodes the outer struct first, reads Type, then decodes Payload into the
// appropriate payload struct. This avoids interface{} while keeping the outer envelope
// decodable without type knowledge.
type Command struct {
	ID        string      `json:"id"` // UUID
	SessionID string      `json:"session_id"`
	Type      CommandType `json:"type"`
	// Payload is the encoded ExecPayload/SignalPayload/ResizePayload,
	// using the same codec as the outer envelope.
	Payload []byte `json:"payload"`
	// nil = use session default channels
	ChannelOverrides []Channel  `json:"channel_overrides,omitempty"`
	Extensions       Extensions `json:"extensions,omitempty"`
}

type CommandStatus struct {
	CommandID  string       `json:"command_id"`
	State      CommandState `json:"state"`
	ExitCode   *int         `json:"exit_code,omitempty"`   // nil until exited
	SignalName *string      `json:"signal_name,omitempty"` // nil unless killed by signal
	Extensions Extensions   `json:"extensions,omitempty"`
}
