package controller

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAllProducts() []entity.Product
	AddProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController) error
	FindActiveProducts() []entity.Product
	FindNotActiveProducts() []entity.Product
	GetProductById(ctx *gin.Context) (entity.Product, error)
	GetProductByCategory(ctx *gin.Context) ([]entity.Product, error)
	EditProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController) error
	ActivateProduct(ctx *gin.Context) error
	DeactivateProduct(ctx *gin.Context) error
	DeleteProduct(ctx *gin.Context) error
	OrderProduct(ctx *gin.Context) error
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		service: service,
	}
}

func (c *productController) FindAllProducts() []entity.Product {
	return c.service.FindAllProducts()
}

func (c *productController) AddProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController) error {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}

	err = validate.Struct(product)
	if err != nil {
		return err
	}

	productCategories := cst.FindAllProductCategories()
	stores := sc.FindAllStores()
	for i := 0; i < len(productCategories); i++ {
		if productCategories[i].ID == product.ProductCategoryId {
			for i := 0; i < len(stores); i++ {
				if stores[i].ID == product.StoreId {

					err = c.service.AddProduct(product)
					return err
				}
			}
		}
	}
	return errors.New("product category or store does not exist")
}

func (c *productController) FindActiveProducts() []entity.Product {
	return c.service.FindActiveProducts()
}

func (c *productController) FindNotActiveProducts() []entity.Product {
	return c.service.FindNotActiveProducts()
}

func (c *productController) GetProductById(ctx *gin.Context) (entity.Product, error) {
	var productId entity.ProductInfoRequest
	var product entity.Product
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return product, err
	}
	product, err = c.service.FindProduct(productId)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (c *productController) EditProduct(ctx *gin.Context, cst ProductCategoryController, sc StoreController) error {
	var productEditInfo entity.ProductEditRequest
	err := ctx.ShouldBindJSON(&productEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productEditInfo)
	if err != nil {
		return err
	}

	productCategories := cst.FindAllProductCategories()
	stores := sc.FindAllStores()
	for i := 0; i < len(productCategories); i++ {
		if productCategories[i].ID == productEditInfo.ProductCategoryId {
			for j := 0; j < len(stores); j++ {
				//!!! TODO::: Make The Product Category Be Related To The Required Store
				if stores[j].ID == productEditInfo.StoreId /*&& (stores[j].ID == productCategories[i].StoreId)*/ {
					err = c.service.EditProduct(productEditInfo)
					return err
				}
			}
		}
	}
	return errors.New("product category or store does not exist")
}

func (c *productController) ActivateProduct(ctx *gin.Context) error {
	var productEditInfo entity.ProductInfoRequest
	err := ctx.ShouldBindJSON(&productEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productEditInfo)
	if err != nil {
		return err
	}

	err = c.service.ActivateProduct(productEditInfo)
	return err
}

func (c *productController) DeactivateProduct(ctx *gin.Context) error {
	var productEditInfo entity.ProductInfoRequest
	err := ctx.ShouldBindJSON(&productEditInfo)
	if err != nil {
		return err
	}
	err = validate.Struct(productEditInfo)
	if err != nil {
		return err
	}

	err = c.service.DeactivateProduct(productEditInfo)
	return err
}

func (c *productController) DeleteProduct(ctx *gin.Context) error {
	var productId entity.ProductInfoRequest
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return err
	}
	err = validate.Struct(productId)
	if err != nil {
		return err
	}

	err = c.service.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}

func (c *productController) GetProductByCategory(ctx *gin.Context) ([]entity.Product, error) {
	var productId entity.ProductByCategoryRequest
	var products []entity.Product
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return products, err
	}
	err = validate.Struct(productId)
	if err != nil {
		return products, err
	}

	products, err = c.service.FindProductsByCategory(productId)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (c *productController) OrderProduct(ctx *gin.Context) error {
	var productId entity.OrderProductRequest
	err := ctx.ShouldBindJSON(&productId)
	if err != nil {
		return err
	}
	err = validate.Struct(productId)
	if err != nil {
		return err
	}

	err = c.service.OrderProduct(productId)
	if err != nil {
		return err
	}
	return nil
}
