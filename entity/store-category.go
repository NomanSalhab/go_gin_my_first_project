package entity

type StoreCategory struct {
	ID     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	Active bool   `json:"active"`
}

type StoreCategoryInfoRequest struct {
	ID int `json:"id"`
}

type StoreCategoryDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}

type StoreCategoryEditRequest struct {
	ID     int    `json:"id" binding:"required"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
