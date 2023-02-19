package entity

type ProductCategory struct {
	ID      int8   `json:"id"`
	Name    string `json:"name"`
	StoreId string `json:"store_id"`
}
