package connect

// Tunnel 通道
type Tunnel interface {
	Conn

	Open() error

	Running() bool

	Online() bool
}
