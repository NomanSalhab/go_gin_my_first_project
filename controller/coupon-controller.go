package controller

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type CouponController interface {
	AddCoupon(ctx *gin.Context) error
	FindAllCoupons() []entity.CouponEditRequest
	EditCoupon(ctx *gin.Context) error
	DeleteCoupon(ctx *gin.Context) error
	GetCouponInfo(idValue int) (entity.CouponAddOrderRequest, error)

	ActivateCoupon(wantedId int) error
	DeactivateCoupon(wantedId int) error
	EnableFreeDelivery(wantedId int) error
	DisableFreeDelivery(wantedId int) error
	EnableFromProducts(wantedId int) error
	DisableFromProducts(wantedId int) error
}

type couponController struct {
	service service.CouponService
}

func NewCouponController(service service.CouponService) CouponController {
	return &couponController{
		service: service,
	}
}

func (c *couponController) FindAllCoupons() []entity.CouponEditRequest {
	return c.service.FindAllCoupons()
}

func (c *couponController) GetCouponInfo(idValue int) (entity.CouponAddOrderRequest, error) {
	var couponInfo entity.CouponAddOrderRequest

	if idValue == 0 {
		return entity.CouponAddOrderRequest{}, errors.New("coupon id can not be zero")
	}

	couponInfo = c.service.GetCouponInfo(idValue)
	return couponInfo, nil
}

func (c *couponController) AddCoupon(ctx *gin.Context) error /*entity.Video*/ {
	var coupon entity.Coupon
	err := ctx.ShouldBindJSON(&coupon)
	if err != nil {
		return err
	}

	err = validate.Struct(coupon)
	if err != nil {
		return err
	}

	err = c.service.AddCoupon(coupon)
	return err /*video*/
}

func (c *couponController) EditCoupon(ctx *gin.Context) error {
	var couponEditInfo entity.CouponEditRequest
	err := ctx.ShouldBindJSON(&couponEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(couponEditInfo)
	if err != nil {
		return err
	}
	err = c.service.EditCoupon(couponEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) DeleteCoupon(ctx *gin.Context) error {
	var couponId entity.CouponEditRequest
	err := ctx.ShouldBindJSON(&couponId)
	if err != nil {
		return err
	}
	err = c.service.DeleteCoupon(couponId)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) ActivateCoupon(wantedId int) error {
	if wantedId == 0 {
		return errors.New("coupon id can not be zero")
	}
	err := c.service.ActivateCoupon(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) DeactivateCoupon(wantedId int) error {
	if wantedId == 0 {
		return errors.New("coupon id can not be zero")
	}
	err := c.service.DeactivateCoupon(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) EnableFreeDelivery(wantedId int) error {
	if wantedId == 0 {
		return errors.New("coupon id can not be zero")
	}
	err := c.service.EnableFreeDelivery(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) DisableFreeDelivery(wantedId int) error {
	if wantedId == 0 {
		return errors.New("coupon id can not be zero")
	}
	err := c.service.DisableFreeDelivery(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) EnableFromProducts(wantedId int) error {
	if wantedId == 0 {
		return errors.New("coupon id can not be zero")
	}
	err := c.service.EnableFromProducts(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (c *couponController) DisableFromProducts(wantedId int) error {
	if wantedId == 0 {
		return errors.New("coupon id can not be zero")
	}
	err := c.service.DisableFromProducts(wantedId)
	if err != nil {
		return err
	}
	return nil
}
