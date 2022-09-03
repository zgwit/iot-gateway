package connect

import (
	"errors"
	"iot-master-gateway/db"
	"iot-master-gateway/model"
	"net"
	"time"
)

//TunnelTcpServer TCP服务器
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

//Open 打开
func (server *TunnelTcpServer) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

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
			_ = mqtt.Publish(fmt.Sprintf("tunnel/%d/online", server.tunnel.Id), nil)

			server.receive()
		}

		server.running = false
	}()

	return nil
}

func (server *TunnelTcpServer) receive() {
	server.online = true
	server.Emit("online")

	buf := make([]byte, 1024)
	for {
		n, err := server.link.Read(buf)
		if err != nil {
			server.onClose()
			break
		}
		if n == 0 {
			continue
		}

		data := buf[:n]
		//过滤心跳包
		if server.tunnel.Heartbeat.Enable && server.tunnel.Heartbeat.Check(data) {
			continue
		}

		//透传转发
		if server.pipe != nil {
			_, err = server.pipe.Write(data)
			if err != nil {
				server.pipe = nil
			} else {
				continue
			}
		}
		server.Emit("data", data)
	}
	server.online = false
	server.Emit("offline")
}

//Close 关闭
func (server *TunnelTcpServer) Close() error {
	if server.listener != nil {
		err := server.listener.Close()
		if err != nil {
			return err
		}
	}
	return server.tunnelBase.Close()
}
