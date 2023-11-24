package products

import "product-variant-svc/stores"

type service struct {
	store stores.Product
}

func New(s stores.Product) service {
	return service{store: s}
}
