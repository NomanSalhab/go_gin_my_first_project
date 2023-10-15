package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/NomanSalhab/go_gin_my_first_project/controller"
	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
	"github.com/NomanSalhab/go_gin_my_first_project/middlewares"
	"github.com/NomanSalhab/go_gin_my_first_project/service"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	userDriver     driver.UserDriver         = driver.NewUserDriver()
	userService    service.UserService       = service.NewUserService(userDriver)
	UserController controller.UserController = controller.NewUserController(userService)

	storeCategoryDriver     driver.StoreCategoryDriver         = driver.NewStoreCategoryDriver()
	storeCategoryService    service.StoreCategoryService       = service.NewStoreCategoryService(storeCategoryDriver)
	StoreCategoryController controller.StoreCategoryController = controller.NewStoreCategoryController(storeCategoryService)

	storeDriver     driver.StoreDriver         = driver.NewStoreDriver()
	storeService    service.StoreService       = service.NewStoreService(storeDriver)
	StoreController controller.StoreController = controller.NewStoreController(storeService)

	productCategoryDriver     driver.ProductCategoryDriver         = driver.NewProductCategoryDriver()
	productCategoryService    service.ProductCategoryService       = service.NewProductCategoryService(productCategoryDriver)
	ProductCategoryController controller.ProductCategoryController = controller.NewProductCategoryController(productCategoryService)

	sliderDriver     driver.SliderDriver         = driver.NewSliderDriver()
	sliderService    service.SliderService       = service.NewSliderService(sliderDriver)
	sliderController controller.SliderController = controller.NewSliderController(sliderService)

	areaDriver     driver.AreaDriver         = driver.NewAreaDriver()
	areaService    service.AreaService       = service.NewAreaService(areaDriver)
	areaController controller.AreaController = controller.NewAreaController(areaService)

	detailsDriver    driver.DetailDriver         = driver.NewDetailDriver()
	detailService    service.DetailService       = service.NewDetailService(detailsDriver)
	detailController controller.DetailController = controller.NewDetailController(detailService)

	complaintsDriver    driver.ComplaintDriver         = driver.NewComplaintDriver()
	complaintService    service.ComplaintService       = service.NewComplaintService(complaintsDriver)
	complaintController controller.ComplaintController = controller.NewComplaintController(complaintService)

	couponsDriver    driver.CouponDriver         = driver.NewCouponDriver()
	couponService    service.CouponService       = service.NewCouponService(couponsDriver)
	couponController controller.CouponController = controller.NewCouponController(couponService)

	productsDriver    driver.ProductDriver         = driver.NewProductDriver()
	productService    service.ProductService       = service.NewProductService(productsDriver)
	ProductController controller.ProductController = controller.NewProductController(productService)

	ordersDriver    driver.OrderDriver         = driver.NewOrderDriver()
	orderService    service.OrderService       = service.NewOrderService(ordersDriver)
	OrderController controller.OrderController = controller.NewOrderController(orderService)

	homePageDriver     driver.HomePageDriver         = driver.NewHomePageDriver()
	homePageService    service.HomePageService       = service.NewHomePageService(homePageDriver)
	homePageController controller.HomePageController = controller.NewHomePageController(homePageService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// func getAllRows(conn *sql.DB) error {
// 	rows, err := conn.Query("select id, user_id, delivery_time, products from orders")
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
// 	var userId, id int
// 	var deliveryTime time.Time
// 	var products string
// 	for rows.Next() {
// 		err := rows.Scan(&id, &userId, &deliveryTime, &products)
// 		if err != nil {
// 			log.Println(err)
// 			return err
// 		}
// 		fmt.Println("Record is:", userId, deliveryTime, products)
// 		if err = rows.Err(); err != nil {
// 			log.Fatal("error Scanning Rows!")
// 		}
// 		fmt.Println("------------------------")
// 	}
// 	return nil
// }
// func addMockData() {
// 	var users []entity.User
// 	users = append(users, entity.User{
// 		ID:       1,
// 		Name:     "Noman Salhab",
// 		Phone:    "0992008516",
// 		Password: "nomanos.net",
// 		Balance:  150,
// 		Active:   true,
// 	})
// 	users = append(users, entity.User{
// 		ID:       2,
// 		Name:     "Fouad Aljundi",
// 		Phone:    "0936425373",
// 		Password: "fouadovich.com",
// 		Balance:  105,
// 		Active:   false,
// 	})
// 	userService.AddMockUsers(users)
// 	var storeCategorys []entity.StoreCategory
// 	storeCategorys = append(storeCategorys, entity.StoreCategory{
// 		ID:     1,
// 		Name:   "Snacks",
// 		Active: true,
// 	})
// 	storeCategorys = append(storeCategorys, entity.StoreCategory{
// 		ID:     2,
// 		Name:   "Super Markets",
// 		Active: true,
// 	})
// 	storeCategoryService.AddMockStoreCategories(storeCategorys)
// 	var stores []entity.Store
// 	stores = append(stores, entity.Store{
// 		ID:              1,
// 		Name:            "The Golden Plate",
// 		StoreCategoryId: 1,
// 		Image:           "Image 1",
// 		Balance:         250000,
// 		Active:          true,
// 		DeliveryRent:    3500,
// 		Discount:        100,
// 	})
// 	stores = append(stores, entity.Store{
// 		ID:              2,
// 		Name:            "Hervy",
// 		StoreCategoryId: 1,
// 		Image:           "Image 2",
// 		Balance:         90000,
// 		Active:          true,
// 		DeliveryRent:    3500,
// 		Discount:        50,
// 	})
// 	storeService.AddMockStores(stores)
// 	var productCategories []entity.ProductCategory
// 	productCategories = append(productCategories, entity.ProductCategory{
// 		ID:      1,
// 		Name:    "Shawarma",
// 		StoreId: 1,
// 		Active:  true,
// 	})
// 	productCategories = append(productCategories, entity.ProductCategory{
// 		ID:      2,
// 		Name:    "Cheese",
// 		StoreId: 2,
// 		Active:  true,
// 	})
// 	productCategortService.AddMockProductCategories(productCategories)
// 	var sliders []entity.Slider
// 	sliders = append(sliders, entity.Slider{
// 		ID:      1,
// 		Image:   "",
// 		StoreId: 1,
// 		Active:  true,
// 	})
// 	sliders = append(sliders, entity.Slider{
// 		ID:      2,
// 		Image:   "",
// 		StoreId: 2,
// 		Active:  true,
// 	})
// 	sliderService.AddMockSliders(sliders)
// 	var products []entity.Product
// 	products = append(products, entity.Product{
// 		ID:                1,
// 		Name:              "Shawrma",
// 		StoreId:           1,
// 		ProductCategoryId: 1,
// 		Image:             "Image 3",
// 		Summary:           "No Cheese",
// 		Price:             7500,
// 		OrderCount:        12,
// 		Active:            true,
// 	})
// 	products = append(products, entity.Product{
// 		ID:                2,
// 		Name:              "Shawrma With Cheese",
// 		StoreId:           2,
// 		ProductCategoryId: 2,
// 		Image:             "Image 4",
// 		Summary:           "With Cheese",
// 		Price:             9000,
// 		OrderCount:        9,
// 		Active:            true,
// 	})
// 	productService.AddMockProducts(products)
// }

func main() {

	setupLogOutput()

	// Connect To DB
	// conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=circle_delivery_app user=postgres password=postgresqlgolangpass")
	// if err != nil {
	// 	log.Fatalf("Unable to connect: %v\n", err)
	// }
	// defer conn.Close()
	// Test Connection
	// err = conn.Ping()
	// if err != nil {
	// 	log.Fatalf("cannot Connect To Database")
	// }
	// Get Rows From Table
	// err = getAllRows(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Insert A row
	// query := `insert into users (name, phone, password, balance, active) values ($1, $2, $3, $4, $5)`
	// _, err = conn.Exec(query, "Rahaf Salhab", "0936425377", "rahafpass", 0, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Re Get Rows
	// err = getAllRows(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Update A Row
	// stmt := `update users set name = $1 where id = $2`
	// _, err = conn.Exec(stmt, "Rahofeh Salhab", "3")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Get Rows
	// Get One Row By ID
	// query := `select name, phone, balance, circles from users where id = $1`
	// var name, phone string
	// var balance, circles int
	// row := conn.QueryRow(query, 1)
	// err = row.Scan(&name, &phone, &balance, &circles)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Query Row Returns:", name, phone, balance, circles)
	// Delete A Row
	// stmt := `delete from users where id = $1`
	// _, err = conn.Exec(stmt, 4)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Get Rows
	// addMockData()

	driver.ConnectSQL("host=localhost port=5432 dbname=circle_delivery_app user=postgres password=postgresqlgolangpass")

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	apiUsersRoutes := server.Group("/api/users")
	{
		apiUsersRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, UserController.FindAllUsers())
		})

		apiUsersRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, UserController.FindActiveUsers())
		})

		apiUsersRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, UserController.FindNotActiveUsers())
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

		// apiUsersRoutes.POST("/info", func(ctx *gin.Context) {
		// 	user, err := UserController.FindUser(ctx, 0)
		// 	if err != nil {
		// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	} else {
		// 		ctx.JSON(http.StatusOK, gin.H{"message": user})
		// 	}
		// })

		apiUsersRoutes.GET("/info/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			user, err := UserController.FindUser(ctx, idValue)
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

		apiUsersRoutes.GET("/user_circles/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			circles, err := UserController.UserCircles(ctx, idValue)
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

		apiUsersRoutes.GET("/user_addresses/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			addresses, err := UserController.UserAddressesList(ctx, idValue)
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

		apiUsersRoutes.POST("/delete_address", func(ctx *gin.Context) {
			err := UserController.UserDeleteAddress(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "address is deleted successfully"})
			}
		})

		apiUsersRoutes.PUT("/activate", func(ctx *gin.Context) {
			err := UserController.ActivateUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user is activated successfully"})
			}
		})

		apiUsersRoutes.PUT("/deactivate", func(ctx *gin.Context) {
			err := UserController.DeactivateUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user is deactivated successfully"})
			}
		})
	}

	apiStoreCategoriesRoutes := server.Group("/api/store_categories")
	{
		apiStoreCategoriesRoutes.GET("/all", func(ctx *gin.Context) {
			value := StoreCategoryController.FindAllStoreCategories()
			if len(value) != 0 {
				ctx.JSON(200, value)
			} else {
				ctx.JSON(200, make([]string, 0))
			}

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

	apiAreasRoutes := server.Group("/api/areas")
	{
		apiAreasRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, areaController.FindAllAreas())
		})

		apiAreasRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, areaController.FindActiveAreas())
		})

		apiAreasRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, areaController.FindNotActiveAreas())
		})

		apiAreasRoutes.POST("/activate", func(ctx *gin.Context) {
			err := areaController.ActivateArea(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "area was activated successfully"})
			}
		})

		apiAreasRoutes.POST("/deactivate", func(ctx *gin.Context) {
			err := areaController.DeactivateArea(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "area was deactivated successfully"})
			}
		})

		apiAreasRoutes.POST("/add", func(ctx *gin.Context) {
			err := areaController.AddArea(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "area added successfully!"})
			}
		})

		apiAreasRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := areaController.EditArea(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "area is edited successfully"})
			}
		})

		apiAreasRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := areaController.DeleteArea(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "area is deleted successfully"})
			}
		})

	}

	apiStoresRoutes := server.Group("/api/stores")
	{
		apiStoresRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindAllStores())
		})

		apiStoresRoutes.POST("/active", func(ctx *gin.Context) {
			stores := StoreController.FindActiveStores(ctx)
			if stores == nil {
				ctx.JSON(http.StatusOK, gin.H{"stores": make([]entity.Store, 0)})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"stores": stores})
			}
		})

		apiStoresRoutes.GET("/not_active", func(ctx *gin.Context) {
			stores := StoreController.FindNotActiveStores()
			if stores == nil {
				ctx.JSON(http.StatusOK, gin.H{"stores": make([]entity.Store, 0)})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"stores": stores})
			}
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

	apiSlidersRoutes := server.Group("/api/sliders")
	{
		apiSlidersRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, sliderController.FindAllSliders())
		})

		apiSlidersRoutes.GET("/active", func(ctx *gin.Context) {
			ctx.JSON(200, sliderController.FindActiveSliders())
		})

		apiSlidersRoutes.GET("/not_active", func(ctx *gin.Context) {
			ctx.JSON(200, sliderController.FindNotActiveSliders())
		})

		apiSlidersRoutes.POST("/store_sliders", func(ctx *gin.Context) {
			sliders, err := sliderController.GetSlidersByStore(ctx, StoreController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"sliders": sliders})
			}
		})

		apiSlidersRoutes.POST("/add", func(ctx *gin.Context) {
			err := sliderController.AddSlider(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "slider added successfully!"})
			}
		})

		apiSlidersRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := sliderController.EditSlider(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "product category is edited successfully"})
			}
		})

		apiSlidersRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := sliderController.DeleteSlider(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "slider is deleted successfully"})
			}
		})

	}

	apiDetailsRoutes := server.Group("/api/details")
	{
		apiDetailsRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, detailController.FindAllDetails())
		})

		apiDetailsRoutes.GET("/addons", func(ctx *gin.Context) {
			ctx.JSON(200, detailController.FindAllAddons())
		})

		apiDetailsRoutes.GET("/flavors", func(ctx *gin.Context) {
			ctx.JSON(200, detailController.FindAllFlavors())
		})

		apiDetailsRoutes.GET("/volumes", func(ctx *gin.Context) {
			ctx.JSON(200, detailController.FindAllVolumes())
		})

		apiDetailsRoutes.POST("/add", func(ctx *gin.Context) {
			err := detailController.AddDetail(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "detail added successfully!"})
			}
		})

		apiDetailsRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := detailController.EditDetail(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "detail is edited successfully"})
			}
		})

		apiDetailsRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := detailController.DeleteDetail(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "detail is deleted successfully"})
			}
		})

	}

	apiCouponsRoutes := server.Group("/api/coupons")
	{
		apiCouponsRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, couponController.FindAllCoupons())
		})

		apiCouponsRoutes.POST("/add", func(ctx *gin.Context) {
			err := couponController.AddCoupon(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "coupon added successfully!"})
			}
		})

		apiCouponsRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := couponController.EditCoupon(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "coupon is edited successfully"})
			}
		})

		apiCouponsRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := couponController.DeleteCoupon(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "coupon is deleted successfully"})
			}
		})

		apiCouponsRoutes.GET("/info/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			coupon, err := couponController.GetCouponInfo(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"coupon_info": coupon})
			}
		})

		apiCouponsRoutes.GET("/activate/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = couponController.ActivateCoupon(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "coupon activated successfully"})
			}
		})

		apiCouponsRoutes.GET("/deactivate/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = couponController.DeactivateCoupon(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "coupon deactivated successfully"})
			}
		})

		apiCouponsRoutes.GET("/enable_free_delivery/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = couponController.EnableFreeDelivery(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "enabled free delivery successfully"})
			}
		})

		apiCouponsRoutes.GET("/disable_free_delivery/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = couponController.DisableFreeDelivery(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "disabled free delivery successfully"})
			}
		})

		apiCouponsRoutes.GET("/enable_from_products_cost/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = couponController.EnableFromProducts(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "enabled from products cost successfully"})
			}
		})

		apiCouponsRoutes.GET("/disable_from_products_cost/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = couponController.DisableFromProducts(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "disabled from products cost successfully"})
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
		apiOrdersRoutes.GET("/all/:page/:limit", func(ctx *gin.Context) {
			page := ctx.Param("page")
			limit := ctx.Param("limit")
			pageValue, err := strconv.Atoi(page)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			limitValue, err := strconv.Atoi(limit)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			offsetValue := (pageValue - 1) * limitValue
			allOrders, paginationInfo, err := OrderController.FindAllOrders(limitValue, offsetValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"all_orders": allOrders, "pagination_info": paginationInfo})
			}
		})

		apiOrdersRoutes.GET("/finished/:page/:limit", func(ctx *gin.Context) {
			page := ctx.Param("page")
			limit := ctx.Param("limit")
			pageValue, err := strconv.Atoi(page)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			limitValue, err := strconv.Atoi(limit)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			offsetValue := (pageValue - 1) * limitValue
			finishedOrders, paginationInfo, err := OrderController.FindFinishedOrders(limitValue, offsetValue)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"finished_orders": finishedOrders, "pagination_info": paginationInfo})
			}
		})

		apiOrdersRoutes.GET("/not_finished/:page/:limit", func(ctx *gin.Context) {
			page := ctx.Param("page")
			limit := ctx.Param("limit")
			pageValue, err := strconv.Atoi(page)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			limitValue, err := strconv.Atoi(limit)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			offsetValue := (pageValue - 1) * limitValue
			notFinishedOrders, paginationInfo, err := OrderController.FindNotFinishedOrders(limitValue, offsetValue)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"not_finished_orders": notFinishedOrders, "pagination_info": paginationInfo})
			}
		})

		apiOrdersRoutes.GET("/user_finished/:user_id", func(ctx *gin.Context) {
			userId := ctx.Param("user_id")
			userIdValue, err := strconv.Atoi(userId)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			userFinishedOrders, err := OrderController.FindUserFinishedOrders(userIdValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"user_finished_orders": userFinishedOrders})
			}
		})

		apiOrdersRoutes.GET("/user_not_finished/:user_id", func(ctx *gin.Context) {
			userId := ctx.Param("user_id")
			userIdValue, err := strconv.Atoi(userId)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			userNotFinishedOrders, err := OrderController.FindUserNotFinishedOrders(userIdValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"user_not_finished_orders": userNotFinishedOrders})
			}
		})

		apiOrdersRoutes.GET("/delivery_worker_not_finished/:delivery_worker_id", func(ctx *gin.Context) {
			deliveryWorkerId := ctx.Param("delivery_worker_id")
			deliveryWorkerIdValue, err := strconv.Atoi(deliveryWorkerId)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			deliveryWorkerNotFinishedOrders, err := OrderController.FindDeliveryWorkerNotFinishedOrders(deliveryWorkerIdValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"user_not_finished_orders": deliveryWorkerNotFinishedOrders})
			}
		})

		apiOrdersRoutes.POST("/add", func(ctx *gin.Context) {
			err := OrderController.AddOrder(ctx, UserController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order added successfully!"})
			}
		})

		// apiOrdersRoutes.POST("/change_state", func(ctx *gin.Context) {
		// 	err := OrderController.ChangeOrderState(ctx, StoreController, ProductController, UserController)
		// 	if err != nil {
		// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	} else {
		// 		ctx.JSON(http.StatusOK, gin.H{"message": "order added successfully!"})
		// 	}
		// })

		apiOrdersRoutes.GET("/finish_order/:id", func(ctx *gin.Context) {
			orderId := ctx.Param("id")
			orderIdValue, err := strconv.Atoi(orderId)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = OrderController.FinishOrder(ctx, orderIdValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order finished successfully!"})
			}
		})

		apiOrdersRoutes.GET("/info/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			order, err := OrderController.FindOrder(ctx, idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"order_info": order})
			}
		})

		apiOrdersRoutes.PUT("/edit", func(ctx *gin.Context) {
			err := OrderController.EditOrder(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order is edited successfully"})
			}
		})

		apiOrdersRoutes.PUT("/change_worker_id", func(ctx *gin.Context) {
			err := OrderController.ChangeOrderWorkerId(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "order worker id is set successfully"})
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

	apiHomePageRoutes := server.Group("/api/home_page")
	{
		apiHomePageRoutes.GET("/get/:limit/:app_version/:area_id", func(ctx *gin.Context) {
			limit := ctx.Param("limit")
			appVersion := ctx.Param("app_version")
			areaId := ctx.Param("area_id")
			limitValue, err := strconv.Atoi(limit)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"limit_error": "limit not valid"})
			}
			appVersionValue, err := strconv.ParseFloat(appVersion, 32)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"app_version_error": "app version not valid"})
			}
			areaIdValue, err := strconv.Atoi(areaId)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"area_id_error": "area id not valid"})
			}
			homePage, err := homePageController.GetHomePage(limitValue, float32(appVersionValue), areaIdValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "home_page": homePage})
			} else {
				ctx.JSON(200, gin.H{"home_page": homePage})
			}
		})

	}

	apiComplaintsRoutes := server.Group("/api/complaints")
	{
		apiComplaintsRoutes.GET("/all", func(ctx *gin.Context) {
			allComplaints, err := complaintController.FindAllComplaints()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"all_complaints": allComplaints})
		})

		apiComplaintsRoutes.GET("/about_delivery", func(ctx *gin.Context) {
			aboutDeliveryComplaints, err := complaintController.FindAboutDeliveryComplaints()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"about_delivery_complaints": aboutDeliveryComplaints})
		})

		apiComplaintsRoutes.GET("/about_the_app", func(ctx *gin.Context) {
			aboutTheAppComplaints, err := complaintController.FindAboutTheAppComplaints()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"about_the_app_complaints": aboutTheAppComplaints})
		})

		apiComplaintsRoutes.GET("/improvement_suggestion", func(ctx *gin.Context) {
			improvementSuggestionComplaints, err := complaintController.FindImprovementSuggestionComplaints()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"improvement_suggestion_complaints": improvementSuggestionComplaints})
		})

		apiComplaintsRoutes.GET("/other_reason", func(ctx *gin.Context) {
			otherReasonComplaints, err := complaintController.FindOtherReasonComplaints()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"other_reason_complaints": otherReasonComplaints})
		})

		apiComplaintsRoutes.GET("/user_complaints/:user_id", func(ctx *gin.Context) {
			id := ctx.Param("user_id")
			idValue, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			userComplaints, err := complaintController.FindUserComplaints(idValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"user_complaints": userComplaints})
		})

		apiComplaintsRoutes.POST("/add", func(ctx *gin.Context) {
			err := complaintController.AddComplaint(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "complaint added successfully!"})
			}
		})

		apiComplaintsRoutes.DELETE("/delete", func(ctx *gin.Context) {
			err := complaintController.DeleteComplaint(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "complaint is deleted successfully"})
			}
		})

	}

	/*viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", UserController.ShowAll)
	}*/

	server.Run(":8016")

}
