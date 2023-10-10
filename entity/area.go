package entity

type Area struct {
	ID     int     `json:"id"`
	Name   string  `json:"name" binding:"required"`
	Lat    float32 `json:"lat" binding:"required"`
	Long   float32 `json:"long" binding:"required"`
	Active bool    `json:"active"`
}

type AreaEditRequest struct {
	ID     int     `json:"id" binding:"required"`
	Name   string  `json:"name"`
	Lat    float32 `json:"lat"`
	Long   float32 `json:"long"`
	Active bool    `json:"active"`
}

type AreaActivateRequest struct {
	ID int `json:"id" binding:"required"`
}

type AreaDeactivateRequest struct {
	ID int `json:"id" binding:"required"`
}
