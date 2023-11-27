package products

import (
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"strconv"

	"product-variant-svc/models"
	"product-variant-svc/stores"
)

type store struct {
	variantStore stores.Variant
}

func New(v stores.Variant) stores.Product {
	return store{variantStore: v}
}

func (s store) Create(ctx *gofr.Context, p models.Product) (int64, error) {
	rows, err := ctx.DB().ExecContext(ctx, insertIntoProducts, p.Name, p.BrandName, p.Details, p.ImageURL)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	id, err := rows.LastInsertId()
	if err != nil || id == 0 {
		return 0, errors.DB{Err: err}
	}

	return id, nil
}

func (s store) GetAllProductsWithVariants(ctx *gofr.Context, filters models.Filters) (products []models.ProductsWithVariantsResponse, err error) {
	whereClause, values := generateWhereClause(filters)

	rows, err := ctx.DB().QueryContext(ctx, getAllProducts+whereClause, values...)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var (
			p   models.ProductsWithVariantsResponse
			pid int
		)

		err = rows.Scan(&p.ID, &p.Name, &p.BrandName, &p.Details, &p.ImageURL)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		pid, err = strconv.Atoi(p.ID)
		if err != nil {
			ctx.Logger.Errorf("Error while converting product ID to int in stores/product/GetAllProductsWithVariants. Err: %v", err)
		}

		if filters.VariantID < 1 {
			p.VariantDetails, err = s.variantStore.GetAllByProductID(ctx, pid)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.Logger.Infof("No variants found for product ID: %v", pid)
				}

				ctx.Logger.Errorf("Error while getting all variant details for product ID: %v in stores/product/GetAllProductsWithVariants. Err: %v", pid, err)
			}
		} else {
			variant, err := s.variantStore.GetByID(ctx, pid, filters.VariantID)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.Logger.Infof("No variant found for product ID: %v", pid)
				}

				ctx.Logger.Errorf("Error while getting variant details for productID: %v & variantID: %v in stores/product/GetAllProductsWithVariants. Err: %v", pid, filters.VariantID, err)
			}

			p.VariantDetails = []models.Variant{variant}
		}

		products = append(products, p)
	}

	return products, nil
}

func (s store) GetByID(ctx *gofr.Context, id int) (p models.Product, err error) {
	err = ctx.DB().QueryRowContext(ctx, getProductByID, id).Scan(&p.ID, &p.Name, &p.BrandName, &p.Details, &p.ImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, sql.ErrNoRows
		}

		return p, errors.DB{Err: err}
	}

	return p, nil
}

func generateWhereClause(filters models.Filters) (clause string, values []interface{}) {
	clause = ""

	if filters.VariantID == 0 && filters.ProductID != 0 && filters.ProductName != "" {
		clause = "WHERE id=? AND name=?"
		values = append(values, filters.ProductID, filters.ProductName)
	} else if filters.VariantID == 0 && filters.ProductID == 0 && filters.ProductName != "" {
		clause = "WHERE name=?"
		values = append(values, filters.ProductName)
	} else if filters.VariantID == 0 && filters.ProductID != 0 && filters.ProductName == "" {
		clause = "WHERE id=?"
		values = append(values, filters.ProductID)
	} else if filters.VariantID == 0 && filters.ProductID == 0 && filters.ProductName == "" {
		clause = ""
		values = []interface{}{}
	}

	return clause, values
}
