package variants

import (
	"gofr.dev/pkg/gofr"
	"product-variant-svc/services"
)

type handler struct {
	service services.Variants
}

func New(s services.Variants) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	return nil, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	return nil, nil
}
