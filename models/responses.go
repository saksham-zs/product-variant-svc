package models

type ProductWithVariantsResponse struct {
	ProductDetails Product   `json:"productDetails"`
	VariantDetails []Variant `json:"variantDetails"`
}

type ProductsWithVariantsResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	BrandName      string    `json:"brandName"`
	Details        string    `json:"details"`
	ImageURL       string    `json:"imageUrl"`
	VariantDetails []Variant `json:"variantDetails"`
}
