package protocol

import "github.com/zgwit/iot-gateway/connect"

type Adapter interface {
	Tunnel() connect.Tunnel

	Mount(device string) error
	Unmount(device string) error

	Get(device, point string) (any, error)
	Set(device, point string, value any) error
	Sync(device string) (map[string]any, error)
}
