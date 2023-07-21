package modbus

import (
	"github.com/zgwit/iot-gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/bin"
	"github.com/zgwit/iot-master/v3/pkg/convert"
)

func encode(mapper *types.Mapper, p *types.Point, data any) []byte {
	if mapper.Code == 1 {
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
	} else if mapper.Code == 3 {
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

	} else {
		//TODO error
	}
	return nil
}

func parse(mapper *types.Mapper, data []byte, values map[string]any) {
	l := len(data)

	//识别位
	if mapper.Code == 1 || mapper.Code == 2 {
		bytes := bin.ExpandBool(data, int(mapper.Size))
		l = len(bytes)
		for _, p := range mapper.Points {
			offset := p.Offset
			if offset >= l {
				continue
			}
			values[p.Name] = bytes[p.Offset] > 0
		}
		return
	}

	//解析16位
	for _, p := range mapper.Points {
		//offset := p.Offset * 2
		offset := p.Offset << 1
		if offset >= l {
			continue
		}
		switch p.Type {
		case "bit", "bool", "boolean":
			var v uint16
			if p.BigEndian {
				v = bin.ParseUint16(data[offset:])
			} else {
				v = bin.ParseUint16LittleEndian(data[offset:])
			}
			values[p.Name] = 1<<(p.Bits-1)&v != 0
		case "short", "int16":
			if p.BigEndian {
				values[p.Name] = int16(bin.ParseUint16(data[offset:]))
			} else {
				values[p.Name] = int16(bin.ParseUint16LittleEndian(data[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				values[p.Name] = float64(values[p.Name].(int16)) * p.Rate
			}
		case "word", "uint16":
			if p.BigEndian {
				values[p.Name] = bin.ParseUint16(data[offset:])
			} else {
				values[p.Name] = bin.ParseUint16LittleEndian(data[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				values[p.Name] = float64(values[p.Name].(uint16)) * p.Rate
			}
		case "int32", "int":
			if p.BigEndian {
				values[p.Name] = int32(bin.ParseUint32(data[offset:]))
			} else {
				values[p.Name] = int32(bin.ParseUint32LittleEndian(data[offset:]))
			}
			if p.Rate != 0 && p.Rate != 1 {
				values[p.Name] = float64(values[p.Name].(int32)) * p.Rate
			}
		case "qword", "uint32", "uint":
			if p.BigEndian {
				values[p.Name] = bin.ParseUint32(data[offset:])
			} else {
				values[p.Name] = bin.ParseUint32LittleEndian(data[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				values[p.Name] = float64(values[p.Name].(uint32)) * p.Rate
			}
		case "float", "float32":
			if p.BigEndian {
				values[p.Name] = bin.ParseFloat32(data[offset:])
			} else {
				values[p.Name] = bin.ParseFloat32LittleEndian(data[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				values[p.Name] = float64(values[p.Name].(float32)) * p.Rate
			}
		case "double", "float64":
			if p.BigEndian {
				values[p.Name] = bin.ParseFloat64(data[offset:])
			} else {
				values[p.Name] = bin.ParseFloat64LittleEndian(data[offset:])
			}
			if p.Rate != 0 && p.Rate != 1 {
				values[p.Name] = values[p.Name].(float64) * p.Rate
			}
		}
	}
}

func lookup(mappers []*types.Mapper, name string) (*types.Mapper, *types.Point) {
	for _, mapper := range mappers {
		for _, point := range mapper.Points {
			if point.Name == name {
				return mapper, point
			}
		}
	}
	return nil, nil
}
