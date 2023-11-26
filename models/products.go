package models

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandName string `json:"brandName"`
	Details   string `json:"details"`
	ImageURL  string `json:"imageUrl"`
}
