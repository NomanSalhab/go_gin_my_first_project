package entity

import "time"

type Order struct {
	ID             int8      `json:"id"`
	Name           string    `json:"name"`
	OrderTime      time.Time `json:"order_time"`
	ProductsIDs    []int8    `json:"products_ids"`
	ProductsCounts []int8    `json:"products_counts"`
	ProductsCost   float32   `json:"products_cost"`
	OrderAddress   Address   `json:"order_address"`
	DeliveryCost   float32   `json:"delivery_cost"`
	AddonsCost     float32   `json:"addons_cost"`
}
