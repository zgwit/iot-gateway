package types

import "strconv"

type Options map[string]any

func (o Options) Float64(name string, def float64) float64 {
	if value, ok := o[name]; ok {
		switch value.(type) {
		case bool:
			if value.(bool) {
				def = 1
			} else {
				def = 0
			}
		case uint8:
			def = float64(value.(uint8))
		case uint16:
			def = float64(value.(uint16))
		case uint32:
			def = float64(value.(uint32))
		case uint64:
			def = float64(value.(uint64))
		case uint:
			def = float64(value.(uint))
		case int8:
			def = float64(value.(int8))
		case int16:
			def = float64(value.(int16))
		case int32:
			def = float64(value.(int32))
		case int64:
			def = float64(value.(int64))
		case int:
			def = float64(value.(int))
		case float32:
			def = float64(value.(float32))
		case float64:
			def = value.(float64)
		case string:
			v, e := strconv.ParseFloat(value.(string), 64)
			if e == nil {
				def = v
			}
		}
	}
	return def
}

func (o Options) Int64(name string, def int64) int64 {
	if value, ok := o[name]; ok {
		switch value.(type) {
		case bool:
			if value.(bool) {
				def = 1
			} else {
				def = 0
			}
		case uint8:
			def = int64(value.(uint8))
		case uint16:
			def = int64(value.(uint16))
		case uint32:
			def = int64(value.(uint32))
		case uint64:
			def = int64(value.(uint64))
		case uint:
			def = int64(value.(uint))
		case int8:
			def = int64(value.(int8))
		case int16:
			def = int64(value.(int16))
		case int32:
			def = int64(value.(int32))
		case int64:
			def = value.(int64)
		case int:
			def = int64(value.(int))
		case float32:
			def = int64(value.(float32))
		case float64:
			def = int64(value.(float64))
		case string:
			v, e := strconv.ParseInt(value.(string), 10, 64)
			if e == nil {
				def = v
			}
		}
	}
	return def
}

func (o Options) Int(name string, def int) int {
	return int(o.Int64(name, int64(def)))
}

func (o Options) Bool(name string, def bool) bool {
	if value, ok := o[name]; ok {
		switch value.(type) {
		case bool:
			def = value.(bool)
		case uint8:
			def = value.(uint8) != 0
		case uint16:
			def = value.(uint16) != 0
		case uint32:
			def = value.(uint32) != 0
		case uint64:
			def = value.(uint64) != 0
		case uint:
			def = value.(uint) != 0
		case int8:
			def = value.(int8) != 0
		case int16:
			def = value.(int16) != 0
		case int32:
			def = value.(int32) != 0
		case int64:
			def = value.(int64) != 0
		case int:
			def = value.(int) != 0
		case float32:
			def = value.(float32) != 0
		case float64:
			def = value.(float64) != 0
		case string:
			switch value.(string) {
			case "true", "TRUE", "True", "1":
				def = true
			default:
				def = false
			}
		}
	}
	return def
}
