package modbus

import (
	"errors"
	"fmt"
	"github.com/iot-master-contrib/gateway/connect"
	"github.com/iot-master-contrib/gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/bin"
	"time"
)

// RTU Modbus-RTU协议
type RTU struct {
	messenger connect.Messenger
	buf       []byte
}

func NewRTU(tunnel connect.Conn, opts types.Options) *RTU {
	rtu := &RTU{
		messenger: connect.Messenger{
			Timeout: time.Millisecond * time.Duration(opts.Int64("timeout", 1000)),
			Tunnel:  tunnel,
		},
		//slave: opts["slave"].(uint8),
		buf: make([]byte, opts.Int("buffer", 256)),
	}

	return rtu
}

func (m *RTU) execute(cmd []byte) ([]byte, error) {

	l, err := m.messenger.AskAtLeast(cmd, m.buf, 5)
	if err != nil {
		return nil, err
	}

	//crc := bin.ParseUint16LittleEndian(m.buf[l-2:])
	//if crc != CRC16(m.buf[:l-2]) {
	//	//检验错误
	//	return nil, errors.New("校验错误")
	//}

	//slave := buf[0]
	fc := m.buf[1]

	//解析错误码
	if fc&0x80 > 0 {
		return nil, fmt.Errorf("错误码：%d", m.buf[2])
	}

	//解析数据
	length := 4
	count := int(m.buf[2])
	switch fc {
	case 1, 2:
		length += 1 + count/8
		if count%8 != 0 {
			length++
		}

		if l < length {
			//长度不够，继续读
			_, err = m.messenger.ReadAtLeast(m.buf[l:], length-l)
			if err != nil {
				return nil, err
			}
			l = len(m.buf)
		}
		b := m.buf[3 : l-2]
		//数组解压
		//b = bin.ExpandBool(b, count)
		return bin.Dup(b), nil
	case 3, 4, 23:
		length += 1 + count
		if l < length {
			//长度不够，继续读
			_, err = m.messenger.ReadAtLeast(m.buf[l:], length-l)
			if err != nil {
				return nil, err
			}
			l = len(m.buf)
			//if n+l < length {
			//	return nil, errors.New("长度不足")
			//}
		}
		b := m.buf[3 : l-2]
		return bin.Dup(b), nil
	case 5, 15, 6, 16:
		//写指令不处理
		return nil, nil
	default:
		return nil, errors.New("不支持的指令")
	}
}

func (m *RTU) Read(station types.Station, addr *types.Address, size int) ([]byte, error) {
	b := make([]byte, 8)
	b[0] = uint8(station["slave"])
	b[1] = uint8(addr.Code)
	bin.WriteUint16(b[2:], uint16(addr.Addr))
	bin.WriteUint16(b[4:], uint16(size))
	bin.WriteUint16LittleEndian(b[6:], CRC16(b[:6]))

	return m.execute(b)
}

func (m *RTU) Write(station types.Station, addr *types.Address, buf []byte) error {
	length := len(buf)
	var code uint8
	switch addr.Code {
	case 1:
		//如果是线圈，需要Shrink
		if length == 1 {
			code = 5
			//数据 转成 0x0000 0xFF00
			if buf[0] > 0 {
				buf = []byte{0xFF, 0}
			} else {
				buf = []byte{0, 0}
			}
		} else {
			code = 15 //0x0F
			//数组压缩
			b := bin.ShrinkBool(buf)
			count := len(b)
			buf = make([]byte, 3+count)
			bin.WriteUint16(buf, uint16(length))
			buf[2] = uint8(count)
			copy(buf[3:], b)
		}
	case 3:
		if length == 2 {
			code = 6
		} else {
			code = 16 //0x10
			b := make([]byte, 3+length)
			bin.WriteUint16(b, uint16(length/2))
			b[2] = uint8(length)
			copy(b[3:], buf)
			buf = b
		}
	default:
		return errors.New("功能码不支持")
	}

	l := 6 + len(buf)
	b := make([]byte, l)
	b[0] = uint8(station["slave"])
	b[1] = code
	bin.WriteUint16(b[2:], uint16(addr.Addr))
	copy(b[4:], buf)
	bin.WriteUint16LittleEndian(b[l-2:], CRC16(b[:l-2]))

	_, err := m.execute(b)
	return err
}

func (m *RTU) Poll(station types.Station, mappers []*types.Mapper) (map[string]any, error) {
	values := make(map[string]any)
	for _, mapper := range mappers {
		data, err := m.Read(station, &mapper.Address, mapper.Size)
		if err != nil {
			return nil, err
		}
		parse(mapper, data, values)
	}
	return values, nil
}

func (m *RTU) Set(station types.Station, mappers []*types.Mapper, name string, value any) error {
	mapper, point := lookup(mappers, name)
	if mapper == nil {
		return errors.New("地址找不到")
	}
	data := encode(mapper, point, value)
	addr := mapper.Address
	addr.Addr += point.Offset //找到指定位置
	return m.Write(station, &mapper.Address, data)
}

func (m *RTU) Get(station types.Station, mappers []*types.Mapper, name string) (any, error) {
	mapper, _ := lookup(mappers, name)
	if mapper == nil {
		return nil, errors.New("地址找不到")
	}
	//此处全部读取了，有些冗余
	data, err := m.Read(station, &mapper.Address, mapper.Size)
	if err != nil {
		return nil, err
	}
	values := make(map[string]any)
	parse(mapper, data, values)
	return values[name], nil
}
