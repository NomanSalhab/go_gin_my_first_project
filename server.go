package main

import (
	"io"
	"net/http"
	"os"

	"github.com/NomanSalhab/go_gin_my_first_project/controller"
	"github.com/NomanSalhab/go_gin_my_first_project/middlewares"
	"github.com/NomanSalhab/go_gin_my_first_project/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	userService             service.UserService                = service.NewUserService()
	UserController          controller.UserController          = controller.NewUserController(userService)
	storeCategoryService    service.StoreCategoryService       = service.NewStoreCategoryService()
	StoreCategoryController controller.StoreCategoryController = controller.NewStoreCategoryController(storeCategoryService)
	storeService            service.StoreService               = service.NewStoreService()
	StoreController         controller.StoreController         = controller.NewStoreController(storeService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	apiUsersRoutes := server.Group("/api/users")
	{
		apiUsersRoutes.GET("/all_users", func(ctx *gin.Context) {
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

		apiUsersRoutes.POST("/user_info", func(ctx *gin.Context) {
			user, err := UserController.FindUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": user})
			}
		})

		apiUsersRoutes.PUT("/edit_user", func(ctx *gin.Context) {
			err := UserController.EditUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user is edited successfully"})
			}
		})

		apiUsersRoutes.DELETE("/delete_user", func(ctx *gin.Context) {
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
		apiStoreCategoriesRoutes.GET("/get_store_categories", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindAllStoreCategories())
		})

		apiStoreCategoriesRoutes.GET("/get_active_store_categories", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindActiveStoreCategories())
		})

		apiStoreCategoriesRoutes.GET("/get_not_active_store_categories", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindNotActiveStoreCategories())
		})

		apiStoreCategoriesRoutes.POST("/add_store_category", func(ctx *gin.Context) {
			err := StoreCategoryController.AddStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Store Category Added Successfully!"})
			}
		})

		apiStoreCategoriesRoutes.POST("/store_category_info", func(ctx *gin.Context) {
			storeCategory, err := StoreCategoryController.GetStoreCategoryById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"store_category_info": storeCategory})
			}
		})

		apiStoreCategoriesRoutes.PUT("/edit_store_category", func(ctx *gin.Context) {
			err := StoreCategoryController.EditStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store category is edited successfully"})
			}
		})

		apiStoreCategoriesRoutes.DELETE("/delete_store_category", func(ctx *gin.Context) {
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
		apiStoresRoutes.GET("/get_stores", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindAllStores())
		})

		apiStoresRoutes.GET("/get_active_stores", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindActiveStores())
		})

		apiStoresRoutes.GET("/get_not_active_stores", func(ctx *gin.Context) {
			ctx.JSON(200, StoreController.FindNotActiveStores())
		})

		apiStoresRoutes.POST("/add_store", func(ctx *gin.Context) {
			err := StoreController.AddStore(ctx, StoreCategoryController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Store Added Successfully!"})
			}
		})

		apiStoresRoutes.POST("/store_info", func(ctx *gin.Context) {
			store, err := StoreController.GetStoreById(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"store_info": store})
			}
		})

		apiStoresRoutes.PUT("/edit_store", func(ctx *gin.Context) {
			err := StoreController.EditStore(ctx, StoreCategoryController)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store is edited successfully"})
			}
		})

		apiStoresRoutes.DELETE("/delete_store", func(ctx *gin.Context) {
			err := StoreController.DeleteStore(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "store is deleted successfully"})
			}
		})

	}

	/*viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", UserController.ShowAll)
	}*/

	server.Run(":8016")

}
