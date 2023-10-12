package service

import (
	"errors"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type OrderService interface {
	AddOrder(order entity.Order) error
	FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindOrder(id entity.OrderInfoRequest) (entity.Order, error)
	FindUserFinishedOrders(userWantedId int) ([]entity.Order, error)
	FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error)

	FinishOrder(orderId int) error
	EditOrder(orderEditInfo entity.OrderEditRequest) error
	DeleteOrder(order entity.OrderDeleteRequest) error

	// AddMockOrders(orders []entity.Order)
}

type orderService struct {
	// orders []entity.Order
	driver driver.OrderDriver
}

func NewOrderService(driver driver.OrderDriver) OrderService {
	return &orderService{
		driver: driver,
	}
}

func (service *orderService) AddOrder(order entity.Order) error {

	orderTime := time.Now()
	order.OrderTime = orderTime
	for i := 0; i < len(order.Products); i++ {
		order.Products[i].OrderTime = order.OrderTime
	}

	err := service.driver.AddOrder(order)
	if err != nil {
		return err
	}
	return nil
	// successOrder := order
	// if len(service.orders) > 0 {
	// 	successOrder.ID = service.orders[len(service.orders)-1].ID + 1
	// } else {
	// 	successOrder.ID = 1
	// }
	// // if len(order.Products) != 0 {
	// // 	successOrder.Products = order.Products
	// // 	price := float32(0.0)
	// // 	for j := 0; j < len(successOrder.Products); j++ {
	// // 		price = price + (float32(successOrder.Products[j].ProductCount) * float32(successOrder.Products[j].ProductPrice))
	// // 	}
	// // 	successOrder.ProductsCost = price
	// // }
	// if order.ProductsCost != 0 {
	// 	successOrder.ProductsCost = order.ProductsCost
	// }
	// successOrder.DeliveryWorkerId = 0
	// service.orders = append(service.orders, successOrder)
	// return nil
}

func (service *orderService) FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	orders, paginationInfo, err := service.driver.FindAllOrders(pageLimit, pageOffset)
	if err != nil {
		return make([]entity.Order, 0), paginationInfo, err
	}
	return orders, paginationInfo, nil
	// return service.orders
}

func (service *orderService) FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	finishedOrders, paginationInfo, err := service.driver.FindFinishedOrders(pageLimit, pageOffset)
	if err != nil {
		return make([]entity.Order, 0), paginationInfo, err
	}
	return finishedOrders, paginationInfo, nil
	// for i := 0; i < len(service.orders); i++ {
	// 	if service.orders[i].Finished {
	// 		finishedOrders = append(finishedOrders, service.orders[i])
	// 	}
	// }
	// return finishedOrders
}

func (service *orderService) FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	notFinishedOrders, paginationInfo, err := service.driver.FindNotFinishedOrders(pageLimit, pageOffset)
	if err != nil {
		return make([]entity.Order, 0), paginationInfo, err
	}
	return notFinishedOrders, paginationInfo, nil
	// var notFinishedOrders []entity.Order
	// for i := 0; i < len(service.orders); i++ {
	// 	if !service.orders[i].Finished {
	// 		notFinishedOrders = append(notFinishedOrders, service.orders[i])
	// 	}
	// }
	// return notFinishedOrders
}

func (service *orderService) FindOrder(id entity.OrderInfoRequest) (entity.Order, error) {
	order, err := service.driver.FindOrder(id.ID)
	if err != nil {
		return entity.Order{}, err
	}
	if order.ID == 0 {
		return entity.Order{}, errors.New("order could not be found")
	}
	return order, nil
	// orders, _ := service.FindAllOrders(pd, dd)
	// var order entity.Order
	// if id.ID != 0 {
	// 	for i := 0; i < len(orders) && len(orders) != 0; i++ {
	// 		if orders[i].ID == id.ID {
	// 			order = orders[i]
	// 			return order, nil
	// 		}
	// 	}
	// } else {
	// 	return order, errors.New("order id cannot be zero")
	// }
	// return order, errors.New("order could not be found")
	// return entity.Order{}, errors.New("order could not be found")
}

func (service *orderService) FindUserFinishedOrders(userWantedId int) ([]entity.Order, error) {
	userFinishedrders, err := service.driver.FindUserFinishedOrders(userWantedId)
	if err != nil {
		return make([]entity.Order, 0), err
	}
	return userFinishedrders, nil
	// return service.orders
}

func (service *orderService) FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	userNotFinishedrders, err := service.driver.FindUserNotFinishedOrders(userWantedId)
	if err != nil {
		return make([]entity.Order, 0), err
	}
	return userNotFinishedrders, nil
	// return service.orders
}

