package driver

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type FileDriver interface {
	AddFile(file entity.File) error
	// FindAllFiles() ([]entity.FileEditRequest, error)
	DeleteFile(wantedId string) error
	// EditFile(couponEditInfo entity.FileEditRequest) (entity.File, error)
	GetFileInfo(wantedId string) (entity.File, error)
	// ActivateFile(wantedId int) error
}

type fileDriver struct {
	// cacheFiles []entity.FileEditRequest
}

func NewFileDriver() FileDriver {
	return &fileDriver{}
}

// func (driver *fileDriver) FindAllCoupons() ([]entity.CouponEditRequest, error) {
// 	coupons := make([]entity.CouponEditRequest, 0)
// 	rows, err := dbConn.SQL.Query(`
// 	select
// 		id, name, code, times_used, free_delivery,
// 		active, times_used_limit, end_date, discount_percentage,
// 		discount_amount, from_products_cost
// 	from coupons`)
// 	if err != nil {
// 		return make([]entity.CouponEditRequest, 0), err
// 	}
// 	defer rows.Close()

// 	var id, timesUsed, timesUsedLimit, discountAmount int
// 	var name, code string
// 	var discountPercentage float32
// 	var freeDelivery, active, fromProductsCost bool
// 	var endDate time.Time

// 	for rows.Next() {
// 		err := rows.Scan(
// 			&id, &name, &code, &timesUsed, &freeDelivery,
// 			&active, &timesUsedLimit, &endDate, &discountPercentage,
// 			&discountAmount, &fromProductsCost)
// 		if err != nil {
// 			return make([]entity.CouponEditRequest, 0), err
// 		}
// 		coupons = append(coupons, entity.CouponEditRequest{
// 			ID:                 id,
// 			Name:               name,
// 			Code:               code,
// 			TimesUsed:          timesUsed,
// 			FreeDelivery:       freeDelivery,
// 			Active:             active,
// 			TimesUsedLimit:     timesUsedLimit,
// 			EndDate:            endDate,
// 			DiscountPercentage: discountPercentage,
// 			DiscountAmount:     discountAmount,
// 			FromProductsCost:   fromProductsCost,
// 		})
// 		if err = rows.Err(); err != nil {
// 			return make([]entity.CouponEditRequest, 0), err
// 		}
// 	}

// 	driver.cacheCoupons = coupons
// 	return coupons, nil
// }

func (driver *fileDriver) AddFile(file entity.File) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO files(
		file_name, uuid)
	VALUES ($1, $2) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt,
		file.Filename, file.UUID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("file could not be added")
	}

	// defer driver.FindAllCoupons()

	return nil
}

func (driver *fileDriver) DeleteFile(wantedUUID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	file, err := driver.GetFileInfo(wantedUUID)
	if err != nil {
		return err
	}
	filePath := filepath.Join("images", file.Filename)
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	stmt := `delete from files where uuid=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedUUID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("file could not be found")
	}
	// defer driver.FindAllCoupons()
	return nil
}

// func (driver *fileDriver) EditCoupon(couponEditInfo entity.CouponEditRequest) (entity.Coupon, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := GetCouponEditStatementString(couponEditInfo) // `UPDATE details SET name = $1 where id = $2`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt)
// 	if err != nil {
// 		return entity.Coupon{}, err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return entity.Coupon{}, errors.New("coupon could not be found")
// 	}
// 	return entity.Coupon{}, nil
// }

// func GetFileEditStatementString(couponEditInfo entity.CouponEditRequest) string {
// 	stmt := `UPDATE coupons SET`
// 	if len(couponEditInfo.Name) != 0 {
// 		stmt = stmt + ` name = '` + couponEditInfo.Name + `',`
// 	}
// 	if couponEditInfo.DiscountAmount != 0 {
// 		stmt = stmt + ` discount_amount = ` + fmt.Sprint(couponEditInfo.DiscountAmount) + `,`
// 	}
// 	if couponEditInfo.DiscountPercentage != 0 {
// 		stmt = stmt + ` discount_percentage = ` + fmt.Sprint(couponEditInfo.DiscountPercentage) + `,`
// 	}
// 	if couponEditInfo.TimesUsed != 0 {
// 		stmt = stmt + ` times_used = ` + fmt.Sprint(couponEditInfo.TimesUsed) + `,`
// 	}
// 	if couponEditInfo.TimesUsedLimit != 0 {
// 		stmt = stmt + ` times_used_limit = ` + fmt.Sprint(couponEditInfo.TimesUsedLimit) + `,`
// 	}
// 	if len(couponEditInfo.Code) != 0 {
// 		stmt = stmt + ` code = ` + fmt.Sprint(couponEditInfo.Code) + `,`
// 	}
// 	stmt = stmt[0:len(stmt)-1] + ` where id = ` + fmt.Sprint(couponEditInfo.ID) + ` returning id`
// 	fmt.Println("Edit Coupon Statement Is:", stmt)

// 	return stmt
// }

func (driver *fileDriver) GetFileInfo(wantedUUID string) (entity.File, error) {
	// fmt.Println("Address ID:", wantedId)
	if len(wantedUUID) == 0 {
		return entity.File{
			Filename: "",
			UUID:     "",
		}, errors.New("no uuid was specified")
	} else {
		// fmt.Println("Here0")
		query := `
		select 
			file_name, uuid 
		from files where uuid = $1`
		var fileName, fileUUID string
		row := dbConn.SQL.QueryRow(query, wantedUUID)
		err := row.Scan(
			&fileName, &fileUUID)
		// fmt.Println("Here1")
		if err != nil {
			fmt.Println("Error:", err)
			return entity.File{
				Filename: "",
				UUID:     "",
			}, err
		}
		// fmt.Println("Here2")
		file := entity.File{
			Filename: fileName,
			UUID:     fileUUID,
		}
		// fmt.Println("File1:", file)
		return file, nil
	}
}

// func (driver *fileDriver) ActivateCoupon(wantedId int) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := `UPDATE coupons SET active = true where id=$1 returning *`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("coupon could not be found")
// 	}
// 	defer driver.FindAllCoupons()
// 	return nil
// }

// func (driver *fileDriver) DeactivateCoupon(wantedId int) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := `UPDATE coupons SET active = false where id=$1 returning *`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("coupon could not be found")
// 	}
// 	defer driver.FindAllCoupons()
// 	return nil
// }

// func (driver *fileDriver) EnableFreeDelivery(wantedId int) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := `UPDATE coupons SET free_delivery = true where id=$1 returning *`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("coupon could not be found")
// 	}
// 	defer driver.FindAllCoupons()
// 	return nil
// }

// func (driver *fileDriver) DisableFreeDelivery(wantedId int) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := `UPDATE coupons SET free_delivery = false where id=$1 returning *`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("coupon could not be found")
// 	}
// 	defer driver.FindAllCoupons()
// 	return nil
// }

// func (driver *fileDriver) EnableFromProducts(wantedId int) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := `UPDATE coupons SET from_products_cost = true where id=$1 returning *`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("coupon could not be found")
// 	}
// 	defer driver.FindAllCoupons()
// 	return nil
// }

// func (driver *fileDriver) DisableFromProducts(wantedId int) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
// 	defer cancel()

// 	stmt := `UPDATE coupons SET from_products_cost = false where id=$1 returning *`

// 	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New("coupon could not be found")
// 	}
// 	defer driver.FindAllCoupons()
// 	return nil
// }
