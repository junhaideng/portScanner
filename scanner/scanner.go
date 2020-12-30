package scanner


type State string

const (
	Open   State = "open"
	Closed State = "closed"

	UDP = "udp"
	TCP = "tcp"

	PortNum int = 1 << 16
)

type Result struct {
	Protocol string  `json:"protocol,omitempty"`
	Port     string `json:"port,omitempty"`
	State    State  `json:"state,omitempty"`
}
