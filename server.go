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

		apiUsersRoutes.POST("/user_info", func(ctx *gin.Context) {
			user, err := UserController.FindUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": user})
			}
		})

	}

	apiStoreCategoriesRoutes := server.Group("/api/store_categories")
	{
		apiStoreCategoriesRoutes.GET("/get", func(ctx *gin.Context) {
			ctx.JSON(200, StoreCategoryController.FindAllStoreCategories())
		})

		apiStoreCategoriesRoutes.POST("/add", func(ctx *gin.Context) {
			err := StoreCategoryController.SaveStoreCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Store Category Added Successfully!"})
			}
		})

	}

	/*viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", UserController.ShowAll)
	}*/

	server.Run(":8016")

}
