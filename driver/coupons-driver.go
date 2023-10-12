package driver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type CouponDriver interface {
	AddCoupon(coupon entity.Coupon) error
	FindAllCoupons() ([]entity.CouponEditRequest, error)
	DeleteCoupon(wantedId int) error
	EditCoupon(couponEditInfo entity.CouponEditRequest) (entity.Coupon, error)
	GetCouponInfo(wantedId int) (entity.CouponAddOrderRequest, error)
	ActivateCoupon(wantedId int) error
	DeactivateCoupon(wantedId int) error
	EnableFreeDelivery(wantedId int) error
	DisableFreeDelivery(wantedId int) error
	EnableFromProducts(wantedId int) error
	DisableFromProducts(wantedId int) error
}

type couponDriver struct {
	cacheCoupons []entity.CouponEditRequest
}

func NewCouponDriver() CouponDriver {
	return &couponDriver{}
}

func (driver *couponDriver) FindAllCoupons() ([]entity.CouponEditRequest, error) {
	coupons := make([]entity.CouponEditRequest, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, name, code, times_used, free_delivery, 
		active, times_used_limit, end_date, discount_percentage, 
		discount_amount, from_products_cost
	from coupons`)
	if err != nil {
		return make([]entity.CouponEditRequest, 0), err
	}
	defer rows.Close()

	var id, timesUsed, timesUsedLimit, discountAmount int
	var name, code string
	var discountPercentage float32
	var freeDelivery, active, fromProductsCost bool
	var endDate time.Time

	for rows.Next() {
		err := rows.Scan(
			&id, &name, &code, &timesUsed, &freeDelivery,
			&active, &timesUsedLimit, &endDate, &discountPercentage,
			&discountAmount, &fromProductsCost)
		if err != nil {
			return make([]entity.CouponEditRequest, 0), err
		}
		coupons = append(coupons, entity.CouponEditRequest{
			ID:                 id,
			Name:               name,
			Code:               code,
			TimesUsed:          timesUsed,
			FreeDelivery:       freeDelivery,
			Active:             active,
			TimesUsedLimit:     timesUsedLimit,
			EndDate:            endDate,
			DiscountPercentage: discountPercentage,
			DiscountAmount:     discountAmount,
			FromProductsCost:   fromProductsCost,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.CouponEditRequest, 0), err
		}
	}

	driver.cacheCoupons = coupons
	return coupons, nil
}

func (driver *couponDriver) AddCoupon(coupon entity.Coupon) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO coupons(
		name, code, times_used, free_delivery, 
		active, times_used_limit, end_date, discount_percentage, 
		discount_amount, from_products_cost)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt,
		coupon.Name, coupon.Code, coupon.TimesUsed, coupon.FreeDelivery,
		coupon.Active, coupon.TimesUsedLimit, coupon.EndDate, coupon.DiscountPercentage,
		coupon.DiscountAmount, coupon.FromProductsCost)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be added")
	}

	defer driver.FindAllCoupons()

	return nil
}

func (driver *couponDriver) DeleteCoupon(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from coupons where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}

func (driver *couponDriver) EditCoupon(couponEditInfo entity.CouponEditRequest) (entity.Coupon, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := GetCouponEditStatementString(couponEditInfo) // `UPDATE details SET name = $1 where id = $2`

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return entity.Coupon{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.Coupon{}, errors.New("coupon could not be found")
	}
	return entity.Coupon{}, nil
}

func GetCouponEditStatementString(couponEditInfo entity.CouponEditRequest) string {
	stmt := `UPDATE coupons SET`
	if len(couponEditInfo.Name) != 0 {
		stmt = stmt + ` name = '` + couponEditInfo.Name + `',`
	}
	if couponEditInfo.DiscountAmount != 0 {
		stmt = stmt + ` discount_amount = ` + fmt.Sprint(couponEditInfo.DiscountAmount) + `,`
	}
	if couponEditInfo.DiscountPercentage != 0 {
		stmt = stmt + ` discount_percentage = ` + fmt.Sprint(couponEditInfo.DiscountPercentage) + `,`
	}
	if couponEditInfo.TimesUsed != 0 {
		stmt = stmt + ` times_used = ` + fmt.Sprint(couponEditInfo.TimesUsed) + `,`
	}
	if couponEditInfo.TimesUsedLimit != 0 {
		stmt = stmt + ` times_used_limit = ` + fmt.Sprint(couponEditInfo.TimesUsedLimit) + `,`
	}
	if len(couponEditInfo.Code) != 0 {
		stmt = stmt + ` code = ` + fmt.Sprint(couponEditInfo.Code) + `,`
	}
	stmt = stmt[0:len(stmt)-1] + ` where id = ` + fmt.Sprint(couponEditInfo.ID) + ` returning id`
	fmt.Println("Edit Coupon Statement Is:", stmt)

	return stmt
}

func (driver *couponDriver) GetCouponInfo(wantedId int) (entity.CouponAddOrderRequest, error) {
	// fmt.Println("Address ID:", wantedId)
	if wantedId == 0 {
		return entity.CouponAddOrderRequest{
			ID:                 0,
			Name:               "",
			Code:               "",
			TimesUsed:          0,
			FreeDelivery:       false,
			Active:             false,
			TimesUsedLimit:     0,
			EndDate:            time.Now(),
			DiscountPercentage: 0.0,
			DiscountAmount:     0,
			FromProductsCost:   false,
		}, nil
	} else {
		query := `
		select 
			id, name, code, times_used, free_delivery, 
			active, times_used_limit, end_date, discount_percentage, 
			discount_amount, from_products_cost 
		from coupons where id = $1`
		var id, timesUsed, timesUsedLimit, discountAmount int
		var name, code string
		var discountPercentage float32
		var freeDelivery, active, fromProductsCost bool
		var endDate time.Time
		row := dbConn.SQL.QueryRow(query, wantedId)
		err := row.Scan(
			&id, &name, &code, &timesUsed, &freeDelivery,
			&active, &timesUsedLimit, &endDate, &discountPercentage,
			&discountAmount, &fromProductsCost)
		if err != nil {
			return entity.CouponAddOrderRequest{
				ID:                 0,
				Name:               "",
				Code:               "",
				TimesUsed:          0,
				FreeDelivery:       false,
				Active:             false,
				TimesUsedLimit:     0,
				EndDate:            time.Now(),
				DiscountPercentage: 0.0,
				DiscountAmount:     0,
				FromProductsCost:   false,
			}, err
		}
		coupon := entity.CouponAddOrderRequest{
			ID:                 id,
			Name:               name,
			Code:               code,
			TimesUsed:          timesUsed,
			FreeDelivery:       freeDelivery,
			Active:             active,
			TimesUsedLimit:     timesUsedLimit,
			EndDate:            endDate,
			DiscountPercentage: discountPercentage,
			DiscountAmount:     discountAmount,
			FromProductsCost:   fromProductsCost,
		}
		return coupon, nil
	}
}

func (driver *couponDriver) ActivateCoupon(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE coupons SET active = true where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}

func (driver *couponDriver) DeactivateCoupon(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE coupons SET active = false where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}

func (driver *couponDriver) EnableFreeDelivery(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE coupons SET free_delivery = true where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}

func (driver *couponDriver) DisableFreeDelivery(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE coupons SET free_delivery = false where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}

func (driver *couponDriver) EnableFromProducts(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE coupons SET from_products_cost = true where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}

func (driver *couponDriver) DisableFromProducts(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE coupons SET from_products_cost = false where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("coupon could not be found")
	}
	defer driver.FindAllCoupons()
	return nil
}
