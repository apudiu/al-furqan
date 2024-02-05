package router

import (
	userRoutes "github.com/apudiu/alfurqan/internal/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user := app.Group("user")
	userRoutes.SetupUserRoutes(user)
}
