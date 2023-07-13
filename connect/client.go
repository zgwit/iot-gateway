package connect

import (
	"errors"
	"fmt"
	"github.com/iot-master-contrib/modbus/types"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"net"
	"time"
)

// Client 网络链接
type Client struct {
	tunnelBase
	model *types.Client
}

func NewClient(client *types.Client) *Client {
	return &Client{
		model: client,
	}
}

// Open 打开
func (client *Client) Open() error {
	if client.running {
		return errors.New("client is opened")
	}
	client.closed = false

	//发起连接
	addr := fmt.Sprintf("%s:%d", client.model.Addr, client.model.Port)
	conn, err := net.Dial(client.model.Net, addr)
	if err != nil {
		client.Retry()
		//time.AfterFunc(time.Minute, client.Retry)
		return err
	}
	client.retry = 0
	client.Conn = &netConn{conn}

	//守护协程
	go func() {
		timeout := client.model.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		for {
			time.Sleep(time.Second * time.Duration(timeout))
			if client.running {
				continue
			}
			if client.closed {
				return
			}
			//如果掉线了，就重新打开
			err := client.Open()
			if err != nil {
				log.Error(err)
			}
			break //Open中，会重新启动协程
		}
	}()

	//启动轮询
	return client.start(&client.model.Tunnel)
}

func (client *Client) Retry() {
	//重连
	retry := &client.model.Retry
	if retry.RetryMaximum == 0 || client.retry < retry.RetryMaximum {
		client.retry++
		timeout := retry.RetryTimeout
		if timeout == 0 {
			timeout = 10
		}
		client.retryTimer = time.AfterFunc(time.Second*time.Duration(timeout), func() {
			client.retryTimer = nil
			err := client.Open()
			if err != nil {
				log.Error(err)
			}
		})
	}
}

// Close 关闭
func (client *Client) Close() error {
	client.running = false

	if client.Conn != nil {
		link := client.Conn
		client.Conn = nil
		return link.Close()
	}
	return errors.New("model is closed")
}
