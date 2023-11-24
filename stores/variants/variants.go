package variants

import (
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
	"product-variant-svc/stores"
)

type store struct{}

func New() stores.Variant {
	return store{}
}

func (s store) Create(ctx *gofr.Context, productID int, variant models.Variant) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetByID(ctx *gofr.Context, productID int, variantID int) (models.Variant, error) {
	//TODO implement me
	panic("implement me")
}
