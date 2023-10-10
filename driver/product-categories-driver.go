package driver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

func FindAllProductCategories() ([]entity.ProductCategory, error) {
	productCategories := make([]entity.ProductCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_id, active from product_categories")
	if err != nil {
		return make([]entity.ProductCategory, 0), err
	}
	defer rows.Close()

	var id, storeId int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &active)
		if err != nil {
			return make([]entity.ProductCategory, 0), err
		}
		productCategories = append(productCategories, entity.ProductCategory{
			ID:      id,
			Name:    name,
			StoreId: storeId,
			Active:  active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.ProductCategory, 0), err
		}
	}

	return productCategories, nil
}

func FindActiveProductCategories() ([]entity.ProductCategory, error) {
	productCategories := make([]entity.ProductCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_id, active from product_categories where active = true")
	if err != nil {
		return make([]entity.ProductCategory, 0), err
	}
	defer rows.Close()

	var id, storeId int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &active)
		if err != nil {
			return make([]entity.ProductCategory, 0), err
		}
		productCategories = append(productCategories, entity.ProductCategory{
			ID:      id,
			Name:    name,
			StoreId: storeId,
			Active:  active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.ProductCategory, 0), err
		}
	}

	return productCategories, nil
}

func FindNotActiveProductCategories() ([]entity.ProductCategory, error) {
	productCategories := make([]entity.ProductCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_id, active from product_categories where active = false")
	if err != nil {
		return make([]entity.ProductCategory, 0), err
	}
	defer rows.Close()

	var id, storeId int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &active)
		if err != nil {
			return make([]entity.ProductCategory, 0), err
		}
		productCategories = append(productCategories, entity.ProductCategory{
			ID:      id,
			Name:    name,
			StoreId: storeId,
			Active:  active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.ProductCategory, 0), err
		}
	}

	return productCategories, nil
}

func FindProductCategoryByStore(wantedId int) ([]entity.ProductCategory, error) {
	productCategories := make([]entity.ProductCategory, 0)
	rows, err := dbConn.SQL.Query("select id, name, store_id, active from product_categories where store_id = $1", wantedId)
	if err != nil {
		return make([]entity.ProductCategory, 0), err
	}
	defer rows.Close()

	var id, storeId int
	var name string
	var active bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &active)
		if err != nil {
			return make([]entity.ProductCategory, 0), err
		}
		productCategories = append(productCategories, entity.ProductCategory{
			ID:      id,
			Name:    name,
			StoreId: storeId,
			Active:  active,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.ProductCategory, 0), err
		}
	}

	return productCategories, nil
}

func FindProductCategory(wantedId int) (entity.ProductCategory, error) {

	query := `select id, name, store_id, active from product_categories where id = $1`
	var id, storeId int
	var name string
	var active bool
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &storeId, &active)
	if err != nil {
		return entity.ProductCategory{
			ID:      0,
			Name:    "",
			StoreId: 0,
			Active:  false,
		}, err
	}
	user := entity.ProductCategory{
		ID:      id,
		Name:    name,
		StoreId: storeId,
		Active:  active,
	}
	return user, nil
}

func AddProductCategory(productCategory entity.ProductCategory) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO product_categories(name, store_id, active)
	VALUES ($1, $2, $3) returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, productCategory.Name, productCategory.StoreId, productCategory.Active)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product category could not be added")
	}

	return nil
}

func DeleteProductCategory(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from product_categories where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product category could not be found")
	}

	return nil
}

func EditProductCategory(productCategoryEditInfo entity.ProductCategoryEditRequest) (entity.ProductCategory, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := GetEditProductCategoryStatementString(productCategoryEditInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return entity.ProductCategory{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.ProductCategory{}, errors.New("product category could not be found")
	}
	productCategory, err := FindProductCategory(productCategoryEditInfo.ID)
	if err != nil {
		return entity.ProductCategory{}, err
	}
	return productCategory, nil
}

func GetEditProductCategoryStatementString(productEditInfo entity.ProductCategoryEditRequest) string {
	stmt := `UPDATE product_categories SET `
	if productEditInfo.Name != "" {
		stmt = stmt + `name = '` + productEditInfo.Name + `', `
	}
	if productEditInfo.StoreId != 0 {
		stmt = stmt + `store_id = ` + fmt.Sprint(productEditInfo.StoreId) + `, `
	}
	if productEditInfo.Active {
		stmt = stmt + `active = true `
	} else {
		stmt = stmt + `active = false `
	}
	stmt = stmt + `where id = ` + fmt.Sprint(productEditInfo.ID) + ` RETURNING *`
	return stmt
}

func ActivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE product_categories SET active = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, productCategoryEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product category could not be found")
	}
	return nil
}

func DeactivateProductCategory(productCategoryEditInfo entity.ProductCategoryInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE product_categories SET active = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, productCategoryEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product category could not be found")
	}
	return nil
}
