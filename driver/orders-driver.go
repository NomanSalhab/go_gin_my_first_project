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

type OrderDriver interface {
	FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindOrder(orderId int) (entity.Order, error)
	AddOrder(order entity.Order) error
	FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	DeleteOrder(wantedId int) error
	FindUserFinishedOrders(userWantedId int) ([]entity.Order, error)
	FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error)
	FinishOrder(orderId int) error
	EditOrder(orderEditInfo entity.OrderEditRequest) error

	FindOrderProducts(wantedIDs []int, singleOrder bool) ([]entity.OrderProduct, error)
	FindAllOrderProducts() ([]entity.OrderProduct, error)
	GetOrderProductsString(orderProducts []int) string
	GetAddressById(wantedId int) (entity.OrderAddress, error)
	AddOrderProducts(orderProducts []entity.OrderProduct) ([]int, int, error)
	GetAddOrderProductsStatement(orderProducts []entity.OrderProduct) (string, int)
	GetMockDetailsSliceFormIDs(ids []int) []entity.DetailEditRequest
	GetTimeStamp(theTime time.Time) string
	GetMonthNumberFromName(name string) string
	AddOrderIdToOrderProducts(orderTime string) error
	AddOrderProductsIDsToOrder(orderProductsIDs []int) error
	GetOrderProductsIDsByOrderId(orderId int) ([]int, error)
	GetUserByIdForOrder(wantedId int) (entity.User, error)
}

type orderDriver struct {
	cacheorders             []entity.Order
	cachedPageLimit         int
	cachedPageOffset        int
	cachedAllOrdersCount    int
	cachedMaximumPagesCount int
	cachedOrderProducts     []entity.OrderProduct
}

func NewOrderDriver() OrderDriver {
	return &orderDriver{}
}

var (
	dd DetailDriver  = NewDetailDriver()
	pd ProductDriver = NewProductDriver()
	ud UserDriver    = NewUserDriver()
	sd StoreDriver   = NewStoreDriver()
)

