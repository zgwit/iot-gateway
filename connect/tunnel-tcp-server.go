package connect

import (
	"errors"
	"fmt"
	"github.com/iot-master-contrib/gateway/db"
	"github.com/iot-master-contrib/gateway/dbus"
	"github.com/zgwit/iot-master/v2/model"
	"net"
	"time"
)

// TunnelTcpServer TCP服务器
type TunnelTcpServer struct {
	tunnelBase

	listener *net.TCPListener
}

func newTunnelTcpServer(tunnel *model.Tunnel) *TunnelTcpServer {
	svr := &TunnelTcpServer{
		tunnelBase: tunnelBase{tunnel: tunnel},
	}
	return svr
}

// Open 打开
func (server *TunnelTcpServer) Open() error {
	if server.running {
		return errors.New("server is opened")
	}

	addr, err := net.ResolveTCPAddr("tcp", resolvePort(server.tunnel.Addr))
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	server.running = true

	//server.running = true
	go func() {
		for {
			conn, err := server.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			server.link = conn
			//上线
			server.tunnel.Last = time.Now()
			server.tunnel.Remote = conn.RemoteAddr().String()
			_ = db.Store().Update(server.tunnel.Id, &server.tunnel)
			_ = dbus.Publish(fmt.Sprintf("tunnel/%d/online", server.tunnel.Id), nil)
		}

		server.running = false
	}()

	return nil
}

// Close 关闭
func (server *TunnelTcpServer) Close() error {
	if server.listener != nil {
		err := server.listener.Close()
		if err != nil {
			return err
		}
	}
	return server.tunnelBase.Close()
}
