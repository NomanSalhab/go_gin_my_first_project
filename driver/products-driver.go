package driver

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/lib/pq"
)

type ProductDriver interface {
	FindAllProducts() ([]entity.Product, error)
	FindActiveProducts() ([]entity.Product, error)
	FindNotActiveProducts() ([]entity.Product, error)
	FindProductByProductCategory(wantedProductStoreId int, wantedProductCategoryId int) ([]entity.Product, error)
	FindProduct(wantedId int) (entity.Product, error)
	FindBestSellingProducts(productsCountLimit int) ([]entity.Product, error)
	FindOffersProducts() ([]entity.Product, error)
	GetIntSliceFromString(theString []string) []int
	AddProduct(product entity.Product) error
	detailsSliceToIdSlice(details []entity.DetailEditRequest) []int
	sliceToString(theSlice []int) string
	DeleteProduct(wantedId int) error
	ActivateProduct(productEditInfo entity.ProductInfoRequest) error
	DeactivateProduct(productEditInfo entity.ProductInfoRequest) error
	EditProduct(productEditInfo entity.ProductEditRequest) (entity.Product, error)
	GetEditProductStatementString(productEditInfo entity.ProductEditRequest) string
	EditProductsSliceOrderCount(productInfo []entity.OrderProduct) error
	EditProductOrderCount(productInfo entity.OrderProduct) error

	// AddMockDetails(details []entity.Detail)
}

type productDriver struct {
	cacheProducts []entity.Product
}

func NewProductDriver() ProductDriver {
	return &productDriver{}
}

// ! Fix Details after details driver is finished
func (driver *productDriver) FindAllProducts() ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	rows, err := dbConn.SQL.Query(`
		select id, name, store_id, product_category_id, 
			image, summary, price, order_count, active, discount_ratio
		from products`)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	defer rows.Close()

	var id, storeId, productCategoryId, price, orderCount int
	var name, image, summary string
	var active bool
	var discountRatio float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &active, &discountRatio)
		if err != nil {
			return make([]entity.Product, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)

		// if err != nil {
		// 	return make([]entity.Product, 0), err
		// }
		products = append(products, entity.Product{
			ID:                id,
			Name:              name,
			StoreId:           storeId,
			ProductCategoryId: productCategoryId,
			Image:             image,
			Summary:           summary,
			Price:             price,
			OrderCount:        orderCount,
			Active:            active,
			Flavors:           make([]entity.DetailEditRequest, 0),
			Volumes:           make([]entity.DetailEditRequest, 0),
			Addons:            make([]entity.DetailEditRequest, 0),
			DiscountRatio:     discountRatio,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Product, 0), err
		}
	}

	driver.cacheProducts = products
	return products, nil
}

func (driver *productDriver) FindActiveProducts() ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	rows, err := dbConn.SQL.Query(`
		select 
			id, name, store_id, product_category_id, 
			image, summary, price, order_count, active, discount_ratio 
		from products where active = true
	`)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	defer rows.Close()

	var id, storeId, productCategoryId, price, orderCount int
	var name, image, summary string
	var active bool
	var discountRatio float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &active, &discountRatio)
		if err != nil {
			return make([]entity.Product, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)

		// if err != nil {
		// 	return make([]entity.Product, 0), err
		// }
		products = append(products, entity.Product{
			ID:                id,
			Name:              name,
			StoreId:           storeId,
			ProductCategoryId: productCategoryId,
			Image:             image,
			Summary:           summary,
			Price:             price,
			OrderCount:        orderCount,
			Active:            active,
			Flavors:           make([]entity.DetailEditRequest, 0),
			Volumes:           make([]entity.DetailEditRequest, 0),
			Addons:            make([]entity.DetailEditRequest, 0),
			DiscountRatio:     discountRatio,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Product, 0), err
		}
	}

	return products, nil
}

func (driver *productDriver) FindNotActiveProducts() ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	rows, err := dbConn.SQL.Query(`
		select id, name, store_id, product_category_id, 
			image, summary, price, order_count, active, discount_ratio 
		from products where active = false`)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	defer rows.Close()

	var id, storeId, productCategoryId, price, orderCount int
	var name, image, summary string
	var active bool
	var discountRatio float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &active, &discountRatio)
		if err != nil {
			return make([]entity.Product, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)

		// if err != nil {
		// 	return make([]entity.Product, 0), err
		// }
		products = append(products, entity.Product{
			ID:                id,
			Name:              name,
			StoreId:           storeId,
			ProductCategoryId: productCategoryId,
			Image:             image,
			Summary:           summary,
			Price:             price,
			OrderCount:        orderCount,
			Active:            active,
			Flavors:           make([]entity.DetailEditRequest, 0),
			Volumes:           make([]entity.DetailEditRequest, 0),
			Addons:            make([]entity.DetailEditRequest, 0),
			DiscountRatio:     discountRatio,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Product, 0), err
		}
	}

	return products, nil
}

