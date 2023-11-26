package models

type Filters struct {
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
	VariantID   int    `json:"variantId"`
}
