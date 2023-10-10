package driver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

func FindAllSliders() ([]entity.Slider, error) {
	sliders := make([]entity.Slider, 0)
	rows, err := dbConn.SQL.Query("select id, image, store_id, product_id, active from sliders")
	if err != nil {
		return make([]entity.Slider, 0), err
	}
	defer rows.Close()

	var id, storeId, productId int
	var image string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &image, &storeId, &productId, &active)
		if err != nil {
			return make([]entity.Slider, 0), err
		}
		sliders = append(sliders, entity.Slider{
			ID:        id,
			Image:     image,
			StoreId:   storeId,
			ProductId: productId,
			Active:    active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Slider, 0), err
		}
	}

	return sliders, nil
}

func FindActiveSliders() ([]entity.Slider, error) {
	sliders := make([]entity.Slider, 0)
	rows, err := dbConn.SQL.Query("select id, image, store_id, product_id, active from sliders where active = true")
	if err != nil {
		return make([]entity.Slider, 0), err
	}
	defer rows.Close()

	var id, storeId, productId int
	var image string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &image, &storeId, &productId, &active)
		if err != nil {
			return make([]entity.Slider, 0), err
		}
		sliders = append(sliders, entity.Slider{
			ID:        id,
			Image:     image,
			StoreId:   storeId,
			ProductId: productId,
			Active:    active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Slider, 0), err
		}
	}

	return sliders, nil
}

func FindNotActiveSliders() ([]entity.Slider, error) {
	sliders := make([]entity.Slider, 0)
	rows, err := dbConn.SQL.Query("select id, image, store_id, product_id, active from sliders where active = false")
	if err != nil {
		return make([]entity.Slider, 0), err
	}
	defer rows.Close()

	var id, storeId, productId int
	var image string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &image, &storeId, &productId, &active)
		if err != nil {
			return make([]entity.Slider, 0), err
		}
		sliders = append(sliders, entity.Slider{
			ID:        id,
			Image:     image,
			StoreId:   storeId,
			ProductId: productId,
			Active:    active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Slider, 0), err
		}
	}

	return sliders, nil
}

func FindSlidersByStore(wantedId int) ([]entity.Slider, error) {
	sliders := make([]entity.Slider, 0)
	rows, err := dbConn.SQL.Query("select id, image, store_id, product_id, active from sliders where store_id = $1", wantedId)
	if err != nil {
		return make([]entity.Slider, 0), err
	}
	defer rows.Close()

	var id, storeId, productId int
	var image string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &image, &storeId, &productId, &active)
		if err != nil {
			return make([]entity.Slider, 0), err
		}
		sliders = append(sliders, entity.Slider{
			ID:        id,
			Image:     image,
			StoreId:   storeId,
			ProductId: productId,
			Active:    active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Slider, 0), err
		}
	}

	return sliders, nil
}

func AddSlider(slider entity.Slider) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO sliders(image, store_id, product_id, active)
	VALUES ($1, $2, $3, $4) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, slider.Image, slider.StoreId, slider.ProductId, slider.Active)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("slider could not be added")
	}

	return nil
}

func DeleteSlider(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from sliders where id = $1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("slider could not be found")
	}

	return nil
}

func EditSlider(sliderEditInfo entity.SliderEditRequest) (entity.Slider, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := GetEditSliderStatementString(sliderEditInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return entity.Slider{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.Slider{}, errors.New("slider could not be found")
	}
	// slider, err := FindSlider(sliderEditInfo.ID)
	// if err != nil {
	// 	return entity.Slider{}, err
	// }
	return entity.Slider{}, nil
}

func GetEditSliderStatementString(sliderEditInfo entity.SliderEditRequest) string {
	stmt := `UPDATE sliders SET `
	if sliderEditInfo.StoreId != 0 {
		stmt = stmt + `store_id = ` + fmt.Sprint(sliderEditInfo.StoreId) + `, `
	}
	if sliderEditInfo.ProductId != 0 {
		stmt = stmt + `product_id = ` + fmt.Sprint(sliderEditInfo.ProductId) + `, `
	}
	if sliderEditInfo.Active {
		stmt = stmt + `active = true `
	} else {
		stmt = stmt + `active = false `
	}
	stmt = stmt + `where id = ` + fmt.Sprint(sliderEditInfo.ID) + ` RETURNING *`
	return stmt
}