func (driver *productDriver) FindProductByProductCategory(wantedProductStoreId int, wantedProductCategoryId int) ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	rows, err := dbConn.SQL.Query(`
		select id, name, store_id, product_category_id, 
			image, summary, price, order_count, active, discount_ratio 
		from products where store_id = $1 and product_category_id = $2`, wantedProductStoreId, wantedProductCategoryId)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	defer rows.Close()

	var id, storeId, productCategoryId, price, orderCount int
	var name, image, summary string
	var active bool
	var discountRatio float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &active, &discountRatio)
		if err != nil {
			return make([]entity.Product, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)

		// if err != nil {
		// 	return make([]entity.Product, 0), err
		// }
		products = append(products, entity.Product{
			ID:                id,
			Name:              name,
			StoreId:           storeId,
			ProductCategoryId: productCategoryId,
			Image:             image,
			Summary:           summary,
			Price:             price,
			OrderCount:        orderCount,
			Active:            active,
			Flavors:           make([]entity.DetailEditRequest, 0),
			Volumes:           make([]entity.DetailEditRequest, 0),
			Addons:            make([]entity.DetailEditRequest, 0),
			DiscountRatio:     discountRatio,
		})
		if err = rows.Err(); err != nil {
			return make([]entity.Product, 0), err
		}
	}

	return products, nil
}

// func FindProductByStore(wantedId int) ([]entity.Product, error) {
// 	var products []entity.Product
// 	rows, err := dbConn.SQL.Query("select id, name, store_id, product_category_id, image, summary, price, order_count, active from products where store_id = $1", wantedId)
// 	if err != nil {
// 		return make([]entity.Product, 0), err
// 	}
// 	defer rows.Close()
// 	var id, storeId, productCategoryId, price, orderCount int
// 	var name, image, summary string
// 	var active bool
// 	for rows.Next() {
// 		err := rows.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &active)
// 		if err != nil {
// 			return make([]entity.Product, 0), err
// 		}
// 		products = append(products, entity.Product{
// 			ID:                id,
// 			Name:              name,
// 			StoreId:           storeId,
// 			ProductCategoryId: productCategoryId,
// 			Image:             image,
// 			Summary:           summary,
// 			Price:             price,
// 			OrderCount:        orderCount,
// 			Active:            active,
// 		})
// 		if err = rows.Err(); err != nil {
// 			return make([]entity.Product, 0), err
// 		}
// 	}
// 	return products, nil
// }

