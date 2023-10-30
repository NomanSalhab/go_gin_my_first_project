package driver

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type StoreDriver interface {
	FindAllStores() ([]entity.Store, error)
	FindActiveStores(wantedId int) ([]entity.Store, error)
	FindNotActiveStores() ([]entity.Store, error)
	FindStoreCategoryStores(wantedId int) ([]entity.Store, error)
	FindStore(wantedId int) (entity.Store, error)
	FindBestSellingStores(storesCountLimit int, wantedAreaId int) ([]entity.Store, error)
	AddStore(store entity.Store) error
	DeleteStore(wantedId int) error
	EditStore(storeEditInfo entity.StoreEditRequest) (entity.Store, error)
	GetEditStoreStatementString(storeEditInfo entity.StoreEditRequest) string
	ActivateStore(storeEditInfo entity.StoreInfoRequest) error
	DeactivateStore(storeEditInfo entity.StoreInfoRequest) error
	EditStoreBalance(storeInfo entity.StoreIncreaseBalance) error
}

type storeDriver struct {
	// cachedUsers []entity.User
}

func NewStoreDriver() StoreDriver {
	return &storeDriver{}
}

func (driver *storeDriver) FindAllStores() ([]entity.Store, error) {
	stores := make([]entity.Store, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_category_id, image, balance, active, delivery_rent, discount, area_id from stores")
	if err != nil {
		return make([]entity.Store, 0), err
	}
	defer rows.Close()

	var id, storeCategoryId, areaId, deliveryRent, balance int
	var name, image string
	var discount float32
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeCategoryId, &image, &balance, &active, &deliveryRent, &discount, &areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.Store, 0), err
		}
		storeDeliveryRent := deliveryRent - int(discount*float32(deliveryRent))
		// imageFile, err := fd.GetFileInfo(image)
		// if err != nil {
		// 	return make([]entity.Store, 0), err
		// }

		stores = append(stores, entity.Store{
			ID:              id,
			Name:            name,
			StoreCategoryId: storeCategoryId,
			Image:           image,
			Balance:         balance,
			Active:          active,
			Discount:        discount,
			DeliveryRent:    storeDeliveryRent,
			AreaID:          areaId,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Store, 0), err
		}
		// fmt.Println("------------------------")
	}

	return stores, nil
}

func (driver *storeDriver) FindActiveStores(wantedId int) ([]entity.Store, error) {
	stores := make([]entity.Store, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_category_id, image, balance, active, delivery_rent, discount, area_id from stores where active = true and area_id = $1", wantedId)
	if err != nil {
		return make([]entity.Store, 0), err
	}
	defer rows.Close()

	var id, storeCategoryId, areaId, deliveryRent, balance int
	var name, image string
	var discount float32
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeCategoryId, &image, &balance, &active, &deliveryRent, &discount, &areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.Store, 0), err
		}
		storeDeliveryRent := deliveryRent - int(discount*float32(deliveryRent))
		// imageFile, err := fd.GetFileInfo(image)
		// if err != nil {
		// 	return make([]entity.Store, 0), err
		// }

		stores = append(stores, entity.Store{
			ID:              id,
			Name:            name,
			StoreCategoryId: storeCategoryId,
			Image:           image,
			Balance:         balance,
			Active:          active,
			Discount:        discount,
			DeliveryRent:    storeDeliveryRent,
			AreaID:          areaId,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Store, 0), err
		}
		// fmt.Println("------------------------")
	}

	return stores, nil
}

func (driver *storeDriver) FindNotActiveStores() ([]entity.Store, error) {
	stores := make([]entity.Store, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_category_id, image, balance, active, delivery_rent, discount, area_id from stores where active = false")
	if err != nil {
		return make([]entity.Store, 0), err
	}
	defer rows.Close()

	var id, storeCategoryId, areaId, deliveryRent, balance int
	var name, image string
	var discount float32
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeCategoryId, &image, &balance, &active, &deliveryRent, &discount, &areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.Store, 0), err
		}
		storeDeliveryRent := deliveryRent - int(discount*float32(deliveryRent))
		// imageFile, err := fd.GetFileInfo(image)
		// if err != nil {
		// 	return make([]entity.Store, 0), err
		// }
		stores = append(stores, entity.Store{
			ID:              id,
			Name:            name,
			StoreCategoryId: storeCategoryId,
			Image:           image,
			Balance:         balance,
			Active:          active,
			Discount:        discount,
			DeliveryRent:    storeDeliveryRent,
			AreaID:          areaId,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Store, 0), err
		}
		// fmt.Println("------------------------")
	}

	return stores, nil
}

func (driver *storeDriver) FindStoreCategoryStores(wantedId int) ([]entity.Store, error) {
	stores := make([]entity.Store, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_category_id, image, balance, active, delivery_rent, discount, area_id from stores where active = true and store_category_id = $1", wantedId)
	if err != nil {
		return make([]entity.Store, 0), err
	}
	defer rows.Close()

	var id, storeCategoryId, areaId, deliveryRent, balance int
	var name, image string
	var discount float32
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeCategoryId, &image, &balance, &active, &deliveryRent, &discount, &areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.Store, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)
		// if err != nil {
		// 	return make([]entity.Store, 0), err
		// }
		stores = append(stores, entity.Store{
			ID:              id,
			Name:            name,
			StoreCategoryId: storeCategoryId,
			Image:           image,
			Balance:         balance,
			Active:          active,
			Discount:        discount,
			DeliveryRent:    deliveryRent,
			AreaID:          areaId,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Store, 0), err
		}
		// fmt.Println("------------------------")
	}

	return stores, nil
}

