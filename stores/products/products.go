package products

import (
	"gofr.dev/pkg/gofr"

	"product-variant-svc/filters"
	"product-variant-svc/models"
	"product-variant-svc/stores"
)

type store struct{}

func New() stores.Product {
	return store{}
}

func (s store) Create(ctx *gofr.Context, product models.Product) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetAll(ctx *gofr.Context, productFilters filters.Products, variantFilters filters.Variants) ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s store) GetByID(ctx *gofr.Context, productID int) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}
