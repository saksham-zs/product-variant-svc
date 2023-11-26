package variants

import (
	"database/sql"
	"fmt"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
	"product-variant-svc/stores"
)

type service struct {
	variantStore stores.Variant
	productStore stores.Product
}

func New(v stores.Variant, p stores.Product) service {
	return service{variantStore: v, productStore: p}
}

func (s service) Create(ctx *gofr.Context, v models.Variant, pid int) (models.Variant, error) {
	err := checkMissingParams(v)
	if err != nil {
		return models.Variant{}, err
	}

	id, err := s.variantStore.Create(ctx, pid, v)
	if err != nil {
		return models.Variant{}, err
	}

	v.ID = strconv.FormatInt(id, 10)

	return v, nil
}

func (s service) GetByID(ctx *gofr.Context, pid, vid string) (models.ProductWithVariantsResponse, error) {
	productID, variantID, err := validateIDs(pid, vid)
	if err != nil {
		ctx.Logger.Errorf("Error while validating product & variant IDs in services/variant/GetByID. Err: %v", err)

		return models.ProductWithVariantsResponse{}, err
	}

	variant, err := s.variantStore.GetByID(ctx, productID, variantID)
	if err != nil {
		ctx.Logger.Errorf("Error while getting variant details by productID in services/variant/GetByID. Err: %v", err)

		if err == sql.ErrNoRows {
			return models.ProductWithVariantsResponse{}, errors.EntityNotFound{Entity: "variant", ID: fmt.Sprint(variantID)}
		}

		return models.ProductWithVariantsResponse{}, err
	}

	product, err := s.productStore.GetByID(ctx, productID)
	if err != nil {
		ctx.Logger.Errorf("Error while getting product details by ID: %v in services/variant/GetByID. Err: %v", productID, err)

		return models.ProductWithVariantsResponse{}, err
	}

	return models.ProductWithVariantsResponse{ProductDetails: product, VariantDetails: []models.Variant{variant}}, nil
}