func (driver *orderDriver) FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	var paginationInfo entity.PaginationInfo
	if len(driver.cacheorders) != 0 && pageLimit == driver.cachedPageLimit && pageOffset == driver.cachedPageOffset && driver.cachedMaximumPagesCount != 0 && driver.cachedAllOrdersCount != 0 {
		// fmt.Println("Got Here")
		driver.cachedPageLimit = pageLimit
		driver.cachedPageOffset = pageOffset
		paginationInfo = driver.FindPaginationInfo()
		return driver.cacheorders, paginationInfo, nil
	} else {
		orders := make([]entity.Order, 0)
		rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, order_time, delivery_time, 
		products_cost, address, 
		delivery_cost, notes, finished, 
		delivery_worker_id, ordered, on_the_way 
	from orders order by order_time desc limit $1 offset $2`, pageLimit, pageOffset) // order_products,
		if err != nil {
			return make([]entity.Order, 0), paginationInfo, err
		}
		defer rows.Close()

		var id, userId, productsCost, address, deliveryCost, deliveryWorkerId int
		var notes string
		var finished, ordered, onTheWay bool
		var orderTime, deliveryTime time.Time
		// var orderProducts []string

		for rows.Next() {
			err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
				&productsCost, &address, // pq.Array(&orderProducts),
				&deliveryCost, &notes, &finished,
				&deliveryWorkerId, &ordered, &onTheWay)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			// fmt.Println("Order Product IDs:", orderProducts)
			//? To Change Wanting Order Products Uncomment Below
			finalOrderProducts, err := driver.FindOrderProductsByOrderId(id)
			// fmt.Println("Final Order Products:", finalOrderProducts)
			//? To Change Wanting Order Products Uncomment Below
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalAddress, err := driver.GetAddressById(address)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalUser, err := driver.GetUserByIdForOrder(userId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			order := entity.Order{
				ID:               id,
				UserID:           userId,
				UserName:         finalUser.Name,
				UserPhone:        finalUser.Phone,
				OrderTime:        orderTime,
				DeliveryTime:     deliveryTime,
				Address:          finalAddress,
				DeliveryCost:     deliveryCost,
				DeliveryWorkerId: deliveryWorkerId,
				Finished:         finished,
				Ordered:          ordered,
				OnTheWay:         onTheWay,
				Notes:            notes,
				ProductsCost:     productsCost,
				Products:/*make([]entity.OrderProduct, 0)*/ finalOrderProducts,
			}
			orders = append(orders, order)
			if err = rows.Err(); err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
		}

		driver.cacheorders = orders
		driver.cachedPageLimit = pageLimit
		driver.cachedPageOffset = pageOffset
		paginationInfo = driver.FindPaginationInfo()
		// fmt.Println("Orders:", orders)
		return orders, paginationInfo, nil
	}
}

func (driver *orderDriver) FindPaginationInfo() entity.PaginationInfo {
	var counter int

	dbConn.SQL.QueryRow("SELECT count(id) FROM orders").Scan(&counter)

	if driver.cachedPageLimit != 0 {
		// fmt.Println("we have", counter, "rows")
		// fmt.Println("we have", int(counter/driver.cachedPageLimit), "rows")
		return entity.PaginationInfo{
			AllItemsCount:     counter,
			MaximumPagesCount: int(math.Ceil(float64(counter) / float64(driver.cachedPageLimit))),
		}
	} else {
		return entity.PaginationInfo{
			AllItemsCount:     counter,
			MaximumPagesCount: int(math.Ceil(float64(float64(counter) / float64(1)))),
		}
	}
}

func (driver *orderDriver) FindFinishedPaginationInfo() entity.PaginationInfo {
	var counter int

	dbConn.SQL.QueryRow("SELECT count(id) FROM orders where finished = true").Scan(&counter)

	if driver.cachedPageLimit != 0 {
		// fmt.Println("we have", counter, "rows")
		// fmt.Println("we have", int(counter/driver.cachedPageLimit), "rows")
		return entity.PaginationInfo{
			AllItemsCount:     counter,
			MaximumPagesCount: int(math.Ceil(float64(counter) / float64(driver.cachedPageLimit))),
		}
	} else {
		return entity.PaginationInfo{
			AllItemsCount:     counter,
			MaximumPagesCount: int(math.Ceil(float64(float64(counter) / float64(1)))),
		}
	}
}

func (driver *orderDriver) FindNotFinishedPaginationInfo() entity.PaginationInfo {
	var counter int

	dbConn.SQL.QueryRow("SELECT count(id) FROM orders where finished = false").Scan(&counter)

	if driver.cachedPageLimit != 0 {
		// fmt.Println("we have", counter, "rows")
		// fmt.Println("we have", int(counter/driver.cachedPageLimit), "rows")
		return entity.PaginationInfo{
			AllItemsCount:     counter,
			MaximumPagesCount: int(math.Ceil(float64(counter) / float64(driver.cachedPageLimit))),
		}
	} else {
		return entity.PaginationInfo{
			AllItemsCount:     counter,
			MaximumPagesCount: int(math.Ceil(float64(float64(counter) / float64(1)))),
		}
	}
}

func (driver *orderDriver) FindAllOrderProducts() ([]entity.OrderProduct, error) {

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

func (driver *orderDriver) FindOrderProducts(wantedIDs []int, singleOrder bool) ([]entity.OrderProduct, error) {
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

func (driver *orderDriver) FindOrderProductsByOrderId(orderID int) ([]entity.OrderProduct, error) {
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

func (driver *orderDriver) GetOrderProductsString(orderProducts []int) string {
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

func (driver *orderDriver) GetAddressById(wantedId int) (entity.OrderAddress, error) {
	// fmt.Println("Address ID:", wantedId)
	query := `select id, name, user_id, latitude, longitude from addresses where id = $1`
	var id, userId int
	var name string
	var latitude, longitude float32
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &userId, &latitude, &longitude)
	if err != nil {
		return entity.OrderAddress{
			ID:        0,
			Name:      "",
			UserId:    0,
			Latitude:  float32(0.0),
			Longitude: float32(0.0),
		}, err
	}
	address := entity.OrderAddress{
		ID:        id,
		Name:      name,
		UserId:    userId,
		Latitude:  latitude,
		Longitude: longitude,
	}
	return address, nil
}

func (driver *orderDriver) GetUserByIdForOrder(wantedId int) (entity.User, error) {
	query := `select id, name, phone from users where id = $1`
	var id int
	var name, phone string
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &phone)
	if err != nil {
		return entity.User{
			ID:    0,
			Name:  "",
			Phone: "",
		}, err
	}
	user := entity.User{
		ID:    id,
		Name:  name,
		Phone: phone,
	}
	return user, nil
}

func (driver *orderDriver) AddOrder(order entity.Order) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	orderTime := time.Now()
	order.OrderTime = orderTime
	for i := 0; i < len(order.Products); i++ {
		order.Products[i].OrderTime = order.OrderTime
	}

	var userDeliveryCost int
	user, err := ud.FindUser(order.UserID)
	if err != nil {
		return err
	}
	storesDeliveryCost := GetOrderDeliveryCost(order)

	userDeliveryCost = storesDeliveryCost - int(user.UserDiscount*float32(storesDeliveryCost))
	// fmt.Println("Delivery Cost:", userDeliveryCost)

	_, productsCost := driver.GetAddOrderProductsStatement(order.Products)

	stmt := `INSERT INTO orders(
		user_id, order_time, delivery_time,
		products_cost, address,
		delivery_cost, notes, delivery_worker_id,
		finished, ordered, on_the_way
	)
	VALUES (
		$1, $2, $3,
		$4, $5, $6,
		$7, $8, $9,
		$10, $11
	) returning *` // order_products, , $12

	result, err := dbConn.SQL.ExecContext(ctx, stmt,
		order.UserID, order.OrderTime, order.DeliveryTime,
		productsCost, order.Address.ID, // op,
		userDeliveryCost, order.Notes, order.DeliveryWorkerId,
		false, false, false)
	if err != nil {
		fmt.Println("Here2")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order could not be added")
	}

	lastOrderID, _ := GetLastOrderId()
	// fmt.Println("Last Order ID:", lastOrderID)
	for i := 0; i < len(order.Products); i++ {
		order.Products[i].OrderId = lastOrderID
	}

	_, _, err = driver.AddOrderProducts(order.Products)
	if err != nil {
		fmt.Println("Here0")
		return err
	}
	// op := pd.sliceToString(o)

	// fmt.Println("Adding Order Result:", result)
	// err = driver.AddOrderIdToOrderProducts(driver.GetTimeStamp(order.OrderTime))
	// if err != nil {
	// 	fmt.Println("Here3")
	// 	return err
	// }

	driver.cacheorders = make([]entity.Order, 0)
	return nil
}

func (driver *orderDriver) AddOrderProducts(orderProducts []entity.OrderProduct) ([]int, int, error) {

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

func (driver *orderDriver) GetAddOrderProductsStatement(orderProducts []entity.OrderProduct) (string, int) {
	var productsCost int
	stmt := `INSERT INTO order_products(
		product_id, product_count, product_price,
		store_id, store_name, store_delivery_cost,
		addons, flavors, volumes, order_time, order_id)
	VALUES`
	for i := 0; i < len(orderProducts)-1; i++ {

		singleProduct, _ := pd.FindProduct(orderProducts[i].ProductID)
		orderProducts[i].ProductPrice = singleProduct.Price
		productsCost = productsCost + (singleProduct.Price * orderProducts[i].ProductCount)
		singleStore, _ := sd.FindStore(orderProducts[i].StoreId)
		orderProducts[i].StoreDeliveryCost = singleStore.DeliveryRent
		orderProducts[i].StoreName = singleStore.Name

		f := make([]int, 0)
		f = append(f, orderProducts[i].Flavors.ID)
		v := make([]int, 0)
		v = append(v, orderProducts[i].Volumes.ID)
		a := make([]int, 0)
		for j := 0; j < len(orderProducts[i].Addons); j++ {
			a = append(a, orderProducts[i].Addons[j].ID)
		}
		// a = append(a, orderProducts[i].Volumes.ID)
		flavors := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(f))
		volumes := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(v))
		addons := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(a))

		// flavorsString := sliceToString(flavors)
		// volumesString := sliceToString(volumes)
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
	orderProducts[lastIndex].ProductPrice = singleProduct.Price
	productsCost = productsCost + (singleProduct.Price * orderProducts[lastIndex].ProductCount)
	fmt.Println("Products Cost:", productsCost)
	singleStore, _ := sd.FindStore(orderProducts[lastIndex].StoreId)
	orderProducts[lastIndex].StoreDeliveryCost = singleStore.DeliveryRent
	orderProducts[lastIndex].StoreName = singleStore.Name

	f := make([]int, 0)
	f = append(f, orderProducts[lastIndex].Flavors.ID)
	v := make([]int, 0)
	v = append(v, orderProducts[lastIndex].Volumes.ID)
	a := make([]int, 0)
	for j := 0; j < len(orderProducts[lastIndex].Addons); j++ {
		a = append(a, orderProducts[lastIndex].Addons[j].ID)
	}
	// a = append(a, orderProducts[lastIndex].Volumes.ID)
	flavors := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(f))
	volumes := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(v))
	addons := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(a))
	// addons := detailsSliceToIdSlice(GetMockDetailsSliceFormIDs(orderProducts[lastIndex].Addons))

	// flavorsString := sliceToString(flavors)
	// volumesString := sliceToString(volumes)
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

func (driver *orderDriver) GetMockDetailsSliceFormIDs(ids []int) []entity.DetailEditRequest {
	details := make([]entity.DetailEditRequest, 0)
	for i := 0; i < len(ids); i++ {
		details = append(details, entity.DetailEditRequest{ID: ids[i]})
	}
	return details
}

func (driver *orderDriver) GetTimeStamp(theTime time.Time) string {
	theString := ""

	theString = theString + fmt.Sprint(theTime.Year()) + "-" + driver.GetMonthNumberFromName(fmt.Sprint(theTime.Month())) + "-" + fmt.Sprint(theTime.Day()) + " " + fmt.Sprint(theTime.Hour()) + ":" + fmt.Sprint(theTime.Minute()) + ":" + fmt.Sprint(theTime.Second()) + "." + fmt.Sprint(theTime.Nanosecond())[0:3]

	// fmt.Println("The Time: ", theString)
	return theString
}

func (driver *orderDriver) GetMonthNumberFromName(name string) string {
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

func (driver *orderDriver) AddOrderIdToOrderProducts(orderTime string) error {
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

func (driver *orderDriver) AddOrderProductsIDsToOrder(orderProductsIDs []int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `UPDATE orders SET order_products = $1 WHERE id = $2`
	orderId, err := GetLastOrderId()
	if err != nil {
		return err
	}

	opi := pd.detailsSliceToIdSlice(driver.GetMockDetailsSliceFormIDs(orderProductsIDs))
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

func (driver *orderDriver) FindOrder(orderId int) (entity.Order, error) {
	order := entity.Order{}
	rows, err := dbConn.SQL.Query(`select 
	id, user_id,
	order_time, delivery_time,
	products_cost, address, delivery_cost,
	notes, delivery_worker_id, ordered, on_the_way, finished
	from orders where id = $1`, orderId) // order_products,
	if err != nil {
		return entity.Order{}, err
	}
	defer rows.Close()

	var id, userId, productsCost, deliveryCost, address, deliveryWorkerId int
	var orderTime, deliveryTime time.Time
	var notes string
	var finished, ordered, onTheWay bool
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId,
			&orderTime, &deliveryTime, // pq.Array(&orderProducts),
			&productsCost, &address, &deliveryCost,
			&notes, &deliveryWorkerId,
			&ordered, &onTheWay, &finished)
		if err != nil {
			return entity.Order{}, err
		}
		finalOrderProducts, err := driver.FindOrderProductsByOrderId(id) // pd.GetIntSliceFromString(orderProducts), true
		if err != nil {
			return entity.Order{}, err
		}
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return entity.Order{}, err
		}
		finalUser, err := driver.GetUserByIdForOrder(userId)
		if err != nil {
			return entity.Order{}, err
		}
		order = entity.Order{
			ID:               id,
			UserID:           userId,
			UserName:         finalUser.Name,
			UserPhone:        finalUser.Phone,
			OrderTime:        orderTime,
			DeliveryTime:     deliveryTime,
			Products:         finalOrderProducts,
			ProductsCost:     productsCost,
			Address:          finalAddress,
			DeliveryCost:     deliveryCost,
			Notes:            notes,
			Finished:         finished,
			Ordered:          ordered,
			OnTheWay:         onTheWay,
			DeliveryWorkerId: deliveryWorkerId,
		}
		if err = rows.Err(); err != nil {
			return entity.Order{}, err
		}
	}

	return order, nil
}

func (driver *orderDriver) FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	var paginationInfo entity.PaginationInfo
	if len(driver.cacheorders) != 0 && pageLimit == driver.cachedPageLimit && pageOffset == driver.cachedPageOffset && driver.cachedMaximumPagesCount != 0 && driver.cachedAllOrdersCount != 0 {
		// fmt.Println("Got Here")
		driver.cachedPageLimit = pageLimit
		driver.cachedPageOffset = pageOffset
		paginationInfo = driver.FindFinishedPaginationInfo()
		return driver.cacheorders, paginationInfo, nil
	} else {
		orders := make([]entity.Order, 0)
		rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, order_time, delivery_time, 
		products_cost, address, 
		delivery_cost, notes, finished, 
		delivery_worker_id, ordered, on_the_way 
	from orders where finished = true order by order_time desc limit $1 offset $2`, pageLimit, pageOffset) // order_products,
		if err != nil {
			return make([]entity.Order, 0), paginationInfo, err
		}
		defer rows.Close()

		var id, userId, productsCost, address, deliveryCost, deliveryWorkerId int
		var notes string
		var finished, ordered, onTheWay bool
		var orderTime, deliveryTime time.Time
		// var orderProducts []string

		for rows.Next() {
			err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
				&productsCost, &address, // pq.Array(&orderProducts),
				&deliveryCost, &notes, &finished,
				&deliveryWorkerId, &ordered, &onTheWay)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			// fmt.Println("Order Product IDs:", orderProducts)
			//? To Change Wanting Order Products Uncomment Below
			finalOrderProducts, err := driver.FindOrderProductsByOrderId(id) // pd.GetIntSliceFromString(orderProducts), false
			// fmt.Println("Final Order Products:", finalOrderProducts)
			//? To Change Wanting Order Products Uncomment Below
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalAddress, err := driver.GetAddressById(address)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalUser, err := driver.GetUserByIdForOrder(userId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			order := entity.Order{
				ID:               id,
				UserID:           userId,
				UserName:         finalUser.Name,
				UserPhone:        finalUser.Phone,
				OrderTime:        orderTime,
				DeliveryTime:     deliveryTime,
				Address:          finalAddress,
				DeliveryCost:     deliveryCost,
				DeliveryWorkerId: deliveryWorkerId,
				Finished:         finished,
				Ordered:          ordered,
				OnTheWay:         onTheWay,
				Notes:            notes,
				ProductsCost:     productsCost,
				Products:/*make([]entity.OrderProduct, 0)*/ finalOrderProducts,
			}
			orders = append(orders, order)
			if err = rows.Err(); err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
		}

		driver.cacheorders = orders
		driver.cachedPageLimit = pageLimit
		driver.cachedPageOffset = pageOffset
		paginationInfo = driver.FindFinishedPaginationInfo()
		// fmt.Println("Orders:", orders)
		return orders, paginationInfo, nil
	}
}

