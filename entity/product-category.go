package entity

type ProductCategory struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	StoreId string `json:"store_id"`
}
