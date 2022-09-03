package connect

import (
	"errors"
	"io"
	"iot-master-gateway/model"
	"net"
	"time"
)

//ServerUdpTunnel UDP链接
type ServerUdpTunnel struct {
	tunnelBase

	conn *net.UDPConn
	addr *net.UDPAddr
}

func newServerUdpTunnel(tunnel *model.Tunnel, conn *net.UDPConn, addr *net.UDPAddr) *ServerUdpTunnel {
	return &ServerUdpTunnel{
		tunnelBase: tunnelBase{
			tunnel: tunnel,
			link:   conn,
		},
		conn: conn,
		addr: addr,
	}
}

func (l *ServerUdpTunnel) Open() error {
	return errors.New("ServerUdpTunnel cannot open")
}

func (l *ServerUdpTunnel) Close() error {
	return errors.New("ServerUdpTunnel cannot close")
}

//Write 写
func (l *ServerUdpTunnel) Write(data []byte) error {
	if !l.running {
		return errors.New("tunnel closed")
	}
	if l.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := l.conn.WriteToUDP(data, l.addr)
	return err
}

func (l *ServerUdpTunnel) Ask(cmd []byte, timeout time.Duration) ([]byte, error) {
	if !l.running {
		return nil, errors.New("tunnel closed")
	}
	//堵塞
	l.lock.Lock()
	defer l.lock.Unlock() //自动解锁

	err := l.Write(cmd)
	if err != nil {
		return nil, err
	}
	return l.wait(timeout)
}

func (l *ServerUdpTunnel) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
	l.pipe = pipe

	//传入空，则关闭
	if l.pipe == nil {
		return
	}

	go func() {
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
			//n, err = l.link.Write(buf[:n])
			_, err = l.conn.WriteToUDP(buf[:n], l.addr)
			if err != nil {
				//发送失败，说明连接失效
				_ = pipe.Close()
				break
			}
		}
		l.pipe = nil
	}()
}

func (l *ServerUdpTunnel) onData(data []byte) {
	l.running = true
	l.online = true

	//透传
	if l.pipe != nil {
		_, err := l.pipe.Write(data)
		if err == nil {
			return
		}
		l.pipe = nil
	}

	l.Emit("data", data)
}
