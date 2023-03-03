package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type ProductService interface {
	AddProduct(entity.Product) error
	FindAllProducts() []entity.Product
	FindActiveProducts() []entity.Product
	FindNotActiveProducts() []entity.Product
	FindProduct(id entity.ProductInfoRequest) (entity.Product, error)
	FindProductsByCategory(id entity.ProductByCategoryRequest) ([]entity.Product, error)
	EditProduct(productEditInfo entity.ProductEditRequest) error
	ActivateProduct(productEditInfo entity.ProductInfoRequest) error
	DeactivateProduct(productEditInfo entity.ProductInfoRequest) error
	DeleteProduct(product entity.ProductInfoRequest) error
	OrderProduct(id entity.OrderProductRequest) error

	AddMockProducts(products []entity.Product)
}

type productService struct {
	products []entity.Product
}

func NewProductService() ProductService {
	return &productService{}
}

func (service *productService) AddProduct(product entity.Product) error {
	successProduct := product
	for i := 0; i < len(service.products); i++ {
		if service.products[i].Name == product.Name {
			return errors.New("products already exisists")
		}
	}
	if len(service.products) > 0 {
		successProduct.ID = service.products[len(service.products)-1].ID + 1
	} else {
		successProduct.ID = 1
	}

	service.products = append(service.products, successProduct)
	return nil
}

func (service *productService) FindAllProducts() []entity.Product {
	return service.products
}

func (service *productService) FindActiveProducts() []entity.Product {
	var activeProducts []entity.Product
	for i := 0; i < len(service.products); i++ {
		if service.products[i].Active {
			activeProducts = append(activeProducts, service.products[i])
		}
	}
	return activeProducts
}

func (service *productService) FindNotActiveProducts() []entity.Product {
	var notActiveProducts []entity.Product
	for i := 0; i < len(service.products); i++ {
		if !service.products[i].Active {
			notActiveProducts = append(notActiveProducts, service.products[i])
		}
	}
	return notActiveProducts
}

func (service *productService) FindProduct(id entity.ProductInfoRequest) (entity.Product, error) {
	products := service.FindAllProducts()
	var product entity.Product
	if id.ID != 0 {
		for i := 0; i < len(products) && len(products) != 0; i++ {
			if products[i].ID == id.ID {
				product = products[i]
				return product, nil
			}
		}
	} else {
		return product, errors.New("product id cannot be zero")
	}
	return product, errors.New("the product couldn't be found")
}

func (service *productService) FindProductsByCategory(id entity.ProductByCategoryRequest) ([]entity.Product, error) {
	allProducts := service.FindAllProducts()
	var categoryProducts []entity.Product
	if id.ID != 0 {
		for i := 0; i < len(allProducts) && len(allProducts) != 0; i++ {
			if allProducts[i].ProductCategoryId == id.ProductCategoryId && allProducts[i].StoreId == id.StoreId {
				categoryProducts = append(categoryProducts, allProducts[i])
			}
		}
	} else {
		return categoryProducts, errors.New("product id cannot be zero")
	}
	// if len(categoryProducts) == 0 {
	// 	return categoryProducts, errors.New("the category is empty")
	// }
	return categoryProducts, nil
}

func (service *productService) EditProduct(productEditInfo entity.ProductEditRequest) error {
	products := service.FindAllProducts()
	if productEditInfo.ID != 0 {
		for i := 0; i < len(products) && len(products) != 0; i++ {
			if products[i].ID == productEditInfo.ID {
				if productEditInfo.Name != "" {
					products[i].Name = productEditInfo.Name
				}
				if productEditInfo.Price != 0 {
					products[i].Price = productEditInfo.Price
				}
				if productEditInfo.ProductCategoryId != 0 {
					products[i].ProductCategoryId = productEditInfo.ProductCategoryId
				}
				if productEditInfo.Image != "" {
					products[i].Image = productEditInfo.Image
				}
				if productEditInfo.Summary != "" {
					products[i].Summary = productEditInfo.Summary
				}
				if productEditInfo.StoreId != 0 {
					products[i].StoreId = productEditInfo.StoreId
				}
				return nil
			}
		}
	} else {
		return errors.New("product id cannot be zero")
	}
	return errors.New("the product couldn't be found")
}

func (service *productService) ActivateProduct(productEditInfo entity.ProductInfoRequest) error {
	products := service.FindAllProducts()
	if productEditInfo.ID != 0 {
		for i := 0; i < len(products) && len(products) != 0; i++ {
			if products[i].ID == productEditInfo.ID {
				products[i].Active = true
				return nil
			}
		}
	} else {
		return errors.New("product id cannot be zero")
	}
	return errors.New("the product couldn't be found")
}

func (service *productService) DeactivateProduct(productEditInfo entity.ProductInfoRequest) error {
	products := service.FindAllProducts()
	if productEditInfo.ID != 0 {
		for i := 0; i < len(products) && len(products) != 0; i++ {
			if products[i].ID == productEditInfo.ID {
				products[i].Active = false
				return nil
			}
		}
	} else {
		return errors.New("product id cannot be zero")
	}
	return errors.New("the product couldn't be found")
}

func (service *productService) DeleteProduct(productId entity.ProductInfoRequest) error {
	products := service.FindAllProducts()
	var tempProducts []entity.Product
	if productId.ID != 0 {
		for i := 0; i < len(products) && len(products) != 0; i++ {
			if products[i].ID != productId.ID {
				tempProducts = append(tempProducts, products[i])
			}
		}
	} else {
		return errors.New("product id cannot be zero")
	}
	if len(products) != len(tempProducts)+1 {
		return errors.New("product could not be found")
	}
	service.products = tempProducts
	return nil
}

func (service *productService) AddMockProducts(products []entity.Product) {
	service.products = append(service.products, products...)
}

func (service *productService) OrderProduct(id entity.OrderProductRequest) error {
	products := service.FindAllProducts()
	if id.ID != 0 {
		for i := 0; i < len(products) && len(products) != 0; i++ {
			if products[i].ID == id.ID {
				products[i].OrderCount = products[i].OrderCount + id.OrderCount
				return nil
			}
		}
	} else {
		return errors.New("product id cannot be zero")
	}
	return errors.New("the product couldn't be found")
}
