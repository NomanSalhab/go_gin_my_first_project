package entity

type Store struct {
	ID              int     `json:"id"`
	AreaID          int     `json:"area_id" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	StoreCategoryId int     `json:"store_category_id" binding:"required"`
	Image           string  `json:"image" binding:"required"`
	Balance         int     `json:"balance"`
	Active          bool    `json:"active"`
	DeliveryRent    int     `json:"delivery_rent"`
	Discount        float32 `json:"discount"`
}

type StoreInfoRequest struct {
	ID int `json:"id" binding:"required"`
}

type StoreEditRequest struct {
	ID              int     `json:"id" binding:"required"`
	AreaID          int     `json:"area_id"`
	Name            string  `json:"name"`
	StoreCategoryId int     `json:"store_category_id"`
	Image           string  `json:"image"`
	Balance         int     `json:"balance"`
	DeliveryRent    int     `json:"delivery_rent"`
	Active          bool    `json:"active"` // binding:"required"
	Discount        float32 `json:"discount"`
}

type StoreDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}

type StoreIncreaseBalance struct {
	ID      int `json:"id" binding:"required"`
	Balance int `json:"balance" binding:"required"`
}
