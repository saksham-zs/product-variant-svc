package models

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandName string `json:"brand_name"`
	Details   string `json:"details"`
	ImageURL  string `json:"image_url"`
}
