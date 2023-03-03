package main

import (
	"io"
	"net/http"
	"os"

	"github.com/NomanSalhab/go_gin_my_first_project/controller"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/middlewares"
	"github.com/NomanSalhab/go_gin_my_first_project/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	userService               service.UserService                  = service.NewUserService()
	UserController            controller.UserController            = controller.NewUserController(userService)
	storeCategoryService      service.StoreCategoryService         = service.NewStoreCategoryService()
	StoreCategoryController   controller.StoreCategoryController   = controller.NewStoreCategoryController(storeCategoryService)
	storeService              service.StoreService                 = service.NewStoreService()
	StoreController           controller.StoreController           = controller.NewStoreController(storeService)
	productCategortService    service.ProductCategoryService       = service.NewProductCategoryService()
	ProductCategoryController controller.ProductCategoryController = controller.NewProductCategoryController(productCategortService)
	productService            service.ProductService               = service.NewProductService()
	ProductController         controller.ProductController         = controller.NewProductController(productService)
	orderService              service.OrderService                 = service.NewOrderService()
	OrderController           controller.OrderController           = controller.NewOrderController(orderService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func addMockData() {
	var users []entity.User
	users = append(users, entity.User{
		ID:       1,
		Name:     "Noman Salhab",
		Phone:    "0992008516",
		Password: "nomanos.net",
		Balance:  150,
		Active:   true,
	})
	users = append(users, entity.User{
		ID:       2,
		Name:     "Fouad Aljundi",
		Phone:    "0936425373",
		Password: "fouadovich.com",
		Balance:  105,
		Active:   true,
	})
	userService.AddMockUsers(users)
	var stores []entity.Store
	stores = append(stores, entity.Store{
		ID:              1,
		Name:            "The Golden Plate",
		StoreCategoryId: 1,
		Image:           "Image 1",
		Balance:         250000,
		Active:          true,
		DeliveryRent:    3500,
	})
	stores = append(stores, entity.Store{
		ID:              2,
		Name:            "Hervy",
		StoreCategoryId: 1,
		Image:           "Image 2",
		Balance:         90000,
		Active:          true,
		DeliveryRent:    3500,
	})
	storeService.AddMockStores(stores)
	var storeCategorys []entity.StoreCategory
	storeCategorys = append(storeCategorys, entity.StoreCategory{
		ID:     1,
		Name:   "Snacks",
		Active: true,
	})
	storeCategorys = append(storeCategorys, entity.StoreCategory{
		ID:     2,
		Name:   "Super Markets",
		Active: true,
	})
	storeCategoryService.AddMockStoreCategories(storeCategorys)
	var products []entity.Product
	products = append(products, entity.Product{
		ID:                1,
		Name:              "Shawrma",
		StoreId:           1,
		ProductCategoryId: 1,
		Image:             "Image 3",
		Summary:           "No Cheese",
		Price:             7500,
		OrderCount:        12,
		Active:            true,
	})
	products = append(products, entity.Product{
		ID:                2,
		Name:              "Shawrma With Cheese",
		StoreId:           2,
		ProductCategoryId: 1,
		Image:             "Image 4",
		Summary:           "With Cheese",
		Price:             9000,
		OrderCount:        9,
		Active:            true,
	})
	productService.AddMockProducts(products)
	var productCategories []entity.ProductCategory
	productCategories = append(productCategories, entity.ProductCategory{
		ID:      1,
		Name:    "Sawarma",
		StoreId: 1,
		Active:  true,
	})
	productCategories = append(productCategories, entity.ProductCategory{
		ID:      2,
		Name:    "Cheese",
		StoreId: 2,
		Active:  true,
	})
	productCategortService.AddMockProductCategories(productCategories)
}

func main() {

	setupLogOutput()

	addMockData()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	apiUsersRoutes := server.Group("/api/users")
	{
		apiUsersRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, UserController.FindAllUsers())
		})

		apiUsersRoutes.POST("/signup", func(ctx *gin.Context) {
			err := UserController.SaveUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User Added Successfully!"})
			}
		})

		apiUsersRoutes.POST("/login", func(ctx *gin.Context) {
			user, err := UserController.LoginUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "User Logged In Successfully!",
					"user":    user,
				})
			}
		})

		apiUsersRoutes.POST("/info", func(ctx *gin.Context) {
			user, err := UserController.FindUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": user})
			}
		})

		apiUsersRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := UserController.EditUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user is edited successfully"})
			}
		})

		apiUsersRoutes.POST("/user_circles", func(ctx *gin.Context) {
			circles, err := UserController.UserCircles(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": circles})
			}
		})

		apiUsersRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := UserController.DeleteUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user is deleted successfully"})
			}
		})

		apiUsersRoutes.POST("/user_addresses", func(ctx *gin.Context) {
			addresses, err := UserController.UserAddressesList(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": addresses})
			}
		})

		apiUsersRoutes.POST("/add_address", func(ctx *gin.Context) {
			err := UserController.UserAddAddress(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "address is added successfully"})
			}
		})

	}

	apiStoreCategoriesRoutes := server.Group("/api/store_categories")
	{
		apiStoreCategoriesRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindAllStoreCategories())
		})

		apiStoreCategoriesRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindActiveStoreCategories())
		})

		apiStoreCategoriesRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindNotActiveStoreCategories())
		})

		apiStoreCategoriesRoutes.POST("/add", func(ctx *gin.Context) {
			err := StoreCategoryController.AddStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Store Category Added Successfully!"})
			}
		})

		apiStoreCategoriesRoutes.POST("/info", func(ctx *gin.Context) {
			storeCategory, err := StoreCategoryController.GetStoreCategoryById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"store_category_info": storeCategory})
			}
		})

		apiStoreCategoriesRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := StoreCategoryController.EditStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store category is edited successfully"})
			}
		})

		apiStoreCategoriesRoutes.PUT("/activate", func(ctx *gin.Context) {
			err := StoreCategoryController.ActivateStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store category is activated successfully"})
			}
		})

		apiStoreCategoriesRoutes.PUT("/deactivate", func(ctx *gin.Context) {
			err := StoreCategoryController.DeactivateStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store category is deactivated successfully"})
			}
		})

		apiStoreCategoriesRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := StoreCategoryController.DeleteStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store category is deleted successfully"})
			}
		})

	}

	apiStoresRoutes := server.Group("/api/stores")
	{
		apiStoresRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindAllStores())
		})

		apiStoresRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindActiveStores())
		})

		apiStoresRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindNotActiveStores())
		})

		apiStoresRoutes.POST("/add", func(ctx *gin.Context) {
			err := StoreController.AddStore(ctx, StoreCategoryController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Store Added Successfully!"})
			}
		})

		apiStoresRoutes.POST("/info", func(ctx *gin.Context) {
			store, err := StoreController.GetStoreById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"store_info": store})
			}
		})

		apiStoresRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := StoreController.EditStore(ctx, StoreCategoryController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store is edited successfully"})
			}
		})

		apiStoresRoutes.PUT("/activate", func(ctx *gin.Context) {
			err := StoreController.ActivateStore(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store is activated successfully"})
			}
		})

		apiStoresRoutes.PUT("/deactivate", func(ctx *gin.Context) {
			err := StoreController.DeactivateStore(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store is deactivated successfully"})
			}
		})

		apiStoresRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := StoreController.DeleteStore(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store is deleted successfully"})
			}
		})

	}

	apiProductCategoriesRoutes := server.Group("/api/product_categories")
	{
		apiProductCategoriesRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, ProductCategoryController.FindAllProductCategories())
		})

		apiProductCategoriesRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, ProductCategoryController.FindActiveProductCategories())
		})

		apiProductCategoriesRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, ProductCategoryController.FindNotActiveProductCategories())
		})

		apiProductCategoriesRoutes.POST("/store_product_categories", func(ctx *gin.Context) {
			storeProductCategories, err := ProductCategoryController.GetProductCategoriesByStore(ctx, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"store_product_categories": storeProductCategories})
			}
		})

		apiProductCategoriesRoutes.POST("/add", func(ctx *gin.Context) {
			err := ProductCategoryController.AddProductCategory(ctx, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product category added successfully!"})
			}
		})

		apiProductCategoriesRoutes.POST("/info", func(ctx *gin.Context) {
			store, err := ProductCategoryController.GetProductCategoryById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"product_category_info": store})
			}
		})

		apiProductCategoriesRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := ProductCategoryController.EditProductCategory(ctx, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product category is edited successfully"})
			}
		})

		apiProductCategoriesRoutes.PUT("/activate", func(ctx *gin.Context) {
			err := ProductCategoryController.ActivateProductCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product category is activated successfully"})
			}
		})

		apiProductCategoriesRoutes.PUT("/deactivate", func(ctx *gin.Context) {
			err := ProductCategoryController.DeactivateProductCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product category is deactivated successfully"})
			}
		})

		apiProductCategoriesRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := ProductCategoryController.DeleteProductCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product category is deleted successfully"})
			}
		})

	}

	apiProductsRoutes := server.Group("/api/products")
	{
		apiProductsRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, ProductController.FindAllProducts())
		})

		apiProductsRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, ProductController.FindActiveProducts())
		})

		apiProductsRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, ProductController.FindNotActiveProducts())
		})

		apiProductsRoutes.POST("/add", func(ctx *gin.Context) {
			err := ProductController.AddProduct(ctx, ProductCategoryController, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product added successfully!"})
			}
		})

		apiProductsRoutes.POST("/info", func(ctx *gin.Context) {
			product, err := ProductController.GetProductById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"product_info": product})
			}
		})

		apiProductsRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := ProductController.EditProduct(ctx, ProductCategoryController, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product is edited successfully"})
			}
		})

		apiProductsRoutes.PUT("/activate", func(ctx *gin.Context) {
			err := ProductController.ActivateProduct(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product is activated successfully"})
			}
		})

		apiProductsRoutes.PUT("/deactivate", func(ctx *gin.Context) {
			err := ProductController.DeactivateProduct(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product is deactivated successfully"})
			}
		})

		apiProductsRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := ProductController.DeleteProduct(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product is deleted successfully"})
			}
		})

		apiProductsRoutes.POST("/product_category_products", func(ctx *gin.Context) {
			products, err := ProductController.GetProductByCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"category_products": products})
			}
		})

	}

	apiOrdersRoutes := server.Group("/api/orders")
	{
		apiOrdersRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, OrderController.FindAllOrders())
		})

		apiOrdersRoutes.GET("/finished", func(ctx *gin.Context) {
			ctx.JSON(200, OrderController.FindFinishedOrders())
		})

		apiOrdersRoutes.GET("/not_finished", func(ctx *gin.Context) {
			ctx.JSON(200, OrderController.FindNotFinishedOrders())
		})

		apiOrdersRoutes.POST("/add", func(ctx *gin.Context) {
			err := OrderController.AddOrder(ctx, UserController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order added successfully!"})
			}
		})

		apiOrdersRoutes.POST("/change_state", func(ctx *gin.Context) {
			err := OrderController.ChangeOrderState(ctx, StoreController, ProductController, UserController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order added successfully!"})
			}
		})

		apiOrdersRoutes.POST("/info", func(ctx *gin.Context) {
			product, err := OrderController.GetOrderById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"order_info": product})
			}
		})

		apiOrdersRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := OrderController.EditOrder(ctx, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order is edited successfully"})
			}
		})

		apiOrdersRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := OrderController.DeleteOrder(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order is deleted successfully"})
			}
		})

	}

	/*viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", UserController.ShowAll)
	}*/

	server.Run(":8016")

}
