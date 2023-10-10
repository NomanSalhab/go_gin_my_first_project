package driver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type DetailDriver interface {
	AddDetail(detail entity.Detail) error
	AddAddon(addon entity.Detail) error
	AddFlavor(flavor entity.Detail) error
	AddVolume(volume entity.Detail) error
	FindAllDetails() ([]entity.DetailEditRequest, error)
	FindAllAddons() ([]entity.Detail, error)
	FindAllFlavors() ([]entity.Detail, error)
	FindAllVolumes() ([]entity.Detail, error)
	DeleteDetail(wantedId int) error
	EditDetail(detailEditInfo entity.DetailEditRequest) (entity.Detail, error)
	FindProductsDetails(productDetails []int) ([]entity.DetailEditRequest, error)
	GetProductDetailsString(productDetails []int) string
	SeparateProductDetails(details []entity.DetailEditRequest) ([]entity.DetailEditRequest, []entity.DetailEditRequest, []entity.DetailEditRequest)

	// AddMockDetails(details []entity.Detail)
}

type detailDriver struct {
	cacheDetails []entity.DetailEditRequest
}

func NewDetailDriver() DetailDriver {
	return &detailDriver{}
}

// cacheDetails := make([]entity.Detail, 0)

func (driver *detailDriver) FindAllDetails() ([]entity.DetailEditRequest, error) {
	details := make([]entity.DetailEditRequest, 0)
	rows, err := dbConn.SQL.Query("select id, name, is_addon, is_flavor, is_volume, price from details")
	if err != nil {
		return make([]entity.DetailEditRequest, 0), err
	}
	defer rows.Close()

	var id, price int
	var name string
	var isAddon, isFlavor, isVolume bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &isAddon, &isFlavor, &isVolume, &price)
		if err != nil {
			return make([]entity.DetailEditRequest, 0), err
		}
		details = append(details, entity.DetailEditRequest{
			ID:       id,
			Name:     name,
			IsAddon:  isAddon,
			IsFlavor: isFlavor,
			IsVolume: isVolume,
			Price:    price,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.DetailEditRequest, 0), err
		}
	}

	driver.cacheDetails = details
	return details, nil
}

func (driver *detailDriver) FindAllAddons() ([]entity.Detail, error) {
	details := make([]entity.Detail, 0)
	rows, err := dbConn.SQL.Query("select id, name, is_addon, is_flavor, is_volume, price from details where is_addon = true")
	if err != nil {
		return make([]entity.Detail, 0), err
	}
	defer rows.Close()

	var id, price int
	var name string
	var isAddon, isFlavor, isVolume bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &isAddon, &isFlavor, &isVolume, &price)
		if err != nil {
			return make([]entity.Detail, 0), err
		}
		details = append(details, entity.Detail{
			ID:       id,
			Name:     name,
			IsAddon:  isAddon,
			IsFlavor: isFlavor,
			IsVolume: isVolume,
			Price:    price,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Detail, 0), err
		}
	}

	return details, nil
}

func (driver *detailDriver) FindAllFlavors() ([]entity.Detail, error) {
	details := make([]entity.Detail, 0)
	rows, err := dbConn.SQL.Query("select id, name, is_addon, is_flavor, is_volume, price from details where is_flavor = true")
	if err != nil {
		return make([]entity.Detail, 0), err
	}
	defer rows.Close()

	var id, price int
	var name string
	var isAddon, isFlavor, isVolume bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &isAddon, &isFlavor, &isVolume, &price)
		if err != nil {
			return make([]entity.Detail, 0), err
		}
		details = append(details, entity.Detail{
			ID:       id,
			Name:     name,
			IsAddon:  isAddon,
			IsFlavor: isFlavor,
			IsVolume: isVolume,
			Price:    price,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Detail, 0), err
		}
	}

	return details, nil
}

func (driver *detailDriver) FindAllVolumes() ([]entity.Detail, error) {
	details := make([]entity.Detail, 0)
	rows, err := dbConn.SQL.Query("select id, name, is_addon, is_flavor, is_volume, price from details where is_volume = true")
	if err != nil {
		return make([]entity.Detail, 0), err
	}
	defer rows.Close()

	var id, price int
	var name string
	var isAddon, isFlavor, isVolume bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &isAddon, &isFlavor, &isVolume, &price)
		if err != nil {
			return make([]entity.Detail, 0), err
		}
		details = append(details, entity.Detail{
			ID:       id,
			Name:     name,
			IsAddon:  isAddon,
			IsFlavor: isFlavor,
			IsVolume: isVolume,
			Price:    price,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Detail, 0), err
		}
	}

	return details, nil
}

func (driver *detailDriver) AddDetail(detail entity.Detail) error {

	if detail.IsAddon {
		err := driver.AddAddon(detail)
		if err != nil {
			return err
		}
	} else if detail.IsFlavor {
		err := driver.AddFlavor(detail)
		if err != nil {
			return err
		}
	} else if detail.IsVolume {
		err := driver.AddVolume(detail)
		if err != nil {
			return err
		}
	} else {
		return errors.New("detail type should be specified")
	}

	defer driver.FindAllDetails()

	// ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// defer cancel()
	// stmt := `INSERT INTO details(name)
	// VALUES ($1) returning *`
	// result, err := dbConn.SQL.ExecContext(ctx, stmt, detail.Name)
	// if err != nil {
	// 	return err
	// }
	// rowsAffected, _ := result.RowsAffected()
	// if rowsAffected == 0 {
	// 	return errors.New("detail could not be added")
	// }

	return nil
}

