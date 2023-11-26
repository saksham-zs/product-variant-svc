package variants

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
	"product-variant-svc/services"
)

type handler struct {
	service services.Variants
}

func New(s services.Variants) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Product ID"}}
	}

	var v models.Variant

	err = ctx.Bind(&v)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Request Body"}}
	}

	resp, err := h.service.Create(ctx, v, pid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	pid := ctx.PathParam("pid")
	vid := ctx.PathParam("vid")

	resp, err := h.service.GetByID(ctx, pid, vid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
