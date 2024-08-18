package rpc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"net/url"
)

const HEADER_SIZE = 8
const CLIENT_BUFFER_SIZE = 1024

type Client struct {
	Url string

	Encoding uint8 //JSON

	id       uint16
	conn     net.Conn
	requests map[uint16]any

	inHeader  [HEADER_SIZE]byte
	outHeader [HEADER_SIZE]byte
	//buffer    [CLIENT_BUFFER_SIZE]byte
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
	//buf := make([]byte, 8)
	buf := make([]byte, CLIENT_BUFFER_SIZE)
	for {
		//n, err := c.conn.Read(buf)
		n, err := io.ReadAtLeast(c.conn, c.inHeader[:], HEADER_SIZE)
		if err != nil {
			break
		}
		if n < HEADER_SIZE {
			break
		}

		if bytes.Compare(c.inHeader[:3], []byte(MAGIC)) != 0 {
			break
		}

		code := buf[3] >> 4
		encoding := buf[3] & 0xf

		l := int(binary.BigEndian.Uint16(c.inHeader[6:])) //长度
		if l > 0 {
			var b []byte
			if l > CLIENT_BUFFER_SIZE {
				b = make([]byte, l)
			} else {
				b = buf
			}

			//_ = c.conn.SetReadDeadline(time.Now().Add(time.Second * 30))
			n, err = io.ReadAtLeast(c.conn, b, l)
			if err != nil {
				break
			}
			if n != l {
				//长度不够，废包
				break
			}

			parse

		}

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
