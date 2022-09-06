package core

import (
	"errors"
	"io"
	"iot-master-gateway/model"
	"net"
)

// TunnelUdpServer UDP服务器
type TunnelUdpServer struct {
	tunnelBase
	addr *net.UDPAddr
	conn *net.UDPConn
}

func newTunnelUdpServer(tunnel *model.Tunnel) *TunnelUdpServer {
	svr := &TunnelUdpServer{
		tunnelBase: tunnelBase{tunnel: tunnel},
	}
	return svr
}

// Open 打开
func (server *TunnelUdpServer) Open() error {
	if server.running {
		return errors.New("server is opened")
	}

	addr, err := net.ResolveUDPAddr("udp", resolvePort(server.tunnel.Addr))
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		//TODO 需要正确处理接收错误
		return err
	}
	server.conn = conn //共用连接
	server.running = true

	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				_ = conn.Close()
				//continue
				break
			}
			server.online = true
			server.addr = addr

			data := buf[:n]
			//过滤心跳包
			if server.tunnel.Heartbeat.Enable && server.tunnel.Heartbeat.Check(data) {
				continue
			}
		}
		server.running = false
		server.online = false
	}()

	return nil
}

// Write 写
func (server *TunnelUdpServer) Write(data []byte) error {
	if server.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := server.conn.WriteToUDP(data, server.addr)
	if err != nil {
		server.onClose()
	}
	return err
}

func (server *TunnelUdpServer) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if server.pipe != nil {
		_ = server.pipe.Close()
	}
	server.pipe = pipe

	//传入空，则关闭
	if server.pipe == nil {
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
			//n, err = server.link.Write(buf[:n])
			_, err = server.conn.WriteToUDP(buf[:n], server.addr)
			if err != nil {
				//发送失败，说明连接失效
				_ = pipe.Close()
				break
			}
		}
		server.pipe = nil
	}()
}

// Close 关闭
func (server *TunnelUdpServer) Close() (err error) {
	if !server.running {
		return errors.New("tunnel closed")
	}
	server.onClose()
	return server.conn.Close()
}
