package userroutes

import (
	userHandler "github.com/apudiu/alfurqan/internal/modules/user/handler"
	"github.com/labstack/echo/v4"
)

func SetupPublic(router *echo.Group) {
	user := router.Group("/user")

	// Read all
	user.GET("/", userHandler.GetUsers)
	// Create one
	user.POST("/", userHandler.CreateUsers)
	// Update one
	user.PATCH("/:userId", userHandler.UpdateUser)
	// Delete one
	user.DELETE("/:userId", userHandler.DeleteUser)
}

func SetupPrivate(router *echo.Group) {
	user := router.Group("user")

	// Read one
	user.GET("/:userId", userHandler.GetUser)
}
