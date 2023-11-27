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
	err = ctx.DB().QueryRowContext(ctx, getVariantByID, productID, variantID).Scan(&v.ID, &v.Name, &v.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			return v, sql.ErrNoRows
		}

		return v, errors.DB{Err: err}
	}

	return v, nil
}

func (s store) GetAllByProductID(ctx *gofr.Context, productID int) (variants []models.Variant, err error) {
	rows, err := ctx.DB().QueryContext(ctx, getVariantsByProductID, productID)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var variant models.Variant

		err = rows.Scan(&variant.ID, &variant.Name, &variant.Details)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		variants = append(variants, variant)
	}

	return variants, nil
}