func (driver *productDriver) FindProduct(wantedId int) (entity.Product, error) {

	query := `
		select 
			id, name, store_id, product_category_id, image, summary, 
			price, order_count, active, flavors, volumes, addons, discount_ratio 
		from products where id = $1`

	var id, storeId, productCategoryId, price, orderCount int
	var name, image, summary string
	var active bool
	var flavors, volumes, addons []string
	var discountRatio float32

	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &active, pq.Array(&flavors), pq.Array(&volumes), pq.Array(&addons), &discountRatio)

	if err != nil {
		fmt.Println("error", err)
		return entity.Product{
			ID:                0,
			Name:              "",
			StoreId:           0,
			ProductCategoryId: 0,
			Image:             "",
			Summary:           "",
			Price:             0,
			OrderCount:        0,
			Active:            active,
			Flavors:           make([]entity.DetailEditRequest, 0),
			Volumes:           make([]entity.DetailEditRequest, 0),
			Addons:            make([]entity.DetailEditRequest, 0),
			DiscountRatio:     0,
		}, err
	}

	// details := make([]entity.Detail, 0)
	detailsIDsStrings := make([]string, 0)
	detailsIDsStrings = append(detailsIDsStrings, flavors...)
	detailsIDsStrings = append(detailsIDsStrings, volumes...)
	detailsIDsStrings = append(detailsIDsStrings, addons...)
	// fmt.Println("detailsIDsStrings:", detailsIDsStrings)
	details, err := dd.FindProductsDetails(driver.GetIntSliceFromString(detailsIDsStrings))
	if err != nil {
		return entity.Product{}, err
	}
	// fmt.Println("details:", details)

	finalFlavors, finalVolumes, finalAddons := dd.SeparateProductDetails(details)

	// finalFlavors, detailErr := FindProductsDetails(GetIntSliceFromString(flavors))
	// fmt.Println("Product Flavors:", flavors)
	// if detailErr != nil {
	// 	fmt.Println("flavors error", err)
	// 	return entity.Product{
	// 		ID:                0,
	// 		Name:              "",
	// 		StoreId:           0,
	// 		ProductCategoryId: 0,
	// 		Image:             "",
	// 		Summary:           "",
	// 		Price:             0,
	// 		OrderCount:        0,
	// 		Active:            active,
	// 		Flavors:           make([]entity.Detail, 0),
	// 		Volumes:           make([]entity.Detail, 0),
	// 		Addons:            make([]entity.Detail, 0),
	// 	}, detailErr
	// }
	// finalVolumes, detailErr := FindProductsDetails(GetIntSliceFromString(volumes))
	// fmt.Println("Product Volumes:", volumes)
	// if detailErr != nil {
	// 	fmt.Println("volumes error", err)
	// 	return entity.Product{
	// 		ID:                0,
	// 		Name:              "",
	// 		StoreId:           0,
	// 		ProductCategoryId: 0,
	// 		Image:             "",
	// 		Summary:           "",
	// 		Price:             0,
	// 		OrderCount:        0,
	// 		Active:            active,
	// 		Flavors:           make([]entity.Detail, 0),
	// 		Volumes:           make([]entity.Detail, 0),
	// 		Addons:            make([]entity.Detail, 0),
	// 	}, detailErr
	// }
	// finalAddons, detailErr := FindProductsDetails(GetIntSliceFromString(addons))
	// fmt.Println("Product Addons:", addons)
	// if detailErr != nil {
	// 	fmt.Println("addons error", err)
	// 	return entity.Product{
	// 		ID:                0,
	// 		Name:              "",
	// 		StoreId:           0,
	// 		ProductCategoryId: 0,
	// 		Image:             "",
	// 		Summary:           "",
	// 		Price:             0,
	// 		OrderCount:        0,
	// 		Active:            active,
	// 		Flavors:           make([]entity.Detail, 0),
	// 		Volumes:           make([]entity.Detail, 0),
	// 		Addons:            make([]entity.Detail, 0),
	// 	}, detailErr
	// }
	// fmt.Println("name", name)
	// fmt.Println("final flavors", finalFlavors)
	// fmt.Println("flavors", flavors)
	// fmt.Println("final volumes", finalVolumes)
	// fmt.Println("volumes", volumes)
	// fmt.Println("final addons", finalAddons)
	// fmt.Println("addons", addons)

	// imageFile, err := fd.GetFileInfo(image)

	// if err != nil {
	// 	return entity.Product{}, err
	// }
	product := entity.Product{
		ID:                id,
		Name:              name,
		StoreId:           storeId,
		ProductCategoryId: productCategoryId,
		Image:             image,
		Summary:           summary,
		Price:             price,
		OrderCount:        orderCount,
		Active:            active,
		Flavors:           finalFlavors,
		Volumes:           finalVolumes,
		Addons:            finalAddons,
		DiscountRatio:     discountRatio,
	}
	return product, nil
}

func (driver *productDriver) GetIntSliceFromString(theString []string) []int {
	// newString := theString[1 : len(theString)-1]
	theSlice := make([]int, 0)

	for i := 0; i < len(theString); i++ {
		ii, err := strconv.Atoi(theString[i])
		if err != nil {
			return make([]int, 0)
		}
		theSlice = append(theSlice, ii)
	}
	// fmt.Println("The Slice", theSlice)
	return theSlice
}

func (driver *productDriver) AddProduct(product entity.Product) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `INSERT INTO products
		(
			name, store_id, product_category_id, 
			image, summary, price, active, 
			flavors, volumes, addons, discount_ratio
		)
	VALUES 
		(
			$1, $2, $3, 
			$4, $5, $6, $7, 
			$8, $9, $10, $11
		) returning *`

	flavors := driver.detailsSliceToIdSlice(product.Flavors)
	volumes := driver.detailsSliceToIdSlice(product.Volumes)
	addons := driver.detailsSliceToIdSlice(product.Addons)
	// for i := 0; i < len(product.Flavors); i++ {
	// 	flavors = append(flavors, product.Flavors[i].ID)
	// }
	// for i := 0; i < len(product.Volumes); i++ {
	// 	volumes = append(volumes, product.Volumes[i].ID)
	// }
	// for i := 0; i < len(product.Addons); i++ {
	// 	addons = append(addons, product.Addons[i].ID)
	// }

	flavorsString := driver.sliceToString(flavors)
	volumesString := driver.sliceToString(volumes)
	addonsString := driver.sliceToString(addons)

	result, err := dbConn.SQL.ExecContext(ctx, stmt, product.Name, product.StoreId, product.ProductCategoryId, product.Image, product.Summary, product.Price, product.Active, flavorsString, volumesString, addonsString, product.DiscountRatio)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product could not be added")
	}

	return nil
}

