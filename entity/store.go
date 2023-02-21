package entity

type Store struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	StoreCategoryId int    `json:"store_category_id"`
	Image           string `json:"image"`
	Balance         string `json:"balance"`
}
