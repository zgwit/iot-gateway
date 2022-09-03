package connect

import (
	"errors"
	"iot-master-gateway/db"
	"iot-master-gateway/model"
	"iot-master-gateway/pkg/events"
	"net"
	"time"
)

//ServerUDP UDP服务器
type ServerUDP struct {
	events.EventEmitter

	server *model.Server

	children map[int64]*ServerUdpTunnel
	tunnels  map[string]*ServerUdpTunnel

	listener *net.UDPConn
	running  bool
}

func newServerUDP(server *model.Server) *ServerUDP {
	svr := &ServerUDP{
		server:   server,
		children: make(map[int64]*ServerUdpTunnel),
		tunnels:  make(map[string]*ServerUdpTunnel),
	}
	return svr
}

//Open 打开
func (server *ServerUDP) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

	addr, err := net.ResolveUDPAddr("udp", resolvePort(server.server.Addr))
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		//TODO 需要正确处理接收错误
		return err
	}
	server.listener = conn //共用连接

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
			data := buf[:n]

			//如果已经保存了链接 TODO 要有超时处理
			tnl, ok := server.tunnels[addr.String()]
			if ok {
				tnl.onData(data)
				continue
			}

			if !server.server.Register.Check(data) {
				_ = conn.Close()
				continue
			}
			sn := string(data)
			tunnel := model.Tunnel{
				ServerId: server.server.Id,
				Addr:     sn,
			}
			has, err := db.Engine.Where("server_id=?", server.server.Id).And("addr", sn).Get(&tunnel)
			if err != nil {
				//return err
				//TODO 日志，关闭连接
				continue
			}

			tunnel.Last = time.Now()
			tunnel.Remote = conn.RemoteAddr().String()
			if !has {
				//保存一条新记录
				tunnel.Type = "server-udp"
				tunnel.Name = sn
				tunnel.Heartbeat = server.server.Heartbeat
				tunnel.Protocol = server.server.Protocol
				_, _ = db.Engine.InsertOne(&tunnel)
			} else {
				//上线
				_, _ = db.Engine.ID(tunnel.Id).Cols("last", "remote").Update(tunnel)
			}

			tnl = newServerUdpTunnel(&tunnel, conn, addr)
			tnl.first = !has
			server.children[tunnel.Id] = tnl

			//启动对应的设备 发消息
			server.Emit("tunnel", tnl)

			tnl.Emit("online")

			tnl.Once("close", func() {
				delete(server.children, tunnel.Id)
				delete(server.tunnels, tnl.addr.String())
			})
		}

		server.running = false
	}()

	return nil
}

//Close 关闭
func (server *ServerUDP) Close() (err error) {
	server.Emit("close")
	//close tunnels
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

//GetTunnel 获取链接
func (server *ServerUDP) GetTunnel(id int64) Tunnel {
	return server.children[id]
}

func (server *ServerUDP) Running() bool {
	return server.running
}
