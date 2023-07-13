package internal

import (
	"github.com/iot-master-contrib/gateway/define"
	"io"
	"sync"
	"time"
)

type Messenger struct {
	Timeout time.Duration
	mu      sync.Mutex
	tunnel  define.Conn
}

func (m *Messenger) Ask(request []byte, response []byte) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//s := bufio.NewReader(m.tunnel)

	//先写
	_, err := m.tunnel.Write(request)
	if err != nil {
		return 0, err
	}

	//读超时
	err = m.tunnel.SetReadTimeout(m.Timeout)
	if err != nil {
		return 0, err
	}

	return m.tunnel.Read(response)
}

func (m *Messenger) AskAtLeast(request []byte, response []byte, min int) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//先写
	_, err := m.tunnel.Write(request)
	if err != nil {
		return 0, err
	}

	//读超时
	err = m.tunnel.SetReadTimeout(m.Timeout)
	if err != nil {
		return 0, err
	}

	return io.ReadAtLeast(m.tunnel, response, min)
}

func (m *Messenger) Read(response []byte) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//读超时
	err := m.tunnel.SetReadTimeout(m.Timeout)
	if err != nil {
		return 0, err
	}
	//读
	return m.tunnel.Read(response)
}

func (m *Messenger) ReadAtLeast(response []byte, min int) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	//读超时
	err := m.tunnel.SetReadTimeout(m.Timeout)
	if err != nil {
		return 0, err
	}

	return io.ReadAtLeast(m.tunnel, response, min)
}
