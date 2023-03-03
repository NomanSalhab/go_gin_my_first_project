package entity

import "time"

type Order struct {
	ID               int            `json:"id"`
	UserID           int            `json:"user_id" binding:"required"`
	OrderTime        time.Time      `json:"order_time"`
	DeliveryTime     time.Time      `json:"delivery_time"`
	Products         []OrderProduct `json:"products" binding:"required"`
	ProductsCost     float32        `json:"products_cost"`
	Address          Address        `json:"address" binding:"required"`
	DeliveryCost     float32        `json:"delivery_cost"` //*  binding:"required"
	Addons           string         `json:"addons"`
	AddonsCost       float32        `json:"addons_cost"`
	State            OrderState     `json:"state"`
	Finished         bool           `json:"finished"`
	DeliveryWorkerId int            `json:"delivery_worker_id"`
}

type OrderProduct struct {
	ProductID    int `json:"product_id"`
	ProductCount int `json:"product_count"`
	ProductPrice int `json:"product_price"`
	StoreId      int `json:"store_id"`
}

type OrderInfoRequest struct {
	ID int `json:"id"`
}

type OrderChangeStateRequest struct {
	ID       int        `json:"id" binding:"required"`
	State    OrderState `json:"state" binding:"required"`
	Finished bool       `json:"finished"`
	UserId   int        `json:"user_id"`
}

type OrderEditRequest struct {
	ID               int            `json:"id" binding:"required"`
	UserID           int            `json:"user_id"`
	OrderTime        time.Time      `json:"order_time"`
	DeliveryTime     time.Time      `json:"delivery_time"`
	Products         []OrderProduct `json:"products[]"`
	ProductsCost     float32        `json:"products_cost"`
	Address          Address        `json:"address"`
	DeliveryCost     float32        `json:"delivery_cost"`
	Addons           string         `json:"addons"`
	AddonsCost       float32        `json:"addons_cost"`
	State            OrderState     `json:"state"`
	Finished         bool           `json:"finished"`
	DeliveryWorkerId int            `json:"delivery_worker_id"`
}

type OrderDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}

type OrderState struct {
	Text   string `json:"text"`
	Number int    `json:"number"`

	//* States:
	//* 1: Ordering
	//* 2: Preparing Order
	//* 3: On The Way
	//* 4: Delivered
}
