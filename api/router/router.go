package router

import (
	userRoutes "github.com/apudiu/alfurqan/internal/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// versioning
	v1 := app.Group("v1")

	// groups

	user := v1.Group("user")
	userRoutes.SetupUserRoutes(user)
}
