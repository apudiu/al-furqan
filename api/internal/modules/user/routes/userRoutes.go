package userroutes

import (
	userHandler "github.com/apudiu/alfurqan/internal/modules/user/handler"
	"github.com/apudiu/alfurqan/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	user := router.Group("user")
	//userProtected := user.Group("", middleware.Protected())

	// Read all
	user.Get("/", userHandler.GetUsers)
	// Create one
	user.Post("/", userHandler.CreateUsers)
	// Read one
	user.Get("/:userId", middleware.Protected(), userHandler.GetUser)
	// Update one
	user.Put("/:userId", userHandler.UpdateUser)
	// Delete one
	user.Delete("/:userId", userHandler.DeleteUser)
}
