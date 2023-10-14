package service

import (
	"errors"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type OrderService interface {
	AddOrder(order entity.Order) error
	FinishOrder(orderId int) error
	EditOrder(orderEditInfo entity.OrderEditRequest) error
	DeleteOrder(order entity.OrderDeleteRequest) error
	ChangeOrderWorkerId(orderChangeWorkerIdInfo entity.OrderChangeWorkerIdRequest) error

	FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindOrder(id entity.OrderInfoRequest) (entity.Order, error)
	FindUserFinishedOrders(userWantedId int) ([]entity.Order, error)
	FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error)
	FindDeliveryWorkerNotFinishedOrders(userWantedId int) ([]entity.Order, error)
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
}

func (service *orderService) FinishOrder(orderId int) error {
	err := service.driver.FinishOrder(orderId)
	if err != nil {
		return err
	}
	return nil
}

func (service *orderService) EditOrder(orderEditInfo entity.OrderEditRequest) error {
	err := service.driver.EditOrder(orderEditInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *orderService) ChangeOrderWorkerId(orderChangeWorkerIdInfo entity.OrderChangeWorkerIdRequest) error {
	err := service.driver.ChangeOrderWorkerId(orderChangeWorkerIdInfo)
	if err != nil {
		return err
	}
	return nil
}

func (service *orderService) DeleteOrder(orderId entity.OrderDeleteRequest) error {
	err := service.driver.DeleteOrder(orderId.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *orderService) FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	orders, paginationInfo, err := service.driver.FindAllOrders(pageLimit, pageOffset)
	if err != nil {
		return make([]entity.Order, 0), paginationInfo, err
	}
	return orders, paginationInfo, nil
}

func (service *orderService) FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	finishedOrders, paginationInfo, err := service.driver.FindFinishedOrders(pageLimit, pageOffset)
	if err != nil {
		return make([]entity.Order, 0), paginationInfo, err
	}
	return finishedOrders, paginationInfo, nil
}

func (service *orderService) FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	notFinishedOrders, paginationInfo, err := service.driver.FindNotFinishedOrders(pageLimit, pageOffset)
	if err != nil {
		return make([]entity.Order, 0), paginationInfo, err
	}
	return notFinishedOrders, paginationInfo, nil
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
}

func (service *orderService) FindUserFinishedOrders(userWantedId int) ([]entity.Order, error) {
	userFinishedrders, err := service.driver.FindUserFinishedOrders(userWantedId)
	if err != nil {
		return make([]entity.Order, 0), err
	}
	return userFinishedrders, nil
}

func (service *orderService) FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	userNotFinishedrders, err := service.driver.FindUserNotFinishedOrders(userWantedId)
	if err != nil {
		return make([]entity.Order, 0), err
	}
	return userNotFinishedrders, nil
}

func (service *orderService) FindDeliveryWorkerNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	deliveryWorkerNotFinishedrders, err := service.driver.FindDeliveryWorkerNotFinishedOrders(userWantedId)
	if err != nil {
		return make([]entity.Order, 0), err
	}
	return deliveryWorkerNotFinishedrders, nil
}
