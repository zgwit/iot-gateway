package connect

// Tunnel 通道
type Tunnel interface {
	Conn

	ID() string

	Open() error

	Close() error

	//Online() bool
}
