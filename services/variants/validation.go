package variants

import (
	"strconv"

	"gofr.dev/pkg/errors"

	"product-variant-svc/models"
)

func checkMissingParams(variant models.Variant) error {
	switch {
	case variant.Name == "":
		return errors.MissingParam{Param: []string{"name"}}
	case variant.Details == "":
		return errors.MissingParam{Param: []string{"details"}}
	}

	return nil
}

func validateIDs(pid, vid string) (int, int, error) {
	productID, err := strconv.Atoi(pid)
	if err != nil {
		return 0, 0, err
	}

	variantID, err := strconv.Atoi(vid)
	if err != nil {
		return 0, 0, err
	}

	return productID, variantID, nil
}
