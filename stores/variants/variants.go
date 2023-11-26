package variants

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"product-variant-svc/models"
	"product-variant-svc/stores"
)

type store struct{}

func New() stores.Variant {
	return store{}
}

func (s store) Create(ctx *gofr.Context, pid int, v models.Variant) (int64, error) {
	rows, err := ctx.DB().ExecContext(ctx, insertIntoVariants, pid, v.Name, v.Details)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	id, err := rows.LastInsertId()
	if err != nil || id == 0 {
		return 0, errors.DB{Err: err}
	}

	return id, nil
}

func (s store) GetByID(ctx *gofr.Context, productID int, variantID int) (v models.Variant, err error) {
	whereClause := generateWhereClause(productID, variantID)

	switch variantID {
	case 0:
		err = ctx.DB().QueryRowContext(ctx, getVariantByID+whereClause, productID).Scan(&v.ID, &v.Name, &v.Details)
	default:
		err = ctx.DB().QueryRowContext(ctx, getVariantByID+whereClause, productID, variantID).Scan(&v.ID, &v.Name, &v.Details)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return v, sql.ErrNoRows
		}

		return v, errors.DB{Err: err}
	}

	return v, nil
}

func generateWhereClause(pid, vid int) (clause string) {
	var flag = false
	clause = " WHERE "

	if pid != 0 {
		clause += "product_id=?"
		flag = true
	}

	// check for when GetByID call's done from another service
	if vid != 0 {
		if flag {
			clause += " AND "
		}

		clause += "id=?"
	}

	return clause
}
