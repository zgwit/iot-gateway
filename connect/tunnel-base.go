package connect

import (
	"errors"
	"github.com/zgwit/iot-master/v2/model"
	"io"
	"sync"
	"time"
)

type tunnelBase struct {
	tunnel *model.Tunnel

	lock sync.Mutex

	link io.ReadWriteCloser

	running bool
	online  bool

	retry      int
	retryTimer *time.Timer

	pipe io.ReadWriteCloser
}

func (l *tunnelBase) Model() *model.Tunnel {
	return l.tunnel
}

func (l *tunnelBase) Running() bool {
	return l.running
}

func (l *tunnelBase) Online() bool {
	return l.online
}

// Close 关闭
func (l *tunnelBase) Close() error {
	if l.retryTimer != nil {
		l.retryTimer.Stop()
	}
	if !l.running {
		return errors.New("tunnel closed")
	}

	l.onClose()
	return l.link.Close()
}

func (l *tunnelBase) onClose() {
	l.running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
}

// Write 写
func (l *tunnelBase) Write(data []byte) (int, error) {
	if !l.running {
		return 0, errors.New("tunnel closed")
	}
	if l.pipe != nil {
		return 0, nil //透传模式下，直接抛弃
	}
	return l.link.Write(data)
}

// Write 写
func (l *tunnelBase) Read(data []byte) (int, error) {
	if !l.running {
		return 0, errors.New("tunnel closed")
	}
	if l.pipe != nil {
		//TODO 先read，然后透传
		return 0, nil //透传模式下，直接抛弃
	}
	return l.link.Read(data)
}

func (l *tunnelBase) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}

	l.pipe = pipe
	//传入空，则关闭
	if pipe == nil {
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := pipe.Read(buf)
		if err != nil {
			//if err == io.EOF {
			//	continue
			//}
			//pipe关闭，则不再透传
			break
		}
		//将收到的数据转发出去
		n, err = l.link.Write(buf[:n])
		if err != nil {
			//发送失败，说明连接失效
			_ = pipe.Close()
			break
		}
	}
	l.pipe = nil

	//TODO 使用io.copy
	//go io.Copy(pipe, l.link)
	//go io.Copy(l.link, pipe)
}
