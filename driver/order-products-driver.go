package driver

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/lib/pq"
)

type OrderProductsDriver interface {
	AddOrderProducts(orderProducts []entity.OrderProduct) ([]int, int, error)
	GetAddOrderProductsStatement(orderProducts []entity.OrderProduct) (string, int)
	GetAddOrderProductDetails(orderProduct entity.OrderProduct, productsCost int) ([]int, []int, []int, int)
	GetAddOrderProductsProductsPriceAfterCoupon(productsCost int, coupon entity.CouponAddOrderRequest) int
	GetMockDetailsSliceFormIDs(ids []int) []entity.DetailEditRequest
	GetTimeStamp(theTime time.Time) string
	GetMonthNumberFromName(name string) string
	AddOrderIdToOrderProducts(orderTime string) error
	AddOrderProductsIDsToOrder(orderProductsIDs []int) error
	GetOrderProductsIDsByOrderId(orderId int) ([]int, error)
	FindOrderProductsByOrderId(orderID int) ([]entity.OrderProduct, error)
	FindAllOrderProducts() ([]entity.OrderProduct, error)
	FindOrderProducts(wantedIDs []int, singleOrder bool) ([]entity.OrderProduct, error)
	GetOrderProductsString(orderProducts []int) string
}

type orderProductsDriver struct {
	cachedOrderProducts []entity.OrderProduct
}

func NewOrderProductsDriver() OrderProductsDriver {
	return &orderProductsDriver{}
}

func (driver *orderProductsDriver) AddOrderProducts(orderProducts []entity.OrderProduct) ([]int, int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt, productsCost := driver.GetAddOrderProductsStatement(orderProducts)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	// lii, _ := result.LastInsertId()
	// fmt.Println("Last Inserted ID:", lii)
	if err != nil {
		fmt.Println("Here1")
		return make([]int, 0), productsCost, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return make([]int, 0), productsCost, errors.New("order products could not be added")
	}

	// fmt.Println("Adding Order Products Result:", result)
	return make([]int, 0), productsCost, nil
}

