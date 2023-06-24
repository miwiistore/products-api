package domain

type Product struct {
	ProductID   int64     `json:"product_id,omitempty"`
	CategoryID  int16     `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pictures    []Picture `json:"pictures"`
	Quantity    int16     `json:"quantity"`
	Price       float32   `json:"price"`
	SKU         string    `json:"sku"`
}
