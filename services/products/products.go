package products

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
	productStore stores.Product
	variantStore stores.Variant
}

func New(p stores.Product, v stores.Variant) service {
	return service{productStore: p, variantStore: v}
}

func (s service) Create(ctx *gofr.Context, p models.Product) (models.Product, error) {
	err := validateParams(p)
	if err != nil {
		return models.Product{}, err
	}

	id, err := s.productStore.Create(ctx, p)
	if err != nil {
		return models.Product{}, err
	}

	p.ID = strconv.FormatInt(id, 10)

	return p, nil
}

func (s service) GetByID(ctx *gofr.Context, pid string) (models.ProductWithVariantsResponse, error) {
	productID, err := validateProductID(pid)
	if err != nil {
		ctx.Logger.Errorf("Error while validating productID: %v, Err: %v, ", pid, err)
	}

	product, err := s.productStore.GetByID(ctx, productID)
	if err != nil {
		ctx.Logger.Errorf("Error while getting product details by ID: %v in services/product/GetByID. Err: %v", productID, err)

		if err == sql.ErrNoRows {
			return models.ProductWithVariantsResponse{}, errors.EntityNotFound{Entity: "product", ID: fmt.Sprint(productID)}
		}

		return models.ProductWithVariantsResponse{}, err
	}

	variant, err := s.variantStore.GetByID(ctx, productID, 0)
	if err != nil {
		ctx.Logger.Errorf("Error while getting variant details by productID: %v in services/product/GetByID. Err: %v", productID, err)

		if err != sql.ErrNoRows {
			return models.ProductWithVariantsResponse{}, err
		}
	}

	return models.ProductWithVariantsResponse{ProductDetails: product, VariantDetails: []models.Variant{variant}}, nil
}

func (s service) GetAll(ctx *gofr.Context, params map[string]string) ([]models.ProductWithVariantsResponse, error) {
	//TODO implement me
	panic("implement me")
}

//func (s service) GetAll(ctx *gofr.Context, params map[string]string) ([]models.ProductWithVariantsResponse, error) {
//	filters, err := populateFilters(ctx, params)
//	if err != nil {
//		ctx.Logger.Errorf("Error while populating filters in services.products/GetAll. Err: %v", err)
//	}
//
//	resp, err := s.productStore.GetProductsWithVariants(ctx, filters)
//	if err != nil {
//		ctx.Logger.Errorf("Error while getting product & variants details in services.products/GetAll. Err: %v", err)
//		return nil, err
//	}
//
//	return resp, nil
//}

//func populateFilters(ctx *gofr.Context, paramsMap map[string]string) (filters models.Filters, err error) {
//	params, err := json.Marshal(paramsMap)
//	if err != nil {
//		ctx.Logger.Errorf("Error while marshalling params map in services/products/GetAll. Err: %v", err)
//		return filters, err
//	}
//
//	err = json.Unmarshal(params, &filters)
//	if err != nil {
//		ctx.Logger.Errorf("Error while un-marshalling params in services/products/GetAll. Err: %v", err)
//		return filters, err
//	}
//
//	return filters, nil
//}