func (driver *productDriver) detailsSliceToIdSlice(details []entity.DetailEditRequest) []int {
	newDetails := make([]int, 0)
	for i := 0; i < len(details); i++ {
		newDetails = append(newDetails, details[i].ID)
	}
	return newDetails
}

func (driver *productDriver) sliceToString(theSlice []int) string {
	s := "" // fmt.Sprint(theSlice)

	for i := 0; i < len(theSlice); i++ {
		s = s + strconv.Itoa(theSlice[i]) + ","
	}

	// fmt.Println("String: ", s)
	if len(s) != 0 {
		s1 := "{" + s[0:len(s)-1] + "}"
		return s1
	}

	// fmt.Println("String1: ", s1)
	return "{}"
}

func (driver *productDriver) DeleteProduct(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from products where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product could not be found")
	}

	return nil
}

func (driver *productDriver) ActivateProduct(productEditInfo entity.ProductInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE products SET active = true WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, productEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product could not be found")
	}
	return nil
}

func (driver *productDriver) DeactivateProduct(productEditInfo entity.ProductInfoRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE products SET active = false WHERE id = $1 RETURNING *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, productEditInfo.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product could not be found")
	}
	return nil
}

func (driver *productDriver) EditProduct(productEditInfo entity.ProductEditRequest) (entity.Product, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	if productEditInfo.ProductCategoryId != 0 {
		if productEditInfo.StoreId == 0 {
			product, err := driver.FindProduct(productEditInfo.ID)
			if err != nil {
				return entity.Product{}, err
			}
			productCategoryToStoreBool, err := IsProductCategoryRelatedToStore(productEditInfo.ProductCategoryId, product.StoreId)
			if err != nil {
				return entity.Product{}, err
			}
			if !productCategoryToStoreBool {
				return entity.Product{}, errors.New("product category is not linked to specified store")
			}
		} else {
			productCategoryToStoreBool, err := IsProductCategoryRelatedToStore(productEditInfo.ProductCategoryId, productEditInfo.StoreId)
			if err != nil {
				return entity.Product{}, err
			}
			if !productCategoryToStoreBool {
				return entity.Product{}, errors.New("product category is not linked to specified store")
			}
		}
	}
	if productEditInfo.StoreId != 0 {
		if productEditInfo.ProductCategoryId == 0 {
			product, err := driver.FindProduct(productEditInfo.ID)
			if err != nil {
				return entity.Product{}, err
			}
			productCategoryToStoreBool, err := IsProductCategoryRelatedToStore(product.ProductCategoryId, productEditInfo.StoreId)
			if err != nil {
				return entity.Product{}, err
			}
			if !productCategoryToStoreBool {
				return entity.Product{}, errors.New("product category is not linked to specified store")
			}
		} else {
			productCategoryToStoreBool, err := IsProductCategoryRelatedToStore(productEditInfo.ProductCategoryId, productEditInfo.StoreId)
			if err != nil {
				return entity.Product{}, err
			}
			if !productCategoryToStoreBool {
				return entity.Product{}, errors.New("product category is not linked to specified store")
			}
		}
	}

	stmt := driver.GetEditProductStatementString(productEditInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return entity.Product{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entity.Product{}, errors.New("product could not be found")
	}
	product, _ := driver.FindProduct(productEditInfo.ID)
	// fmt.Println("Product ID:", productEditInfo.ID)
	// fmt.Println("Error:", err.Error())
	// if err != nil {
	// 	return entity.Product{}, err
	// }
	return product, nil
}

func (driver *productDriver) GetEditProductStatementString(productEditInfo entity.ProductEditRequest) string {
	stmt := `UPDATE products SET `
	if productEditInfo.Name != "" {
		stmt = stmt + `name = '` + productEditInfo.Name + `', `
	}
	if productEditInfo.StoreId != 0 {
		stmt = stmt + `store_id = ` + fmt.Sprint(productEditInfo.StoreId) + `, `
	}
	if productEditInfo.ProductCategoryId != 0 {
		stmt = stmt + `product_category_id = ` + fmt.Sprint(productEditInfo.ProductCategoryId) + `, `
	}
	if productEditInfo.Image != "" {
		stmt = stmt + `image = '` + productEditInfo.Image + `', `
	}
	if productEditInfo.Summary != "" {
		stmt = stmt + `summary = '` + fmt.Sprint(productEditInfo.Summary) + `', `
	}
	if productEditInfo.Price != 0 {
		stmt = stmt + `price = ` + fmt.Sprint(productEditInfo.Price) + `, `
	}
	if productEditInfo.DiscountRatio != 0 {
		stmt = stmt + `discount_ratio = ` + fmt.Sprint(productEditInfo.DiscountRatio) + `, `
	}
	if len(productEditInfo.Flavors) != 0 {
		flavors := driver.detailsSliceToIdSlice(productEditInfo.Flavors)
		flavorsString := driver.sliceToString(flavors)

		stmt = stmt + `flavors = '` + flavorsString + `', `
	}
	if len(productEditInfo.Volumes) != 0 {
		volumes := driver.detailsSliceToIdSlice(productEditInfo.Volumes)
		volumesString := driver.sliceToString(volumes)

		stmt = stmt + `volumes = '` + fmt.Sprint(volumesString) + `', `
	}
	if len(productEditInfo.Addons) != 0 {
		addons := driver.detailsSliceToIdSlice(productEditInfo.Addons)
		addonsString := driver.sliceToString(addons)

		stmt = stmt + `addons = '` + fmt.Sprint(addonsString) + `', `
	}
	if productEditInfo.Active {
		stmt = stmt + `active = true `
	} else {
		stmt = stmt + `active = false `
	}
	stmt = stmt + `where id = ` + fmt.Sprint(productEditInfo.ID) + ` RETURNING *`
	return stmt
}

func (driver *productDriver) EditProductsSliceOrderCount(productInfo []entity.OrderProduct) error {
	for _, item := range productInfo {
		err := driver.EditProductOrderCount(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (driver *productDriver) EditProductOrderCount(productInfo entity.OrderProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE products SET order_count = order_count + $1 where id = $2 RETURNING id`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, productInfo.ProductCount, productInfo.ProductID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product could not be found")
	}
	_, err = driver.FindProduct(productInfo.ProductID)
	if err != nil {
		return err
	}
	return nil
}

func (driver *productDriver) FindBestSellingProducts(productsCountLimit int) ([]entity.Product, error) {
	bestSellingProducts := make([]entity.Product, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, name, store_id, product_category_id, image, summary, 
		price, order_count, discount_ratio 
	from products order by order_count desc limit $1`, productsCountLimit)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	defer rows.Close()

	var id, storeId, productCategoryId, price, orderCount int
	var name, image, summary string
	var discountRatio float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &productCategoryId, &image, &summary, &price, &orderCount, &discountRatio)
		if err != nil {
			// log.Println(err)
			return make([]entity.Product, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)

		// if err != nil {
		// 	return make([]entity.Product, 0), err
		// }
		bestSellingProducts = append(bestSellingProducts, entity.Product{
			ID:                id,
			Name:              name,
			StoreId:           storeId,
			ProductCategoryId: productCategoryId,
			Image:             image,
			Summary:           summary,
			Price:             price,
			OrderCount:        orderCount,
			DiscountRatio:     discountRatio,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Product, 0), err
		}
		// fmt.Println("------------------------")
	}

	return bestSellingProducts, nil
}

func (driver *productDriver) FindOffersProducts() ([]entity.Product, error) {
	offersProducts := make([]entity.Product, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, name, store_id, image, 
		price, discount_ratio 
	from products where discount_ratio != 0.0 order by discount_ratio desc`)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	defer rows.Close()

	var id, storeId, price int
	var name, image string
	var discountRatio float32

	for rows.Next() {
		err := rows.Scan(&id, &name, &storeId, &image, &price, &discountRatio)
		if err != nil {
			// log.Println(err)
			return make([]entity.Product, 0), err
		}
		// imageFile, err := fd.GetFileInfo(image)

		// if err != nil {
		// 	return make([]entity.Product, 0), err
		// }
		offersProducts = append(offersProducts, entity.Product{
			ID:            id,
			Name:          name,
			StoreId:       storeId,
			Image:         image,
			Price:         price,
			DiscountRatio: discountRatio,
		})
		// fmt.Println("Record is:", userId, deliveryTime, products)
		if err = rows.Err(); err != nil {
			// log.Fatal("error Scanning Rows!")
			return make([]entity.Product, 0), err
		}
		// fmt.Println("------------------------")
	}

	return offersProducts, nil
}

func IsProductCategoryRelatedToStore(productCategoryId int, storeId int) (bool, error) {
	productCategory, err := pcd.FindProductCategory(productCategoryId)
	if err != nil {
		return false, err
	}
	return productCategory.StoreId == storeId, nil
}
