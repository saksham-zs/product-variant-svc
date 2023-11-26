package products

import (
	"strconv"

	"gofr.dev/pkg/errors"

	"product-variant-svc/models"
)

func validateParams(product models.Product) error {
	if err := checkMissingParams(product); err != nil {
		return err
	}

	return nil
}

func checkMissingParams(product models.Product) error {
	switch {
	case product.Name == "":
		return errors.MissingParam{Param: []string{"name"}}
	case product.BrandName == "":
		return errors.MissingParam{Param: []string{"brand_name"}}
	case product.Details == "":
		return errors.MissingParam{Param: []string{"details"}}
	case product.ImageURL == "":
		return errors.MissingParam{Param: []string{"image_url"}}
	}

	return nil
}

func validateProductID(pid string) (int, error) {
	id, err := strconv.Atoi(pid)
	if err != nil {
		return 0, err
	}

	return id, nil
}
