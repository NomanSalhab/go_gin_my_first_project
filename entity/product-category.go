package entity

type ProductCategory struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	StoreId int    `json:"store_id" binding:"required"`
	Active  bool   `json:"active"`
}

type ProductCategoryInfoRequest struct {
	ID int `json:"id" binding:"required"`
}

type ProductCategoriesByStoreInfoRequest struct {
	StoreId int `json:"store_id" binding:"required"`
}

type ProductCategoryEditRequest struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name"`
	StoreId int    `json:"store_id" binding:"required"`
	Active  bool   `json:"active"` // binding:"required"
}

type ProductCategoryDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}
