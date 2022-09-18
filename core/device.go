package core

import "github.com/zgwit/iot-master/v2/model"

type Point struct {
	model.Point
}

type Product struct {
	model.Product
}

type Device struct {
	model.Product

	product *Product
}
