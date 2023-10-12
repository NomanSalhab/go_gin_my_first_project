package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type CouponService interface {
	AddCoupon(entity.Coupon) error
	FindAllCoupons() []entity.CouponEditRequest
	EditCoupon(couponEditInfo entity.CouponEditRequest) error
	DeleteCoupon(couponDeleteInfo entity.CouponEditRequest) error
	GetCouponInfo(idValue int) entity.CouponAddOrderRequest

	ActivateCoupon(wantedId int) error
	DeactivateCoupon(wantedId int) error
	EnableFreeDelivery(wantedId int) error
	DisableFreeDelivery(wantedId int) error
	EnableFromProducts(wantedId int) error
	DisableFromProducts(wantedId int) error

	// AddMockDetails(details []entity.Detail)
}

type couponService struct {
	driver driver.CouponDriver
}

func NewCouponService(driver driver.CouponDriver) CouponService {
	return &couponService{
		driver: driver,
	}
}

func (service *couponService) AddCoupon(coupon entity.Coupon) error {
	couponsList, err := service.driver.FindAllCoupons()
	if err != nil {
		return err
	}
	for i := 0; i < len(couponsList); i++ {
		if couponsList[i].Name == coupon.Name && couponsList[i].Code == coupon.Code {
			return errors.New("coupon name and code already exist")
		}
	}
	if len(coupon.Code) != 0 {
		err = service.driver.AddCoupon(coupon)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("coupon code should be specified")
	}

	// successDetail := detail
	// if len(service.details) > 0 {
	// 	successDetail.ID = service.details[len(service.details)-1].ID + 1
	// } else {
	// 	successDetail.ID = 1
	// }
	// service.details = append(service.details, successDetail)
	// return nil
}

func (service *couponService) FindAllCoupons() []entity.CouponEditRequest {
	allCoupons, err := service.driver.FindAllCoupons()
	if err != nil {
		return make([]entity.CouponEditRequest, 0)
	}
	return allCoupons
	// return service.details
}

func (service *couponService) GetCouponInfo(idValue int) entity.CouponAddOrderRequest {
	coupon, err := service.driver.GetCouponInfo(idValue)
	if err != nil {
		return entity.CouponAddOrderRequest{}
	}
	return coupon
	// return service.details
}

func (service *couponService) EditCoupon(couponEditInfo entity.CouponEditRequest) error {
	_, err := service.driver.EditCoupon(couponEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) DeleteCoupon(couponDeleteInfo entity.CouponEditRequest) error {
	err := service.driver.DeleteCoupon(couponDeleteInfo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) ActivateCoupon(wantedId int) error {
	err := service.driver.ActivateCoupon(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) DeactivateCoupon(wantedId int) error {
	err := service.driver.DeactivateCoupon(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) EnableFreeDelivery(wantedId int) error {
	err := service.driver.EnableFreeDelivery(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) DisableFreeDelivery(wantedId int) error {
	err := service.driver.DisableFreeDelivery(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) EnableFromProducts(wantedId int) error {
	err := service.driver.EnableFromProducts(wantedId)
	if err != nil {
		return err
	}
	return nil
}

func (service *couponService) DisableFromProducts(wantedId int) error {
	err := service.driver.DisableFromProducts(wantedId)
	if err != nil {
		return err
	}
	return nil
}
