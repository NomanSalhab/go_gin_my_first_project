package driver

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
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
	FindDeliveryWorkerNotFinishedOrders(userWantedId int) ([]entity.Order, error)
	FinishOrder(orderId int) error
	EditOrder(orderEditInfo entity.OrderEditRequest) error
	ChangeOrderWorkerId(orderChangeWorkerIdInfo entity.OrderChangeWorkerIdRequest) error

	GetAddressById(wantedId int) (entity.OrderAddress, error)

	FindPaginationInfo() entity.PaginationInfo
	FindFinishedPaginationInfo() entity.PaginationInfo
	FindNotFinishedPaginationInfo() entity.PaginationInfo
}

type orderDriver struct {
	cacheorders             []entity.Order
	cachedPageLimit         int
	cachedPageOffset        int
	cachedAllOrdersCount    int
	cachedMaximumPagesCount int
}

func NewOrderDriver() OrderDriver {
	return &orderDriver{}
}

var (
	dd  DetailDriver        = NewDetailDriver()
	pd  ProductDriver       = NewProductDriver()
	ud  UserDriver          = NewUserDriver()
	sd  StoreDriver         = NewStoreDriver()
	cd  CouponDriver        = NewCouponDriver()
	opd OrderProductsDriver = NewOrderProductsDriver()
	ad  AreaDriver          = NewAreaDriver()
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
		delivery_worker_id, ordered, on_the_way, coupon_id 
	from orders order by order_time desc limit $1 offset $2`, pageLimit, pageOffset) // order_products,
		if err != nil {
			return make([]entity.Order, 0), paginationInfo, err
		}
		defer rows.Close()

		var id, userId, productsCost, address, deliveryCost, deliveryWorkerId, couponId int
		var notes string
		var finished, ordered, onTheWay bool
		var orderTime, deliveryTime time.Time
		// var orderProducts []string

		for rows.Next() {
			err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
				&productsCost, &address, // pq.Array(&orderProducts),
				&deliveryCost, &notes, &finished,
				&deliveryWorkerId, &ordered, &onTheWay, &couponId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			//? finalOrderProducts, err := opd.FindOrderProductsByOrderId(id)
			//? if err != nil {
			//? 	return make([]entity.Order, 0), paginationInfo, err
			//? }
			finalAddress, err := driver.GetAddressById(address)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalUser, err := ud.FindUser(userId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			coupon, err := cd.GetCouponInfo(couponId)
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
				Coupon:           coupon,
				Products:         make([]entity.OrderProduct, 0), /*finalOrderProducts*/
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

func (driver *orderDriver) GetAddressById(wantedId int) (entity.OrderAddress, error) {
	// fmt.Println("Address ID:", wantedId)
	query := `select id, name, description, user_id, latitude, longitude from addresses where id = $1`
	var id, userId int
	var name, description string
	var latitude, longitude float32
	row := dbConn.SQL.QueryRow(query, wantedId)
	err := row.Scan(&id, &name, &description, &userId, &latitude, &longitude)
	if err != nil {
		return entity.OrderAddress{
			ID:           0,
			Name:         "",
			Descripition: "",
			UserId:       0,
			Latitude:     float32(0.0),
			Longitude:    float32(0.0),
		}, err
	}
	address := entity.OrderAddress{
		ID:           id,
		Name:         name,
		Descripition: description,
		UserId:       userId,
		Latitude:     latitude,
		Longitude:    longitude,
	}
	return address, nil
}

func (driver *orderDriver) AddOrder(order entity.Order) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	var userDeliveryCost int
	user, err := ud.FindUser(order.UserID)
	if err != nil {
		return err
	}
	if user.Active {
		storesDeliveryCost := GetOrderDeliveryCost(order)

		userDeliveryCost = storesDeliveryCost - int(user.UserDiscount*float32(storesDeliveryCost))

		_, productsCost := opd.GetAddOrderProductsStatement(order.Products)
		coupon, err := cd.GetCouponInfo(order.Coupon.ID)
		if err != nil {
			return err
		}
		if !coupon.FreeDelivery {
			productsCost = opd.GetAddOrderProductsProductsPriceAfterCoupon(productsCost, coupon)
		} else {
			userDeliveryCost = 0
		}

		stmt := `INSERT INTO orders(
		user_id, order_time, delivery_time,
		products_cost, address,
		delivery_cost, notes, delivery_worker_id,
		finished, ordered, on_the_way, coupon_id
	)
	VALUES (
		$1, $2, $3,
		$4, $5, $6,
		$7, $8, $9,
		$10, $11, $12
	) returning *` // order_products, , $12

		result, err := dbConn.SQL.ExecContext(ctx, stmt,
			order.UserID, order.OrderTime, order.DeliveryTime,
			productsCost, order.Address.ID, // op,
			userDeliveryCost, order.Notes, order.DeliveryWorkerId,
			false, false, false, order.Coupon.ID)
		if err != nil {
			return err
		}
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return errors.New("order could not be added")
		}

		lastOrderID, _ := GetLastOrderId()
		for i := 0; i < len(order.Products); i++ {
			order.Products[i].OrderId = lastOrderID
		}

		_, _, err = opd.AddOrderProducts(order.Products)
		if err != nil {
			return err
		}

		driver.cacheorders = make([]entity.Order, 0)
		return nil
	} else {
		return errors.New("user is not active")
	}
}

func (driver *orderDriver) FindOrder(orderId int) (entity.Order, error) {
	order := entity.Order{}
	rows, err := dbConn.SQL.Query(`select 
	id, user_id,
	order_time, delivery_time,
	products_cost, address, delivery_cost,
	notes, delivery_worker_id, ordered, on_the_way, finished, coupon_id 
	from orders where id = $1`, orderId) // order_products,
	if err != nil {
		return entity.Order{}, err
	}
	defer rows.Close()

	var id, userId, productsCost, deliveryCost, address, deliveryWorkerId, couponId int
	var orderTime, deliveryTime time.Time
	var notes string
	var finished, ordered, onTheWay bool
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId,
			&orderTime, &deliveryTime, // pq.Array(&orderProducts),
			&productsCost, &address, &deliveryCost,
			&notes, &deliveryWorkerId,
			&ordered, &onTheWay, &finished, &couponId)
		if err != nil {
			return entity.Order{}, err
		}
		finalOrderProducts, err := opd.FindOrderProductsByOrderId(id) // pd.GetIntSliceFromString(orderProducts), true
		if err != nil {
			return entity.Order{}, err
		}
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return entity.Order{}, err
		}
		finalUser, err := ud.FindUser(userId)
		if err != nil {
			return entity.Order{}, err
		}
		coupon, err := cd.GetCouponInfo(couponId)
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
			Coupon:           coupon,
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
		delivery_worker_id, ordered, on_the_way, coupon_id  
	from orders where finished = true order by order_time desc limit $1 offset $2`, pageLimit, pageOffset) // order_products,
		if err != nil {
			return make([]entity.Order, 0), paginationInfo, err
		}
		defer rows.Close()

		var id, userId, productsCost, address, deliveryCost, deliveryWorkerId, couponId int
		var notes string
		var finished, ordered, onTheWay bool
		var orderTime, deliveryTime time.Time
		// var orderProducts []string

		for rows.Next() {
			err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
				&productsCost, &address, // pq.Array(&orderProducts),
				&deliveryCost, &notes, &finished,
				&deliveryWorkerId, &ordered, &onTheWay, &couponId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			// finalOrderProducts, err := opd.FindOrderProductsByOrderId(id)
			// if err != nil {
			// 	return make([]entity.Order, 0), paginationInfo, err
			// }
			finalAddress, err := driver.GetAddressById(address)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalUser, err := ud.FindUser(userId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			coupon, err := cd.GetCouponInfo(couponId)
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
				Coupon:           coupon,
				Products:         make([]entity.OrderProduct, 0), /*finalOrderProducts*/
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
		delivery_worker_id, ordered, on_the_way, coupon_id  
	from orders where finished = false order by order_time desc limit $1 offset $2`, pageLimit, pageOffset) //order_products,
		if err != nil {
			return make([]entity.Order, 0), paginationInfo, err
		}
		defer rows.Close()

		var id, userId, productsCost, address, deliveryCost, deliveryWorkerId, couponId int
		var notes string
		var finished, ordered, onTheWay bool
		var orderTime, deliveryTime time.Time
		// var orderProducts []string

		for rows.Next() {
			err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
				&productsCost, &address, // pq.Array(&orderProducts),
				&deliveryCost, &notes, &finished,
				&deliveryWorkerId, &ordered, &onTheWay, &couponId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}

			// finalOrderProducts, err := opd.FindOrderProductsByOrderId(id)
			// if err != nil {
			// 	return make([]entity.Order, 0), paginationInfo, err
			// }
			finalAddress, err := driver.GetAddressById(address)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			finalUser, err := ud.FindUser(userId)
			if err != nil {
				return make([]entity.Order, 0), paginationInfo, err
			}
			coupon, err := cd.GetCouponInfo(couponId)
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
				Coupon:           coupon,
				Products:         make([]entity.OrderProduct, 0), /*finalOrderProducts*/
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
		delivery_worker_id, ordered, on_the_way, coupon_id  
	from orders where user_id = $1 and finished = true order by order_time desc`, userWantedId) //order_products,
	if err != nil {
		return make([]entity.Order, 0), err
	}
	defer rows.Close()

	var id, userId, productsCost, address, deliveryCost, deliveryWorkerId, couponId int
	var notes string
	var finished, ordered, onTheWay bool
	var orderTime, deliveryTime time.Time
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
			&productsCost, &address, // pq.Array(&orderProducts),
			&deliveryCost, &notes, &finished,
			&deliveryWorkerId, &ordered, &onTheWay, &couponId)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		// finalOrderProducts, err := opd.FindOrderProductsByOrderId(id)
		// if err != nil {
		// 	return make([]entity.Order, 0), err
		// }
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalUser, err := ud.FindUser(userId)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		coupon, err := cd.GetCouponInfo(couponId)
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
			Coupon:           coupon,
			Products:         make([]entity.OrderProduct, 0), /*finalOrderProducts*/
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
		delivery_worker_id, ordered, on_the_way, coupon_id  
	from orders where user_id = $1 and finished = false order by order_time desc`, userWantedId) //order_products,
	if err != nil {
		return make([]entity.Order, 0), err
	}
	defer rows.Close()

	var id, userId, productsCost, address, deliveryCost, deliveryWorkerId, couponId int
	var notes string
	var finished, ordered, onTheWay bool
	var orderTime, deliveryTime time.Time
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
			&productsCost, &address, // pq.Array(&orderProducts),
			&deliveryCost, &notes, &finished,
			&deliveryWorkerId, &ordered, &onTheWay, &couponId)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		// finalOrderProducts, err := opd.FindOrderProductsByOrderId(id)
		// if err != nil {
		// 	return make([]entity.Order, 0), err
		// }
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalUser, err := ud.FindUser(userId)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		coupon, err := cd.GetCouponInfo(couponId)
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
			Coupon:           coupon,
			Products:         make([]entity.OrderProduct, 0), /*finalOrderProducts*/
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

func (driver *orderDriver) FindDeliveryWorkerNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	orders := make([]entity.Order, 0)
	rows, err := dbConn.SQL.Query(`
	select 
		id, user_id, order_time, delivery_time, 
		products_cost, address, 
		delivery_cost, notes, finished, 
		delivery_worker_id, ordered, on_the_way, coupon_id  
	from orders where delivery_worker_id = $1 and finished = false order by order_time desc`, userWantedId) //order_products,
	if err != nil {
		return make([]entity.Order, 0), err
	}
	defer rows.Close()

	var id, userId, productsCost, address, deliveryCost, deliveryWorkerId, couponId int
	var notes string
	var finished, ordered, onTheWay bool
	var orderTime, deliveryTime time.Time
	// var orderProducts []string

	for rows.Next() {
		err := rows.Scan(&id, &userId, &orderTime, &deliveryTime,
			&productsCost, &address, // pq.Array(&orderProducts),
			&deliveryCost, &notes, &finished,
			&deliveryWorkerId, &ordered, &onTheWay, &couponId)
		if err != nil {
			return make([]entity.Order, 0), err
		}

		// finalOrderProducts, err := opd.FindOrderProductsByOrderId(id)
		// if err != nil {
		// 	return make([]entity.Order, 0), err
		// }
		finalAddress, err := driver.GetAddressById(address)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		finalUser, err := ud.FindUser(userId)
		if err != nil {
			return make([]entity.Order, 0), err
		}
		coupon, err := cd.GetCouponInfo(couponId)
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
			Coupon:           coupon,
			Products:         make([]entity.OrderProduct, 0), /*finalOrderProducts*/
		}
		orders = append(orders, order)
		if err = rows.Err(); err != nil {
			return make([]entity.Order, 0), err
		}
	}
	return orders, nil
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
	if order.DeliveryWorkerId == 0 {
		return errors.New("order is not set to delivery worker")
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
	if order.DeliveryWorkerId != 0 {
		err = ud.EditDeliveryWorkerBalanceAndCircles(order.DeliveryWorkerId, order.DeliveryCost)
		if err != nil {
			return err
		}
	} else {
		return errors.New("order is not set to delivery worker")
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
	// fmt.Println("Edit Order Statement Is:", stmt)

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
	// fmt.Println("Stores Icrease Balance Items is:", deliveryCost)
	// fmt.Println("Stores Delivery Cost is:", deliveryCost)
	return deliveryCost
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

func (driver *orderDriver) ChangeOrderWorkerId(orderChangeWorkerIdInfo entity.OrderChangeWorkerIdRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	worker, err := ud.FindUser(orderChangeWorkerIdInfo.DeliveryWorkerId)
	if err != nil {
		return err
	}
	if worker.Role != 1 {
		return errors.New("user id is not for a delivery worker")
	}
	stmt := GetOrderChangeWorkerIdStatementString(orderChangeWorkerIdInfo)

	result, err := dbConn.SQL.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("order could not be found")
	}
	driver.cacheorders = make([]entity.Order, 0)
	if orderChangeWorkerIdInfo.Finished {
		driver.FinishOrder(orderChangeWorkerIdInfo.ID)
	}

	return nil
}

func GetOrderChangeWorkerIdStatementString(order entity.OrderChangeWorkerIdRequest) string {
	stmt := `UPDATE orders SET`
	if order.Ordered {
		stmt = stmt + ` ordered = true,`
	}
	if order.OnTheWay {
		stmt = stmt + ` on_the_way = true,`
	}
	if order.DeliveryWorkerId != 0 {
		stmt = stmt + ` delivery_worker_id = ` + fmt.Sprint(order.DeliveryWorkerId) + `,`
	}
	stmt = stmt[0:len(stmt)-1] + ` where id = ` + fmt.Sprint(order.ID) + ` returning id`

	return stmt
}
