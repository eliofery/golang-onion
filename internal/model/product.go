package model

//const ProductTableName = "products"

type Product struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"image,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
