package cast

func ToBool(value interface{}) bool {
	switch value.(type) {
	case bool:
		return value.(bool)
	case uint8:
		return value.(uint8) > 0
	case uint16:
		return value.(uint16) > 0
	case uint32:
		return value.(uint32) > 0
	case uint64:
		return value.(uint64) > 0
	case uint:
		return value.(uint) > 0
	case int8:
		return value.(int8) > 0
	case int16:
		return value.(int16) > 0
	case int32:
		return value.(int32) > 0
	case int64:
		return value.(int64) > 0
	case int:
		return value.(int) > 0
	case float32:
		return value.(float32) > 0
	case float64:
		return value.(float64) > 0
	}
	return false
}

func ToUint8(value interface{}) uint8 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint8(value.(uint8))
	case uint16:
		return uint8(value.(uint16))
	case uint32:
		return uint8(value.(uint32))
	case uint64:
		return uint8(value.(uint64))
	case uint:
		return uint8(value.(uint))
	case int8:
		return uint8(value.(int8))
	case int16:
		return uint8(value.(int16))
	case int32:
		return uint8(value.(int32))
	case int64:
		return uint8(value.(int64))
	case int:
		return uint8(value.(int))
	case float32:
		return uint8(value.(float32))
	case float64:
		return uint8(value.(float64))
	}
	return 0
}

func ToUint16(value interface{}) uint16 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint16(value.(uint8))
	case uint16:
		return uint16(value.(uint16))
	case uint32:
		return uint16(value.(uint32))
	case uint64:
		return uint16(value.(uint64))
	case uint:
		return uint16(value.(uint))
	case int8:
		return uint16(value.(int8))
	case int16:
		return uint16(value.(int16))
	case int32:
		return uint16(value.(int32))
	case int64:
		return uint16(value.(int64))
	case int:
		return uint16(value.(int))
	case float32:
		return uint16(value.(float32))
	case float64:
		return uint16(value.(float64))
	}
	return 0
}

func ToUint32(value interface{}) uint32 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint32(value.(uint8))
	case uint16:
		return uint32(value.(uint16))
	case uint32:
		return uint32(value.(uint32))
	case uint64:
		return uint32(value.(uint64))
	case uint:
		return uint32(value.(uint))
	case int8:
		return uint32(value.(int8))
	case int16:
		return uint32(value.(int16))
	case int32:
		return uint32(value.(int32))
	case int64:
		return uint32(value.(int64))
	case int:
		return uint32(value.(int))
	case float32:
		return uint32(value.(float32))
	case float64:
		return uint32(value.(float64))
	}
	return 0
}

func ToUint64(value interface{}) uint64 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return uint64(value.(uint8))
	case uint16:
		return uint64(value.(uint16))
	case uint32:
		return uint64(value.(uint32))
	case uint64:
		return uint64(value.(uint64))
	case uint:
		return uint64(value.(uint))
	case int8:
		return uint64(value.(int8))
	case int16:
		return uint64(value.(int16))
	case int32:
		return uint64(value.(int32))
	case int64:
		return uint64(value.(int64))
	case int:
		return uint64(value.(int))
	case float32:
		return uint64(value.(float32))
	case float64:
		return uint64(value.(float64))
	}
	return 0
}

func ToInt8(value interface{}) int8 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int8(value.(uint8))
	case uint16:
		return int8(value.(uint16))
	case uint32:
		return int8(value.(uint32))
	case uint64:
		return int8(value.(uint64))
	case uint:
		return int8(value.(uint))
	case int8:
		return int8(value.(int8))
	case int16:
		return int8(value.(int16))
	case int32:
		return int8(value.(int32))
	case int64:
		return int8(value.(int64))
	case int:
		return int8(value.(int))
	case float32:
		return int8(value.(float32))
	case float64:
		return int8(value.(float64))
	}
	return 0
}

func ToInt16(value interface{}) int16 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int16(value.(uint8))
	case uint16:
		return int16(value.(uint16))
	case uint32:
		return int16(value.(uint32))
	case uint64:
		return int16(value.(uint64))
	case uint:
		return int16(value.(uint))
	case int8:
		return int16(value.(int8))
	case int16:
		return int16(value.(int16))
	case int32:
		return int16(value.(int32))
	case int64:
		return int16(value.(int64))
	case int:
		return int16(value.(int))
	case float32:
		return int16(value.(float32))
	case float64:
		return int16(value.(float64))
	}
	return 0
}

func ToInt32(value interface{}) int32 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int32(value.(uint8))
	case uint16:
		return int32(value.(uint16))
	case uint32:
		return int32(value.(uint32))
	case uint64:
		return int32(value.(uint64))
	case uint:
		return int32(value.(uint))
	case int8:
		return int32(value.(int8))
	case int16:
		return int32(value.(int16))
	case int32:
		return int32(value.(int32))
	case int64:
		return int32(value.(int64))
	case int:
		return int32(value.(int))
	case float32:
		return int32(value.(float32))
	case float64:
		return int32(value.(float64))
	}
	return 0
}

func ToInt64(value interface{}) int64 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return int64(value.(uint8))
	case uint16:
		return int64(value.(uint16))
	case uint32:
		return int64(value.(uint32))
	case uint64:
		return int64(value.(uint64))
	case uint:
		return int64(value.(uint))
	case int8:
		return int64(value.(int8))
	case int16:
		return int64(value.(int16))
	case int32:
		return int64(value.(int32))
	case int64:
		return int64(value.(int64))
	case int:
		return int64(value.(int))
	case float32:
		return int64(value.(float32))
	case float64:
		return int64(value.(float64))
	}
	return 0
}

func ToFloat32(value interface{}) float32 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return float32(value.(uint8))
	case uint16:
		return float32(value.(uint16))
	case uint32:
		return float32(value.(uint32))
	case uint64:
		return float32(value.(uint64))
	case uint:
		return float32(value.(uint))
	case int8:
		return float32(value.(int8))
	case int16:
		return float32(value.(int16))
	case int32:
		return float32(value.(int32))
	case int64:
		return float32(value.(int64))
	case int:
		return float32(value.(int))
	case float32:
		return value.(float32)
	case float64:
		return float32(value.(float64))
	}
	return 0
}

func ToFloat64(value interface{}) float64 {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1
		} else {
			return 0
		}
	case uint8:
		return float64(value.(uint8))
	case uint16:
		return float64(value.(uint16))
	case uint32:
		return float64(value.(uint32))
	case uint64:
		return float64(value.(uint64))
	case uint:
		return float64(value.(uint))
	case int8:
		return float64(value.(int8))
	case int16:
		return float64(value.(int16))
	case int32:
		return float64(value.(int32))
	case int64:
		return float64(value.(int64))
	case int:
		return float64(value.(int))
	case float32:
		return float64(value.(float32))
	case float64:
		return value.(float64)
	}
	return 0
}
