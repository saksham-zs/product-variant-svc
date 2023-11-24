package variants

import "product-variant-svc/stores"

type service struct {
	store stores.Variant
}

func New(s stores.Variant) service {
	return service{store: s}
}
