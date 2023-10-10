package entity

type Detail struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	IsFlavor bool   `json:"is_flavor"`
	IsVolume bool   `json:"is_volume"`
	IsAddon  bool   `json:"is_addon"`
	Price    int    `json:"price"`
}

type DetailEditRequest struct {
	ID       int    `json:"id" binding:"required"`
	Name     string `json:"name"`
	IsFlavor bool   `json:"is_flavor"`
	IsVolume bool   `json:"is_volume"`
	IsAddon  bool   `json:"is_addon"`
	Price    int    `json:"price"`
}
