package entity

type Store struct {
	ID              int8   `json:"id"`
	Name            string `json:"name"`
	StoreCategoryId int8   `json:"store_category_id"`
	Image           string `json:"image"`
	Balance         string `json:"balance"`
}
