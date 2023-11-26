package main

import (
	"gofr.dev/pkg/gofr"

	productsHandlers "product-variant-svc/http/products"
	productsServices "product-variant-svc/services/products"
	productsStores "product-variant-svc/stores/products"

	variantsHandlers "product-variant-svc/http/variants"
	variantsServices "product-variant-svc/services/variants"
	variantsStores "product-variant-svc/stores/variants"
)

func main() {
	app := gofr.New()

	// initialize products & variants instance
	variantStore := variantsStores.New()
	productStore := productsStores.New()

	variantService := variantsServices.New(variantStore, productStore)
	productService := productsServices.New(productStore, variantStore)

	variantHandler := variantsHandlers.New(variantService)
	productHandler := productsHandlers.New(productService)

	app.POST("/product", productHandler.Create)                     // Create Product
	app.GET("/product/{pid}", productHandler.GetByID)               // Get Product Details (including all variants) by ID
	app.GET("/product", productHandler.GetAll)                      // Filter and list Products
	app.POST("/product/{pid}/variant", variantHandler.Create)       // Create Variant for Product
	app.GET("/product/{pid}/variant/{vid}", variantHandler.GetByID) // Get Variant by ID

	app.Start()
}