func (driver *orderDriver) FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {

	var paginationInfo entity.PaginationInfo
	if len(driver.cacheorders) != 0 && pageLimit == driver.cachedPageLimit && pageOffset == driver.cachedPageOffset && driver.cachedMaximumPagesCount != 0 && driver.cachedAllOrdersCount != 0 {
		// fmt.Println("Got Here")
		driver.cachedPageLimit = pageLimit
		driver.cachedPageOffset = pageOffset
		paginationInfo = driver.FindNotFinishedPaginationInfo()
		return driver.cacheorders, paginationInfo, nil
	} else {
		notFinishedOrders := make([]entity.Order, 0)
		rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, order_time, delivery_time, 
		products_cost, address, 
		delivery_cost, notes, finished, 
		delivery_worker_id, ordered, on_the_way 
	from orders where finished = false order by order_time desc limit $1 offset $2`, pageLimit, pageOffset) //order_products,
		if err != nil {
			return make([]entity.Order, 0), paginationInfo, err
		}
		defer rows.Close()

		var id, userId, productsCost, address, deliveryCost, deliveryWorkerId int
		var notes string
		var finished, ordered, onTheWay bool
		var orderTime, deliveryTime time.Time
		// var orderProducts []string

		for rows.Next() {
			err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
				&productsCost, &address, // pq.Array(&orderProducts),
				&deliveryCost, &notes, &finished,
				&deliveryWorkerId, &ordered, &onTheWay)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			// fmt.Println("Order Product IDs:", orderProducts)
			//? To Change Wanting Order Products Uncomment Below
			finalOrderProducts, err := driver.FindOrderProductsByOrderId(id) // pd.GetIntSliceFromString(orderProducts), false
			// fmt.Println("Final Order Products:", finalOrderProducts)
			//? To Change Wanting Order Products Uncomment Below
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalAddress, err := driver.GetAddressById(address)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalUser, err := driver.GetUserByIdForOrder(userId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			order := entity.Order{
				ID:               id,
				UserID:           userId,
				UserName:         finalUser.Name,
				UserPhone:        finalUser.Phone,
				OrderTime:        orderTime,
				DeliveryTime:     deliveryTime,
				Address:          finalAddress,
				DeliveryCost:     deliveryCost,
				DeliveryWorkerId: deliveryWorkerId,
				Finished:         finished,
				Ordered:          ordered,
				OnTheWay:         onTheWay,
				Notes:            notes,
				ProductsCost:     productsCost,
				Products:/*make([]entity.OrderProduct, 0)*/ finalOrderProducts,
			}
			notFinishedOrders = append(notFinishedOrders, order)
			if err = rows.Err(); err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
		}

		driver.cacheorders = notFinishedOrders
		driver.cachedPageLimit = pageLimit
		driver.cachedPageOffset = pageOffset
		paginationInfo = driver.FindNotFinishedPaginationInfo()
		// fmt.Println("Orders:", orders)
		return notFinishedOrders, paginationInfo, nil
	}
}

func (driver *orderDriver) FindUserFinishedOrders(userWantedId int) ([]entity.Order, error) {
	orders := make([]entity.Order, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, order_time, delivery_time, 
		products_cost, address, 
		delivery_cost, notes, finished, 
		delivery_worker_id, ordered, on_the_way 
	from orders where user_id = $1 and finished = true order by order_time desc`, userWantedId) //order_products,
	if err != nil {
		return make([]entity.Order, 0), err
	}
	defer rows.Close()

	var id, userId, productsCost, address, deliveryCost, deliveryWorkerId int
	var notes string
	var finished, ordered, onTheWay bool
	var orderTime, deliveryTime time.Time
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
			&productsCost, &address, // pq.Array(&orderProducts),
			&deliveryCost, &notes, &finished,
			&deliveryWorkerId, &ordered, &onTheWay)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		// fmt.Println("Order Product IDs:", orderProducts)
		//? To Change Wanting Order Products Uncomment Below
		finalOrderProducts, err := driver.FindOrderProductsByOrderId(id) // pd.GetIntSliceFromString(orderProducts), true
		// fmt.Println("Final Order Products:", finalOrderProducts)
		//? To Change Wanting Order Products Uncomment Below
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalUser, err := driver.GetUserByIdForOrder(userId)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		order := entity.Order{
			ID:               id,
			UserID:           userId,
			UserName:         finalUser.Name,
			UserPhone:        finalUser.Phone,
			OrderTime:        orderTime,
			DeliveryTime:     deliveryTime,
			Address:          finalAddress,
			DeliveryCost:     deliveryCost,
			DeliveryWorkerId: deliveryWorkerId,
			Finished:         finished,
			Ordered:          ordered,
			OnTheWay:         onTheWay,
			Notes:            notes,
			ProductsCost:     productsCost,
			Products:/*make([]entity.OrderProduct, 0)*/ finalOrderProducts,
		}
		orders = append(orders, order)
		if err = rows.Err(); err != nil {
			return make([]entity.Order, 0), err
		}
	}

	// driver.cacheorders = orders
	// driver.cachedPageLimit = pageLimit
	// driver.cachedPageOffset = pageOffset
	// paginationInfo = driver.FindPaginationInfo()
	// fmt.Println("Orders:", orders)
	return orders, nil
}

