package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type OrderService interface {
	AddOrder(entity.Order) error
	ChangeOrderState(entity.OrderChangeStateRequest) error
	FindAllOrders() []entity.Order
	FindFinishedOrders() []entity.Order
	FindNotFinishedOrders() []entity.Order
	FindOrder(id entity.OrderInfoRequest) (entity.Order, error)
	EditOrder(orderEditInfo entity.OrderEditRequest) error
	DeleteOrder(order entity.OrderDeleteRequest) error
	FinishOrder(orderInfo entity.OrderInfoRequest) error

	AddMockOrders(orders []entity.Order)
}

type orderService struct {
	orders []entity.Order
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (service *orderService) AddOrder(order entity.Order) error {
	successOrder := order
	if len(service.orders) > 0 {
		successOrder.ID = service.orders[len(service.orders)-1].ID + 1
	} else {
		successOrder.ID = 1
	}
	if len(order.Products) != 0 {
		successOrder.Products = order.Products
		price := float32(0.0)
		for j := 0; j < len(successOrder.Products); j++ {
			price = price + (float32(successOrder.Products[j].ProductCount) * float32(successOrder.Products[j].ProductPrice))
		}
		successOrder.ProductsCost = price
	}
	if order.ProductsCost != 0 {
		successOrder.ProductsCost = order.ProductsCost
	}
	successOrder.DeliveryWorkerId = 0
	service.orders = append(service.orders, successOrder)
	return nil
}

func (service *orderService) FindAllOrders() []entity.Order {
	return service.orders
}

func (service *orderService) FindFinishedOrders() []entity.Order {
	var finishedOrders []entity.Order
	for i := 0; i < len(service.orders); i++ {
		if service.orders[i].Finished {
			finishedOrders = append(finishedOrders, service.orders[i])
		}
	}
	return finishedOrders
}

func (service *orderService) FindNotFinishedOrders() []entity.Order {
	var notFinishedOrders []entity.Order
	for i := 0; i < len(service.orders); i++ {
		if !service.orders[i].Finished {
			notFinishedOrders = append(notFinishedOrders, service.orders[i])
		}
	}
	return notFinishedOrders
}

func (service *orderService) FindOrder(id entity.OrderInfoRequest) (entity.Order, error) {
	orders := service.FindAllOrders()
	var order entity.Order
	if id.ID != 0 {
		for i := 0; i < len(orders) && len(orders) != 0; i++ {
			if orders[i].ID == id.ID {
				order = orders[i]
				return order, nil
			}
		}
	} else {
		return order, errors.New("order id cannot be zero")
	}
	return order, errors.New("order could not be found")
}

func (service *orderService) EditOrder(orderEditInfo entity.OrderEditRequest) error {
	orders := service.FindAllOrders()
	if orderEditInfo.ID != 0 {
		for i := 0; i < len(orders) && len(orders) != 0; i++ {
			if orders[i].ID == orderEditInfo.ID {
				if orderEditInfo.Finished {
					orders[i].Finished = true
				}
				if orderEditInfo.State.Text != "" {
					orders[i].State = entity.OrderState{
						Text:   orderEditInfo.State.Text,
						Number: orderEditInfo.State.Number,
					}
				}
				if orderEditInfo.UserID != 0 {
					orders[i].UserID = orderEditInfo.UserID
				}
				if orderEditInfo.DeliveryTime.String() != "0001-01-01T00:00:00Z" {
					orders[i].DeliveryTime = orderEditInfo.DeliveryTime
				}
				if len(orderEditInfo.Products) != 0 {
					orders[i].Products = orderEditInfo.Products
					price := float32(0.0)
					for j := 0; j < len(orders[i].Products); j++ {
						price = price + (float32(orders[i].Products[j].ProductCount) * float32(orders[i].Products[j].ProductPrice))
					}
					orders[i].ProductsCost = price
				}
				if orderEditInfo.ProductsCost != 0 {
					orders[i].ProductsCost = orderEditInfo.ProductsCost
				}
				if orderEditInfo.DeliveryCost != 0 {
					orders[i].DeliveryCost = orderEditInfo.DeliveryCost
				}
				if orderEditInfo.Addons != "" {
					orders[i].Addons = orderEditInfo.Addons
				}
				if orderEditInfo.AddonsCost != 0 {
					orders[i].AddonsCost = orderEditInfo.AddonsCost
				}
				if orderEditInfo.Address.ID != 0 {
					orders[i].Address = orderEditInfo.Address
				}
				if orderEditInfo.DeliveryWorkerId != 0 {
					orders[i].DeliveryWorkerId = orderEditInfo.DeliveryWorkerId
				}
				if orderEditInfo.Finished {
					orders[i].Finished = orderEditInfo.Finished
				}
				return nil
			}
		}
	} else {
		return errors.New("order id cannot be zero")
	}
	return errors.New("order could not be found")
}

func (service *orderService) ChangeOrderState(orderChangeStateInfo entity.OrderChangeStateRequest) error {
	orders := service.FindAllOrders()

	if orderChangeStateInfo.ID != 0 {
		for i := 0; i < len(orders) && len(orders) != 0; i++ {
			if orders[i].ID == orderChangeStateInfo.ID {
				orders[i].State = entity.OrderState{
					Text:   orderChangeStateInfo.State.Text,
					Number: orderChangeStateInfo.State.Number,
				}
				orders[i].Finished = orderChangeStateInfo.Finished
				return nil
			}
		}
	} else {
		return errors.New("order id cannot be zero")
	}
	return errors.New("order could not be found")
}

func (service *orderService) FinishOrder(orderInfo entity.OrderInfoRequest) error {
	orders := service.FindAllOrders()
	if orderInfo.ID != 0 {
		for i := 0; i < len(orders) && len(orders) != 0; i++ {
			if orders[i].ID == orderInfo.ID {
				orders[i].State = entity.OrderState{
					Text:   "Delivered",
					Number: 4,
				}
				orders[i].Finished = true
				return nil
			}
		}
	} else {
		return errors.New("order id cannot be zero")
	}
	return errors.New("order could not be found")
}

func (service *orderService) DeleteOrder(orderId entity.OrderDeleteRequest) error {
	orders := service.FindAllOrders()
	var tempOrder []entity.Order
	if orderId.ID != 0 {
		for i := 0; i < len(orders) && len(orders) != 0; i++ {
			if orders[i].ID != orderId.ID {
				tempOrder = append(tempOrder, orders[i])
			}
		}
	} else {
		return errors.New("order id cannot be zero")
	}
	if len(orders) != len(tempOrder)+1 {
		return errors.New("order could not be found")
	}
	service.orders = tempOrder
	return nil
}

func (service *orderService) AddMockOrders(orders []entity.Order) {
	service.orders = append(service.orders, orders...)
}
