package rpc

const (
	SERIAL_LIST uint8 = iota
	SERIAL_OPEN
	SERIAL_CLOSE
	SERIAL_WATCH
	SERIAL_PIPE
)

type SerialId struct {
	Id string `json:"id"`
}

type SerialItem struct {
	Id      string        `json:"id"` //COM1 COM2 /dev/ttyS1
	Name    string        `json:"name,omitempty"`
	Options SerialOptions `json:"options"`
}

type SerialOptions struct {
	BaudRate   int    `json:"baud_rate,omitempty"`   // 9600 115200 ...
	DataBits   int    `json:"data_bits,omitempty"`   // 7 8
	StopBits   int    `json:"stop_bits,omitempty"`   // 1 2
	ParityMode string `json:"parity_mode,omitempty"` // N E O 0 1
}

type SerialListRequest struct{}

type SerialListResponse []SerialItem

type SerialOpenRequest SerialItem

type SerialOpenResponse struct{}

type SerialCloseRequest SerialId

type SerialCloseResponse struct{}

type SerialWatchRequest SerialId

type SerialWatchResponse StreamId

type SerialPipeRequest SerialId

type SerialPipeResponse StreamId
