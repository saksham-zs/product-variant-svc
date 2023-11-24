package products

import (
	"gofr.dev/pkg/gofr"
	"product-variant-svc/services"
)

type handler struct {
	service services.Products
}

func New(s services.Products) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	return nil, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	return nil, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	return nil, nil
}