func (driver *detailDriver) AddAddon(addon entity.Detail) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO details(name, is_addon, price)
	VALUES ($1, $2, $3) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, addon.Name, addon.IsAddon, addon.Price)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("addon could not be added")
	}

	return nil
}

func (driver *detailDriver) AddFlavor(flavor entity.Detail) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO details(name, is_flavor, price)
	VALUES ($1, $2, $3) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, flavor.Name, flavor.IsFlavor, flavor.Price)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("flavor could not be added")
	}

	return nil
}

func (driver *detailDriver) AddVolume(volume entity.Detail) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO details(name, is_volume, price)
	VALUES ($1, $2, $3) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, volume.Name, volume.IsVolume, volume.Price)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("volume could not be added")
	}

	return nil
}

func (driver *detailDriver) DeleteDetail(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from details where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("detail could not be found")
	}
	defer driver.FindAllDetails()
	return nil
}

func (driver *detailDriver) EditDetail(detailEditInfo entity.DetailEditRequest) (entity.Detail, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := GetDetailEditStatementString(detailEditInfo) // `UPDATE details SET name = $1 where id = $2`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, detailEditInfo.Name, detailEditInfo.ID)
	if err != nil {
		return entity.Detail{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.Detail{}, errors.New("detail could not be found")
	}
	return entity.Detail{}, nil
}

func GetDetailEditStatementString(detailEditInfo entity.DetailEditRequest) string {
	stmt := `UPDATE details SET`
	if len(detailEditInfo.Name) != 0 {
		stmt = stmt + ` name = '` + detailEditInfo.Name + `',`
	}
	if detailEditInfo.Price != 0 {
		stmt = stmt + ` price = ` + fmt.Sprint(detailEditInfo.Price) + `,`
	}
	stmt = stmt[0:len(stmt)-1] + ` where id = ` + fmt.Sprint(detailEditInfo.ID) + ` returning id`
	fmt.Println("Edit Detail Statement Is:", stmt)

	return stmt
}

func (driver *detailDriver) FindProductsDetails(productDetails []int) ([]entity.DetailEditRequest, error) {
	if len(driver.cacheDetails) != 0 {
		wantedDetails := make([]entity.DetailEditRequest, 0)
		// fmt.Println("cached Details:", driver.cacheDetails)
		for _, item2 := range productDetails {
			for _, item := range driver.cacheDetails {
				if item.ID == item2 {
					item1 := entity.DetailEditRequest{
						ID:       item.ID,
						Name:     item.Name,
						IsFlavor: item.IsFlavor,
						IsVolume: item.IsVolume,
						IsAddon:  item.IsAddon,
						Price:    item.Price,
					}
					wantedDetails = append(wantedDetails, item1)
					break
				}
			}
		}
		// fmt.Println("wanted Details:", wantedDetails)
		return wantedDetails, nil
	} else {
		var details []entity.DetailEditRequest
		query := driver.GetProductDetailsString(productDetails)
		// fmt.Println("Query is:", query)
		rows, err := dbConn.SQL.Query(query)
		if err != nil {
			// fmt.Println("scan0 error:", err)
			return make([]entity.DetailEditRequest, 0), err
		}
		defer rows.Close()

		var id, price int
		var name string
		var isAddon, isFlavor, isVolume bool

		for rows.Next() {
			err := rows.Scan(&id, &name, &isAddon, &isFlavor, &isVolume, &price)
			// fmt.Println("running")
			if err != nil {
				// fmt.Println("scan1 error:", err)
				return make([]entity.DetailEditRequest, 0), err
			}
			details = append(details, entity.DetailEditRequest{
				ID:       id,
				Name:     name,
				IsAddon:  isAddon,
				IsFlavor: isFlavor,
				IsVolume: isVolume,
				Price:    price,
			})
			// fmt.Println(id, "-", name)
			if err = rows.Err(); err != nil {
				return make([]entity.DetailEditRequest, 0), err
			}
		}

		defer driver.FindAllDetails()
		// fmt.Println("Details Are:", details)
		return details, nil
	}

}

func (driver *detailDriver) GetProductDetailsString(productDetails []int) string {
	stmt := `select id, name, is_addon, is_flavor, is_volume, price from details where `

	for i := 0; i < len(productDetails)-1; i++ {
		stmt = stmt + `id = ` + fmt.Sprint(productDetails[i]) + ` or `
	}
	stmt = stmt + `id = ` + fmt.Sprint(productDetails[len(productDetails)-1]) + ``

	// fmt.Println("Statement is:", stmt)
	return stmt
}

func (driver *detailDriver) SeparateProductDetails(details []entity.DetailEditRequest) ([]entity.DetailEditRequest, []entity.DetailEditRequest, []entity.DetailEditRequest) {
	finalFlavors := make([]entity.DetailEditRequest, 0)
	finalVolumes := make([]entity.DetailEditRequest, 0)
	finalAddons := make([]entity.DetailEditRequest, 0)
	for _, detail := range details {
		if detail.IsFlavor {
			finalFlavors = append(finalFlavors, detail)
			continue
		}
		if detail.IsVolume {
			finalVolumes = append(finalVolumes, detail)
			continue
		}
		if detail.IsAddon {
			finalAddons = append(finalAddons, detail)
		}
	}
	return finalFlavors, finalVolumes, finalAddons
}
