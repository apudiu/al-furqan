package main

import (
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	// connect to DB
	database.ConnectDB()

	// setup router
	router.SetupRoutes(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Alhum-du-lillah")
	})

	log.Fatalln(
		app.Listen(":3001"),
	)
}
