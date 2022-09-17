package connect

import (
	"errors"
	"github.com/zgwit/iot-master/v2/model"
	"io"
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

	return nil
}

func (server *TunnelUdpServer) Read(data []byte) (int, error) {
	n, addr, err := server.conn.ReadFromUDP(data)
	server.addr = addr
	return n, err
}

// Write 写
func (server *TunnelUdpServer) Write(data []byte) (int, error) {
	if server.pipe != nil {
		return 0, nil //透传模式下，直接抛弃
	}
	return server.conn.WriteToUDP(data, server.addr)
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