func (service *orderService) EditOrder(orderEditInfo entity.OrderEditRequest) error {
	err := service.driver.EditOrder(orderEditInfo)
	if err != nil {
		return err
	}
	return nil
	// orders, _ := service.FindAllOrders(pd, dd)
	// if orderEditInfo.ID != 0 {
	// 	for i := 0; i < len(orders) && len(orders) != 0; i++ {
	// 		if orders[i].ID == orderEditInfo.ID {
	// 			if orderEditInfo.Finished {
	// 				orders[i].Finished = orderEditInfo.Finished // true
	// 				orders[i].DeliveryTime = time.Now()
	// 			}
	// 			// if orderEditInfo.State.Text != "" {
	// 			// 	orders[i].State = entity.OrderState{
	// 			// 		Text:   orderEditInfo.State.Text,
	// 			// 		Number: orderEditInfo.State.Number,
	// 			// 	}
	// 			// }
	// 			if orderEditInfo.UserID != 0 {
	// 				orders[i].UserID = orderEditInfo.UserID
	// 			}
	// 			if orderEditInfo.DeliveryTime.String() != "0001-01-01T00:00:00Z" {
	// 				orders[i].DeliveryTime = orderEditInfo.DeliveryTime
	// 			}
	// 			if len(orderEditInfo.Products) != 0 {
	// 				orders[i].Products = orderEditInfo.Products
	// 				price := 0
	// 				for j := 0; j < len(orders[i].Products); j++ {
	// 					price = price + (orders[i].Products[j].ProductCount * orders[i].Products[j].ProductPrice)
	// 				}
	// 				orders[i].ProductsCost = price
	// 			}
	// 			if orderEditInfo.ProductsCost != 0 {
	// 				orders[i].ProductsCost = orderEditInfo.ProductsCost
	// 			}
	// 			if orderEditInfo.DeliveryCost != 0 {
	// 				orders[i].DeliveryCost = orderEditInfo.DeliveryCost
	// 			}
	// 			if orderEditInfo.Notes != "" {
	// 				orders[i].Notes = orderEditInfo.Notes
	// 			}
	// 			// if orderEditInfo.AddonsCost != 0 {
	// 			// 	orders[i].AddonsCost = orderEditInfo.AddonsCost
	// 			// }
	// 			if orderEditInfo.Address.ID != 0 {
	// 				orders[i].Address = orderEditInfo.Address
	// 			}
	// 			if orderEditInfo.DeliveryWorkerId != 0 {
	// 				orders[i].DeliveryWorkerId = orderEditInfo.DeliveryWorkerId
	// 			}
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("order id cannot be zero")
	// }
	// return errors.New("order could not be found")
}

func (service *orderService) FinishOrder(orderId int) error {
	err := service.driver.FinishOrder(orderId)
	if err != nil {
		return err
	}
	return nil
	// orders, _ := service.FindAllOrders(pd, dd)
	// if orderChangeStateInfo.ID != 0 {
	// 	for i := 0; i < len(orders) && len(orders) != 0; i++ {
	// 		if orders[i].ID == orderChangeStateInfo.ID {
	// 			// orders[i].State = entity.OrderState{
	// 			// 	Text:   orderChangeStateInfo.State.Text,
	// 			// 	Number: orderChangeStateInfo.State.Number,
	// 			// }
	// 			orders[i].Finished = orderChangeStateInfo.Finished
	// 			if orderChangeStateInfo.Finished {
	// 				orders[i].Finished = orderChangeStateInfo.Finished // true
	// 				orders[i].DeliveryTime = time.Now()
	// 			}
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("order id cannot be zero")
	// }
	// return errors.New("order could not be found")
}

// func (service *orderService) FinishOrder(orderInfo entity.OrderInfoRequest) error {
// 	// orders, _ := service.FindAllOrders(pd, dd)
// 	// if orderInfo.ID != 0 {
// 	// 	for i := 0; i < len(orders) && len(orders) != 0; i++ {
// 	// 		if orders[i].ID == orderInfo.ID {
// 	// 			// orders[i].State = entity.OrderState{
// 	// 			// 	Text:   "Delivered",
// 	// 			// 	Number: 4,
// 	// 			// }
// 	// 			orders[i].Finished = true
// 	// 			orders[i].DeliveryTime = time.Now()
// 	// 			return nil
// 	// 		}
// 	// 	}
// 	// } else {
// 	// 	return errors.New("order id cannot be zero")
// 	// }
// 	return errors.New("order could not be found")
// }

func (service *orderService) DeleteOrder(orderId entity.OrderDeleteRequest) error {
	err := service.driver.DeleteOrder(orderId.ID)
	if err != nil {
		return err
	}
	return nil
	// orders, _ := service.FindAllOrders(pd, dd)
	// var tempOrder []entity.Order
	// if orderId.ID != 0 {
	// 	for i := 0; i < len(orders) && len(orders) != 0; i++ {
	// 		if orders[i].ID != orderId.ID {
	// 			tempOrder = append(tempOrder, orders[i])
	// 		}
	// 	}
	// } else {
	// 	return errors.New("order id cannot be zero")
	// }
	// if len(orders) != len(tempOrder)+1 {
	// 	return errors.New("order could not be found")
	// }
	// // service.orders = tempOrder
}

// func (service *orderService) AddMockOrders(orders []entity.Order) {
// 	service.orders = append(service.orders, orders...)
// }
