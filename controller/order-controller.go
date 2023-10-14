package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	AddOrder(ctx *gin.Context, uc UserController) error
	// ChangeOrderState(ctx *gin.Context, sc StoreController, pc ProductController, uc UserController) error
	FinishOrder(ctx *gin.Context, orderIdValue int) error
	ChangeOrderWorkerId(ctx *gin.Context) error
	FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error)
	FindOrder(ctx *gin.Context, idValue int) (entity.Order, error)
	FindUserFinishedOrders(userWantedId int) ([]entity.Order, error)
	FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error)
	FindDeliveryWorkerNotFinishedOrders(userWantedId int) ([]entity.Order, error)
	EditOrder(ctx *gin.Context) error
	DeleteOrder(ctx *gin.Context) error
}

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return &orderController{
		service: service,
	}
}

func (c *orderController) FindAllOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	return c.service.FindAllOrders(pageLimit, pageOffset)
}

func (c *orderController) AddOrder(ctx *gin.Context, uc UserController) error {
	var order entity.Order
	productsCost := 0
	// storesCount := make([]int, 0)
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		fmt.Println("Its Here")
		return err
	}

	err = validate.Struct(order)
	if err != nil {
		return err
	}

	order.Finished = false
	// order.State = entity.OrderState{
	// 	Text:   "Ordering",
	// 	Number: 1,
	// }
	if len(order.Products) == 0 {
		return errors.New("products list should not be empty")
	}
	for i := 0; i < len(order.Products); i++ {
		productsCost = productsCost + (order.Products[i].ProductPrice * order.Products[i].ProductCount)
	}

	// seen := make(map[int]bool)

	//? Loop through the slice, adding elements to the map if they haven't been seen before
	// for _, val := range order.Products {
	// 	if _, ok := seen[val.StoreId]; !ok {
	// 		seen[val.StoreId] = true
	// 		storesCount = append(storesCount, val.StoreId)
	// 	}
	// }
	// fmt.Println(storesCount)
	// order.DeliveryCost = float32(len(storesCount) * 4000)

	// order.AddonsCost = 0
	order.OrderTime = time.Now()
	order.ProductsCost = productsCost

	err = c.service.AddOrder(order)
	return err
	// productCategories := cst.FindAllOrderCategories()
	// stores := sc.FindAllStores()
	// for i := 0; i < len(productCategories); i++ {
	// 	if productCategories[i].ID == product.OrderCategoryId {
	// 		for i := 0; i < len(stores); i++ {
	// 			if stores[i].ID == product.StoreId {
	// 				err = c.service.AddProduct(product)
	// 				return err
	// 			}
	// 		}
	// 	}
	// }
}

func (c *orderController) FindFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	return c.service.FindFinishedOrders(pageLimit, pageOffset)
}

func (c *orderController) FindNotFinishedOrders(pageLimit int, pageOffset int) ([]entity.Order, entity.PaginationInfo, error) {
	return c.service.FindNotFinishedOrders(pageLimit, pageOffset)
}

func (c *orderController) FindUserFinishedOrders(userWantedId int) ([]entity.Order, error) {
	return c.service.FindUserFinishedOrders(userWantedId)
}

func (c *orderController) FindUserNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	return c.service.FindUserNotFinishedOrders(userWantedId)
}

func (c *orderController) FindDeliveryWorkerNotFinishedOrders(userWantedId int) ([]entity.Order, error) {
	return c.service.FindDeliveryWorkerNotFinishedOrders(userWantedId)
}

func (c *orderController) FindOrder(ctx *gin.Context, idValue int) (entity.Order, error) {
	if idValue == 0 {
		var order entity.Order
		return order, errors.New("order id cannot be zero")
	} else {
		orderId := entity.OrderInfoRequest{
			ID: idValue,
		}
		var order entity.Order
		// err := ctx.ShouldBindJSON(&orderId)
		// if err != nil {
		// 	return order, err
		// }

		err := validate.Struct(order)
		if err != nil {
			return order, err
		}

		order, err = c.service.FindOrder(orderId)
		if err != nil {
			return order, err
		}
		return order, nil
	}
}

func (c *orderController) EditOrder(ctx *gin.Context) error {
	var orderEditInfo entity.OrderEditRequest
	err := ctx.ShouldBindJSON(&orderEditInfo)
	if err != nil {
		return err
	}
	// productCategories := cst.FindAllOrderCategories()
	// stores := sc.FindAllStores()
	// for i := 0; i < len(productCategories); i++ {
	// 	if productCategories[i].ID == productEditInfo.ProductCategoryId {
	// 		for i := 0; i < len(stores); i++ {
	// 			//!!! TODO::: Make The Product Category Be Related To The Required Store
	// 			if stores[i].ID == productEditInfo.StoreId {
	// 				err = c.service.EditProduct(productEditInfo)
	// 				return err
	// 			}
	// 		}
	// 	}
	// }

	err = validate.Struct(orderEditInfo)
	if err != nil {
		return err
	}

	err = c.service.EditOrder(orderEditInfo)
	return err
}

func (c *orderController) ChangeOrderWorkerId(ctx *gin.Context) error {
	var orderChangeWorkerIdInfo entity.OrderChangeWorkerIdRequest
	err := ctx.ShouldBindJSON(&orderChangeWorkerIdInfo)
	if err != nil {
		return err
	}

	err = validate.Struct(orderChangeWorkerIdInfo)
	if err != nil {
		return err
	}

	err = c.service.ChangeOrderWorkerId(orderChangeWorkerIdInfo)
	return err
}

func (c *orderController) FinishOrder(ctx *gin.Context, orderIdValue int) error {
	err := c.service.FinishOrder(orderIdValue)
	return err
}

func (c *orderController) DeleteOrder(ctx *gin.Context) error {
	var orderId entity.OrderDeleteRequest
	err := ctx.ShouldBindJSON(&orderId)
	if err != nil {
		return err
	}
	err = c.service.DeleteOrder(orderId)
	if err != nil {
		return err
	}
	return nil
}
