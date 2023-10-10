package entity

import "time"

type Coupon struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Code               string    `json:"code" binding:"required"`
	TimesUsed          int       `json:"times_used"`
	FreeDelivery       bool      `json:"free_delivery"`
	Active             bool      `json:"active"`
	TimesUsedLimit     bool      `json:"times_used_limit"`
	EndDate            time.Time `json:"end_date"`
	DiscountPercentage float32   `json:"discount_percentage"`
	DiscountAmount     int       `json:"discount_amount"`
	FromProductsCost   bool      `json:"from_products_cost"`
}

type CouponEditRequest struct {
	ID                 int       `json:"id" binding:"required"`
	Name               string    `json:"name"`
	Code               string    `json:"code"`
	TimesUsed          int       `json:"times_used"`
	FreeDelivery       bool      `json:"free_delivery"`
	Active             bool      `json:"active"`
	TimesUsedLimit     bool      `json:"times_used_limit"`
	EndDate            time.Time `json:"end_date"`
	DiscountPercentage float32   `json:"discount_percentage"`
	DiscountAmount     int       `json:"discount_amount"`
	FromProductsCost   bool      `json:"from_products_cost"`
}
