package entity

type Product struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	StoreId           int     `json:"store_id"`
	StoreName         string  `json:"store_name"`
	ProductCategoryId string  `json:"product_category_id"`
	Image             string  `json:"image"`
	Summary           string  `json:"summary"`
	Price             float32 `json:"price"`
	OrderCount        string  `json:"order_count"`
}
