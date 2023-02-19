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
	userService    service.UserService       = service.New()
	UserController controller.UserController = controller.New(userService)
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
		apiUsersRoutes.GET("/get", func(ctx *gin.Context) {
			ctx.JSON(200, UserController.FindAll())
		})

		apiUsersRoutes.POST("/signup", func(ctx *gin.Context) {
			err := UserController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User Input Accepted!"})
			}
		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", UserController.ShowAll)
	}

	server.Run(":8016")

}
