package stores

import (
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
)

type Product interface {
	Create(ctx *gofr.Context, product models.Product) (int64, error)
	GetAllProductsWithVariants(ctx *gofr.Context, filters models.Filters) ([]models.ProductsWithVariantsResponse, error)
	GetByID(ctx *gofr.Context, productID int) (models.Product, error)
}

type Variant interface {
	Create(ctx *gofr.Context, productID int, variant models.Variant) (int64, error)
	GetByID(ctx *gofr.Context, productID int, variantID int) (models.Variant, error)
	GetAllByProductID(ctx *gofr.Context, productID int) (variants []models.Variant, err error)
}
