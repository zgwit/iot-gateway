package rpc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"net/url"
)

type Client struct {
	Url string

	Encoding uint8 //JSON

	id       uint16
	conn     net.Conn
	requests map[uint16]any
}

func (c *Client) Write(pack *Pack) error {
	c.id++

	c.conn.Write(pack.Header)
}

func (c *Client) Write2(typ uint8, data any) error {
	header := make([]byte, 8)

}

func (c *Client) Request(request any) (response any, err error) {
	header := make([]byte, 8)

}

func (c *Client) receive() {
	buf := make([]byte, 8)
	for {
		//_ = c.conn.SetReadDeadline(time.Time{})
		//n, err := c.conn.Read(buf)
		n, err := io.ReadAtLeast(c.conn, buf, 8)
		if err != nil {
			break
		}
		if n < 8 {
			break
		}

		if bytes.Compare(buf[:3], []byte(MAGIC)) != 0 {
			break
		}

		l := int(binary.BigEndian.Uint16(buf[6:])) //长度
		if l > 0 {
			b := make([]byte, l)
			//_ = c.conn.SetReadDeadline(time.Now().Add(time.Second * 30))
			n, err = io.ReadAtLeast(c.conn, b, l)
			if err != nil {
				break
			}
			if n != l {
				//长度不够，费包
				break
			}
		}

		tp := buf[3] >> 4
		enc := buf[3] & 0xf

		var pack Pack
		pack.Decode(b)

	}

	//关闭连接
	_ = c.conn.Close()

}

func (c *Client) Open() error {
	//rpc://username:password@127.0.0.1:719
	u, err := url.Parse(c.Url)
	if err != nil {
		return err
	}

	if u.Scheme != "rpc" {
		return errors.New("协议错误")
	}

	if u.User == nil {
		return errors.New("缺少用户名和密码")
	}

	c.conn, err = net.Dial("tcp", u.Hostname())
	if err != nil {
		return err
	}

	pack := Pack{
		Type:     CONNECT,
		Encoding: JSON,
		Content: map[string]any{
			"username": u.User.Username(),
			"password": u.User.Password(),
		},
	}

	err = c.Write(&pack)
	if err != nil {
		return err
	}

	//持续接收消息
	go c.receive()

	return nil
}