func (driver *storeDriver) FindStore(wantedId int) (entity.Store, error) {

	query := `select id, name, store_category_id, image, balance, active, delivery_rent, discount, area_id from stores where id = $1`
	var id, storeCategoryId, areaId, deliveryRent, balance int
	var name, image string
	var discount float32
	var active bool
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &storeCategoryId, &image, &balance, &active, &deliveryRent, &discount, &areaId)
	if err != nil {
		return entity.Store{
			ID:     0,
			Name:   "",
			Active: false,
		}, err
	}
	storeDeliveryRent := deliveryRent - int(discount*float32(deliveryRent))
	// imageFile, err := fd.GetFileInfo(image)
	// if err != nil {
	// 	return entity.Store{}, err
	// }

	store := entity.Store{
		ID:              id,
		Name:            name,
		StoreCategoryId: storeCategoryId,
		Image:           image,
		Balance:         balance,
		Active:          active,
		Discount:        discount,
		DeliveryRent:    storeDeliveryRent,
		AreaID:          areaId,
	}
	return store, nil
}

func (driver *storeDriver) AddStore(store entity.Store) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO stores(name, store_category_id, image, balance, active, delivery_rent, discount, area_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, store.Name, store.StoreCategoryId, store.Image, store.Balance, store.Active, store.DeliveryRent, store.Discount, store.AreaID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store could not be added")
	}

	return nil
}

func (driver *storeDriver) DeleteStore(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from stores where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store could not be found")
	}

	return nil
}

func (driver *storeDriver) EditStore(storeEditInfo entity.StoreEditRequest) (entity.Store, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := driver.GetEditStoreStatementString(storeEditInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeEditInfo.ID)
	if err != nil {
		return entity.Store{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.Store{}, errors.New("store could not be found")
	}
	store, _ := driver.FindStore(storeEditInfo.ID)
	// if err != nil {
	// 	return entity.Store{}, err
	// }
	return store, nil
}

func (driver *storeDriver) GetEditStoreStatementString(storeEditInfo entity.StoreEditRequest) string {
	stmt := `UPDATE stores SET `
	if storeEditInfo.Name != "" {
		stmt = stmt + `name = '` + storeEditInfo.Name + `', `
	}
	if storeEditInfo.StoreCategoryId != 0 {
		stmt = stmt + `store_category_id = ` + fmt.Sprint(storeEditInfo.StoreCategoryId) + `, `
	}
	if storeEditInfo.Image != "" {
		stmt = stmt + `image = '` + storeEditInfo.Image + `', `
	}
	if storeEditInfo.Balance != 0 {
		stmt = stmt + `balance = ` + fmt.Sprint(storeEditInfo.Balance) + `, `
	}
	if storeEditInfo.DeliveryRent != 0 {
		stmt = stmt + `delivery_rent = ` + fmt.Sprint(storeEditInfo.DeliveryRent) + `, `
	}
	if storeEditInfo.Discount != 0 {
		stmt = stmt + `discount = '` + fmt.Sprint(storeEditInfo.Discount) + `', `
	}
	if storeEditInfo.AreaID != 0 {
		stmt = stmt + `area_id = ` + fmt.Sprint(storeEditInfo.AreaID) + ` ` // ,
	}
	// if storeEditInfo.Active {
	// 	stmt = stmt + `active = true `
	// } else {
	// 	stmt = stmt + `active = false `
	// }
	stmt = stmt + `where id = ` + /*fmt.Sprint(storeEditInfo.ID)  */ `$1` + ` RETURNING *`
	return stmt
}

func (driver *storeDriver) ActivateStore(storeEditInfo entity.StoreInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE stores SET active = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store could not be found")
	}
	return nil
}

func (driver *storeDriver) DeactivateStore(storeEditInfo entity.StoreInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE stores SET active = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store could not be found")
	}
	return nil
}

func (driver *storeDriver) EditStoreBalance(storeInfo entity.StoreIncreaseBalance) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE stores SET balance = balance + $1 WHERE id = $2 RETURNING *`
	balanceIncrease := int(math.Floor(float64(storeInfo.Balance) / float64(1000)))

	result, err := dbConn.SQL.ExecContext(ctx, stmt, balanceIncrease, storeInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store could not be found")
	}
	return nil
}

func (driver *storeDriver) FindBestSellingStores(storesCountLimit int, wantedAreaId int) ([]entity.Store, error) {
	stores := make([]entity.Store, 0)
	rows, err := dbConn.SQL.Query(`select 
	id, name, store_category_id, 
	image, balance, active, 
	delivery_rent, discount, area_id 
	from stores where area_id = $1 order by balance desc limit $2`, wantedAreaId, storesCountLimit)
	if err != nil {
		return make([]entity.Store, 0), err
	}
	defer rows.Close()

	var id, storeCategoryId, areaId, deliveryRent, balance int
	var name, image string
	var discount float32
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeCategoryId, &image, &balance, &active, &deliveryRent, &discount, &areaId)
		if err != nil {
			// log.Println(err)
			return make([]entity.Store, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)
		// if err != nil {
		// 	return make([]entity.Store, 0), err
		// }
		stores = append(stores, entity.Store{
			ID:              id,
			Name:            name,
			StoreCategoryId: storeCategoryId,
			Image:           image,
			Balance:         balance,
			Active:          active,
			Discount:        discount,
			DeliveryRent:    deliveryRent,
			AreaID:          areaId,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Store, 0), err
		}
		// fmt.Println("------------------------")
	}

	return stores, nil
}
