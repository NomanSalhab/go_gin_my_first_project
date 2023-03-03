package controller

import (
	"errors"
	"time"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	FindAllOrders() []entity.Order
	AddOrder(ctx *gin.Context, uc UserController) error
	ChangeOrderState(ctx *gin.Context, sc StoreController, pc ProductController, uc UserController) error
	FindFinishedOrders() []entity.Order
	FindNotFinishedOrders() []entity.Order
	GetOrderById(ctx *gin.Context) (entity.Order, error)
	EditOrder(ctx *gin.Context, sc StoreController) error
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

func (c *orderController) FindAllOrders() []entity.Order {
	return c.service.FindAllOrders()
}

func (c *orderController) AddOrder(ctx *gin.Context, uc UserController) error {
	var order entity.Order
	productsCost := float32(0.0)
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		return err
	}

	err = validate.Struct(order)
	if err != nil {
		return err
	}

	order.Finished = false
	order.State = entity.OrderState{
		Text:   "Ordering",
		Number: 1,
	}
	if len(order.Products) == 0 {
		return errors.New("products list shold not be empty")
	}
	for i := 0; i < len(order.Products); i++ {
		productsCost = productsCost + (float32(order.Products[i].ProductPrice) * float32(order.Products[i].ProductCount))
	}
	order.AddonsCost = 0
	order.OrderTime = time.Now()

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

func (c *orderController) FindFinishedOrders() []entity.Order {
	return c.service.FindFinishedOrders()
}

func (c *orderController) FindNotFinishedOrders() []entity.Order {
	return c.service.FindNotFinishedOrders()
}

func (c *orderController) GetOrderById(ctx *gin.Context) (entity.Order, error) {
	var orderId entity.OrderInfoRequest
	var order entity.Order
	err := ctx.ShouldBindJSON(&orderId)
	if err != nil {
		return order, err
	}

	err = validate.Struct(order)
	if err != nil {
		return order, err
	}

	order, err = c.service.FindOrder(orderId)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (c *orderController) EditOrder(ctx *gin.Context, sc StoreController) error {
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

func (c *orderController) ChangeOrderState(ctx *gin.Context, sc StoreController, pc ProductController, uc UserController) error {
	var orderStateInfo entity.OrderChangeStateRequest
	err := ctx.ShouldBindJSON(&orderStateInfo)
	if err != nil {
		return err
	}

	err = validate.Struct(orderStateInfo)
	if err != nil {
		return err
	}

	products := pc.FindAllProducts()
	orderProducts, err := c.service.FindOrder(entity.OrderInfoRequest{ID: orderStateInfo.ID})
	users := uc.FindAllUsers()
	stores := sc.FindAllStores()
	if err != nil {
		return err
	}
	err = c.service.ChangeOrderState(orderStateInfo)
	if err == nil {
		if orderStateInfo.Finished {
			for i := 0; i < len(orderProducts.Products); i++ {
				for j := 0; j < len(products); j++ {
					if products[j].ID == orderProducts.Products[i].ProductID {
						products[j].OrderCount = products[j].OrderCount + orderProducts.Products[i].ProductCount
					}
				}
			}
			for i := 0; i < len(orderProducts.Products); i++ {
				for j := 0; j < len(stores); j++ {
					if stores[j].ID == orderProducts.Products[i].StoreId {
						stores[j].Balance = stores[j].Balance + (float32(orderProducts.Products[i].ProductPrice) * float32(orderProducts.Products[i].ProductCount))
					}
				}
			}
			userBalance := int(orderProducts.DeliveryCost) + int(orderProducts.AddonsCost)
			for i := 0; i < len(orderProducts.Products); i++ {
				userBalance = userBalance + (orderProducts.Products[i].ProductPrice * orderProducts.Products[i].ProductCount)
			}
			for j := 0; j < len(users); j++ {
				if users[j].ID == orderProducts.UserID {
					users[j].Circles = users[j].Circles + int(orderProducts.DeliveryCost/1000)
					users[j].Balance = users[j].Balance + userBalance
				}
			}
		}
	}
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