func (driver *orderProductsDriver) GetAddOrderProductsStatement(orderProducts []entity.OrderProduct) (string, int) {
	var productsCost int
	stmt := `INSERT INTO order_products(
		product_id, product_count, product_price,
		store_id, store_name, store_delivery_cost,
		addons, flavors, volumes, order_time, order_id)
	VALUES`
	for i := 0; i < len(orderProducts)-1; i++ {

		singleProduct, _ := pd.FindProduct(orderProducts[i].ProductID)
		if orderProducts[i].ProductPrice == 0 || orderProducts[i].ProductPrice <= singleProduct.Price {
			orderProducts[i].ProductPrice = int(math.Ceil(float64(singleProduct.Price) * float64(1-singleProduct.DiscountRatio)))
			productsCost = productsCost + (int(math.Ceil(float64(singleProduct.Price)*float64(1-singleProduct.DiscountRatio))) * orderProducts[i].ProductCount)
		} else {
			productsCost = productsCost + (orderProducts[i].ProductPrice * orderProducts[i].ProductCount)
		}
		// fmt.Println("Products Cost:", productsCost)
		singleStore, _ := sd.FindStore(orderProducts[i].StoreId)
		orderProducts[i].StoreDeliveryCost = singleStore.DeliveryRent
		orderProducts[i].StoreName = singleStore.Name

		// f := make([]int, 0)
		// f = append(f, orderProducts[i].Flavors.ID)
		// productsCost = productsCost + (orderProducts[i].Flavors.Price * orderProducts[i].ProductCount)
		// v := make([]int, 0)
		// v = append(v, orderProducts[i].Volumes.ID)
		// productsCost = productsCost + (orderProducts[i].Volumes.Price * orderProducts[i].ProductCount)
		// a := make([]int, 0)
		// for j := 0; j < len(orderProducts[i].Addons); j++ {
		// 	a = append(a, orderProducts[i].Addons[j].ID)
		// 	productsCost = productsCost + (orderProducts[i].Addons[j].Price * orderProducts[i].ProductCount)
		// }
		// flavors := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(f))
		// volumes := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(v))
		// addons := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(a))

		var flavors, volumes, addons []int
		flavors, volumes, addons, productsCost = driver.GetAddOrderProductDetails(orderProducts[i], productsCost)
		addonsString := pd.sliceToString(addons) // `(` +
		stmt = stmt + ` (` +
			fmt.Sprint(orderProducts[i].ProductID) + ` , ` +
			fmt.Sprint(orderProducts[i].ProductCount) + ` , ` +
			fmt.Sprint(orderProducts[i].ProductPrice) + ` , ` +
			fmt.Sprint(orderProducts[i].StoreId) + ` , '` +
			orderProducts[i].StoreName + `' , ` +
			fmt.Sprint(orderProducts[i].StoreDeliveryCost) + ` , '` +
			addonsString + `' , ` + fmt.Sprint(flavors[0]) + ` , ` +
			fmt.Sprint(volumes[0]) + ` , '` +
			driver.GetTimeStamp(orderProducts[i].OrderTime) + `' , ` +
			fmt.Sprint(orderProducts[i].OrderId) + `), `
	}
	lastIndex := len(orderProducts) - 1

	singleProduct, _ := pd.FindProduct(orderProducts[lastIndex].ProductID)
	if orderProducts[lastIndex].ProductPrice == 0 || orderProducts[lastIndex].ProductPrice <= singleProduct.Price {
		orderProducts[lastIndex].ProductPrice = int(math.Ceil(float64(singleProduct.Price) * float64(1-singleProduct.DiscountRatio)))
		productsCost = productsCost + (int(math.Ceil(float64(singleProduct.Price)*float64(1-singleProduct.DiscountRatio))) * orderProducts[lastIndex].ProductCount)
	} else {
		productsCost = productsCost + (orderProducts[lastIndex].ProductPrice * orderProducts[lastIndex].ProductCount)
	}
	singleStore, _ := sd.FindStore(orderProducts[lastIndex].StoreId)
	orderProducts[lastIndex].StoreDeliveryCost = singleStore.DeliveryRent
	orderProducts[lastIndex].StoreName = singleStore.Name

	// f := make([]int, 0)
	// f = append(f, orderProducts[lastIndex].Flavors.ID)
	// productsCost = (productsCost + orderProducts[lastIndex].Flavors.Price*orderProducts[lastIndex].ProductCount)
	// v := make([]int, 0)
	// v = append(v, orderProducts[lastIndex].Volumes.ID)
	// productsCost = (productsCost + orderProducts[lastIndex].Volumes.Price*orderProducts[lastIndex].ProductCount)
	// a := make([]int, 0)
	// for j := 0; j < len(orderProducts[lastIndex].Addons); j++ {
	// 	a = append(a, orderProducts[lastIndex].Addons[j].ID)
	// 	productsCost = (productsCost + orderProducts[lastIndex].Addons[j].Price*orderProducts[lastIndex].ProductCount)
	// }
	// flavors := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(f))
	// volumes := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(v))
	// addons := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(a))

	var flavors, volumes, addons []int
	flavors, volumes, addons, productsCost = driver.GetAddOrderProductDetails(orderProducts[lastIndex], productsCost)
	addonsString := pd.sliceToString(addons) //  `(` +
	stmt = stmt + ` (` + fmt.Sprint(orderProducts[lastIndex].ProductID) + ` , ` +
		fmt.Sprint(orderProducts[lastIndex].ProductCount) + ` , ` +
		fmt.Sprint(orderProducts[lastIndex].ProductPrice) + ` , ` +
		fmt.Sprint(orderProducts[lastIndex].StoreId) + ` , '` + orderProducts[lastIndex].StoreName + `' , ` +
		fmt.Sprint(orderProducts[lastIndex].StoreDeliveryCost) + ` , '` +
		addonsString + `' , ` + fmt.Sprint(flavors[0]) + ` , ` + fmt.Sprint(volumes[0]) + ` , '` +
		driver.GetTimeStamp(orderProducts[lastIndex].OrderTime) + `' , ` +
		fmt.Sprint(orderProducts[lastIndex].OrderId) + `),`

	stmt = stmt[0:len(stmt)-1] + ` returning id`
	fmt.Println("Adding Order Products Statement", stmt)
	return stmt, productsCost
}

func (driver *orderProductsDriver) GetAddOrderProductDetails(orderProduct entity.OrderProduct, productsCost int) ([]int, []int, []int, int) {
	newProductsCost := productsCost
	f := make([]int, 0)
	f = append(f, orderProduct.Flavors.ID)
	newProductsCost = newProductsCost + (orderProduct.Flavors.Price * orderProduct.ProductCount)
	v := make([]int, 0)
	v = append(v, orderProduct.Volumes.ID)
	newProductsCost = newProductsCost + (orderProduct.Volumes.Price * orderProduct.ProductCount)
	a := make([]int, 0)
	for j := 0; j < len(orderProduct.Addons); j++ {
		a = append(a, orderProduct.Addons[j].ID)
		newProductsCost = newProductsCost + (orderProduct.Addons[j].Price * orderProduct.ProductCount)
	}
	// a = append(a, orderProducts[i].Volumes.ID)
	flavors := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(f))
	volumes := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(v))
	addons := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(a))

	return flavors, volumes, addons, newProductsCost
}

