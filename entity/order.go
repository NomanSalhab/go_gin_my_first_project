package entity

import "time"

type Order struct {
	ID               int            `json:"id"`
	UserID           int            `json:"user_id" binding:"required"`
	UserName         string         `json:"user_name"`  //  binding:"required"
	UserPhone        string         `json:"user_phone"` //  binding:"required"
	OrderTime        time.Time      `json:"order_time"`
	DeliveryTime     time.Time      `json:"delivery_time"`
	Products         []OrderProduct `json:"products" binding:"required"`
	ProductsCost     int            `json:"products_cost"`
	Address          OrderAddress   `json:"address" binding:"required"`
	DeliveryCost     int            `json:"delivery_cost"` // binding:"required"
	Notes            string         `json:"notes"`         //* binding:"required"
	DeliveryWorkerId int            `json:"delivery_worker_id"`
	Ordered          bool           `json:"ordered"`
	OnTheWay         bool           `json:"on_the_way"`
	Finished         bool           `json:"finished"`
	CouponID         int            `json:"coupon_id"`
}

// Addons       [][]int        `json:"addons" binding:"required"`
// Flavors      [][]int        `json:"flavors" binding:"required"`
// Volumes      [][]int        `json:"volumes" binding:"required"`
// AddonsCost       float32        `json:"addons_cost"`
// State            OrderState `json:"state"`

type OrderProduct struct {
	ID                int                 `json:"id"`
	ProductID         int                 `json:"product_id" binding:"required"`
	ProductCount      int                 `json:"product_count" binding:"required"`
	ProductPrice      int                 `json:"product_price" binding:"required"`
	StoreId           int                 `json:"store_id" binding:"required"`
	StoreName         string              `json:"store_name" binding:"required"`
	StoreDeliveryCost int                 `json:"store_delivery_cost" binding:"required"`
	OrderId           int                 `json:"order_id" binding:"required"`
	Addons            []DetailEditRequest `json:"addons"` //  binding:"required"
	Flavors           DetailEditRequest   `json:"flavors" binding:"required"`
	Volumes           DetailEditRequest   `json:"volumes" binding:"required"`
	OrderTime         time.Time           `json:"order_time"`
}

type OrderInfoRequest struct {
	ID int `json:"id" binding:"required"`
}

type OrderChangeStateRequest struct {
	ID       int  `json:"id" binding:"required"`
	UserId   int  `json:"user_id" binding:"required"`
	Finished bool `json:"finished"`
	Ordered  bool `json:"ordered"`
	OnTheWay bool `json:"on_the_way"`
}

type OrderEditRequest struct {
	ID               int            `json:"id" binding:"required"`
	UserID           int            `json:"user_id"`
	OrderTime        time.Time      `json:"order_time"`
	DeliveryTime     time.Time      `json:"delivery_time"`
	Address          Address        `json:"address"`
	Products         []OrderProduct `json:"products"`
	ProductsCost     int            `json:"products_cost"`
	DeliveryCost     int            `json:"delivery_cost"`
	DeliveryWorkerId int            `json:"delivery_worker_id"`
	Notes            string         `json:"notes"`
	Ordered          bool           `json:"ordered"`
	OnTheWay         bool           `json:"on_the_way"`
	Finished         bool           `json:"finished"`
	CouponID         int            `json:"coupon_id"`
}

// Addons       string         `json:"addons"`
// AddonsCost       float32        `json:"addons_cost"`
// State            OrderState `json:"state"`

type OrderDeleteRequest struct {
	ID int `json:"id" binding:"required"`
}

// type OrderState struct {
// 	Text   string `json:"text"`
// 	Number int    `json:"number"`
// 	//* States:
// 	//* 1: Ordering
// 	//* 2: Preparing Order
// 	//* 3: On The Way
// 	//* 4: Delivered
// }
