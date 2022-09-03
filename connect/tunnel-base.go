package connect

import (
	"errors"
	"io"
	"iot-master-gateway/model"
	"iot-master-gateway/pkg/events"
	"sync"
	"time"
)

type tunnelBase struct {
	events.EventEmitter

	tunnel *model.Tunnel

	lock sync.Mutex

	link io.ReadWriteCloser

	running bool
	online  bool
	first   bool

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

func (l *tunnelBase) First() bool {
	return l.first
}

// Close 关闭
func (l *tunnelBase) Close() error {
	if l.retryTimer != nil {
		l.retryTimer.Stop()
	}
	if !l.running {
		return errors.New("tunnel closed")
	}
	l.Emit("close")
	l.onClose()
	return l.link.Close()
}

func (l *tunnelBase) onClose() {
	l.running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
	l.Emit("close")
}

// Write 写
func (l *tunnelBase) Write(data []byte) error {
	if !l.running {
		return errors.New("tunnel closed")
	}
	if l.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := l.link.Write(data)
	return err
}

func (l *tunnelBase) wait(duration time.Duration) ([]byte, error) {
	resp := make(chan []byte, 1)
	l.Once("data", func(data []byte) {
		resp <- data
	})
	select {
	case <-time.After(duration):
		return nil, errors.New("超时")
	case buf := <-resp:
		return buf, nil
	}
}

func (l *tunnelBase) Ask(cmd []byte, timeout time.Duration) ([]byte, error) {
	if !l.running {
		return nil, errors.New("tunnel closed")
	}

	//堵塞
	l.lock.Lock()
	defer l.lock.Unlock() //自动解锁

	_, err := l.link.Write(cmd)
	if err != nil {
		return nil, err
	}
	return l.wait(timeout)
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
}
