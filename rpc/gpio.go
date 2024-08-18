package rpc

const (
	GPIO_LIST uint8 = iota
	GPIO_WRITE
	GPIO_READ
	GPIO_STATUS
)

type GpioPath struct {
	Path string `json:"path"`
}

type GpioItem struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	Value    int64  `json:"value"`
	Writable bool   `json:"writable"`
}

type GpioValue struct {
	Path  string `json:"path"`
	Value int64  `json:"value"`
}

type GpioSearchRequest GpioPath

type GpioSearchResponse []GpioItem

type GpioReadRequest GpioPath

type GpioReadResponse GpioValue

type GpioWriteRequest GpioValue

type GpioWriteResponse struct{}
