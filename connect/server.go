package connect

import (
	"errors"
	"fmt"
	"github.com/iot-master-contrib/modbus/define"
	"github.com/iot-master-contrib/modbus/types"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"net"
)

// Server TCP服务器
type Server struct {
	model *types.Server

	children map[string]*Link

	listener *net.TCPListener

	running bool
}

func NewServer(model *types.Server) *Server {
	s := &Server{
		model:    model,
		children: make(map[string]*Link),
	}
	return s
}

// Open 打开
func (s *Server) Open() error {
	if s.running {
		return errors.New("s is opened")
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.model.Port))
	if err != nil {
		return err
	}
	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	//defer s.listener.Close()

	s.running = true
	go func() {
		for {
			c, err := s.listener.AcceptTCP()
			if err != nil {
				//TODO 需要正确处理接收错误
				log.Error(err)
				break
			}

			//单例模式，关闭之前的连接
			if s.model.Standalone {
				const k = "internal"
				if cc, ok := s.children[k]; ok {
					_ = cc.Close()
				}

				lnk := newLink(&types.Link{
					Tunnel:   s.model.Tunnel,
					ServerId: s.model.Id,
					Remote:   c.RemoteAddr().String(),
				}, c)
				s.children[k] = lnk

				//启动轮询
				err = lnk.start(&lnk.model.Tunnel)
				if err != nil {
					log.Error(err)
					continue
					//return
				}

				//以ServerID保存
				links.Store(s.model.Id, lnk)
				continue
			}

			buf := make([]byte, 128)
			n, err := c.Read(buf)
			if err != nil {
				_ = c.Close()
				continue
			}
			data := buf[:n]
			sn := string(data)

			var link types.Link
			//get, err := db.Engine.Where("server_id=?", s.model.Id).And("sn=?", sn).Get(&conn)
			get, err := db.Engine.ID(sn).Get(&link)
			if err != nil {
				_, _ = c.Write([]byte(err.Error()))
				_ = c.Close()
				continue
			}
			if !get {
				link = types.Link{
					Tunnel:   s.model.Tunnel,
					ServerId: s.model.Id,
					Remote:   c.RemoteAddr().String(),
				}
				link.Id = sn
				_, err := db.Engine.InsertOne(&link)
				if err != nil {
					_, _ = c.Write([]byte(err.Error()))
					_ = c.Close()
					continue
				}
			}

			lnk := newLink(&link, c)
			s.children[sn] = lnk

			//启动轮询
			err = lnk.start(&link.Tunnel)
			if err != nil {
				log.Error(err)
				continue
			}

			links.Store(link.Id, lnk)
		}

		s.running = false
	}()

	return nil
}

// Close 关闭
func (s *Server) Close() (err error) {
	//close tunnels
	if s.children != nil {
		for _, l := range s.children {
			_ = l.Close()
		}
	}
	return s.listener.Close()
}

// GetTunnel 获取连接
func (s *Server) GetTunnel(id string) define.Tunnel {
	return s.children[id]
}

func (s *Server) Running() bool {
	return s.running
}