func (driver *orderProductsDriver) GetAddOrderProductsProductsPriceAfterCoupon(productsCost int, coupon entity.CouponAddOrderRequest) int {
	var newProductsCost int
	newProductsCost = productsCost
	if coupon.Active && coupon.EndDate.After(time.Now()) {
		if coupon.FromProductsCost {
			if coupon.TimesUsed < coupon.TimesUsedLimit {
				if coupon.DiscountPercentage != 0 {
					newProductsCost = (newProductsCost - int(math.Floor(float64(newProductsCost)*float64(coupon.DiscountPercentage))))
				} else {
					fmt.Println("Not Discount Percentage")
					newProductsCost = newProductsCost - coupon.DiscountAmount
				}
			} else {
				fmt.Println("Exeeded Limt")
			}
		} else {
			fmt.Println("Not From Products Cost")
		}
	} else {
		fmt.Println("Not Active Or After End Date")
	}
	fmt.Println("Products Cost:", productsCost)
	fmt.Println("Coupon:", coupon)
	fmt.Println("New Products Cost:", newProductsCost)
	return newProductsCost
}

func (driver *orderProductsDriver) GetMockDetailsSliceFormIDs(ids []int) []entity.DetailEditRequest {
	details := make([]entity.DetailEditRequest, 0)
	for i := 0; i < len(ids); i++ {
		details = append(details, entity.DetailEditRequest{ID: ids[i]})
	}
	return details
}

func (driver *orderProductsDriver) GetTimeStamp(theTime time.Time) string {
	theString := ""

	theString = theString + fmt.Sprint(theTime.Year()) + "-" + driver.GetMonthNumberFromName(fmt.Sprint(theTime.Month())) + "-" + fmt.Sprint(theTime.Day()) + " " + fmt.Sprint(theTime.Hour()) + ":" + fmt.Sprint(theTime.Minute()) + ":" + fmt.Sprint(theTime.Second()) + "." + fmt.Sprint(theTime.Nanosecond())[0:3]

	// fmt.Println("The Time: ", theString)
	return theString
}

func (driver *orderProductsDriver) GetMonthNumberFromName(name string) string {
	number := "0"
	if name == "January" {
		number = "1"
	} else if name == "February" {
		number = "2"
	} else if name == "March" {
		number = "3"
	} else if name == "April" {
		number = "4"
	} else if name == "May" {
		number = "5"
	} else if name == "June" {
		number = "6"
	} else if name == "July" {
		number = "7"
	} else if name == "August" {
		number = "8"
	} else if name == "September" {
		number = "9"
	} else if name == "October" {
		number = "10"
	} else if name == "November" {
		number = "11"
	} else if name == "December" {
		number = "12"
	}
	return number
}

func (driver *orderProductsDriver) AddOrderIdToOrderProducts(orderTime string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE order_products SET order_id = $1 WHERE order_time = $2`
	orderId, err := GetLastOrderId()
	if err != nil {
		return err
	}

	result, err := dbConn.SQL.ExecContext(ctx, stmt, orderId, orderTime)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order id could not be added")
	}
	orderProductsIDs, err := driver.GetOrderProductsIDsByOrderId(orderId)
	if err != nil {
		return err
	}
	err = driver.AddOrderProductsIDsToOrder(orderProductsIDs)
	if err != nil {
		return err
	}

	return nil
}

func (driver *orderProductsDriver) AddOrderProductsIDsToOrder(orderProductsIDs []int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE orders SET order_products = $1 WHERE id = $2`
	orderId, err := GetLastOrderId()
	if err != nil {
		return err
	}

	opi := pd.detailsSliceToIdSlice(opd.GetMockDetailsSliceFormIDs(orderProductsIDs))
	opiString := pd.sliceToString(opi)

	result, err := dbConn.SQL.ExecContext(ctx, stmt, opiString, orderId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order id could not be added")
	}

	return nil
}

