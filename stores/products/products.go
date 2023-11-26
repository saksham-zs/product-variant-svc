package products

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
	"product-variant-svc/stores"
)

type store struct{}

func New() stores.Product {
	return store{}
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

func (s store) GetProductsWithVariants(ctx *gofr.Context, filters models.Filters) ([]models.ProductWithVariantsResponse, error) {
	//	query := "SELECT id, name, brand_name, details, image_url FROM products "
	//	whereClause, values := generateWhereClause(filters)
	//
	return nil, nil
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

//func generateWhereClause(filters models.Filters) (string, []interface{}) {
//	clause := "WHERE "
//	flag := false
//
//	if filters.ProductID > 0 {
//		clause += "id = ?"
//		flag = true
//	}
//
//	if filters.VariantID > 0 {
//		if flag {
//			clause += ", """
//		}
//	}
//
//	var values []interface{}
//
//	if filters.ProductID > 1 {
//		values = append(values, filters.ProductID)
//		column := ""
//		if filters.VariantID > 1 {
//			column = "id"
//		}
//		if filters.ProductName != "" {
//			column = "name"
//		}
//
//		clause += and + column + "=$"
//	}
//
//	return clause, values
//}
