package services

import (
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
)

type Products interface {
	Create(ctx *gofr.Context, product models.Product) (models.Product, error)
	GetByID(ctx *gofr.Context, productID string) (models.ProductWithVariantsResponse, error)
	GetAll(ctx *gofr.Context, params map[string]string) ([]models.ProductsWithVariantsResponse, error)
}

type Variants interface {
	Create(ctx *gofr.Context, v models.Variant, pid int) (models.Variant, error)
	GetByID(ctx *gofr.Context, pid, vid string) (models.ProductWithVariantsResponse, error)
}
