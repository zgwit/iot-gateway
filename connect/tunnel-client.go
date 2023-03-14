package connect

import (
	"errors"
	"fmt"
	"github.com/iot-master-contrib/gateway/db"
	"github.com/iot-master-contrib/gateway/dbus"
	"github.com/zgwit/iot-master/v2/model"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"net"
	"time"
)

// TunnelClient 网络链接
type TunnelClient struct {
	tunnelBase
	net string
}

func newTunnelClient(tunnel *model.Tunnel, net string) *TunnelClient {
	return &TunnelClient{
		tunnelBase: tunnelBase{tunnel: tunnel},
		net:        net,
	}
}

// Open 打开
func (client *TunnelClient) Open() error {
	if client.running {
		return errors.New("client is opened")
	}

	//发起连接
	conn, err := net.Dial(client.net, client.tunnel.Addr)
	if err != nil {
		client.Retry()
		return err
	}
	client.retry = 0
	client.link = conn

	//上线
	client.tunnel.Last = time.Now()
	client.tunnel.Remote = conn.LocalAddr().String()
	_ = db.Store().Update(client.tunnel.Id, &client.tunnel)
	_ = dbus.Publish(fmt.Sprintf("tunnel/%d/online", client.tunnel.Id), nil)

	return nil
}

func (client *TunnelClient) Retry() {
	//重连
	retry := &client.tunnel.Retry
	if retry.Enable && (retry.Maximum == 0 || client.retry < retry.Maximum) {
		client.retry++
		client.retryTimer = time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			client.retryTimer = nil
			err := client.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}

// Close 关闭
func (client *TunnelClient) Close() error {
	client.running = false

	if client.link != nil {
		link := client.link
		client.link = nil
		return link.Close()
	}
	return errors.New("tunnel is closed")
}
