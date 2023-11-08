package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
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

	// AddMockProducts(products []entity.Product)
}

type productService struct {
	// products []entity.Product
	driver driver.ProductDriver
}

func NewProductService(driver driver.ProductDriver) ProductService {
	return &productService{
		driver: driver,
	}
}

func (service *productService) AddProduct(product entity.Product) error {
	productsList, err := service.driver.FindAllProducts()
	if err != nil {
		return err
	}
	for i := 0; i < len(productsList); i++ {
		if productsList[i].Name == product.Name && productsList[i].StoreId == product.StoreId {
			return errors.New("product name with store_id already exists")
		}
	}
	err = service.driver.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
	// successProduct := product
	// for i := 0; i < len(service.products); i++ {
	// 	if service.products[i].Name == product.Name && service.products[i].StoreId == product.StoreId {
	// 		return errors.New("product already exisists")
	// 	}
	// }
	// if len(service.products) > 0 {
	// 	successProduct.ID = service.products[len(service.products)-1].ID + 1
	// } else {
	// 	successProduct.ID = 1
	// }
	// service.products = append(service.products, successProduct)
	// return nil
}

func (service *productService) FindAllProducts() []entity.Product {
	allProducts, err := service.driver.FindAllProducts()
	if err != nil {
		return make([]entity.Product, 0)
	}
	return allProducts
	// return service.products
}

func (service *productService) FindActiveProducts() []entity.Product {
	activeProducts, err := service.driver.FindActiveProducts()
	if err != nil {
		return make([]entity.Product, 0)
	}
	return activeProducts
	// var activeProducts []entity.Product
	// for i := 0; i < len(service.products); i++ {
	// 	if service.products[i].Active {
	// 		activeProducts = append(activeProducts, service.products[i])
	// 	}
	// }
	// return activeProducts
}

func (service *productService) FindNotActiveProducts() []entity.Product {
	notAProducts, err := service.driver.FindNotActiveProducts()
	if err != nil {
		return make([]entity.Product, 0)
	}
	return notAProducts
	// var notActiveProducts []entity.Product
	// for i := 0; i < len(service.products); i++ {
	// 	if !service.products[i].Active {
	// 		notActiveProducts = append(notActiveProducts, service.products[i])
	// 	}
	// }
	// return notActiveProducts
}

func (service *productService) FindProduct(id entity.ProductInfoRequest) (entity.Product, error) {
	product, err := service.driver.FindProduct(id.ID)
	if err != nil {
		return entity.Product{}, err
	}
	if product.Name == "" {
		return product, errors.New("the product couldn't be found")
	}
	return product, nil
	// products := service.FindAllProducts()
	// var product entity.Product
	// if id.ID != 0 {
	// 	for i := 0; i < len(products) && len(products) != 0; i++ {
	// 		if products[i].ID == id.ID {
	// 			product = products[i]
	// 			return product, nil
	// 		}
	// 	}
	// } else {
	// 	return product, errors.New("product id cannot be zero")
	// }
	// return product, errors.New("the product couldn't be found")
}

func (service *productService) FindProductsByCategory(id entity.ProductByCategoryRequest) ([]entity.Product, error) {
	productCategoryProducts, err := service.driver.FindProductByProductCategory( /*id.StoreId, */ id.ProductCategoryId)
	if err != nil {
		return make([]entity.Product, 0), err
	}
	return productCategoryProducts, nil
	// allProducts := service.FindAllProducts()
	// var categoryProducts []entity.Product
	// // if id.ID != 0 {
	// for i := 0; i < len(allProducts) && len(allProducts) != 0; i++ {
	// 	if allProducts[i].ProductCategoryId == id.ProductCategoryId && allProducts[i].StoreId == id.StoreId {
	// 		categoryProducts = append(categoryProducts, allProducts[i])
	// 	}
	// }
	// // } else {
	// // return categoryProducts, errors.New("product id cannot be zero")
	// // }
	// // if len(categoryProducts) == 0 {
	// // 	return categoryProducts, errors.New("the category is empty")
	// // }
	// return categoryProducts, nil
}

func (service *productService) EditProduct(productEditInfo entity.ProductEditRequest) error {
	_, err := service.driver.EditProduct(productEditInfo)
	if err != nil {
		return err
	}
	return nil
	// products := service.FindAllProducts()
	// if productEditInfo.ID != 0 {
	// 	for i := 0; i < len(products) && len(products) != 0; i++ {
	// 		if products[i].ID == productEditInfo.ID {
	// 			if productEditInfo.Name != "" {
	// 				products[i].Name = productEditInfo.Name
	// 			}
	// 			if productEditInfo.Price != 0 {
	// 				products[i].Price = productEditInfo.Price
	// 			}
	// 			if productEditInfo.ProductCategoryId != 0 {
	// 				products[i].ProductCategoryId = productEditInfo.ProductCategoryId
	// 			}
	// 			if productEditInfo.Image != "" {
	// 				products[i].Image = productEditInfo.Image
	// 			}
	// 			if productEditInfo.Summary != "" {
	// 				products[i].Summary = productEditInfo.Summary
	// 			}
	// 			if productEditInfo.StoreId != 0 {
	// 				products[i].StoreId = productEditInfo.StoreId
	// 			}
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product id cannot be zero")
	// }
	// return errors.New("the product couldn't be found")
}

func (service *productService) ActivateProduct(productEditInfo entity.ProductInfoRequest) error {
	err := service.driver.ActivateProduct(productEditInfo)
	if err != nil {
		return err
	}
	return nil
	// products := service.FindAllProducts()
	// if productEditInfo.ID != 0 {
	// 	for i := 0; i < len(products) && len(products) != 0; i++ {
	// 		if products[i].ID == productEditInfo.ID {
	// 			products[i].Active = true
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product id cannot be zero")
	// }
	// return errors.New("the product couldn't be found")
}

func (service *productService) DeactivateProduct(productEditInfo entity.ProductInfoRequest) error {
	err := service.driver.DeactivateProduct(productEditInfo)
	if err != nil {
		return err
	}
	return nil
	// products := service.FindAllProducts()
	// if productEditInfo.ID != 0 {
	// 	for i := 0; i < len(products) && len(products) != 0; i++ {
	// 		if products[i].ID == productEditInfo.ID {
	// 			products[i].Active = false
	// 			return nil
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product id cannot be zero")
	// }
	// return errors.New("the product couldn't be found")
}

func (service *productService) DeleteProduct(productId entity.ProductInfoRequest) error {
	err := service.driver.DeleteProduct(productId.ID)
	if err != nil {
		return err
	}
	return nil
	// products := service.FindAllProducts()
	// var tempProducts []entity.Product
	// if productId.ID != 0 {
	// 	for i := 0; i < len(products) && len(products) != 0; i++ {
	// 		if products[i].ID != productId.ID {
	// 			tempProducts = append(tempProducts, products[i])
	// 		}
	// 	}
	// } else {
	// 	return errors.New("product id cannot be zero")
	// }
	// if len(products) != len(tempProducts)+1 {
	// 	return errors.New("product could not be found")
	// }
	// service.products = tempProducts
	// return nil
}

// func (service *productService) AddMockProducts(products []entity.Product) {
// 	service.products = append(service.products, products...)
// }

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
