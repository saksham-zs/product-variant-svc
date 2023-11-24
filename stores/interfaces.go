package stores

import (
	"gofr.dev/pkg/gofr"

	"product-variant-svc/filters"
	"product-variant-svc/models"
)

type Product interface {
	Create(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetAll(ctx *gofr.Context, productFilters filters.Products, variantFilters filters.Variants) ([]models.Product, error)
	GetByID(ctx *gofr.Context, productID int) (models.Product, error)
}

type Variant interface {
	Create(ctx *gofr.Context, productID int, variant models.Variant) (models.Product, error)
	GetByID(ctx *gofr.Context, productID int, variantID int) (models.Variant, error)
}
