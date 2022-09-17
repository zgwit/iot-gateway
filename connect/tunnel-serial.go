package connect

import (
	"errors"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/zgwit/iot-master/v2/model"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"iot-master-gateway/db"
	"iot-master-gateway/dbus"
	"time"
)

// TunnelSerial 串口
type TunnelSerial struct {
	tunnelBase
}

func newTunnelSerial(tunnel *model.Tunnel) *TunnelSerial {
	return &TunnelSerial{
		tunnelBase: tunnelBase{tunnel: tunnel},
	}
}

// Open 打开
func (s *TunnelSerial) Open() error {
	if s.running {
		return errors.New("serial is opened")
	}

	mode := serial.OpenOptions{
		PortName:              s.tunnel.Serial.Port,
		BaudRate:              s.tunnel.Serial.BaudRate,
		DataBits:              s.tunnel.Serial.DataBits,
		StopBits:              s.tunnel.Serial.StopBits,
		ParityMode:            serial.ParityMode(s.tunnel.Serial.Parity),
		MinimumReadSize:       4,   //避免单字节读出
		InterCharacterTimeout: 100, //100ms分包
	}

	port, err := serial.Open(mode)
	if err != nil {
		//TODO 串口重试
		s.Retry()
		return err
	}
	s.running = true
	s.online = true
	s.link = port
	s.retry = 0

	//清空重连计数
	s.retry = 0

	go s.receive()

	//上线
	s.tunnel.Last = time.Now()
	_ = db.Store().Update(s.tunnel.Id, &s.tunnel)
	_ = dbus.Publish(fmt.Sprintf("tunnel/%d/online", s.tunnel.Id), nil)

	return nil
}
func (s *TunnelSerial) Retry() {
	retry := &s.tunnel.Retry
	if retry.Enable && (retry.Maximum == 0 || s.retry < retry.Maximum) {
		s.retry++
		s.retryTimer = time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			s.retryTimer = nil
			err := s.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}

func (s *TunnelSerial) receive() {
	s.running = true
	s.online = true

	buf := make([]byte, 1024)
	for {
		n, err := s.link.Read(buf)
		if err != nil {
			s.onClose()
			break
		}
		if n == 0 {
			continue
		}
		//透传转发
		if s.pipe != nil {
			_, err = s.pipe.Write(buf[:n])
			if err != nil {
				s.pipe = nil
			} else {
				continue
			}
		}
	}
	s.running = false
	s.online = false

	s.Retry()
}
