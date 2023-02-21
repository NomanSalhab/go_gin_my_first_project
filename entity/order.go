package entity

import "time"

type Order struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	OrderTime      time.Time `json:"order_time"`
	ProductsIDs    []int     `json:"products_ids"`
	ProductsCounts []int     `json:"products_counts"`
	ProductsCost   float32   `json:"products_cost"`
	OrderAddress   Address   `json:"order_address"`
	DeliveryCost   float32   `json:"delivery_cost"`
	AddonsCost     float32   `json:"addons_cost"`
}