func (driver *orderProductsDriver) GetOrderProductsIDsByOrderId(orderId int) ([]int, error) {
	ids := make([]int, 0)
	rows, err := dbConn.SQL.Query("select id from order_products where order_id = $1", orderId)
	if err != nil {
		return make([]int, 0), err
	}
	defer rows.Close()

	var id int

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return make([]int, 0), err
		}
		ids = append(ids, id)
		if err = rows.Err(); err != nil {
			return make([]int, 0), err
		}
	}

	return ids, nil
}

func (driver *orderProductsDriver) FindOrderProductsByOrderId(orderID int) ([]entity.OrderProduct, error) {
	if len(driver.cachedOrderProducts) != 0 {
		wantedOrderProducts := make([]entity.OrderProduct, 0)
		// fmt.Println("cached Order Products:", driver.cachedOrderProducts)
		for _, item := range driver.cachedOrderProducts {
			if item.OrderId == orderID {
				wantedOrderProducts = append(wantedOrderProducts, item)
				// break
			}
		}
		if len(wantedOrderProducts) == 0 {
			driver.cachedOrderProducts = make([]entity.OrderProduct, 0)
			driver.FindOrderProductsByOrderId(orderID)
		}
		// fmt.Println("wanted Order Products:", wantedOrderProducts)
		return wantedOrderProducts, nil
	} else {
		orderProducts := make([]entity.OrderProduct, 0)
		query := `
			select 
				id, product_id, product_count, product_price, 
				store_id, store_delivery_cost, order_id,
				addons, flavors, volumes, order_time, store_name
			from order_products where order_id = $1
		`
		//driver.GetOrderProductsString(wantedIDs)
		// fmt.Println("Query is:", query)
		rows, err := dbConn.SQL.Query(query, orderID)
		if err != nil {
			// fmt.Println("scan0 error:", err)
			return make([]entity.OrderProduct, 0), err
		}
		defer rows.Close()

		var id, productId, productCount, productPrice, storeId, store_deliveryCost, orderId, flavors, volumes int
		var orderTime time.Time
		var addons []string
		var storeName string

		for rows.Next() {
			err := rows.Scan(
				&id, &productId, &productCount, &productPrice,
				&storeId, &store_deliveryCost, &orderId, pq.Array(&addons),
				&flavors, &volumes, &orderTime, &storeName)
			// fmt.Println("run")
			if err != nil {
				// fmt.Println("scan1 error:", err)
				return make([]entity.OrderProduct, 0), err
			}
			detailsIDsStrings := make([]string, 0)
			detailsIDsStrings = append(detailsIDsStrings, strconv.Itoa(flavors))
			detailsIDsStrings = append(detailsIDsStrings, strconv.Itoa(volumes))
			detailsIDsStrings = append(detailsIDsStrings, addons...)
			// fmt.Println("detailsIDsStrings:", detailsIDsStrings)
			details, err := dd.FindProductsDetails(pd.GetIntSliceFromString(detailsIDsStrings))
			if err != nil {
				return make([]entity.OrderProduct, 0), err
			}
			// fmt.Println("details:", details)

			finalFlavors, finalVolumes, finalAddons := dd.SeparateProductDetails(details)
			// fmt.Println("finalFlavors:", finalFlavors)
			// fmt.Println("finalVolumes:", finalVolumes)
			// fmt.Println("finalAddons:", finalAddons)
			// f := make([]int, 0)
			// f = append(f, flavors)
			// v := make([]int, 0)
			// v = append(v, volumes)
			// finalFlavors, detailErr := FindProductsDetails(f)
			// if detailErr != nil {
			// 	fmt.Println("flavors error", err)
			// 	return make([]entity.OrderProduct, 0), detailErr
			// }
			// finalVolumes, detailErr := FindProductsDetails(v)
			// if detailErr != nil {
			// 	fmt.Println("volumes error", err)
			// 	return make([]entity.OrderProduct, 0), detailErr
			// }
			// fmt.Println("Addons:", addons)
			// finalAddons, detailErr := FindProductsDetails(GetIntSliceFromString(addons))
			// fmt.Println("Final Addons:", finalAddons)
			// if detailErr != nil {
			// 	fmt.Println("addons error", err)
			// 	return make([]entity.OrderProduct, 0), detailErr
			// }
			orderProducts = append(orderProducts, entity.OrderProduct{
				ID:                id,
				ProductID:         productId,
				ProductCount:      productCount,
				ProductPrice:      productPrice,
				StoreId:           storeId,
				StoreName:         storeName,
				StoreDeliveryCost: store_deliveryCost,
				OrderId:           orderId,
				Addons:            finalAddons,
				Flavors:           finalFlavors[0],
				Volumes:           finalVolumes[0],
				OrderTime:         orderTime,
			})
			// fmt.Println(id, "-", name)
			if err = rows.Err(); err != nil {
				return make([]entity.OrderProduct, 0), err
			}
		}

		// fmt.Println("Time Called the Func:++++++++++++++++++++++++++")
		defer driver.FindAllOrderProducts()
		return orderProducts, nil
	}
}

