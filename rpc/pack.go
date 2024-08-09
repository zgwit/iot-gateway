package rpc

import (
	"encoding/binary"
	"encoding/xml"
	"errors"
	"github.com/gocarina/gocsv"
	"github.com/goccy/go-json"
	"github.com/shamaton/msgpack/v2"
	"gopkg.in/yaml.v3"
)

var ErrEncoding = errors.New("编码不支持")
var ErrNotEnough = errors.New("长度不足")

type Encoder func(any) ([]byte, error)
type Decoder func([]byte, any) error

const MAGIC = "rpc"

const (
	DISCONNECT uint8 = iota
	CONNECT
	CONNECT_ACK
	HEARTBEAT
	REQUEST
	REQUEST_END
	RESPONSE
	RESPONSE_END
	STREAM
	STREAM_END
	PUBLISH
	PUBLISH_END
	PUBLISH_ACK
	SUBSCRIBE
	SUBSCRIBE_ACK
	UNSUBSCRIBE
)

const (
	BINARY uint8 = iota
	JSON
	XML
	YAML
	CSV
	MSGPACK
	PROTOBUF
)

var encoders = map[uint8]Encoder{
	JSON:    json.Marshal,
	XML:     xml.Marshal,
	YAML:    yaml.Marshal,
	CSV:     gocsv.MarshalBytes,
	MSGPACK: msgpack.Marshal,
}

var decoders = map[uint8]Decoder{
	JSON:    json.Unmarshal,
	XML:     xml.Unmarshal,
	YAML:    yaml.Unmarshal,
	CSV:     gocsv.UnmarshalBytes,
	MSGPACK: msgpack.Unmarshal,
}

func RegisterEncoding(typ uint8, encoder Encoder, decoder Decoder) {
	encoders[typ] = encoder
	decoders[typ] = decoder
}

type Pack struct {
	Type     uint8
	Encoding uint8
	Id       uint16
	Length   uint16
	Data     []byte
	Content  any
}

func (p *Pack) Encode() (buf []byte, err error) {
	if p.Content != nil && p.Encoding > 0 {
		if encoder, ok := encoders[p.Encoding]; ok {
			p.Data, err = encoder(p.Content)
		} else {
			err = ErrEncoding
		}

		if err != nil {
			return
		}
	}

	//构建包
	p.Length = uint16(len(p.Data))
	buf = make([]byte, 8+len(p.Data))
	copy(buf, MAGIC)
	buf[3] = p.Type<<4 + p.Encoding&0xF0
	binary.BigEndian.PutUint16(buf[4:], p.Id)
	binary.BigEndian.PutUint16(buf[6:], p.Length)
	if p.Data != nil {
		copy(buf[8:], p.Data) //内存复制了
	}
	return
}

func (p *Pack) Decode(buf []byte) (err error) {
	p.Id = binary.BigEndian.Uint16(buf[4:])
	p.Length = binary.BigEndian.Uint16(buf[6:])
	p.Type = buf[3] >> 4
	p.Encoding = buf[3] & 0xF0

	if p.Length > 0 {
		if len(buf) < 8+int(p.Length) {
			return ErrNotEnough
		}
		p.Data = buf[8 : 8+p.Length]

		if p.Encoding > 0 {
			if decoder, ok := decoders[p.Encoding]; ok {
				err = decoder(p.Data, p.Content)
			} else {
				err = ErrEncoding
			}
		} else {
			p.Content = p.Data
		}
		if err != nil {
			return
		}
	}
	return
}
