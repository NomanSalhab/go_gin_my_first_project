package entity

type HomePage struct {
	Sliders             []Slider  `json:"sliders"`
	BestSellingProducts []Product `json:"best_selling_products"`
	BestSellingStores   []Store   `json:"best_selling_stores"`
	OffersProducts      []Product `json:"offers_products"`
	IsUpdatedVersion    bool      `json:"is_updated_version"`
	VersionLink         string    `json:"link"`
}
