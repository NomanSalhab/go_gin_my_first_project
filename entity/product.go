package entity

type Product struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name" binding:"required"`
	StoreId           int                 `json:"store_id" binding:"required"`
	ProductCategoryId int                 `json:"product_category_id" binding:"required"`
	Image             string              `json:"image" binding:"required"`
	Summary           string              `json:"summary" binding:"required"`
	Price             int                 `json:"price" binding:"required"`
	OrderCount        int                 `json:"order_count"`
	Active            bool                `json:"active"`
	Flavors           []DetailEditRequest `json:"flavors"`
	Volumes           []DetailEditRequest `json:"volumes"`
	Addons            []DetailEditRequest `json:"addons"`
	DiscountRatio     float32             `json:"discount_ratio"`
}

type ProductInfoRequest struct {
	ID int `json:"id" binding:"required"`
}

type ProductByCategoryRequest struct {
	ID                int `json:"id"`
	StoreId           int `json:"store_id"` //  binding:"required"
	ProductCategoryId int `json:"product_category_id" binding:"required"`
}

type OrderProductRequest struct {
	ID         int `json:"id" binding:"required"`
	OrderCount int `json:"order_count"`
}

type ProductEditRequest struct {
	ID                int                 `json:"id" binding:"required"`
	Name              string              `json:"name"`
	StoreId           int                 `json:"store_id"`
	ProductCategoryId int                 `json:"product_category_id"`
	Image             string              `json:"image"`
	Summary           string              `json:"summary"`
	Price             int                 `json:"price"`
	Active            bool                `json:"active"` // binding:"required"
	Flavors           []DetailEditRequest `json:"flavors"`
	Volumes           []DetailEditRequest `json:"volumes"`
	Addons            []DetailEditRequest `json:"addons"`
	DiscountRatio     float32             `json:"discount_ratio"`
}
