package device

import "github.com/god-jason/bucket/boot"

func init() {
	boot.Register("mqtt", &boot.Task{
		Startup:  Startup,
		Shutdown: Shutdown,
		Depends:  []string{"config", "mqtt"},
	})
}

func Startup() error {
	return nil
}

func Shutdown() error {
	return nil
}
