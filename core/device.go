package core

import (
	"github.com/zgwit/go-plc/protocol"
	"github.com/zgwit/iot-master/v2/model"
)

type Point struct {
	*model.Point

	addr protocol.Addr
}

type Poller struct {
	*model.Poller

	addr protocol.Addr
}

type Product struct {
	*model.Product

	pollers []Poller
	points  []Point
}

func (p *Product) Resolve(resolver protocol.AddressResolver) error {
	p.points = make([]Point, 0)
	for _, pt := range p.Points {
		addr, err := resolver(pt.Area, pt.Address)
		if err != nil {
			return err
		}
		p.points = append(p.points, Point{
			Point: pt,
			addr:  addr,
		})
	}

	p.pollers = make([]Poller, 0)
	for _, pt := range p.Pollers {
		addr, err := resolver(pt.Area, pt.Address)
		if err != nil {
			return err
		}
		p.pollers = append(p.pollers, Poller{
			Poller: pt,
			addr:   addr,
		})
	}

	return nil
}

func (p *Product) Parse(buf []byte, start protocol.Addr) (map[string]any, error) {
	values := make(map[string]any)
	length := len(buf)
	for _, pt := range p.points {
		offset, has := pt.addr.Diff(start)
		if !has {
			continue
		}

		if offset >= length {
			continue
		}

		val, err := pt.Type.Decode(buf[offset:], pt.LE, pt.Dot)
		if err != nil {
			return nil, err
		}

		values[pt.Name] = val
	}
	return values, nil
}

type Device struct {
	*model.Device

	product *Product
}
