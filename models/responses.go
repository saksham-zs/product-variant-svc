package models

type ProductWithVariantsResponse struct {
	ProductDetails Product   `json:"productDetails"`
	VariantDetails []Variant `json:"variantDetails"`
}