func (driver *orderProductsDriver) FindAllOrderProducts() ([]entity.OrderProduct, error) {

	orderProducts := make([]entity.OrderProduct, 0)
	query := `select 
		id, product_id, product_count, product_price, 
		store_id, store_delivery_cost, order_id,
		addons, flavors, volumes, order_time, store_name
	from order_products`
	// fmt.Println("Query is:", query)
	rows, err := dbConn.SQL.Query(query)
	if err != nil {
		// fmt.Println("scan0 error:", err)
		return make([]entity.OrderProduct, 0), err
	}
	defer rows.Close()

	var id, productId, productCount, productPrice, storeId, store_deliveryCost, orderId, flavors, volumes int
	var orderTime time.Time
	var addons []string
	var storeName string

	for rows.Next() {
		err := rows.Scan(&id, &productId, &productCount, &productPrice, &storeId, &store_deliveryCost, &orderId, pq.Array(&addons), &flavors, &volumes, &orderTime, &storeName)
		// fmt.Println("run")
		if err != nil {
			// fmt.Println("scan1 error:", err)
			return make([]entity.OrderProduct, 0), err
		}
		detailsIDsStrings := make([]string, 0)
		detailsIDsStrings = append(detailsIDsStrings, strconv.Itoa(flavors))
		detailsIDsStrings = append(detailsIDsStrings, strconv.Itoa(volumes))
		detailsIDsStrings = append(detailsIDsStrings, addons...)
		// fmt.Println("detailsIDsStrings:", detailsIDsStrings)
		details, err := dd.FindProductsDetails(pd.GetIntSliceFromString(detailsIDsStrings))
		if err != nil {
			return make([]entity.OrderProduct, 0), err
		}
		// fmt.Println("details:", details)

		finalFlavors, finalVolumes, finalAddons := dd.SeparateProductDetails(details)
		// fmt.Println("finalFlavors:", finalFlavors)
		// fmt.Println("finalVolumes:", finalVolumes)
		// fmt.Println("finalAddons:", finalAddons)
		// f := make([]int, 0)
		// f = append(f, flavors)
		// v := make([]int, 0)
		// v = append(v, volumes)
		// finalFlavors, detailErr := FindProductsDetails(f)
		// if detailErr != nil {
		// 	fmt.Println("flavors error", err)
		// 	return make([]entity.OrderProduct, 0), detailErr
		// }
		// finalVolumes, detailErr := FindProductsDetails(v)
		// if detailErr != nil {
		// 	fmt.Println("volumes error", err)
		// 	return make([]entity.OrderProduct, 0), detailErr
		// }
		// fmt.Println("Addons:", addons)
		// finalAddons, detailErr := FindProductsDetails(GetIntSliceFromString(addons))
		// fmt.Println("Final Addons:", finalAddons)
		// if detailErr != nil {
		// 	fmt.Println("addons error", err)
		// 	return make([]entity.OrderProduct, 0), detailErr
		// }
		orderProducts = append(orderProducts, entity.OrderProduct{
			ID:                id,
			ProductID:         productId,
			ProductCount:      productCount,
			ProductPrice:      productPrice,
			StoreId:           storeId,
			StoreDeliveryCost: store_deliveryCost,
			OrderId:           orderId,
			Addons:            finalAddons,
			Flavors:           finalFlavors[0],
			Volumes:           finalVolumes[0],
			OrderTime:         orderTime,
			StoreName:         storeName,
		})
		// fmt.Println(id, "-", name)
		if err = rows.Err(); err != nil {
			return make([]entity.OrderProduct, 0), err
		}
	}

	// fmt.Println("Details Are:", details)
	driver.cachedOrderProducts = orderProducts
	return orderProducts, nil
}

