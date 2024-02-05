package userRoutes

import (
	userHandler "github.com/apudiu/alfurqan/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("user")

	// Read all
	user.Get("/", userHandler.GetUsers)
	// Create a Note
	user.Post("/", userHandler.CreateUsers)
	// Read one
	user.Get("/:userId", userHandler.GetUser)
	// Update one
	user.Put("/:userId", userHandler.UpdateUser)
	// Delete one
	user.Delete("/:userId", userHandler.DeleteUser)
}
