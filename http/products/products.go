package products

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
	"product-variant-svc/services"
)

type handler struct {
	service services.Products
}

func New(s services.Products) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var p models.Product

	err := ctx.Bind(&p)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Request Body"}}
	}

	resp, err := h.service.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	params := ctx.Params()

	resp, err := h.service.GetAll(ctx, params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	pid := ctx.PathParam("pid")

	resp, err := h.service.GetByID(ctx, pid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