func (driver *orderProductsDriver) FindOrderProducts(wantedIDs []int, singleOrder bool) ([]entity.OrderProduct, error) {
	if len(driver.cachedOrderProducts) != 0 && (!singleOrder) {
		wantedOrderProducts := make([]entity.OrderProduct, 0)
		// fmt.Println("cached Order Products:", driver.cachedOrderProducts)
		for _, item2 := range wantedIDs {
			for _, item := range driver.cachedOrderProducts {
				if item.ID == item2 {
					wantedOrderProducts = append(wantedOrderProducts, item)
					break
				}
			}
		}
		// fmt.Println("wanted Order Products:", wantedOrderProducts)
		return wantedOrderProducts, nil
	} else {
		orderProducts := make([]entity.OrderProduct, 0)
		query := driver.GetOrderProductsString(wantedIDs)
		// fmt.Println("Query is:", query)
		rows, err := dbConn.SQL.Query(query)
		if err != nil {
			// fmt.Println("scan0 error:", err)
			return make([]entity.OrderProduct, 0), err
		}
		defer rows.Close()

		var id, productId, productCount, productPrice, storeId, store_deliveryCost, orderId, flavors, volumes int
		var orderTime time.Time
		var addons []string
		var storeName string

		for rows.Next() {
			err := rows.Scan(
				&id, &productId, &productCount, &productPrice,
				&storeId, &store_deliveryCost, &orderId, pq.Array(&addons),
				&flavors, &volumes, &orderTime, &storeName)
			// fmt.Println("run")
			if err != nil {
				// fmt.Println("scan1 error:", err)
				return make([]entity.OrderProduct, 0), err
			}
			detailsIDsStrings := make([]string, 0)
			detailsIDsStrings = append(detailsIDsStrings, strconv.Itoa(flavors))
			detailsIDsStrings = append(detailsIDsStrings, strconv.Itoa(volumes))
			detailsIDsStrings = append(detailsIDsStrings, addons...)
			// fmt.Println("detailsIDsStrings:", detailsIDsStrings)
			details, err := dd.FindProductsDetails(pd.GetIntSliceFromString(detailsIDsStrings))
			if err != nil {
				return make([]entity.OrderProduct, 0), err
			}
			// fmt.Println("details:", details)

			finalFlavors, finalVolumes, finalAddons := dd.SeparateProductDetails(details)
			// fmt.Println("finalFlavors:", finalFlavors)
			// fmt.Println("finalVolumes:", finalVolumes)
			// fmt.Println("finalAddons:", finalAddons)
			// f := make([]int, 0)
			// f = append(f, flavors)
			// v := make([]int, 0)
			// v = append(v, volumes)
			// finalFlavors, detailErr := FindProductsDetails(f)
			// if detailErr != nil {
			// 	fmt.Println("flavors error", err)
			// 	return make([]entity.OrderProduct, 0), detailErr
			// }
			// finalVolumes, detailErr := FindProductsDetails(v)
			// if detailErr != nil {
			// 	fmt.Println("volumes error", err)
			// 	return make([]entity.OrderProduct, 0), detailErr
			// }
			// fmt.Println("Addons:", addons)
			// finalAddons, detailErr := FindProductsDetails(GetIntSliceFromString(addons))
			// fmt.Println("Final Addons:", finalAddons)
			// if detailErr != nil {
			// 	fmt.Println("addons error", err)
			// 	return make([]entity.OrderProduct, 0), detailErr
			// }
			orderProducts = append(orderProducts, entity.OrderProduct{
				ID:                id,
				ProductID:         productId,
				ProductCount:      productCount,
				ProductPrice:      productPrice,
				StoreId:           storeId,
				StoreName:         storeName,
				StoreDeliveryCost: store_deliveryCost,
				OrderId:           orderId,
				Addons:            finalAddons,
				Flavors:           finalFlavors[0],
				Volumes:           finalVolumes[0],
				OrderTime:         orderTime,
			})
			// fmt.Println(id, "-", name)
			if err = rows.Err(); err != nil {
				return make([]entity.OrderProduct, 0), err
			}
		}

		// fmt.Println("Time Called the Func:++++++++++++++++++++++++++")
		defer driver.FindAllOrderProducts()
		return orderProducts, nil
	}
}

func (driver *orderProductsDriver) GetOrderProductsString(orderProducts []int) string {
	stmt := `select 
		id, product_id, product_count, product_price, 
		store_id, store_delivery_cost, order_id,
		addons, flavors, volumes, order_time, store_name
	from order_products where `

	for i := 0; i < len(orderProducts)-1; i++ {
		stmt = stmt + `id = ` + fmt.Sprint(orderProducts[i]) + ` or `
	}
	stmt = stmt + `id = ` + fmt.Sprint(orderProducts[len(orderProducts)-1]) + ``

	// fmt.Println("Statement is:", stmt)
	return stmt
}
