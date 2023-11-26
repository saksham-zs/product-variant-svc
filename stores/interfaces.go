package stores

import (
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
)

type Product interface {
	Create(ctx *gofr.Context, product models.Product) (int64, error)
	GetProductsWithVariants(ctx *gofr.Context, filters models.Filters) ([]models.ProductWithVariantsResponse, error)
	GetByID(ctx *gofr.Context, productID int) (models.Product, error)
}

type Variant interface {
	Create(ctx *gofr.Context, productID int, variant models.Variant) (int64, error)
	GetByID(ctx *gofr.Context, productID int, variantID int) (models.Variant, error)
}
