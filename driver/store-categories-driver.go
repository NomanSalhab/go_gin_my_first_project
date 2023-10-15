package driver

import (
	"context"
	"errors"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type StoreCategoryDriver interface {
	FindAllStoreCategories() ([]entity.StoreCategory, error)
	FindActiveStoreCategories() ([]entity.StoreCategory, error)
	FindNotActiveStoreCategories() ([]entity.StoreCategory, error)
	FindStoreCategory(wantedId int) (entity.StoreCategory, error)
	AddStoreCategory(storeCategory entity.StoreCategory) error
	DeleteStoreCategory(wantedId int) error
	EditStoreCategory(storeCategoryEditInfo entity.StoreCategoryEditRequest) (entity.StoreCategory, error)
	ActivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error
	DeactivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error
}

type storeCategoryDriver struct {
}

func NewStoreCategoryDriver() StoreCategoryDriver {
	return &storeCategoryDriver{}
}

func (driver *storeCategoryDriver) FindAllStoreCategories() ([]entity.StoreCategory, error) {
	storeCategories := make([]entity.StoreCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, active from store_categories")
	if err != nil {
		return make([]entity.StoreCategory, 0), err
	}
	defer rows.Close()

	var id int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &active)
		if err != nil {
			// log.Println(err)
			return make([]entity.StoreCategory, 0), err
		}
		storeCategories = append(storeCategories, entity.StoreCategory{
			ID:     id,
			Name:   name,
			Active: active,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.StoreCategory, 0), err
		}
		// fmt.Println("------------------------")
	}

	return storeCategories, nil
}

func (driver *storeCategoryDriver) FindActiveStoreCategories() ([]entity.StoreCategory, error) {
	storeCategories := make([]entity.StoreCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, active from store_categories where active = true")
	if err != nil {
		return make([]entity.StoreCategory, 0), err
	}
	defer rows.Close()

	var id int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &active)
		if err != nil {
			// log.Println(err)
			return make([]entity.StoreCategory, 0), err
		}
		storeCategories = append(storeCategories, entity.StoreCategory{
			ID:     id,
			Name:   name,
			Active: active,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.StoreCategory, 0), err
		}
		// fmt.Println("------------------------")
	}

	return storeCategories, nil
}

func (driver *storeCategoryDriver) FindNotActiveStoreCategories() ([]entity.StoreCategory, error) {
	storeCategories := make([]entity.StoreCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, active from store_categories where active = false")
	if err != nil {
		return make([]entity.StoreCategory, 0), err
	}
	defer rows.Close()

	var id int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &active)
		if err != nil {
			// log.Println(err)
			return make([]entity.StoreCategory, 0), err
		}
		storeCategories = append(storeCategories, entity.StoreCategory{
			ID:     id,
			Name:   name,
			Active: active,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.StoreCategory, 0), err
		}
		// fmt.Println("------------------------")
	}

	return storeCategories, nil
}

func (driver *storeCategoryDriver) FindStoreCategory(wantedId int) (entity.StoreCategory, error) {

	query := `select id, name, active from store_categories where id = $1`
	var id int
	var name string
	var active bool
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &active)
	if err != nil {
		return entity.StoreCategory{
			ID:     0,
			Name:   "",
			Active: false,
		}, err
	}
	user := entity.StoreCategory{
		ID:     id,
		Name:   name,
		Active: active,
	}
	return user, nil
}

func (driver *storeCategoryDriver) AddStoreCategory(storeCategory entity.StoreCategory) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO store_categories(name, active)
	VALUES ($1, $2) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeCategory.Name, storeCategory.Active)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store category could not be added")
	}

	return nil
}

func (driver *storeCategoryDriver) DeleteStoreCategory(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from store_categories where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store category could not be found")
	}

	return nil
}

func (driver *storeCategoryDriver) EditStoreCategory(storeCategoryEditInfo entity.StoreCategoryEditRequest) (entity.StoreCategory, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE store_categories SET name = $1 WHERE id = $2 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeCategoryEditInfo.Name, storeCategoryEditInfo.ID)
	if err != nil {
		return entity.StoreCategory{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.StoreCategory{}, errors.New("store category could not be found")
	}
	storeCategory, err := driver.FindStoreCategory(storeCategoryEditInfo.ID)
	if err != nil {
		return entity.StoreCategory{}, err
	}
	return storeCategory, nil
}

func (driver *storeCategoryDriver) ActivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE store_categories SET active = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeCategoryEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store category could not be found")
	}
	return nil
}

func (driver *storeCategoryDriver) DeactivateStoreCategory(storeCategoryEditInfo entity.StoreCategoryInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE store_categories SET active = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, storeCategoryEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("store category could not be found")
	}
	return nil
}
