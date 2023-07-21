package internal

import (
	"github.com/PaesslerAG/gval"
	"github.com/zgwit/iot-gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/calc"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"xorm.io/xorm"
)

type Product struct {
	*types.Product

	filters     []gval.Evaluable
	calculators []gval.Evaluable
}

var Products lib.Map[Product]

func LoadProducts() error {
	var products []*types.Product
	err := db.Engine.Find(&products)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range products {
		err := LoadProduct(m)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func LoadProduct(m *types.Product) error {
	p := &Product{Product: m}

	for _, v := range m.Filters {
		expr, err := calc.New(v.Expression)
		if err != nil {
			return err
		}
		p.filters = append(p.filters, expr)
	}

	for _, v := range m.Calculators {
		expr, err := calc.New(v.Expression)
		if err != nil {
			return err
		}
		p.calculators = append(p.calculators, expr)
	}

	Products.Store(m.Id, p)
	return nil
}

func GetProduct(id string) *Product {
	return Products.Load(id)
}
