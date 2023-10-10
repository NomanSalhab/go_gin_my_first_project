package entity

type PaginationInfo struct {
	MaximumPagesCount int `json:"maximum_pages_count"`
	AllItemsCount     int `json:"all_items_count"`
}
