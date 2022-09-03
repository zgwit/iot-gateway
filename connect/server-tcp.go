package connect

import (
	"errors"
	"iot-master-gateway/db"
	"iot-master-gateway/model"
	"iot-master-gateway/pkg/events"
	"net"
	"time"
)

// ServerTCP TCP服务器
type ServerTCP struct {
	events.EventEmitter

	server *model.Server

	children map[uint64]*ServerTcpTunnel

	listener *net.TCPListener

	running bool
}

func newServerTCP(server *model.Server) *ServerTCP {
	svr := &ServerTCP{
		server:   server,
		children: make(map[uint64]*ServerTcpTunnel),
	}
	return svr
}

// Open 打开
func (server *ServerTCP) Open() error {
	if server.running {
		return errors.New("server is opened")
	}
	server.Emit("open")

	addr, err := net.ResolveTCPAddr("tcp", resolvePort(server.server.Addr))
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	server.running = true
	go func() {
		for {
			c, err := server.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				break
			}

			buf := make([]byte, 128)
			n := 0
			n, err = c.Read(buf)
			if err != nil {
				_ = c.Close()
				continue
			}
			data := buf[:n]
			if !server.server.Register.Check(data) {
				_ = c.Close()
				continue
			}

			sn := string(data)
			tunnel := model.Tunnel{
				ServerId: server.server.Id,
				Addr:     sn,
			}

			err = db.Store().FindOne(&tunnel, bolthold.Where("ServerId").Eq(server.server.Id).And("SN").Eq(sn))
			has := err == bolthold.ErrNotFound
			//has, err := db.Engine.Where("server_id=?", server.server.Id).And("addr", sn).Get(&tunnel)
			if err != nil {
				_ = mqtt.Publish(fmt.Sprintf("server/%d/error", server.server.Id), []byte(err.Error()))
				continue
			}

			tunnel.Last = time.Now()
			tunnel.Remote = c.RemoteAddr().String()
			if !has {
				//保存一条新记录
				tunnel.Type = "server-tcp"
				tunnel.Name = sn
				tunnel.SN = sn
				tunnel.Addr = server.server.Addr
				tunnel.Heartbeat = server.server.Heartbeat
				tunnel.Protocol = server.server.Protocol
				//_, _ = db.Engine.InsertOne(&tunnel)
				tunnel.Created = time.Now()
				_ = db.Store().Insert(bolthold.NextSequence(), &tunnel)
			} else {
				//上线
				//_, _ = db.Engine.ID(tunnel.Id).Cols("last", "remote").Update(tunnel)
				_ = db.Store().Update(tunnel.Id, &tunnel)
			}
			_ = mqtt.Publish(fmt.Sprintf("tunnel/%d/online", tunnel.Id), nil)

			tnl := newServerTcpTunnel(&tunnel, c)
			tnl.first = !has
			go tnl.receive()
			server.children[tunnel.Id] = tnl

			//启动对应的设备 发消息
			server.Emit("tunnel", tnl)

			tnl.Once("close", func() {
				delete(server.children, tunnel.Id)
			})
		}

		server.running = false
	}()

	return nil
}

//Close 关闭
func (server *ServerTCP) Close() (err error) {
	server.Emit("close")
	//close tunnels
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

//GetTunnel 获取连接
func (server *ServerTCP) GetTunnel(id uint64) link.Tunnel {
	return server.children[id]
}

func (server *ServerTCP) Running() bool {
	return server.running
}
