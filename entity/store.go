package entity

type Store struct {
	ID              int     `json:"id"`
	Name            string  `json:"name" binding:"required"`
	StoreCategoryId int     `json:"store_category_id" binding:"required"`
	Image           string  `json:"image" binding:"required"`
	Balance         float32 `json:"balance"`
	Active          bool    `json:"active"`
	DeliveryRent    int     `json:"delivery_rent"`
}

type StoreInfoRequest struct {
	ID int `json:"id" binding:"required"`
}

type StoreEditRequest struct {
	ID              int     `json:"id" binding:"required"`
	Name            string  `json:"name"`
	StoreCategoryId int     `json:"store_category_id"`
	Image           string  `json:"image"`
	Balance         float32 `json:"balance"`
	DeliveryRent    int     `json:"delivery_rent"`
	// Active          bool    `json:"active"` binding:"required"
}

type StoreDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}
