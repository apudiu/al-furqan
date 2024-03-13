package router

import (
	authRoutes "github.com/apudiu/alfurqan/internal/modules/auth/routes"
	userRoutes "github.com/apudiu/alfurqan/internal/modules/user/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// versioning
	apiVer := app.Group("v1")

	// Register routes

	userRoutes.Setup(apiVer)
	authRoutes.Setup(apiVer)
}
