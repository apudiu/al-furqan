package main

import (
	"github.com/apudiu/alfurqan/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// connect to DB
	database.ConnectDB()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Alhum-du-lillah")
	})

	log.Fatalln(
		app.Listen(":3001"),
	)
}
