package define

import (
	"github.com/iot-master-contrib/gateway/connect"
	"github.com/iot-master-contrib/gateway/types"
)

type Factory func(conn connect.Conn, opts map[string]any) Adapter

type Enum struct {
	Name  string
	Label string
}

type StationOptions struct {
	Name    string
	Label   string
	Default int
	Enums   []Enum
}

type Protocol struct {
	Name    string
	Label   string
	Codes   []types.Code
	Station []StationOptions

	Factory Factory
}

type Adapter interface {
	Get(station types.Station, mappers []*types.Mapper, name string) (any, error)
	Set(station types.Station, mappers []*types.Mapper, name string, value any) error
	Poll(station types.Station, mappers []*types.Mapper) (map[string]any, error)
}
