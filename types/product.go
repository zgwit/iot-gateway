package types

import (
	"github.com/zgwit/iot-master/v3/pkg/bin"
	"github.com/zgwit/iot-master/v3/pkg/convert"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"time"
)

type Product struct {
	Id   string `json:"id" xorm:"pk"`
	Name string `json:"name,omitempty"` //名称
	Desc string `json:"desc,omitempty"` //说明
	//model.Product `xorm:"extends"`
	Mappers     []Mapper     `json:"mappers" xorm:"json"`
	Filters     []Filter     `json:"filters" xorm:"json"`
	Calculators []Calculator `json:"calculators" xorm:"json"`
	Created     time.Time    `json:"created" xorm:"created"` //创建时间
}

type Filter struct {
	Name       string `json:"name"`       //字段
	Expression string `json:"expression"` //表达式
	//Entire     bool   `json:"entire"`
}

type Calculator struct {
	Name       string `json:"name"`       //赋值
	Expression string `json:"expression"` //表达式
}

type Mapper struct {
	Code   uint8   `json:"code"`   //指令
	Addr   uint16  `json:"addr"`   //地址
	Size   uint16  `json:"size"`   //长度
	Points []Point `json:"points"` //数据点
}

type Point struct {
	Name      string  `json:"name"`           //名称
	Type      string  `json:"type"`           //类型
	Offset    uint16  `json:"offset"`         //偏移
	Bits      uint16  `json:"bits,omitempty"` //位，1 2 3...
	BigEndian bool    `json:"be,omitempty"`   //大端模式
	Rate      float64 `json:"rate,omitempty"` //倍率
}

func (m *Mapper) Encode(name string, data any) []byte {
	//TODO 支持 布尔，数值，数组
	for _, p := range m.Points {
		if p.Name == name {
			continue
		}

		switch m.Code {
		case 1:
			//convert.ToBool(data) 太范了
			val, ok := data.(bool)
			if !ok {
				//TODO error
				return nil
			}
			if val {
				return []byte{0xFF, 00}
			} else {
				return []byte{0x00, 00}
			}
		case 3:
			var ret []byte

			//倍率逆转换
			if p.Rate != 0 && p.Rate != 1 {
				val, ok := data.(float64)
				if ok {
					data = val / p.Rate
				} else {
					//error ?
				}
			}

			switch p.Type {
			case "short", "int16":
				ret = make([]byte, 2)
				val := convert.ToInt16(data)
				if p.BigEndian {
					bin.WriteUint16(ret, uint16(val))
				} else {
					bin.WriteUint16LittleEndian(ret, uint16(val))
				}
			case "word", "uint16":
				ret = make([]byte, 2)
				val := convert.ToUint16(data)
				if p.BigEndian {
					bin.WriteUint16(ret, val)
				} else {
					bin.WriteUint16LittleEndian(ret, val)
				}
			case "int32", "int":
				ret = make([]byte, 4)
				val := convert.ToInt32(data)
				if p.BigEndian {
					bin.WriteUint32(ret, uint32(val))
				} else {
					bin.WriteUint32LittleEndian(ret, uint32(val))
				}
			case "qword", "uint32", "uint":
				ret = make([]byte, 4)
				val := convert.ToUint32(data)
				if p.BigEndian {
					bin.WriteUint32(ret, val)
				} else {
					bin.WriteUint32LittleEndian(ret, val)
				}
			case "float", "float32":
				ret = make([]byte, 4)
				val := convert.ToFloat32(data)
				if p.BigEndian {
					bin.WriteFloat32(ret, val)
				} else {
					bin.WriteFloat32LittleEndian(ret, val)
				}
			case "double", "float64":
				ret = make([]byte, 8)
				val := convert.ToFloat64(data)
				if p.BigEndian {
					bin.WriteFloat64(ret, val)
				} else {
					bin.WriteFloat64LittleEndian(ret, val)
				}
			}

			return ret
		}
	}

	return nil
}

func (m *Mapper) Parse(buf []byte, ret map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	l := uint16(len(buf))

	//识别位
	if m.Code == 1 || m.Code == 2 {
		bytes := bin.ExpandBool(buf, int(m.Size))
		l = uint16(len(bytes))
		for _, p := range m.Points {
			offset := p.Offset
			if offset >= l {
				continue
			}
			ret[p.Name] = bytes[p.Offset] > 0
		}
		return
	}

	//解析16位
	for _, p := range m.Points {
		//offset := p.Offset * 2
		offset := p.Offset << 1
		if offset >= l {
			continue
		}
		switch p.Type {
		case "bit", "bool", "boolean":
			var v uint16
			if p.BigEndian {
				v = bin.ParseUint16(buf[offset:])
			} else {
				v = bin.ParseUint16LittleEndian(buf[offset:])
			}
			ret[p.Name] = 1<<(p.Bits-1)&v != 0
		case "short", "int16":
			if p.BigEndian {
				ret[p.Name] = int16(bin.ParseUint16(buf[offset:]))
			} else {
				ret[p.Name] = int16(bin.ParseUint16LittleEndian(buf[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int16)) * p.Rate
			}
		case "word", "uint16":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint16(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseUint16LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint16)) * p.Rate
			}
		case "int32", "int":
			if p.BigEndian {
				ret[p.Name] = int32(bin.ParseUint32(buf[offset:]))
			} else {
				ret[p.Name] = int32(bin.ParseUint32LittleEndian(buf[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(int32)) * p.Rate
			}
		case "qword", "uint32", "uint":
			if p.BigEndian {
				ret[p.Name] = bin.ParseUint32(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseUint32LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(uint32)) * p.Rate
			}
		case "float", "float32":
			if p.BigEndian {
				ret[p.Name] = bin.ParseFloat32(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseFloat32LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = float64(ret[p.Name].(float32)) * p.Rate
			}
		case "double", "float64":
			if p.BigEndian {
				ret[p.Name] = bin.ParseFloat64(buf[offset:])
			} else {
				ret[p.Name] = bin.ParseFloat64LittleEndian(buf[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				ret[p.Name] = ret[p.Name].(float64) * p.Rate
			}
		}
	}
}