func (driver *orderDriver) FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	orders := make([]entity.Order, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, order_time, delivery_time, 
		products_cost, address, 
		delivery_cost, notes, finished, 
		delivery_worker_id, ordered, on_the_way 
	from orders where user_id = $1 and finished = false order by order_time desc`, userWantedId) //order_products,
	if err != nil {
		return make([]entity.Order, 0), err
	}
	defer rows.Close()

	var id, userId, productsCost, address, deliveryCost, deliveryWorkerId int
	var notes string
	var finished, ordered, onTheWay bool
	var orderTime, deliveryTime time.Time
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
			&productsCost, &address, // pq.Array(&orderProducts),
			&deliveryCost, &notes, &finished,
			&deliveryWorkerId, &ordered, &onTheWay)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		// fmt.Println("Order Product IDs:", orderProducts)
		//? To Change Wanting Order Products Uncomment Below
		finalOrderProducts, err := driver.FindOrderProductsByOrderId(id) // pd.GetIntSliceFromString(orderProducts), true
		fmt.Println("Final Order Products:", finalOrderProducts)
		//? To Change Wanting Order Products Uncomment Below
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalUser, err := driver.GetUserByIdForOrder(userId)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		order := entity.Order{
			ID:               id,
			UserID:           userId,
			UserName:         finalUser.Name,
			UserPhone:        finalUser.Phone,
			OrderTime:        orderTime,
			DeliveryTime:     deliveryTime,
			Address:          finalAddress,
			DeliveryCost:     deliveryCost,
			DeliveryWorkerId: deliveryWorkerId,
			Finished:         finished,
			Ordered:          ordered,
			OnTheWay:         onTheWay,
			Notes:            notes,
			ProductsCost:     productsCost,
			Products:/*make([]entity.OrderProduct, 0)*/ finalOrderProducts,
		}
		orders = append(orders, order)
		if err = rows.Err(); err != nil {
			return make([]entity.Order, 0), err
		}
	}

	// driver.cacheorders = orders
	// driver.cachedPageLimit = pageLimit
	// driver.cachedPageOffset = pageOffset
	// paginationInfo = driver.FindPaginationInfo()
	// fmt.Println("Orders:", orders)
	return orders, nil
}

func (driver *orderDriver) GetOrderProductsIDsByOrderId(orderId int) ([]int, error) {
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

func (driver *orderDriver) DeleteOrder(wantedId int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := `delete from orders where id=$1 returning *`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, wantedId)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order could not be found")
	}
	driver.cacheorders = make([]entity.Order, 0)

	return nil
}

func (driver *orderDriver) FinishOrder(orderId int) error {

	order, err := driver.FindOrder(orderId)
	if err != nil {
		return err
	}
	if order.Finished {
		return errors.New("order is already finished")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	deliveryTime := time.Now()

	stmt := `UPDATE orders SET 
	finished = true, ordered = true, on_the_way = true, delivery_time = $1
	where id=$2 returning id`

	result, err := dbConn.SQL.ExecContext(ctx, stmt, deliveryTime, order.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order could not be found")
	}
	driver.cacheorders = make([]entity.Order, 0)
	err = pd.EditProductsSliceOrderCount(order.Products)
	if err != nil {
		return err
	}
	err = ud.EditUserBalanceAndCircles(order.UserID, order.DeliveryCost, order.ProductsCost)
	if err != nil {
		return err
	}
	storesMap := make(map[int]int)
	for _, item := range order.Products {
		storesMap[item.StoreId] = storesMap[item.StoreId] + (item.ProductCount * item.ProductPrice)
	}
	storeIncreaseBalanceItems := make([]entity.StoreIncreaseBalance, 0, len(storesMap))
	for k, v := range storesMap {
		storeIncreaseBalanceItems = append(storeIncreaseBalanceItems, entity.StoreIncreaseBalance{
			ID:      k,
			Balance: v,
		})
	}
	for _, storeIncreaseBalanceItem := range storeIncreaseBalanceItems {
		err = sd.EditStoreBalance(storeIncreaseBalanceItem)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetLastOrderId() (int, error) {
	query := `select id from orders order by id desc limit 1`
	var id int
	row := dbConn.SQL.QueryRow(query)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (driver *orderDriver) EditOrder(orderEditInfo entity.OrderEditRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	stmt := GetOrderEditStatementString(orderEditInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order could not be found")
	}
	driver.cacheorders = make([]entity.Order, 0)
	if orderEditInfo.Finished {
		driver.FinishOrder(orderEditInfo.ID)
	}

	return nil
}

func GetOrderEditStatementString(order entity.OrderEditRequest) string {
	stmt := `UPDATE orders SET`
	if order.Ordered {
		stmt = stmt + ` ordered = true,`
	} else {
		stmt = stmt + ` ordered = false,`
	}
	if order.OnTheWay {
		stmt = stmt + ` on_the_way = true,`
	}
	if len(order.Notes) != 0 {
		stmt = stmt + ` notes = '` + order.Notes + `',`
	}
	if order.DeliveryWorkerId != 0 {
		stmt = stmt + ` delivery_worker_id = ` + fmt.Sprint(order.DeliveryWorkerId) + `,`
	}
	if order.DeliveryCost != 0 {
		stmt = stmt + ` delivery_cost = ` + fmt.Sprint(order.DeliveryCost) + `,`
	}
	if order.ProductsCost != 0 {
		stmt = stmt + ` products_cost = ` + fmt.Sprint(order.ProductsCost) + `,`
	}
	if order.Address.ID != 0 {
		stmt = stmt + ` address = ` + fmt.Sprint(order.Address.ID) + `,`
	}
	stmt = stmt[0:len(stmt)-1] + ` where id = ` + fmt.Sprint(order.ID) + ` returning id`
	fmt.Println("Edit Order Statement Is:", stmt)

	return stmt
}

func GetOrderDeliveryCost(order entity.Order) int {
	var deliveryCost int
	storesMap := make(map[int]int)
	for _, item := range order.Products {
		store, err := sd.FindStore(item.StoreId)
		if err != nil {
			return 5000
		}
		// storeDiscount := store.Discount
		storeDeliveryCost := store.DeliveryRent
		// fmt.Println("Store Discount is:", storeDiscount)
		// fmt.Println("Store Discount is:", storeDeliveryCost)
		storesMap[item.StoreId] = storeDeliveryCost //  - int(storeDiscount*float32(storeDeliveryCost))
		fmt.Println("Stores Map is:", storesMap)
	}
	storeIncreaseBalanceItems := make([]entity.StoreIncreaseBalance, 0, len(storesMap))
	for k, v := range storesMap {
		storeIncreaseBalanceItems = append(storeIncreaseBalanceItems, entity.StoreIncreaseBalance{
			ID:      k,
			Balance: v,
		})
	}
	for _, item := range storeIncreaseBalanceItems {
		deliveryCost = deliveryCost + item.Balance
	}
	fmt.Println("Stores Icrease Balance Items is:", deliveryCost)
	fmt.Println("Stores Delivery Cost is:", deliveryCost)
	return deliveryCost
}
