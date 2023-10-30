package entity

type Slider struct {
	ID        int    `json:"id"`
	Image     string `json:"image"  binding:"required"` /*File*/
	StoreId   int    `json:"store_id"  binding:"required"`
	ProductId int    `json:"product_id"  binding:"required"`
	Active    bool   `json:"active"`
}

type SliderEditRequest struct {
	ID        int  `json:"id" binding:"required"`
	StoreId   int  `json:"store_id"`
	ProductId int  `json:"product_id"`
	Active    bool `json:"active"`
}

type StoreSliders struct {
	StoreId int `json:"store_id"  binding:"required"`
}
