package internal

type Modbus interface {
	Read(slave uint8, code uint8, addr uint16, size uint16) ([]byte, error)
	Write(slave uint8, code uint8, addr uint16, buf []byte) error
}
