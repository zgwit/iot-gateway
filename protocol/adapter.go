package protocol

import (
	"github.com/god-jason/bucket/types"
)

type Adapter interface {
	//Tunnel() connect.Tunnel

	//设备动态添加
	Mount(device string, product string, station types.Options) error
	Unmount(device string) error

	//设备数据操作
	Get(device, point string) (any, error)
	Set(device, point string, value any) error
	Sync(device string) (map[string]any, error)
}
